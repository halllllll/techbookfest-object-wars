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
	number int
}

func NewCard(num int) Card {
	if num < min_number || max_number < num {
		log.Fatalf("invalid number: %d", num)
	}
	return Card{
		number: num,
	}
}

func (c Card) Equals(other Card) bool {
	return c.number == other.GetNumber()
}
func (c Card) LessThan(other Card) bool {
	return c.number < other.GetNumber()
}

func (c Card) String() string {
	return fmt.Sprintf("Card(%d)", c.number)
}

func (c Card) GetNumber() int {
	return c.number
}

func AllCards() []Card {
	cards := make([]Card, 0, max_number-min_number+1)
	for i := min_number; i <= max_number; i++ {
		c := NewCard(i)
		cards = append(cards, c)
	}

	return cards
}
