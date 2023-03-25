package deck

import "github.com/secmohammed/deck-poc/internal/entity"

type SuccessResponse struct {
	Data interface{} `json:"data"`
}
type CardResponse struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
	Order int    `json:"order"`
}

func transformCardsResponse(cards []entity.CardDeck) []CardResponse {
	var cardsResponse []CardResponse

	for _, cards := range cards {
		cardsResponse = append(cardsResponse, CardResponse{
			Value: string(cards.Card.Value),
			Code:  string(cards.Card.Code),
			Suit:  string(cards.Card.Suit),
			Order: cards.Order,
		})
	}
	return cardsResponse
}
