package geo

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	PosUp    Pos = Pos{0, -1}
	PosRight Pos = Pos{1, 0}
	PosDown  Pos = Pos{0, 1}
	PosLeft  Pos = Pos{-1, 0}

	PosUDLR = []Pos{PosUp, PosDown, PosLeft, PosRight}
)

type Pos struct {
	X int
	Y int
}

func NewPos(x, y int) Pos {
	return Pos{x, y}
}

func ParsePos(s string) (Pos, error) {
	a := strings.Split(s, ",")
	if len(a) != 2 {
		return Pos{}, fmt.Errorf("invalid Pos: %s", s)
	}
	x, err := strconv.Atoi(a[0])
	if err != nil {
		return Pos{}, fmt.Errorf("invalid Pos: %s", s)
	}
	y, err := strconv.Atoi(a[1])
	if err != nil {
		return Pos{}, fmt.Errorf("invalid Pos: %s", s)
	}
	return Pos{X: x, Y: y}, nil
}

func (p Pos) Add(v Pos) Pos {
	return Pos{p.X + v.X, p.Y + v.Y}
}
