package ai

import (
	"chess/game"
	"chess/game/board"
)

func AlphaBeta(current_board board.Board, depth int, side int) float32 {
	if depth == 0 {
		return Evalute(current_board)
	}

	all_turns := game.GetAllTurns(current_board, game.White)

	for _, turn := range all_turns {
		game.MakeTurn(&current_board, turn)
		game.UnmakeTurn(&current_board, turn)
	}

	return 0.0

}
