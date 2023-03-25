package deck

import (
	"errors"
	"github.com/google/uuid"
	"github.com/secmohammed/deck-poc/internal/app/dto"
	"github.com/secmohammed/deck-poc/internal/entity"
)

var ErrDeckNotFound = errors.New("deck not found")

type DeckRepository interface {
	Create(in *dto.CreateDeckDTO, cards []*entity.CardTemplate) (*CreateDeckResponse, error)
	Get(id uuid.UUID) (*GetDeckResponse, error)
	DrawDeckCards(id uuid.UUID, count int) (*entity.Deck, error)
}
