package card

type Deal struct {
	player0_hand Hand
	player1_hand Hand
	rest_card    Card
}

func NewDeal(player0_hand Hand, player1_hand Hand, rest_card Card) Deal {
	return Deal{
		player0_hand: player0_hand,
		player1_hand: player1_hand,
		rest_card:    rest_card,
	}
}

func (d Deal) Player0Hand() Hand {
	return d.player0_hand
}

func (d Deal) Player1Hand() Hand {
	return d.player1_hand
}

func (d Deal) RestCard() Card {
	return d.rest_card
}
