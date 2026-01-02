package player

import (
	"fmt"
	"log"
	"os"
	"sample/action"
	"sample/card"
	"sample/terminal"
	"slices"
	"strconv"
	"strings"
)

type HumanPlayer struct {
	name     string
	hand     card.Hand
	terminal terminal.Terminal
}

func NewHumanPlayer(name string, hand card.Hand, terminal terminal.Terminal) Player {
	return &HumanPlayer{
		name:     name,
		hand:     hand,
		terminal: terminal,
	}
}

// SelectAction implements Player.
func (h *HumanPlayer) SelectAction(actions action.ActionList) action.Actioner {
	for {
		h.printHelp(actions)
		command, args := h.getCommand()
		if command == "" {
			h.terminal.Print("Empty Command")
			h.terminal.EmptyLine()
			continue
		}
		action := h.parseCommand(command, args)
		if action == nil {
			h.terminal.Print("Parse Error")
			h.terminal.EmptyLine()
			continue
		}

		// isContains := slices.ContainsFunc(actions.AllActions(), func(a actions.Actioner) bool {
		// 	return a.Equals(action)
		// })
		var exists bool
		for _, candidate := range actions.AllActions() {
			if candidate.Equals(action) {
				exists = true
			}
		}

		if !exists {
			h.terminal.Print(fmt.Sprintf("unavailable action: %v\n", action))

			continue
		}

		return action
	}
}

// Name implements Player.
func (h HumanPlayer) Name() string {
	return h.name
}

func (h HumanPlayer) printHelp(actions action.ActionList) {
	hand_str := h.formatCard(h.hand.Cards)
	h.terminal.Print(fmt.Sprintf("Your hand: %v\n", hand_str))
	h.terminal.Print("Available command:\n")

	var ask_cards []card.Card
	for _, v := range actions.AskActions() {
		ask_cards = append(ask_cards, v.GetCard())
	}
	ask_str := h.formatCard(ask_cards)
	if ask_str != "" {
		h.terminal.Print(fmt.Sprintf("   ask <card>    (<card>: %v)\n", ask_str))
	}

	var guess_cards []card.Card
	for _, v := range actions.GuessActions() {
		guess_cards = append(guess_cards, v.GetCard())
	}
	guess_str := h.formatCard(guess_cards)
	if guess_str != "" {
		h.terminal.Print(fmt.Sprintf("   guess <card>    (<card>: %v)\n", guess_str))
	}

	h.terminal.Print("   exit")
	h.terminal.EmptyLine()
}

func (h HumanPlayer) formatCard(cards []card.Card) string {
	var card_numbrers []string
	for _, c := range cards {
		card_numbrers = append(card_numbrers, strconv.Itoa(c.GetNumber()))
	}
	return strings.Join(card_numbrers, ", ")
}

func (h HumanPlayer) getCommand() (string, []string) {
	input, err := h.terminal.Prompt(h.Name() + ":> ")
	if err != nil {
		log.Fatal(err)
	}
	args := strings.Fields(input)
	if len(args) < 1 {
		return "", nil
	}
	command := strings.ToLower(args[0])
	return command, args
}

func (h HumanPlayer) parseCommand(command string, args []string) action.Actioner {
	if command == "exit" {
		h.terminal.Print("Bye :)")
		os.Exit(1)
	}
	if !slices.Contains([]string{"ask", "guess"}, command) {
		h.terminal.Print(fmt.Sprintf("Unknown command: %s\n", command))
		return nil
	}
	if len(args) < 1 {
		h.terminal.Print("Card is not specified")
		return nil
	}
	num, err := strconv.Atoi(args[1])
	if err != nil {
		h.terminal.Print(err.Error())
		return nil
	}
	card := card.NewCard(num)

	if command == "ask" {
		return action.AskAction{
			Card: card,
		}
	} else {
		return action.GuessAction{
			Card: card,
		}
	}
}
