package player

import (
	"sample/action"
)

type Player interface {
	Name() string
	SelectAction(action.ActionList) action.Actioner
}
