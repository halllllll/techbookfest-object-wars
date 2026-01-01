package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"sample/card"
	"sample/game"
	"sample/player"
	"sample/terminal"
)

func main() {
	seed := rand.NewPCG(1, 5)
	dealer := card.NewDealer(*seed)
	deal := dealer.DealStart()

	terminal := terminal.New(os.Stdin, os.Stdout)
	player0 := player.NewHumanPlayer("Player0", deal.Player0Hand(), *terminal)
	player1 := player.NewRandomAI("Player1", *seed)
	game := game.NewGame(deal, player0, player1, *terminal)

	win_player := game.Start()

	terminal.Print(fmt.Sprintf("%s won\n", win_player.Name()))
}
