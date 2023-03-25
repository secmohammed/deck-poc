package deck

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/secmohammed/deck-poc/internal/app/dto"
	"github.com/secmohammed/deck-poc/internal/app/usecase/deck"
	"github.com/secmohammed/deck-poc/utils"
	"strconv"
	"strings"

	"net/http"
)

type restHandler struct {
	du deck.UseCase
}

func (h *restHandler) GetDeck(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.NewBadRequest(err.Error()))
		return

	}
	res, err := h.du.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.NewBadRequest(err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"deck_id":   res.ID,
		"shuffled":  res.Shuffle,
		"remaining": res.Remaining,
		"cards":     transformCardsResponse(res.Cards),
	})

}
func (h *restHandler) DrawCards(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.NewBadRequest(err.Error()))
		return

	}
	count, err := strconv.Atoi(c.Query("count"))
	if err != nil || count <= 0 {
		count = 1
	}

	result, err := h.du.Draw(id, count)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": utils.NewBadRequest(err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"cards": transformCardsResponse(result.Cards),
	})
}
func (h *restHandler) CreateDeck(c *gin.Context) {

	body := dto.CreateDeckDTO{}
	if ok := utils.BindData(c, &body); !ok {
		return
	}
	cards := c.Query("cards")

	if cards != "" {
		cardSlice := strings.Split(cards, ",")
		body.CardFilters = cardSlice
	}
	result, err := h.du.Create(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": utils.NewBadRequest(err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, SuccessResponse{Data: map[string]interface {
	}{
		"id":        result.ID,
		"remaining": result.Remaining,
		"shuffled":  result.Shuffle,
	}})

}

func NewDeckHandler(du deck.UseCase) DeckRestHandler {
	return &restHandler{du}
}
