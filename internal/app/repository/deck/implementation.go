package deck

import (
	"errors"
	"github.com/google/uuid"
	"github.com/secmohammed/deck-poc/container"
	"github.com/secmohammed/deck-poc/internal/app/dto"
	"github.com/secmohammed/deck-poc/internal/entity"
	"gorm.io/gorm"
)

type cr struct {
	c container.Container
}

type DeckResponse struct {
	*entity.Deck
	Remaining int
}
type CreateDeckResponse struct {
	DeckResponse
}
type GetDeckResponse struct {
	DeckResponse
}

func (cr cr) DrawDeckCards(id uuid.UUID, count int) (*entity.Deck, error) {
	var (
		cds []entity.CardDeck
	)
	// get the first item that's drawn, so we can go lesser than its order if exists.
	result := cr.c.Database().Get().Where("deck_id = ? AND drawn = ?",
		id,
		true,
	).Order("\"order\" ASC").Limit(1).Find(&cds)
	if result.Error != nil {
		return nil, result.Error
	}
	// if record is not found, it means that we are at the entrypoint of the deck and this is the first draw.
	if len(cds) == 0 {
		result = cr.c.Database().Get().Where("deck_id = ?", id).Order("\"order\" DESC").Limit(count).Find(&cds)
		if result.Error != nil {
			return nil, result.Error
		}
		if len(cds) != count {
			return nil, errors.New("cannot draw cards more than the available ones")
		}

	} else {
		// since cds will always return one item because of LIMIT 1, we can safely access the first index
		// we are going to get the ones that have less order than current drawn card and limit the selected ones with the passed parameter of count,
		// so we can have access of the next ones to draw them later.
		result = cr.c.Database().Get().Where("deck_id = ? AND \"order\" < ?", id, cds[0].Order).Order("\"order\" DESC").Limit(count).Find(&cds)
		if result.Error != nil {
			return nil, result.Error
		}
		if len(cds) != count {
			return nil, errors.New("all cards are already drawn")
		}
	}
	var orders []int
	for _, c := range cds {
		orders = append(orders, c.Order)
	}
	// update the selected ones by their orders to be drawn
	result = cr.c.Database().Get().Model(&entity.CardDeck{}).Where("deck_id = ? AND \"order\" IN ?", id, orders).Update("drawn", true)
	if result.Error != nil {
		return nil, result.Error
	}
	// get drawn cards
	result = cr.c.Database().Get().Preload("Card").Where("deck_id = ? AND drawn = ?", id, true).Order("\"order\" DESC").Find(&cds)
	return &entity.Deck{
		BaseModel: entity.BaseModel{
			ID: id,
		},
		Cards: cds,
	}, result.Error

}
func (cr cr) Get(id uuid.UUID) (*GetDeckResponse, error) {
	var d entity.Deck
	result := cr.c.Database().Get().Preload("Cards", func(db *gorm.DB) *gorm.DB {
		return db.Where("drawn = false")
	}).Preload("Cards.Card").Where("id = ?", id).First(&d)
	return &GetDeckResponse{
		DeckResponse{
			&d,
			len(d.Cards),
		},
	}, result.Error

}
func (cr cr) Create(in *dto.CreateDeckDTO, cards []*entity.CardTemplate) (*CreateDeckResponse, error) {
	d := &entity.Deck{
		Shuffle: *in.Shuffle,
	}
	result := cr.c.Database().Get().Create(d)
	if result.Error != nil {
		return nil, result.Error
	}
	var cardDecks []*entity.CardDeck
	for i, card := range cards {
		cardDecks = append(cardDecks, &entity.CardDeck{
			CardID: card.ID,
			DeckID: d.ID,
			Order:  i + 1,
		})
	}
	result = cr.c.Database().Get().CreateInBatches(cardDecks, 52)
	return &CreateDeckResponse{
		DeckResponse{
			d,
			len(cardDecks),
		},
	}, result.Error
}

func NewDeckRepository(c container.Container) DeckRepository {
	return &cr{c}
}
