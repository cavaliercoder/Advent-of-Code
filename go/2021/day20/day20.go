package day20

import (
	"bytes"
	"fmt"
	"io"

	"aoc/internal/geo"
)

type State struct {
	m        map[geo.Pos]byte
	infinity byte
	bounds   geo.Rect
}

func NewState() *State {
	return &State{
		m: make(map[geo.Pos]byte, 4096),
	}
}

func (c *State) Set(p geo.Pos) {
	c.m[p] = 1
	c.bounds = c.bounds.Fit(p)
}

func (c *State) Get(p geo.Pos) byte {
	if c.bounds.Contains(p) {
		return c.m[p]
	}
	return c.infinity
}

func (c *State) Format(w io.Writer) {
	r := c.bounds.Expand(3)
	for y := r.A.Y; y >= r.B.Y; y-- {
		for x := r.A.X; x <= r.B.X; x++ {
			b := c.Get(geo.Pos{X: x, Y: y})
			switch b {
			case 0:
				fmt.Fprint(w, ".")
			default:
				fmt.Fprint(w, "#")
			}
		}
		fmt.Fprint(w, "\n")
	}
}

func (c *State) String() string {
	b := new(bytes.Buffer)
	c.Format(b)
	return b.String()
}

var paramBounds = []geo.Pos{
	{X: -1, Y: 1},
	{X: 0, Y: 1},
	{X: 1, Y: 1},
	{X: -1, Y: 0},
	{X: 0, Y: 0},
	{X: 1, Y: 0},
	{X: -1, Y: -1},
	{X: 0, Y: -1},
	{X: 1, Y: -1},
}

func (c *State) nextPixel(p geo.Pos, algo []byte) byte {
	n := 0
	for _, offset := range paramBounds {
		n <<= 1
		if b := c.Get(p.Add(offset)); b != 0 {
			n |= 1
		}
	}
	return algo[n]
}

func (c *State) Step(algo []byte) (next *State) {
	next = NewState()
	if c.infinity == 0 && algo[0] == 1 {
		next.infinity = 1
	}
	if c.infinity == 1 && algo[511] == 1 {
		next.infinity = 1
	}
	iter := c.bounds.Expand(1).Iter()
	for iter.Next() {
		p := iter.Pos()
		if b := c.nextPixel(p, algo); b == 1 {
			next.Set(p)
		}
	}
	return
}

func (c *State) LitCount() int {
	if c.infinity > 0 {
		return -1 // infinity
	}
	return len(c.m)
}
