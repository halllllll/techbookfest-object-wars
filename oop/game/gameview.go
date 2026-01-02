package game

import (
	"fmt"
	"sample/action"
	"sample/player"
	"sample/terminal"
)

type GameView struct {
	terminal *terminal.Terminal
}

func NewGameView(terminal *terminal.Terminal) *GameView {
	return &GameView{
		terminal: terminal,
	}
}

func (g *GameView) PlayerAsked(player player.Player, ask action.AskAction, hit bool) {
	g.terminal.Print(fmt.Sprintf("%s: %v\n", player.Name(), ask))
	var result string
	if hit {
		result = "Hit."
	} else {
		result = "Miss."
	}

	g.terminal.Print(result)
	g.terminal.EmptyLine()

}

func (g *GameView) PlayerGuessed(player player.Player, guess action.GuessAction, hit bool) {
	g.terminal.Print(fmt.Sprintf("%s: %v\n", player.Name(), guess))
	var result string
	if hit {
		result = "Hit."
	} else {
		result = "Miss."
	}

	g.terminal.Print(result)
	g.terminal.EmptyLine()

}
