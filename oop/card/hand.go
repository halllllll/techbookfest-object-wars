package card

import (
	"slices"
)

type Hand struct {
	Cards []Card
}

func (h Hand) HasCard(card Card) bool {
	return slices.Contains(h.Cards, card)
}

func (h Hand) GetCards() []Card {
	return h.Cards
}
