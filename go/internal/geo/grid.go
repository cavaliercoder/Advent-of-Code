package geo

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

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

func (c *Grid) Reset(b byte) *Grid {
	for i := range c.Data {
		c.Data[i] = b
	}
	return c
}

func (c *Grid) Normalize(b byte) *Grid {
	for i, a := range c.Data {
		c.Data[i] = a - b
	}
	return c
}

// Len returns the total number of cells in the grid.
func (c *Grid) Len() int { return len(c.Data) }

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
	if i < 0 || i >= len(c.Data) {
		panic(fmt.Sprintf("index out of bounds: %d", i))
	}
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

// Get returns the value of the cell at p or 0 if it is out of bounds.
func (c *Grid) Get(p Pos) byte {
	i := c.Index(p)
	if i < 0 {
		return 0
	}
	return c.Data[i]
}

func (c *Grid) GetWithDefault(p Pos, value byte) byte {
	i := c.Index(p)
	if i < 0 {
		return value
	}
	return c.Data[i]
}

func (c *Grid) MaybeGet(p Pos) (b byte, ok bool) {
	i := c.Index(p)
	if i < 0 {
		return
	}
	return c.Data[i], true
}

func (c *Grid) MustGet(p Pos) byte {
	i := c.Index(p)
	if i < 0 {
		panic(fmt.Sprintf("position out of bounds: %v", p))
	}
	return c.Data[i]
}

func (c *Grid) Set(p Pos, b byte) {
	i := c.Index(p)
	if i < 0 {
		panic(fmt.Sprintf("position out of bounds: %v", p))
	}
	c.Data[i] = b
}

func (c *Grid) Format(w io.Writer) {
	newline := []byte{'\n'}
	for y := 0; y < c.Height; y++ {
		i := y * c.Width
		w.Write(c.Data[i : i+c.Width])
		w.Write(newline)
	}
}

func (c *Grid) String() string {
	b := new(bytes.Buffer)
	c.Format(b)
	return b.String()
}
