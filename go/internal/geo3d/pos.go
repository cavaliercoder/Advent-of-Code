package geo3d

import (
	"fmt"
	"strconv"
	"strings"

	"aoc/internal/util"
)

// Tranform is a 3D rotation matrix.
type Transform [9]int

// Rotations contains all possible 3D rotations with increments of 90 degrees.
// Courtesy: https://www.euclideanspace.com/maths/algebra/matrix/transforms/examples/index.htm
var Rotations = []Transform{
	{1, 0, 0, 0, 1, 0, 0, 0, 1}, // identity
	{0, 0, 1, 0, 1, 0, -1, 0, 0},
	{-1, 0, 0, 0, 1, 0, 0, 0, -1},
	{0, 0, -1, 0, 1, 0, 1, 0, 0},

	{0, -1, 0, 1, 0, 0, 0, 0, 1},
	{0, 0, 1, 1, 0, 0, 0, 1, 0},
	{0, 1, 0, 1, 0, 0, 0, 0, -1},
	{0, 0, -1, 1, 0, 0, 0, -1, 0},

	{0, 1, 0, -1, 0, 0, 0, 0, 1},
	{0, 0, 1, -1, 0, 0, 0, -1, 0},
	{0, -1, 0, -1, 0, 0, 0, 0, -1},
	{0, 0, -1, -1, 0, 0, 0, 1, 0},

	{1, 0, 0, 0, 0, -1, 0, 1, 0},
	{0, 1, 0, 0, 0, -1, -1, 0, 0},
	{-1, 0, 0, 0, 0, -1, 0, -1, 0},
	{0, -1, 0, 0, 0, -1, 1, 0, 0},

	{1, 0, 0, 0, -1, 0, 0, 0, -1},
	{0, 0, -1, 0, -1, 0, -1, 0, 0},
	{-1, 0, 0, 0, -1, 0, 0, 0, 1},
	{0, 0, 1, 0, -1, 0, 1, 0, 0},

	{1, 0, 0, 0, 0, 1, 0, -1, 0},
	{0, -1, 0, 0, 0, 1, -1, 0, 0},
	{-1, 0, 0, 0, 0, 1, 0, 1, 0},
	{0, 1, 0, 0, 0, 1, 1, 0, 0},
}

type Pos struct {
	X, Y, Z int
}

func NewPos(x, y, z int) Pos { return Pos{X: x, Y: y, Z: z} }

func ParsePos(s string) (p Pos, err error) {
	tokens := strings.Split(s, ",")
	if len(tokens) != 3 {
		return Pos{}, fmt.Errorf("invalid position: %s", s)
	}
	p.X, err = strconv.Atoi(tokens[0])
	if err != nil {
		return Pos{}, fmt.Errorf("invalid position: %s", s)
	}
	p.Y, err = strconv.Atoi(tokens[1])
	if err != nil {
		return Pos{}, fmt.Errorf("invalid position: %s", s)
	}
	p.Z, err = strconv.Atoi(tokens[2])
	if err != nil {
		return Pos{}, fmt.Errorf("invalid position: %s", s)
	}
	return
}

func (p Pos) Add(other Pos) Pos {
	return Pos{
		X: p.X + other.X,
		Y: p.Y + other.Y,
		Z: p.Z + other.Z,
	}
}

func (p Pos) Sub(other Pos) Pos {
	return Pos{
		X: p.X - other.X,
		Y: p.Y - other.Y,
		Z: p.Z - other.Z,
	}
}

func (p Pos) Transform(t Transform) Pos {
	return Pos{
		X: p.X*t[0] + p.Y*t[1] + p.Z*t[2],
		Y: p.X*t[3] + p.Y*t[4] + p.Z*t[5],
		Z: p.X*t[6] + p.Y*t[7] + p.Z*t[8],
	}
}

func (p Pos) In(c Cube) bool {
	if p.X < c.A.X || p.Y < c.A.Y || p.Z < c.A.Z {
		return false
	}
	if p.X > c.B.X || p.Y > c.B.Y || p.Z > c.B.Z {
		return false
	}
	return true
}

func (c Pos) Manhattan() int {
	return util.Abs(c.X) + util.Abs(c.Y) + util.Abs(c.Z)
}

func (c Pos) String() string {
	return fmt.Sprintf("<%d,%d,%d>", c.X, c.Y, c.Z)
}
