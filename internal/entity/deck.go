package entity

type Deck struct {
	BaseModel
	Shuffle bool       `json:"shuffle" gorm:"default:false"`
	Cards   []CardDeck `json:"cards,omitempty"`
}
