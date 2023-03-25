package deck

import "github.com/gin-gonic/gin"

type DeckRestHandler interface {
	CreateDeck(c *gin.Context)
	GetDeck(c *gin.Context)
	DrawCards(c *gin.Context)
}
