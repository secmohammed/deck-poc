package card_template

import (
    "github.com/secmohammed/deck-poc/container"
    "github.com/secmohammed/deck-poc/internal/entity"
)

type cr struct {
    c container.Container
}

func (cr cr) GetAll(shuffle bool) ([]*entity.CardTemplate, error) {
    var cards []*entity.CardTemplate
    result := cr.c.Database().Get().Order(getOrder(shuffle)).Find(&cards)
    return cards, result.Error
}

func (cr cr) FindWhereCodeIn(in []string, shuffle bool) ([]*entity.CardTemplate, error) {
    var cards []*entity.CardTemplate
    s := make([]interface{}, len(in))
    for i, v := range in {
        s[i] = v
    }
    result := cr.c.Database().Get().Debug().Where("code in ?", s).Order(getOrder(shuffle)).Find(&cards)

    return cards, result.Error
}
func getOrder(shuffle bool) string {
    order := "\"order\" ASC"
    if shuffle {
        order = "random()"
    }
    return order
}
func NewCardTemplateRepository(c container.Container) CardTemplateRepository {
    return &cr{c}
}
