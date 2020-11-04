package day24

import (
	"bytes"
	"fmt"
	"io"
	"math"

	"aoc"
)

type Eris struct {
	a, b          *aoc.Grid
	parent, child *Eris
	recursive     bool
	depth         int
}

func NewEris(recursive bool) *Eris {
	g := aoc.NewGrid(5, 5)
	for i := 0; i < len(g.Data); i++ {
		g.Data[i] = '.'
	}
	if recursive {
		g.Set(aoc.NewPos(2, 2), '?')
	}
	return &Eris{
		a:         g,
		b:         g.Copy(),
		recursive: recursive,
	}
}

func ReadEris(r io.Reader, recursive bool) *Eris {
	g, err := aoc.ReadGrid(r)
	if err != nil {
		panic(err)
	}
	if recursive {
		g.Set(aoc.NewPos(2, 2), '?')
	}
	return &Eris{a: g, b: g.Copy(), recursive: recursive}
}

func (c *Eris) SHA256() string { return c.a.SHA256() }

func (c *Eris) TopLevel() *Eris {
	for ; c.parent != nil; c = c.parent {
	}
	return c
}

func (c *Eris) isBug(pos aoc.Pos) bool {
	if c == nil {
		return false
	}
	b := c.a.GetWithDefault(pos, '.')
	if b != '#' {
		return false
	}
	return true
}

func (c *Eris) countBugs(positions ...aoc.Pos) (n int) {
	for _, pos := range positions {
		if c.recursive && pos == (aoc.NewPos(2, 2)) {
			// middle cells don't count in recursive mode
			continue
		}
		if c.isBug(pos) {
			n++
		}
	}
	return
}

func (c *Eris) countAdjacentBugs(pos aoc.Pos) (n int) {
	// start with neighbors on this plane
	npos := pos.URDL()
	n = c.countBugs(npos[:]...)
	if !c.recursive {
		return
	}

	// outer edges, adjacent to parent plane
	if pos.Y == 0 && c.parent.isBug(aoc.NewPos(2, 1)) {
		n++
	}
	if pos.X == 0 && c.parent.isBug(aoc.NewPos(1, 2)) {
		n++
	}
	if pos.Y == c.a.Height-1 && c.parent.isBug(aoc.NewPos(2, 3)) {
		n++
	}
	if pos.X == c.a.Width-1 && c.parent.isBug(aoc.NewPos(3, 2)) {
		n++
	}

	// inner edges, adjacent to child plane
	if pos == aoc.NewPos(2, 1) {
		for x := 0; x < c.a.Width; x++ {
			y := 0
			if c.child.isBug(aoc.NewPos(x, y)) {
				n++
			}
		}
	}
	if pos == aoc.NewPos(3, 2) {
		x := c.a.Width - 1
		for y := 0; y < c.a.Height; y++ {
			if c.child.isBug(aoc.NewPos(x, y)) {
				n++
			}
		}
	}
	if pos == aoc.NewPos(2, 3) {
		y := c.a.Height - 1
		for x := 0; x < c.a.Width; x++ {
			if c.child.isBug(aoc.NewPos(x, y)) {
				n++
			}
		}
	}
	if pos == aoc.NewPos(1, 2) {
		x := 0
		for y := 0; y < c.a.Height; y++ {
			if c.child.isBug(aoc.NewPos(x, y)) {
				n++
			}
		}
	}
	return
}

func (c *Eris) Step() {
	if !c.recursive {
		c.prepareState()
		c.switchState()
		return
	}

	// move to top plane
	c = c.TopLevel()

	// do we need a higher plane?
	if c.a.Count('#') > 0 {
		c.parent = NewEris(true)
		c.parent.child = c
		c.parent.depth = c.depth - 1
		c = c.parent
	}

	// prepare next states
	for plane := c; plane != nil; plane = plane.child {
		plane.prepareState()

		// do we need a lower plane?
		if plane.child == nil && plane.a.Count('#') > 0 {
			plane.child = NewEris(true)
			plane.child.parent = plane
			plane.child.depth = plane.depth + 1
		}
	}

	// switch states
	for plane := c; plane != nil; plane = plane.child {
		plane.switchState()
	}
}

func (c *Eris) prepareState() {
	for y := 0; y < c.a.Height; y++ {
		for x := 0; x < c.a.Width; x++ {
			if c.recursive && x == 2 && y == 2 {
				// never change middle tile in recursive mode
				continue
			}
			pos := aoc.NewPos(x, y)
			isBug := c.isBug(pos)
			adjacentBugs := c.countAdjacentBugs(pos)
			if isBug {
				// bugs die unless there is exactly one adjacent bug
				if adjacentBugs != 1 {
					isBug = false
				}
			} else {
				// bugs appear if 1 or 2 adjacent bugs
				if adjacentBugs == 1 || adjacentBugs == 2 {
					isBug = true
				}
			}
			if isBug {
				c.b.Set(pos, '#')
			} else {
				c.b.Set(pos, '.')
			}
		}
	}
}

func (c *Eris) switchState() {
	c.a, c.b = c.b, c.a
}

func (c *Eris) BiodiversityRating() (n int) {
	for i := 0; i < len(c.a.Data); i++ {
		if c.a.Data[i] != '#' {
			continue
		}
		n += int(math.Pow(2, float64(i)))
	}
	return
}

func (c *Eris) RecursiveBugCount() (n int) {
	for c = c.TopLevel(); c != nil; c = c.child {
		n += c.a.Count('#')
	}
	return
}

func (c *Eris) String() string {
	b := &bytes.Buffer{}
	fmt.Fprintf(b, "Depth %d:\n", c.depth)
	c.a.Print(b)
	return b.String()
}
