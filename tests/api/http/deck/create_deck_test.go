package deck

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/secmohammed/deck-poc/internal/app/dto"
	"github.com/secmohammed/deck-poc/internal/entity"
	"github.com/secmohammed/deck-poc/tests"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	gin.SetMode(gin.TestMode)
	c := setup()
	exitCode := m.Run()
	teardown(c)

	os.Exit(exitCode)
}

func TestItShouldCreateDeckWithoutShuffling(t *testing.T) {
	val := false
	payload := &dto.CreateDeckDTO{
		Shuffle: &val,
	}

	writer := tests.MakeRequest("POST", "/api/decks", payload, router(app))
	assert.Equal(t, http.StatusOK, writer.Code)
	var res struct{ Data map[string]interface{} }
	err := json.Unmarshal(writer.Body.Bytes(), &res)
	assert.NoError(t, err)
	assert.Equal(t, val, res.Data["shuffled"])
	var cards []entity.CardTemplate
	// since it shouldn't be shuffled, the first card should be Ace Spade.
	result := app.Database().Get().Where("code = 'AS'").Find(&cards)
	assert.NoError(t, result.Error)

	var deck entity.Deck
	result = app.Database().Get().Preload("Cards", func(db *gorm.DB) *gorm.DB {
		return db.Limit(1)
	}).First(&deck)
	assert.NoError(t, result.Error)
	assert.Equal(t, deck.Cards[0].CardID, cards[0].ID)
}

func TestItShouldCreateDeckWithShuffling(t *testing.T) {
	val := true
	payload := &dto.CreateDeckDTO{
		Shuffle: &val,
	}

	writer := tests.MakeRequest("POST", "/api/decks", payload, router(app))
	assert.Equal(t, http.StatusOK, writer.Code)
	var res struct{ Data map[string]interface{} }
	err := json.Unmarshal(writer.Body.Bytes(), &res)
	assert.NoError(t, err)
	assert.Equal(t, val, res.Data["shuffled"].(bool))
	// testing that it really shuffles is non-determined, might be good idea to take a sample and validate against.
}
func TestItShouldCreateDeckWithAllDecks(t *testing.T) {
	val := true
	payload := &dto.CreateDeckDTO{
		Shuffle: &val,
	}

	writer := tests.MakeRequest("POST", "/api/decks", payload, router(app))
	assert.Equal(t, http.StatusOK, writer.Code)
	var res struct{ Data map[string]interface{} }
	err := json.Unmarshal(writer.Body.Bytes(), &res)
	assert.NoError(t, err)
	assert.Equal(t, float64(52), res.Data["remaining"].(float64))
}
func TestItShouldCreateDeckWithPasssedDecks(t *testing.T) {
	val := true
	payload := &dto.CreateDeckDTO{
		Shuffle: &val,
	}

	writer := tests.MakeRequest("POST", "/api/decks?cards=AS,KD,AC,2C,KH", payload, router(app))
	assert.Equal(t, http.StatusOK, writer.Code)
	var res struct{ Data map[string]interface{} }
	err := json.Unmarshal(writer.Body.Bytes(), &res)
	assert.NoError(t, err)
	assert.Equal(t, float64(5), res.Data["remaining"].(float64))
}
