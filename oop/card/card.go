package card

import (
	"fmt"
	"log"
)

const (
	min_number int = 1
	max_number int = 9
)

type Card struct {
	Number int
}

func NewCard(num int) Card {
	if num < min_number || max_number < num {
		log.Fatalf("invalid number: %d", num)
	}
	return Card{
		Number: num,
	}
}

func (c Card) Equals(other Card) bool {
	return c.Number == other.Number
}
func (c Card) LessThan(other Card) bool {
	return c.Number < other.Number
}

func (c Card) String() string {
	return fmt.Sprintf("Card(%d)", c.Number)
}

func (c Card) GetNumber() int {
	return c.Number
}

func AllCards() []Card {
	cards := make([]Card, 0, max_number-min_number+1)
	for i := min_number; i <= max_number; i++ {
		c := NewCard(i)
		cards = append(cards, c)
	}

	return cards
}

