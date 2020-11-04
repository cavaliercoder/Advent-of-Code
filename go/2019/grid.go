package aoc

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

var (
	DirectionUp    Pos = Pos{0, -1}
	DirectionRight Pos = Pos{1, 0}
	DirectionDown  Pos = Pos{0, 1}
	DirectionLeft  Pos = Pos{-1, 0}
)

type Pos struct {
	X int
	Y int
}

func NewPos(x, y int) Pos {
	return Pos{x, y}
}

func (p Pos) IsZero() bool {
	return p == Pos{}
}

func (p Pos) Add(v Pos) Pos {
	return Pos{p.X + v.X, p.Y + v.Y}
}

func (p Pos) Subtract(v Pos) Pos {
	return Pos{p.X - v.X, p.Y - v.Y}
}

// URDL returns the positions one step up, right, down and left of the position.
func (p Pos) URDL() [4]Pos {
	return [4]Pos{
		p.Add(DirectionUp),
		p.Add(DirectionRight),
		p.Add(DirectionDown),
		p.Add(DirectionLeft),
	}
}

func (p Pos) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
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

func OpenGrid(name string) (*Grid, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ReadGrid(f)
}

func (c *Grid) Copy() *Grid {
	b := make([]byte, len(c.Data))
	g := &Grid{
		Data:   b,
		Width:  c.Width,
		Height: c.Height,
	}
	copy(g.Data, c.Data)
	return g
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

func (c *Grid) Get(pos Pos) byte {
	i := c.Index(pos)
	if i < 0 {
		panic(fmt.Sprintf("out of bounds: %v", pos))
	}
	return c.Data[i]
}

func (c *Grid) GetWithDefault(pos Pos, def byte) byte {
	i := c.Index(pos)
	if i < 0 {
		return def
	}
	return c.Data[i]
}

func (c *Grid) Set(pos Pos, b byte) {
	i := c.Index(pos)
	if i < 0 {
		println(fmt.Sprintf("out of bounds: %v", pos))
	}
	// fmt.Printf("set %v = %c (was %c)\n", pos, b, c.Data[i])
	c.Data[i] = b
}

func (c *Grid) FindOne(b byte) int {
	for i, a := range c.Data {
		if a == b {
			return i
		}
	}
	return -1
}

func (c *Grid) FindAll(b byte) []int {
	v := make([]int, 0)
	for i, a := range c.Data {
		if a == b {
			v = append(v, i)
		}
	}
	return v
}

func (c *Grid) Print(w io.Writer) {
	newline := []byte{'\n'}
	for y := 0; y < c.Height; y++ {
		i := y * c.Width
		w.Write(c.Data[i : i+c.Width])
		w.Write(newline)
	}
}
