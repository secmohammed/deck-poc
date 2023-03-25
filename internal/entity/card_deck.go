package entity

import "github.com/google/uuid"

type CardDeck struct {
	CardID uuid.UUID    `json:"card_id" gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	DeckID uuid.UUID    `json:"deck_id" gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	Order  int          `json:"order" gorm:"default:0"`
	Drawn  bool         `json:"drawn" gorm:"default:false"`
	Deck   Deck         `json:"deck,omitempty" gorm:"foreignKey:DeckID"`
	Card   CardTemplate `json:"card,omitempty" gorm:"foreignKey:CardID"`
}
