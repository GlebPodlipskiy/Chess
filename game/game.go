package game

import (
	"chess/game/board"
	"chess/game/fen"
	"chess/game/turn"
	"fmt"
)

var pawn = 1
var knight = 2
var bishop = 3
var rook = 4
var queen = 5
var king = 6
var white = 1
var black = -1
var none = 0

var fen_settings = fen.Settings{

	FenMap: map[int]int{
		'p': pawn,
		'n': knight,
		'b': bishop,
		'r': rook,
		'q': queen,
		'k': king,
	},
	White: white,
	Black: black,
}

func get_side(cell int) int {
	if cell > 0 {
		return white
	} else if cell < 0 {
		return black
	} else {
		return none
	}
}
func in_board(rank int, file int) bool {
	return rank >= 0 && file >= 0 && rank <= 7 && file <= 7
}
func get_line(current_board board.Board, rank int, file int, directions [][3]int) []turn.Turn {
	turns := []turn.Turn{}

	piece := current_board.Get(rank, file)
	side := get_side(piece)

	old_cords := turn.Cords{
		Rank: rank,
		File: file,
	}

	for _, direction := range directions {
		dx, dy := direction[0], direction[1]
		for distance := 1; distance <= direction[2]; distance++ {
			next_rank := rank + distance*dx
			next_file := file + distance*dy

			if in_board(next_rank, next_file) {
				next_cell := current_board.Get(next_rank, next_file)
				next_side := get_side(next_cell)
				next_cords := turn.Cords{
					File: next_file,
					Rank: next_rank,
				}
				if next_side == 0 {
					move_turn := turn.Turn{
						Claimed:  false,
						OldCords: old_cords,
						NewCords: next_cords,
					}
					turns = append(turns, move_turn)
				} else if next_side == -side {
					claimed_turn := turn.Turn{
						Claimed:  true,
						OldCords: old_cords,
						NewCords: next_cords,
					}
					turns = append(turns, claimed_turn)
					break
				} else if next_side == side {
					break
				}
			} else {
				break
			}
		}
	}
	return turns
}
func get_turns(current_board board.Board, rank int, file int) []turn.Turn {
	var piece = current_board.Get(rank, file)
	side := get_side(piece)

	if piece == pawn*side {
		old_cords := turn.Cords{
			File: file,
			Rank: rank,
		}
		turns := []turn.Turn{}

		//Default by white side
		direction := -1
		jump_file := 6

		if side == black {
			direction = 1
			jump_file = 1
		}

		all_move_cords := []turn.Cords{
			{File: direction + file, Rank: rank},
		}

		if jump_file == file {
			all_move_cords = append(all_move_cords, turn.Cords{
				File: file + direction*2,
				Rank: rank,
			})
		}

		all_claimed_cords := []turn.Cords{
			{File: file + direction, Rank: rank + 1},
			{File: file + direction, Rank: rank - 1},
		}

		for _, move_cords := range all_move_cords {
			if current_board.Get(move_cords.Rank, move_cords.File) != 0 {
				break
			}
			move_turn := turn.Turn{
				OldCords: old_cords,
				NewCords: move_cords,
				Claimed:  false,
			}
			turns = append(turns, move_turn)
		}

		for _, claimed_cords := range all_claimed_cords {
			if in_board(claimed_cords.Rank, claimed_cords.File) {
				next_cell := current_board.Get(claimed_cords.Rank, claimed_cords.File)
				next_side := get_side(next_cell)

				if side == -next_side {
					turns = append(turns, turn.Turn{
						OldCords: old_cords,
						NewCords: claimed_cords,
						Claimed:  true,
					})
				}
			}
		}

		return turns

	} else if piece == knight*side {
		directions := [][3]int{
			{2, 1, 1},
			{1, 2, 1},
			{-1, 2, 1},
			{2, -1, 1},
			{-1, -2, 1},
			{-2, -1, 1},
			{1, -2, 1},
			{-2, 1, 1},
		}
		return get_line(current_board, rank, file, directions)
	} else if piece == bishop*side {
		directions := [][3]int{
			{1, 1, 8},
			{1, -1, 8},
			{-1, 1, 8},
			{-1, -1, 8},
		}
		return get_line(current_board, rank, file, directions)
	} else if piece == rook*side {
		directions := [][3]int{
			{1, 0, 8},
			{0, -1, 8},
			{-1, 0, 8},
			{0, 1, 8},
		}
		return get_line(current_board, rank, file, directions)
	} else if piece == queen*side {
		directions := [][3]int{
			{1, 1, 8},
			{1, -1, 8},
			{-1, 1, 8},
			{-1, -1, 8},
			{1, 0, 8},
			{0, -1, 8},
			{-1, 0, 8},
			{0, 1, 8},
		}
		return get_line(current_board, rank, file, directions)
	} else if piece == king*side {
		directions := [][3]int{
			{1, 1, 1},
			{1, -1, 1},
			{-1, 1, 1},
			{-1, -1, 1},
			{1, 0, 1},
			{0, -1, 1},
			{-1, 0, 1},
			{0, 1, 1},
		}
		return get_line(current_board, rank, file, directions)
	}
	return []turn.Turn{}
}
func GetAllTurns(current_board board.Board, current_side int) []turn.Turn {
	all_turns := []turn.Turn{}
	for _, piece := range current_board.GetPieces() {
		side := get_side(piece.Value)
		if side == current_side {
			turns := get_turns(current_board, piece.Rank, piece.File)
			fmt.Println(len(turns))
			all_turns = append(all_turns, turns...)
		}

	}
	return all_turns
}
