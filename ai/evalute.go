package ai

import (
	"chess/game"
	"chess/game/board"
	"math"
)

func Evalute(current_board board.Board) float32 {
	var pieces_cost = map[int]float32{
		1: 1,
		2: 3.0,
		3: 3.25,
		4: 5.0,
		5: 9.0,
		6: 100.0,
	}

	var result float32 = 0
	for _, piece := range current_board.GetPieces() {
		index := int(math.Abs(float64(piece.Value)))
		cost := pieces_cost[index]
		side := game.GetSide(piece.Value)

		result += cost * float32(side)
	}

	return result
}
