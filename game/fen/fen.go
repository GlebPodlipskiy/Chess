package fen

import (
	"chess/game/board"
	"errors"
	"fmt"
	"unicode"
)

func FENToBoard(fen string, settings Settings) (board.Board, error) {

	current_board := board.Board{}

	file := 0
	rank := 0

	for i, char := range fen {

		lower_rune_char := unicode.ToLower(char)

		if char == '/' {
			file++
			rank = 0
		} else if char > '0' && char < '9' {
			rank += int(char) - '0'
		} else if _, ok := settings.FenMap[int(lower_rune_char)]; ok {

			piece, _ := settings.FenMap[int(lower_rune_char)]
			side := settings.Black

			if char != lower_rune_char {
				side = settings.White
			}
			current_board.Set(rank, file, piece*side)
			rank++
		} else {
			return current_board, errors.New("Uncorect FEN. Undefined symbol at position " + fmt.Sprint(i) + ".")
		}
		i += 1
	}

	return current_board, nil
}
func BoardToString(current_board board.Board, settigns Settings) (string, error) {
	res := ""

	for file := 0; file < 8; file++ {
		row_string := ""

		for rank := 0; rank < 8; rank++ {
			row_string += "|"
			var char int = '_'
			current_piece := current_board.Get(rank, file)

			for key := range settigns.FenMap {
				piece, _ := settigns.FenMap[key]
				if piece == current_piece*settigns.Black {
					char = key
					break
				} else if piece == current_piece*settigns.White {
					char = int(unicode.ToUpper(rune(key)))
					break
				}
			}
			row_string += string(rune(char))
		}

		row_string += "|\n"

		res += row_string
	}

	return res, nil
}
