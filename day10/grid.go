package day10

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
)

const (
	CellSpace    = '.'
	CellAsteroid = '#'
)

var (
	ErrOutOfBound = errors.New("out of bounds")
)

type Grid struct {
	width, height int
	data          []byte
}

// ReadGrid decodes a grid.
func ReadGrid(r io.Reader) (*Grid, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	grid := &Grid{
		data: make([]byte, 4096),
	}
	j := 0
	for i := 0; i < len(b); i++ {
		if b[i] == '\n' {
			grid.height++
			if grid.width == 0 {
				grid.width = i
			}
		} else {
			grid.data[j] = b[i]
			j++
		}
	}
	grid.data = grid.data[:j]
	return grid, nil
}

// OpenGrid decodes a grid from a file.
func OpenGrid(name string) (*Grid, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ReadGrid(f)
}

// Contains returns true if the given coordinate is within the bounds of the
// grid.
func (c *Grid) Contains(p Coord) bool {
	x := int(p.x)
	y := int(p.y)
	return x >= 0 && x < c.width && y >= 0 && y < c.height
}

// IndexOf returns the index of the given coordinate within the grid's encoded
// data.
func (c *Grid) IndexOf(p Coord) int {
	return (c.width * int(p.y)) + int(p.x)
}

// CoordOf returns the cartesian coordinate of the position stored at index i
// of the grid's encoded data.
func (c *Grid) CoordOf(i int) Coord {
	return Coord{
		x: float64(i % c.width),
		y: float64(i / c.height),
	}
}

// Get returns the state of the space at coordinate p.
func (c *Grid) Get(p Coord) (byte, error) {
	if !c.Contains(p) {
		return 0, ErrOutOfBound
	}
	return c.data[c.IndexOf(p)], nil
}

// Count returns the number of positions in the grid with the given state.
func (c *Grid) Count(b byte) (n int) {
	for i := 0; i < len(c.data); i++ {
		if c.data[i] == b {
			n++
		}
	}
	return
}

func (c *Grid) String() string {
	b := make([]byte, (c.width*c.height)+c.height)
	j := 0
	for i := 0; i < len(c.data); i++ {
		if i > 0 && i%c.width == 0 {
			b[j] = '\n'
			j++
		}
		b[j] = c.data[i]
		j++
	}
	b[j] = '\n'
	return string(b)
}
