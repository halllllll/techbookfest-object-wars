package action

import (
	"fmt"
	"sample/card"
	"slices"
)

type Actioner interface {
	GetCard() card.Card
	Equals(Actioner) bool
	String() string
}

type AskAction struct {
	Card card.Card
}


func (a AskAction) GetCard() card.Card {
	return a.Card
}

func (a AskAction) Hit(hand card.Hand) bool {
	return hand.HasCard(a.Card)
}

func (a AskAction) Equals(other Actioner) bool {
	if _, ok := other.(AskAction); !ok {
		return false
	}
	return a.Card.Equals(other.GetCard())
}

func (a AskAction) String() string {
	return fmt.Sprintf("ASK (%v)\n", a.Card)
}

type GuessAction struct {
	Card card.Card
}


func (a GuessAction) GetCard() card.Card {
	return a.Card
}

func (a GuessAction) Hit(card card.Card) bool {
	return a.Card.Equals(card)
}

func (a GuessAction) Equals(other Actioner) bool {
	if _, ok := other.(GuessAction); !ok {
		return false
	}
	return a.Card.Equals(other.GetCard())
}

func (a GuessAction) String() string {
	return fmt.Sprintf("Guess (%v)\n", a.Card)
}

type ActionList struct {
	ask_actions   []Actioner
	guess_actions []Actioner
}

func NewActionList(ask_actions []Actioner, guess_actins []Actioner) ActionList {
	return ActionList{
		ask_actions:   ask_actions,
		guess_actions: guess_actins,
	}
}

func (al ActionList) AskActions() []Actioner {
	return al.ask_actions
}

func (al ActionList) GuessActions() []Actioner {
	return al.guess_actions
}

func (al ActionList) AllActions() []Actioner {
	return append(al.guess_actions, al.ask_actions...)
}

func (al ActionList) Contains(action Actioner) bool {
	return slices.Contains(al.AllActions(), action)
}

func (al ActionList) String() string {
	return fmt.Sprintf("%v\n", al.AllActions())
}

func AvailableActions(hand card.Hand, prev_action *AskAction) ActionList {
	var ask_actions []Actioner
	// ask_actions := make([]Actioner, 0, len(card.AllCards()))
	for _, v := range card.AllCards() {
		ask_actions = append(ask_actions, AskAction{
			Card: v,
		})
	}

	var guess_actions []Actioner

	if prev_action != nil {
		// idx := slices.Index(ask_actions, *prev_action)
		idx := slices.IndexFunc(ask_actions, func(a Actioner) bool {
			return a.Equals(prev_action)
		})
		if idx != -1 {
			ask_actions = slices.Delete(ask_actions, idx, idx+1)
		}
		for _, v := range card.AllCards() {
			if !hand.HasCard(v) {
				guess_actions = append(guess_actions, GuessAction{
					Card: v,
				})
			}
		}
	}

	return ActionList{
		ask_actions:   ask_actions,
		guess_actions: guess_actions,
	}
}
