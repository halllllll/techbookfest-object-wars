package card

import (
	"math/rand/v2"
)

type Dealer struct {
	random_state rand.PCG
}

func NewDealer(random_state rand.PCG) *Dealer {
	return &Dealer{
		random_state: random_state,
	}
}

func (d *Dealer) DealStart() Deal {
	// seed
	r := rand.New(&d.random_state)

	all_cards := AllCards()
	r.Shuffle(len(all_cards), func(i, j int) {
		all_cards[i], all_cards[j] = all_cards[j], all_cards[i]
	})
	player0_hand := Hand{
		Cards: all_cards[:4],
	}
	player1_hand := Hand{
		Cards: all_cards[4:8],
	}

	rest_card := all_cards[len(all_cards)-1]

	return NewDeal(player0_hand, player1_hand, rest_card)
}
