package player

import (
	"math/rand/v2"
	"sample/action"
)

type RandomAI struct {
	name         string
	random_state *rand.Rand
}

func NewRandomAI(name string, random_state rand.PCG) Player {
	return &RandomAI{
		name:         name,
		random_state: rand.New(&random_state),
	}
}

// Name implements Player.
func (r RandomAI) Name() string {
	return r.name
}

// SelectAction implements Player.
func (r *RandomAI) SelectAction(action action.ActionList) action.Actioner {
	all_actions := action.AllActions()

	idx := r.random_state.IntN(len(all_actions))
	return all_actions[idx]
}
