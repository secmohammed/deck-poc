package deck

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/secmohammed/deck-poc/internal/entity"
	"github.com/secmohammed/deck-poc/tests"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCannotDrawDeckCardIfIDIsInvalid(t *testing.T) {
	writer := tests.MakeRequest("PATCH", "/api/decks/12312312", nil, router(app))
	assert.Equal(t, http.StatusUnprocessableEntity, writer.Code)

}
func TestItShouldNotDrawDeckCardByIDIfNotFound(t *testing.T) {
	writer := tests.MakeRequest("PATCH", fmt.Sprintf("/api/decks/%s", uuid.New().String()), nil, router(app))
	assert.Equal(t, http.StatusNotFound, writer.Code)
}

func TestItShouldFallbackToOneIfPassedCountIsInvalid(t *testing.T) {
	var deck entity.Deck
	var cards []*entity.CardTemplate

	result := app.Database().Get().Create(&deck)
	assert.NoError(t, result.Error)
	result = app.Database().Get().Limit(4).Find(&cards)
	assert.NoError(t, result.Error)
	var cardDecks []*entity.CardDeck
	// associate the opened deck with the 4 cards
	for i, card := range cards {
		cardDecks = append(cardDecks, &entity.CardDeck{
			CardID: card.ID,
			DeckID: deck.ID,
			Order:  i + 1,
		})
	}
	result = app.Database().Get().CreateInBatches(cardDecks, 4)
	assert.NoError(t, result.Error)

	writer := tests.MakeRequest("PATCH", fmt.Sprintf("/api/decks/%s?count=hello", deck.ID.String()), nil, router(app))
	assert.Equal(t, http.StatusOK, writer.Code)
	var res map[string]interface{}
	err := json.Unmarshal(writer.Body.Bytes(), &res)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(res["cards"].([]interface{})))
}

func TestItShouldDrawCardsWithNCount(t *testing.T) {
	var deck entity.Deck
	var cards []*entity.CardTemplate

	result := app.Database().Get().Create(&deck)
	assert.NoError(t, result.Error)
	result = app.Database().Get().Limit(4).Find(&cards)
	assert.NoError(t, result.Error)
	var cardDecks []*entity.CardDeck
	// associate the opened deck with the 4 cards
	for i, card := range cards {
		cardDecks = append(cardDecks, &entity.CardDeck{
			CardID: card.ID,
			DeckID: deck.ID,
			Order:  i + 1,
		})
	}
	result = app.Database().Get().CreateInBatches(cardDecks, 4)
	assert.NoError(t, result.Error)

	writer := tests.MakeRequest("PATCH", fmt.Sprintf("/api/decks/%s?count=2", deck.ID.String()), nil, router(app))
	assert.Equal(t, http.StatusOK, writer.Code)
	var res map[string]interface{}
	err := json.Unmarshal(writer.Body.Bytes(), &res)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(res["cards"].([]interface{})))
}
func TestItShouldFallbackToOneIfPassedCountIsLessThanOne(t *testing.T) {
	var deck entity.Deck
	var cards []*entity.CardTemplate

	result := app.Database().Get().Create(&deck)
	assert.NoError(t, result.Error)
	result = app.Database().Get().Limit(4).Find(&cards)
	assert.NoError(t, result.Error)
	var cardDecks []*entity.CardDeck
	// associate the opened deck with the 4 cards
	for i, card := range cards {
		cardDecks = append(cardDecks, &entity.CardDeck{
			CardID: card.ID,
			DeckID: deck.ID,
			Order:  i + 1,
		})
	}
	result = app.Database().Get().CreateInBatches(cardDecks, 4)
	assert.NoError(t, result.Error)

	writer := tests.MakeRequest("PATCH", fmt.Sprintf("/api/decks/%s?count=-1", deck.ID.String()), nil, router(app))
	assert.Equal(t, http.StatusOK, writer.Code)
	var res map[string]interface{}
	err := json.Unmarshal(writer.Body.Bytes(), &res)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(res["cards"].([]interface{})))
}
func TestItShouldNotDrawCardsWithMoreThanAvailableAtOpenDeck(t *testing.T) {
	var deck entity.Deck
	var cards []*entity.CardTemplate

	result := app.Database().Get().Create(&deck)
	assert.NoError(t, result.Error)
	result = app.Database().Get().Limit(4).Find(&cards)
	assert.NoError(t, result.Error)
	var cardDecks []*entity.CardDeck
	// associate the opened deck with the 4 cards
	for i, card := range cards {
		cardDecks = append(cardDecks, &entity.CardDeck{
			CardID: card.ID,
			DeckID: deck.ID,
			Order:  i + 1,
		})
	}
	result = app.Database().Get().CreateInBatches(cardDecks, 4)
	assert.NoError(t, result.Error)

	writer := tests.MakeRequest("PATCH", fmt.Sprintf("/api/decks/%s?count=10", deck.ID.String()), nil, router(app))
	assert.Equal(t, http.StatusNotFound, writer.Code)
	assert.Contains(t, writer.Body.String(), "cannot draw cards more than the available ones")
	//var res map[string]interface{}
	//err := json.Unmarshal(writer.Body.Bytes(), &res)
	//assert.NoError(t, err)
	//assert.Equal(t, 2, len(res["cards"].([]interface{})))
}
