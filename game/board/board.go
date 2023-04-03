package board

type Board struct {
	field [8][8]int
}

func (this *Board) Set(rank int, file int, value int) {
	this.field[file][rank] = value
}
func (this Board) Get(rank int, file int) int {
	return this.field[file][rank]
}
