package aoc2021

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

var (
	PosUp    Pos = Pos{0, -1}
	PosRight Pos = Pos{1, 0}
	PosDown  Pos = Pos{0, 1}
	PosLeft  Pos = Pos{-1, 0}
)

type Pos struct {
	X int
	Y int
}

func NewPos(x, y int) Pos {
	return Pos{x, y}
}

func (p Pos) Add(v Pos) Pos {
	return Pos{p.X + v.X, p.Y + v.Y}
}

type Grid struct {
	Data   []byte
	Width  int
	Height int
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

func (c *Grid) Index(pos Pos) int {
	if !c.Contains(pos) {
		return -1
	}
	return (pos.Y * c.Width) + pos.X
}

func (c *Grid) Pos(i int) Pos {
	return Pos{i % c.Width, i / c.Width}
}

func (c *Grid) Contains(pos Pos) bool {
	return pos.X >= 0 && pos.X < c.Width && pos.Y >= 0 && pos.Y < c.Height
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
		panic(fmt.Sprintf("out of bounds: %v", p))
	}
	return b
}
