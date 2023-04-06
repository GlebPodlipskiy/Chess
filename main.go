package main

import (
	"chess/game"
	"chess/game/fen"
	"fmt"
)

func main() {
	FEN := "2R3k1/3r1ppp/8/8/8/8/5PPP/6K1"

	current_board, _ := fen.FENToBoard(FEN, game.FenSettings)

	board_string, _ := fen.BoardToString(current_board, game.FenSettings)

	fmt.Println(board_string)

	turns := game.GetAllTurns(current_board, game.Black)
	fmt.Println(len(turns))
}
