package game

import (
	"chess/game/board"
	"chess/game/fen"
	"chess/game/turn"
	"chess/library"
)

var Pawn = 1
var Knight = 2
var Bishop = 3
var Rook = 4
var Queen = 5
var King = 6
var White = 1
var Black = -1
var None = 0

var FenSettings = fen.Settings{

	FenMap: map[int]int{
		'p': Pawn,
		'n': Knight,
		'b': Bishop,
		'r': Rook,
		'q': Queen,
		'k': King,
	},
	White: White,
	Black: Black,
}

func GetSide(cell int) int {
	if cell > 0 {
		return White
	} else if cell < 0 {
		return Black
	} else {
		return None
	}
}
func GetOpositeSide(side int) int {
	if side == White {
		return Black
	} else if side == Black {
		return White
	}
	return None
}
func in_board(rank int, file int) bool {
	return rank >= 0 && file >= 0 && rank <= 7 && file <= 7
}
func get_line(current_board board.Board, rank int, file int, directions [][3]int) []turn.Turn {
	turns := []turn.Turn{}

	piece := current_board.Get(rank, file)
	side := GetSide(piece)

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
				next_side := GetSide(next_cell)
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
	side := GetSide(piece)

	if piece == Pawn*side {
		old_cords := turn.Cords{
			File: file,
			Rank: rank,
		}
		turns := []turn.Turn{}

		//Default by white side
		direction := -1
		jump_file := 6

		if side == Black {
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
				next_side := GetSide(next_cell)

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

	} else if piece == Knight*side {
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
	} else if piece == Bishop*side {
		directions := [][3]int{
			{1, 1, 8},
			{1, -1, 8},
			{-1, 1, 8},
			{-1, -1, 8},
		}
		return get_line(current_board, rank, file, directions)
	} else if piece == Rook*side {
		directions := [][3]int{
			{1, 0, 8},
			{0, -1, 8},
			{-1, 0, 8},
			{0, 1, 8},
		}
		return get_line(current_board, rank, file, directions)
	} else if piece == Queen*side {
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
	} else if piece == King*side {
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
		side := GetSide(piece.Value)
		if side == current_side {
			turns := get_turns(current_board, piece.Rank, piece.File)
			validated_turn := []turn.Turn{}
			for _, current_turn := range turns {
				if ValidateTurn(&current_board, current_turn, current_side) {
					validated_turn = append(validated_turn, current_turn)
				}
			}
			all_turns = append(all_turns, validated_turn...)
		}

	}
	return all_turns
}
func MakeTurn(current_board *board.Board, turn turn.Turn) {
	piece := current_board.Get(turn.OldCords.Rank, turn.OldCords.File)

	current_board.Set(turn.OldCords.Rank, turn.OldCords.File, 0)
	current_board.Set(turn.NewCords.Rank, turn.NewCords.File, piece)
}
func UnmakeTurn(current_board *board.Board, turn turn.Turn) {
	piece := current_board.Get(turn.NewCords.Rank, turn.NewCords.File)

	current_board.Set(turn.NewCords.Rank, turn.NewCords.File, 0)
	current_board.Set(turn.OldCords.Rank, turn.OldCords.File, piece)
}
func CheckLine(current_board board.Board, start turn.Cords, direction [2]int, target_value int) bool {
	current_value := current_board.Get(start.Rank, start.File)
	distance := 1

	dx, dy := direction[0], direction[1]
	for distance < 8 {
		next_rank := start.Rank + dx*distance
		next_file := start.File + dy*distance

		if in_board(next_rank, next_file) {
			current_value = current_board.Get(next_rank, next_file)

			if current_value == target_value {
				return true
			}
			if current_value != None {
				return false
			}
			distance++
		} else {
			return false
		}

	}
	return false
}
func IsUnderAttack(current_board board.Board, piece board.Piece, target board.Piece) bool {
	direction_x := target.Rank - piece.Rank
	direction_y := target.File - piece.File

	var dx int
	var dy int

	if direction_x != 0 {
		dx = library.Abs(direction_x) / direction_x
	} else {
		dx = 0
	}
	if direction_y != 0 {
		dy = library.Abs(direction_y) / direction_y
	} else {
		dy = 0
	}

	if piece.Value == Pawn {
		if direction_y == 0 && library.Abs(direction_x) == 1 {
			return true
		}
	} else if piece.Value == Knight {
		dist := direction_x + direction_y
		if dist == 3 && direction_x != 0 && direction_y != 0 {
			return true
		}
	} else if piece.Value == Bishop {
		on_diagonal := library.Abs(direction_x) == library.Abs(direction_y)
		if on_diagonal {
			return CheckLine(current_board, turn.Cords{File: piece.File, Rank: piece.Rank}, [2]int{dx, dy}, target.Value)
		} else {
			return false
		}
	} else if piece.Value == Rook {
		on_vert := direction_x == 0 || direction_y == 0
		if on_vert {
			return CheckLine(current_board, turn.Cords{File: piece.File, Rank: piece.Rank}, [2]int{dx, dy}, target.Value)
		} else {
			return false
		}
	} else if piece.Value == Queen {
		on_diagonal := library.Abs(direction_x) == library.Abs(direction_y)
		on_vert := direction_x == 0 || direction_y == 0

		if on_diagonal || on_vert {
			return CheckLine(current_board, turn.Cords{File: piece.File, Rank: piece.Rank}, [2]int{dx, dy}, target.Value)
		}
	}
	return false
}
func IsChecked(current_board board.Board, side int) bool {
	king := current_board.GetPiece(side * King)

	for _, piece := range current_board.GetPieces() {
		piece_side := GetSide(piece.Value)
		if side == -piece_side {
			if IsUnderAttack(current_board, piece, king) {
				return true
			}
		}
	}

	return false
}
func ValidateTurn(current_board *board.Board, posible_turn turn.Turn, side int) bool {
	MakeTurn(current_board, posible_turn)
	//fmt.Println(fen.BoardToString(*current_board, FenSettings))
	validated := !IsChecked(*current_board, side)
	UnmakeTurn(current_board, posible_turn)

	return validated
}
