package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Action struct {
	kind string
	card int
}

func deal() ([]int, []int, int) {
	var all_cards []int

	for i := 1; i < 10; i++ {
		all_cards = append(all_cards, i)
	}

	rand.Shuffle(len(all_cards), func(i, j int) {
		all_cards[i], all_cards[j] = all_cards[j], all_cards[i]
	})
	shuffle_cards := all_cards
	player0_hand, player1_hand, rest := shuffle_cards[:4], shuffle_cards[4:8], shuffle_cards[8]
	sort.Ints(player0_hand)
	sort.Ints(player1_hand)
	return player0_hand, player1_hand, rest
}
func start_game(player0_hand *[]int, player1_hand *[]int, rest_card int) int {
	var turn_player int = 0
	var win_player int
	var action Action
	var prev_action *Action

	for {
		if turn_player == 0 {
			action = select_action_human(player0_hand, prev_action)
			win_player = check_action(0, action, player1_hand, rest_card)
		} else {
			action = select_action_ai(player1_hand, prev_action)
			win_player = check_action(1, action, player0_hand, rest_card)
		}

		if win_player != -1 {
			return win_player
		}
		prev_action = &action
		turn_player = (turn_player + 1) % 2
	}
}

func get_available_actions(hand *[]int, prev_action *Action) []Action {
	// 可能な行動一覧
	// - 入力1: 手番プレイやーの手札（整数のスライス）
	// - 入力2: 直前の行動
	// 出力: 可能な行動一覧

	var actions []Action
	for i := 1; i < 10; i++ {
		a := Action{
			kind: "ask",
			card: i,
		}
		actions = append(actions, a)
	}
	if prev_action != nil {

		idx := slices.Index(actions, *prev_action)
		if idx != -1 {
			actions = slices.Delete(actions, idx, idx+1)
		}
		for i := 1; i < 10; i++ {
			if !slices.Contains(*hand, i) {
				actions = append(actions, Action{
					kind: "guess", card: i,
				})
			}
		}

	}
	return actions
}

func select_action_human(hand *[]int, prev_action *Action) Action {
	available_actions := get_available_actions(hand, prev_action)
	var ask_cards []int
	var guess_cards []int

	for _, action := range available_actions {
		if action.kind == "ask" {
			ask_cards = append(ask_cards, action.card)
		} else {
			guess_cards = append(guess_cards, action.card)
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	var action Action
	for {
		fmt.Fprintf(os.Stdout, "Your hand: %v\n", hand)
		fmt.Fprint(os.Stdout, "Available commands:\n")
		if len(ask_cards) > 0 {
			fmt.Fprintf(os.Stdout, " ask <card>		(<card>: %v)\n", ask_cards)
		}
		if len(guess_cards) > 0 {
			fmt.Fprintf(os.Stdout, " guess <card>		(<card>: %v)\n", guess_cards)
		}
		fmt.Fprint(os.Stdout, "		exit\n")
		fmt.Fprint(os.Stdout, "player>: ")
		if !scanner.Scan() {
			break
		}
		args := strings.Fields(scanner.Text())

		if len(args) < 1 {
			fmt.Fprint(os.Stdout, "Empty command.\n\n")
			continue
		}
		command := strings.ToLower(args[0])
		var num int = -1
		if len(args) > 1 {
			card := args[1]
			_num, err := strconv.Atoi(card)
			if err != nil {
				fmt.Fprint(os.Stderr, "error: %w\n", err)
				continue
			}
			num = _num
		}

		if command == "ask" {
			if num < 0 {
				fmt.Fprint(os.Stdout, "Card is not specified\n")
				continue
			}
			action = Action{
				kind: "ask",
				card: num,
			}
		} else if command == "guess" {
			if num < 0 {
				fmt.Fprint(os.Stdout, "Card is not specified\n")
				continue
			}
			action = Action{
				kind: "guess",
				card: num,
			}
		} else if command == "exit" {
			fmt.Fprint(os.Stdout, "exit game.\n")
			os.Exit(1)
		} else {
			fmt.Fprintf(os.Stderr, "unknown command: %s\n\n", command)
			continue
		}

		if !slices.Contains(available_actions, action) {
			fmt.Fprintf(os.Stderr, "unavaialble action: %v\n\n", action)
			continue
		} else {
			fmt.Fprintf(os.Stdout, "your select %v\n", action)
			break
		}
	}
	return action
}

func select_action_ai(hand *[]int, prev_action *Action) Action {
	available_actions := get_available_actions(hand, prev_action)
	// 適当に取る
	idx := rand.Intn(len(available_actions))
	action := available_actions[idx]
	fmt.Fprintf(os.Stdout, "ai select %v\n", action)
	return action
}

func check_action(player int, action Action, oppenent_hand *[]int, rest_card int) int {
	var win_player int = -1
	if action.kind == "ask" {
		if slices.Contains(*oppenent_hand, action.card) {
			fmt.Fprint(os.Stdout, "Hit.\n")
		} else {
			fmt.Fprint(os.Stdout, "Miss.\n")
		}
	} else {
		// guess
		if action.card == rest_card {
			fmt.Fprint(os.Stdout, "Hit!")
			win_player = player
		} else {
			fmt.Fprintf(os.Stdout, "Miss... rest card: %d\n", rest_card)
			opponent_player := (player + 1) % 2
			win_player = opponent_player
		}
	}
	fmt.Fprint(os.Stdout, "\n")
	return win_player
}

func show_result(win_player int) {
	if win_player == 0 {
		fmt.Fprint(os.Stdout, "You won!\n")
	} else {
		fmt.Fprint(os.Stdout, "You lost...\n")
	}
}
