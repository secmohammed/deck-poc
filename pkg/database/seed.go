package database

import (
	"fmt"
	"github.com/secmohammed/deck-poc/internal/entity"
	"gorm.io/gorm"
	"strconv"
)

func SeedTemplates(tx *gorm.DB) error {
	var batch []*entity.CardTemplate
	for i := 1; i <= 13; i++ {
		value := strconv.Itoa(i)
		code := value
		if i == 1 {
			value = "ACE"
			code = "A"
		}
		if i == 11 {
			value = "JACK"
			code = "J"
		}
		if i == 12 {
			value = "QUEEN"
			code = "Q"
		}
		if i == 13 {
			value = "KING"
			code = "K"
		}
		batch = append(batch, &entity.CardTemplate{
			Value: entity.Value(value),
			Code:  entity.Code(fmt.Sprintf("%sC", code)),
			Suit:  entity.CLUB,
			Order: i + 26,
		}, &entity.CardTemplate{
			Value: entity.Value(value),
			Code:  entity.Code(fmt.Sprintf("%sH", code)),
			Suit:  entity.HEART,
			Order: i + 39,
		}, &entity.CardTemplate{
			Value: entity.Value(value),
			Code:  entity.Code(fmt.Sprintf("%sS", code)),
			Suit:  entity.SPADE,
			Order: i,
		}, &entity.CardTemplate{
			Value: entity.Value(value),
			Code:  entity.Code(fmt.Sprintf("%sD", code)),
			Suit:  entity.DIAMOND,
			Order: i + 13,
		})

	}
	return tx.Model(&entity.CardTemplate{}).CreateInBatches(batch, 52).Error
}
