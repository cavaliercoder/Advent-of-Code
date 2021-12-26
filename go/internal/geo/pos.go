package geo

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	PosUp        Pos = Pos{0, -1}
	PosUpRight   Pos = Pos{1, -1}
	PosRight     Pos = Pos{1, 0}
	PosDownRight Pos = Pos{1, 1}
	PosDown      Pos = Pos{0, 1}
	PosDownLeft  Pos = Pos{-1, 1}
	PosLeft      Pos = Pos{-1, 0}
	PosUpLeft    Pos = Pos{-1, -1}

	// PosUDLR is constains adjacent horizontal and vertical positions.
	PosUDLR = []Pos{
		PosUp,
		PosDown,
		PosLeft,
		PosRight,
	}

	// PosAdj is all adjacent positions, including diagonals.
	PosAdj = []Pos{
		PosUp,
		PosUpRight,
		PosRight,
		PosDownRight,
		PosDown,
		PosDownLeft,
		PosLeft,
		PosUpLeft,
	}
)

// Pos is a grid position where the Y-axis increases as it gets lower.
type Pos struct {
	X int
	Y int
}

func NewPos(x, y int) Pos { return Pos{x, y} }

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

func (p Pos) Sub(v Pos) Pos {
	return Pos{p.X - v.X, p.Y - v.Y}
}

// Orient returns the direction toward v.
func (p Pos) Orient(v Pos) Pos {
	o := Pos{}
	if p.X < v.X {
		o.X = 1
	} else if p.X > v.X {
		o.X = -1
	}
	if p.Y < v.Y {
		o.Y = 1
	} else if p.Y > v.Y {
		o.Y = -1
	}
	return o
}

func MapPos(f func(Pos) Pos, positions ...Pos) []Pos {
	a := make([]Pos, len(positions))
	for i, p := range positions {
		a[i] = f(p)
	}
	return a
}
