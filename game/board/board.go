package board

import "chess/library"

type Piece struct {
	Value int
	File  int
	Rank  int
}

type Board struct {
	field  [8][8]int
	pieces []Piece
}

func (this Board) GetPieces() []Piece {
	return this.pieces
}
func (this *Board) Set(rank int, file int, value int) {
	if value != 0 {
		founded := false
		for _, piece := range this.pieces {
			if piece.File == file && piece.Rank == rank {
				piece.Value = value
				founded = true
			}
		}
		if !founded {
			this.pieces = append(this.pieces, Piece{
				Value: value,
				File:  file,
				Rank:  rank,
			})
		}
	} else {
		this.pieces = library.Filter(this.pieces, func(el Piece) bool {
			return !(el.File == file && el.Rank == rank)
		})
	}

	this.field[file][rank] = value
}
func (this Board) Get(rank int, file int) int {
	return this.field[file][rank]
}
