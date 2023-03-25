package deck

import (
	"github.com/google/uuid"
	"github.com/secmohammed/deck-poc/internal/app/dto"
	"github.com/secmohammed/deck-poc/internal/app/repository/deck"
	"github.com/secmohammed/deck-poc/internal/entity"
)

type UseCase interface {
	Create(in *dto.CreateDeckDTO) (*deck.CreateDeckResponse, error)
	Get(id uuid.UUID) (*deck.GetDeckResponse, error)
	Draw(id uuid.UUID, count int) (*entity.Deck, error)
}
