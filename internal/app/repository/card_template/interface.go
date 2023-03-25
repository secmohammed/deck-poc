package card_template

import (
	"errors"
	"github.com/secmohammed/deck-poc/internal/entity"
)

var ErrDeckNotFound = errors.New("deck not found")

type CardTemplateRepository interface {
	FindWhereCodeIn(in []string, shuffle bool) ([]*entity.CardTemplate, error)
	GetAll(shuffle bool) ([]*entity.CardTemplate, error)
}
