package aoc2021

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
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

type Rect struct {
	A, B Pos
}

func (r Rect) ContainsX(x int) bool { return r.A.X <= x && x <= r.B.X }
func (r Rect) ContainsY(y int) bool { return r.A.Y >= y && y >= r.B.Y }
func (r Rect) Contains(p Pos) bool  { return r.ContainsX(p.X) && r.ContainsY(p.Y) }

func (r Rect) Expand(n int) Rect {
	return Rect{
		A: Pos{X: r.A.X - n, Y: r.A.Y + n},
		B: Pos{X: r.B.X + n, Y: r.B.Y - n},
	}
}

func (r Rect) Fit(p Pos) Rect {
	if p.X < r.A.X {
		r.A.X = p.X
	}
	if p.X > r.B.X {
		r.B.X = p.X
	}
	if p.Y > r.A.Y {
		r.A.Y = p.Y
	}
	if p.Y < r.B.Y {
		r.B.Y = p.Y
	}
	return r
}

func (r Rect) Iter() *RectIter {
	return &RectIter{r: r}
}

type RectIter struct {
	r Rect
	p Pos
	i int
}

func (iter *RectIter) Next() bool {
	if iter.i == 0 {
		iter.p = iter.r.A
	} else {
		iter.p.X++
		if iter.p.X > iter.r.B.X {
			iter.p.X = iter.r.A.X
			iter.p.Y--
		}
	}
	iter.i++
	return iter.p.Y >= iter.r.B.Y
}

func (iter *RectIter) Pos() Pos { return iter.p }

type Grid struct {
	Data   []byte
	Width  int
	Height int
}

func NewGrid(width, height int) *Grid {
	return &Grid{
		Data:   make([]byte, width*height),
		Width:  width,
		Height: height,
	}
}

func ReadGrid(r io.Reader) (*Grid, error) {
	grid := &Grid{}
	buf := bytes.NewBuffer(nil)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		b := scanner.Bytes()
		if grid.Width == 0 {
			grid.Width = len(b)
		} else {
			if grid.Width != len(b) {
				return nil, fmt.Errorf("bad line width")
			}
		}
		grid.Height++
		buf.Write(scanner.Bytes())
	}
	grid.Data = buf.Bytes()
	return grid, nil
}

// Len returns the total number of cells in the grid.
func (c *Grid) Len() int { return len(c.Data) }

func (c *Grid) Reset(b byte) *Grid {
	for i := 0; i < len(c.Data); i++ {
		c.Data[i] = b
	}
	return c
}

func (c *Grid) Index(pos Pos) int {
	if !c.Contains(pos) {
		return -1
	}
	return (pos.Y * c.Width) + pos.X
}

func (c *Grid) Indexes(positions ...Pos) []int {
	a := make([]int, len(positions))
	for i, pos := range positions {
		a[i] = c.Index(pos)
	}
	return a
}

func (c *Grid) Pos(i int) Pos {
	return Pos{i % c.Width, i / c.Width}
}

func (c *Grid) Contains(p Pos) bool {
	return p.X >= 0 && p.X < c.Width && p.Y >= 0 && p.Y < c.Height
}

// Adj returns a slice of all adjascent positions to the given position,
// including diagonals.
func (c *Grid) Adj(p Pos) []Pos {
	a := make([]Pos, 0, 9)
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			p2 := Pos{
				X: p.X - 1 + x,
				Y: p.Y - 1 + y,
			}
			if p == p2 {
				continue
			}
			if c.Contains(p2) {
				a = append(a, p2)
			}
		}
	}
	return a
}

// UDLR (up, down, left, right) returns a slice of all adjacent positions to the
// given position, excluding diagonals.
func (c *Grid) UDLR(p Pos) []Pos {
	a := make([]Pos, 0, 4)
	for _, dir := range PosUDLR {
		if p2 := p.Add(dir); c.Contains(p2) {
			a = append(a, p2)
		}
	}
	return a
}

func (c *Grid) Get(p Pos) (b byte, ok bool) {
	i := c.Index(p)
	if i < 0 {
		return 0, false
	}
	return c.Data[i], true
}

func (c *Grid) GetWithDefault(p Pos, value byte) byte {
	b, ok := c.Get(p)
	if !ok {
		return value
	}
	return b
}

func (c *Grid) MustGet(p Pos) byte {
	b, ok := c.Get(p)
	if !ok {
		panic(fmt.Sprintf("out of bounds: %v (%dx%d)", p, c.Width, c.Height))
	}
	return b
}

func (c *Grid) Set(pos Pos, b byte) {
	i := c.Index(pos)
	if i < 0 {
		panic(fmt.Sprintf("out of bounds: %v", pos))
	}
	c.Data[i] = b
}

func (c *Grid) Print(w io.Writer) {
	newline := []byte{'\n'}
	for y := 0; y < c.Height; y++ {
		i := y * c.Width
		w.Write(c.Data[i : i+c.Width])
		w.Write(newline)
	}
}
