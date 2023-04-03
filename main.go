package main

import (
	"chess/game/fen"
	"fmt"
)

func main() {
	FEN := "6k1/3r1ppp/8/8/8/8/5PPP/2R3K1"

	settings := fen.Settings{
		FenMap: map[int]int{
			'p': 1,
			'n': 2,
			'b': 3,
			'r': 4,
			'q': 5,
			'k': 6,
		},
		White: 1,
		Black: -1,
	}
	current_board, _ := fen.FENToBoard(FEN, settings)

	board_string, _ := fen.BoardToString(current_board, settings)
	fmt.Println(board_string)
}
