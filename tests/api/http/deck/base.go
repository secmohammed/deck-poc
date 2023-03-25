package deck

import (
    "github.com/gin-gonic/gin"
    "github.com/secmohammed/deck-poc/config"
    "github.com/secmohammed/deck-poc/container"
    "github.com/secmohammed/deck-poc/internal/app/repository/card_template"
    deck3 "github.com/secmohammed/deck-poc/internal/app/repository/deck"
    deck2 "github.com/secmohammed/deck-poc/internal/app/usecase/deck"
    "github.com/secmohammed/deck-poc/internal/controller/http/deck"
    "github.com/secmohammed/deck-poc/internal/entity"
    "github.com/secmohammed/deck-poc/pkg/database"

    "log"
)

var c = config.Factory("local")
var app = container.NewApplication(c)

func setup() container.Container {
    if err := app.Database().Get().Migrator().AutoMigrate(&entity.CardTemplate{}, &entity.CardDeck{}, &entity.Deck{}); err != nil {
        log.Fatal(err)
    }

    err := database.SeedTemplates(app.Database().Get())
    if err != nil {
        log.Fatalf("failed to seed data reason :%s", err)
        return nil
    }
    return app

}
func router(c container.Container) *gin.Engine {

    r := gin.New()
    dr := deck3.NewDeckRepository(c)
    ctr := card_template.NewCardTemplateRepository(c)
    du := deck2.NewUseCase(dr, ctr)
    dh := deck.NewDeckHandler(du)

    rg := r.Group("/api/decks")
    rg.POST("", dh.CreateDeck)
    rg.GET("/:id", dh.GetDeck)
    rg.PATCH("/:id", dh.DrawCards)
    return r

}
func teardown(c container.Container) {
    if err := c.Database().Get().Migrator().DropTable(&entity.CardTemplate{}, &entity.CardDeck{}, &entity.Deck{}); err != nil {
        log.Fatal(err)
    }
}
