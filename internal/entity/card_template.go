package entity

type Suit string
type Code string
type Value string

const (
	DIAMOND Suit = "DIAMOND"
	CLUB    Suit = "CLUB"
	HEART   Suit = "HEART"
	SPADE   Suit = "SPADE"
)

type CardTemplate struct {
	BaseModel
	Value Value `json:"value" gorm:"type:varchar(255);not null"`
	Code  Code  `json:"code" gorm:"type:varchar(255);not null"`
	Suit  Suit  `json:"suit" sql:"type:ENUM('CLUB', 'HEART', 'SPADE', 'DIAMOND')" gorm:"index"`
	Order int   `json:"order" gorm:"type:int; not null"`
}
