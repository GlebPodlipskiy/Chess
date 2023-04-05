package turn

type Cords struct {
	File int
	Rank int
}
type Turn struct {
	OldCords Cords
	NewCords Cords
	Claimed  bool
}
