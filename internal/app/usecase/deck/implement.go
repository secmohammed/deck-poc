package deck

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/secmohammed/deck-poc/internal/app/dto"
	"github.com/secmohammed/deck-poc/internal/app/repository/card_template"
	"github.com/secmohammed/deck-poc/internal/app/repository/deck"
	"github.com/secmohammed/deck-poc/internal/entity"
)

type usecase struct {
	dr  deck.DeckRepository
	ctr card_template.CardTemplateRepository
}

func (uc *usecase) Get(id uuid.UUID) (*deck.GetDeckResponse, error) {
	return uc.dr.Get(id)
}

func NewUseCase(dr deck.DeckRepository, ctr card_template.CardTemplateRepository) UseCase {
	return &usecase{dr: dr, ctr: ctr}
}
func (uc *usecase) Draw(id uuid.UUID, count int) (*entity.Deck, error) {
	return uc.dr.DrawDeckCards(id, count)
}
func (uc *usecase) Create(in *dto.CreateDeckDTO) (*deck.CreateDeckResponse, error) {
	var (
		filteredCards []*entity.CardTemplate
		err           error
	)
	if in.CardFilters != nil {
		filteredCards, err = uc.ctr.FindWhereCodeIn(in.CardFilters, *in.Shuffle)
		if err != nil {
			return nil, err
		}
		if len(filteredCards) != len(in.CardFilters) {
			return nil, ErrInvalidDeckCardFilter
		}
	} else {
		filteredCards, err = uc.ctr.GetAll(*in.Shuffle)
		if err != nil {
			return nil, err
		}
	}

	data, err := uc.dr.Create(in, filteredCards)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnexpected, err.Error())
	}
	return data, nil
}
