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

func TestCannotGetOpenDeckIfIDIsInvalid(t *testing.T) {
	writer := tests.MakeRequest("GET", "/api/decks/12312312", nil, router(app))
	assert.Equal(t, http.StatusUnprocessableEntity, writer.Code)

}

func TestItShouldNotGetOpenDeckByIDIfNotFound(t *testing.T) {
	writer := tests.MakeRequest("GET", fmt.Sprintf("/api/decks/%s", uuid.New().String()), nil, router(app))
	assert.Equal(t, http.StatusNotFound, writer.Code)
}

func TestItShouldGetOpenDeckByIDIfFound(t *testing.T) {
	var deck entity.Deck
	result := app.Database().Get().Create(&deck)
	assert.NoError(t, result.Error)
	writer := tests.MakeRequest("GET", fmt.Sprintf("/api/decks/%s", deck.ID.String()), nil, router(app))
	assert.Equal(t, http.StatusOK, writer.Code)
	var res map[string]interface{}
	err := json.Unmarshal(writer.Body.Bytes(), &res)
	assert.NoError(t, err)
	assert.Equal(t, deck.ID.String(), res["deck_id"].(string))
}

func TestItShouldGetOpenDeckByIdWithNonDrawnCardsAsRemaining(t *testing.T) {
	var deck entity.Deck
	var cards []*entity.CardTemplate
	result := app.Database().Get().Create(&deck)
	assert.NoError(t, result.Error)
	// find the first 4 cards
	result = app.Database().Get().Limit(4).Find(&cards)
	assert.NoError(t, result.Error)
	var cardDecks []*entity.CardDeck
	// associate the opened deck with the 4 cards
	for i, card := range cards {
		drawn := false
		// draw the last 2 cards
		if i >= 2 {
			drawn = true
		}
		cardDecks = append(cardDecks, &entity.CardDeck{
			CardID: card.ID,
			DeckID: deck.ID,
			Order:  i + 1,
			Drawn:  drawn,
		})
	}
	result = app.Database().Get().CreateInBatches(cardDecks, 4)
	assert.NoError(t, result.Error)

	writer := tests.MakeRequest("GET", fmt.Sprintf("/api/decks/%s", deck.ID.String()), nil, router(app))
	assert.Equal(t, http.StatusOK, writer.Code)
	var res map[string]interface{}
	err := json.Unmarshal(writer.Body.Bytes(), &res)
	assert.NoError(t, err)
	assert.Equal(t, float64(2), res["remaining"].(float64))
	assert.Equal(t, 2, len(res["cards"].([]interface{})))
}
