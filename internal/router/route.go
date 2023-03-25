package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/secmohammed/deck-poc/container"
	"github.com/secmohammed/deck-poc/internal/app/repository/card_template"
	deck3 "github.com/secmohammed/deck-poc/internal/app/repository/deck"
	deck2 "github.com/secmohammed/deck-poc/internal/app/usecase/deck"
	"github.com/secmohammed/deck-poc/internal/controller/http/deck"
	"github.com/secmohammed/deck-poc/utils"
	"github.com/siruspen/logrus"
	"net/http"
	"time"
)

type rest struct {
	r *gin.Engine
	c container.Container
}

func NewRestRepository(c container.Container) *rest {
	env := c.Config().GetString("app.env")
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	logEnabled := c.Config().GetBool("app.log.debug")
	r := gin.New()
	if logEnabled {
		r.Use(gin.Logger())

	}
	return &rest{c: c, r: r}
}
func setupDefaults(r *gin.Engine) {

	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		logMessage := fmt.Sprintf("%s |%s %d %s| %s |%s %s %s %s | %s | %s | %s\n",
			param.TimeStamp.Format(time.RFC1123),
			param.StatusCodeColor(),
			param.StatusCode,
			param.ResetColor(),
			param.ClientIP,
			param.MethodColor(),
			param.Method,
			param.ResetColor(),
			param.Path,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
		logrus.Info(fmt.Sprintf("%s | %d | %s | %s | %s | %s | %s | %s", param.TimeStamp.Format(time.RFC1123), param.StatusCode, param.ClientIP, param.Method, param.Path, param.Latency, param.Request.UserAgent(), param.ErrorMessage))
		return logMessage
	}))

	r.ForwardedByClientIP = true
	// recover from error when server fails to start and retry.
	r.Use(gin.Recovery())
	// Health check API
	r.GET("/api/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"health": "OK"})
	})
	// If the user goes to a route that's not defined, we show the user that this route is not found.
	// fallback route.
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, utils.NewNotFound(c.Request.URL.Path))
	})
}
func (r *rest) registerDeckRoutes() {
	dr := deck3.NewDeckRepository(r.c)
	ctr := card_template.NewCardTemplateRepository(r.c)
	du := deck2.NewUseCase(dr, ctr)
	dh := deck.NewDeckHandler(du)

	rg := r.r.Group("/api/decks")
	rg.POST("/", dh.CreateDeck)
	rg.GET("/:id", dh.GetDeck)
	rg.PATCH("/:id", dh.DrawCards)
}

func (r *rest) Expose() error {
	port := r.c.Config().GetString("app.rest.port")
	setupDefaults(r.r)
	r.registerDeckRoutes()

	return r.r.Run(fmt.Sprintf(":%s", port))
}
