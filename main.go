package main

import (
	"chess/game"
	"chess/game/fen"
	"fmt"
)

func main() {
	FEN := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"

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

	turns_count := len(game.GetAllTurns(current_board, settings.White))
	fmt.Println(turns_count)
}
