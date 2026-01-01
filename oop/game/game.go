package game

import (
	"fmt"
	"sample/action"
	"sample/card"
	"sample/player"
	"sample/terminal"
)

type Game struct {
	deal     card.Deal
	player0  player.Player
	player1  player.Player
	terminal terminal.Terminal
}

func NewGame(deal card.Deal, player0 player.Player, player1 player.Player, terminal terminal.Terminal) *Game {
	return &Game{
		deal:     deal,
		player0:  player0,
		player1:  player1,
		terminal: terminal,
	}
}

func (g *Game) Start() player.Player {
	turn_player := g.player0
	turn_hand := g.deal.Player0Hand()

	opponent_player := g.player1
	opponent_hand := g.deal.Player1Hand()
	rest_card := g.deal.RestCard()

	var prev_action *action.AskAction
	for {
		available_actions := action.AvailableActions(turn_hand, prev_action)
		selected_action := turn_player.SelectAction(available_actions)
		g.terminal.Print(fmt.Sprintf("%s: %v\n", turn_player.Name(), selected_action))

		var hit bool
		var win_player *player.Player
		var ask_action *action.AskAction
		// action が AskAction か GuessAction かで分岐したい
		switch act := selected_action.(type) {
		case action.AskAction:
			ask_action = &action.AskAction{
				Card: act.GetCard(),
			}
			hit = ask_action.Hit(opponent_hand)
		case action.GuessAction:
			guess_action := action.GuessAction{
				Card: act.GetCard(),
			}
			hit = guess_action.Hit(rest_card)
			if hit {
				win_player = &turn_player
			} else {
				win_player = &opponent_player
			}
		}

		var result string
		if hit {
			result = "Hit."
		} else {
			result = "Miss."
		}

		g.terminal.Print(result)
		g.terminal.EmptyLine()

		if win_player != nil {
			return *win_player
		}

		prev_action = ask_action
		turn_player, opponent_player = opponent_player, turn_player
		turn_hand, opponent_hand = opponent_hand, turn_hand
	}
}
