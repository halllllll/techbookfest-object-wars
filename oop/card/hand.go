package card

import (
	"slices"
)

type Hand struct {
	cards []Card
}

func NewHand(cards []Card) Hand {
	return Hand{
		cards: cards,
	}
}

func (h Hand) HasCard(card Card) bool {
	return slices.Contains(h.cards, card)
}

func (h Hand) GetCards() []Card {
	return h.cards
}
