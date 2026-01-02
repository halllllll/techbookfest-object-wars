package game

import (
	"sample/action"
	"sample/player"
)

type GameObserver interface {
	PlayerAsked(player player.Player, ask action.AskAction, hit bool)
	PlayerGuessed(player player.Player, guess action.GuessAction, hit bool)
}

