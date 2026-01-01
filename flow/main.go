package main

func main() {
	p0h, p1h, r := deal()
	win_player := start_game(&p0h, &p1h, r)
	show_result(win_player)
}
