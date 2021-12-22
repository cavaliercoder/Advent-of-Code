package day22

import (
	"fmt"
)

type Coord struct {
	X, Y, Z int
}

func (p Coord) In(c Cube) bool {
	if p.X < c.A.X || p.Y < c.A.Y || p.Z < c.A.Z {
		return false
	}
	if p.X > c.B.X || p.Y > c.B.Y || p.Z > c.B.Z {
		return false
	}
	return true
}

func (p Coord) String() string { return fmt.Sprintf("%d,%d,%d", p.X, p.Y, p.Z) }

type Cube struct {
	A Coord // close, bottom, left
	B Coord // far, top, right
}

func (c Cube) Intersects(other Cube) bool {
	a, b := c, other
	if a.B.X < b.A.X || b.B.X < a.A.X {
		return false
	}
	if a.B.Y < b.A.Y || b.B.Y < a.A.Y {
		return false
	}
	if a.B.Z < b.A.Z || b.B.Z < a.A.Z {
		return false
	}
	return true
}

func (c Cube) Split(other Cube) []Cube {
	a := make([]Cube, 1, 8)
	a[0] = c
	if !c.Intersects(other) {
		return a
	}
	a = split(a, other.A)
	a = split(a, other.B)
	return a
}

func split(a []Cube, p Coord) []Cube {
	for i, c := range a {
		if p.X > c.A.X && p.X < c.B.X {
			c2 := c
			c.B.X, c2.A.X = p.X, p.X
			a[i] = c
			a = append(a, c2)
		}
	}
	for i, c := range a {
		if p.Y > c.A.Y && p.Y < c.B.Y {
			c2 := c
			c.B.Y, c2.A.Y = p.Y, p.Y
			a[i] = c
			a = append(a, c2)
		}
	}
	for i, c := range a {
		if p.Z > c.A.Z && p.Z < c.B.Z {
			c2 := c
			c.B.Z, c2.A.Z = p.Z, p.Z
			a[i] = c
			a = append(a, c2)
		}
	}
	return a
}

func (c Cube) Width() int         { return c.B.X - c.A.X }
func (c Cube) Height() int        { return c.B.Y - c.A.Y }
func (c Cube) Depth() int         { return c.B.Z - c.A.Z }
func (c Cube) Volume() int        { return c.Width() * c.Height() * c.Depth() }
func (c Cube) In(other Cube) bool { return c.A.In(other) && c.B.In(other) }

func (c Cube) String() string {
	return fmt.Sprintf(
		"x=%d..%d,y=%d..%d,z=%d..%d",
		c.A.X, c.B.X-1,
		c.A.Y, c.B.Y-1,
		c.A.Z, c.B.Z-1,
	)
}

type Op struct {
	On   bool
	Cube Cube
}

type Reactor struct {
	m map[Cube]struct{}
}

func NewReactor() *Reactor { return &Reactor{} }

func (r *Reactor) do(op Op) {
	newCubes := make(map[Cube]struct{})
	for c := range r.m {
		for _, child := range c.Split(op.Cube) {
			if child.In(op.Cube) {
				continue
			}
			newCubes[child] = struct{}{}
		}
	}
	r.m = newCubes
	if op.On {
		r.m[op.Cube] = struct{}{}
	}
}

var initBounds = Cube{
	A: Coord{X: -50, Y: -50, Z: -50},
	B: Coord{X: 51, Y: 51, Z: 51},
}

func (r *Reactor) Init(ops ...Op) int {
	for _, op := range ops {
		if op.Cube.In(initBounds) {
			r.do(op)
		}
	}
	n := 0
	for cube := range r.m {
		n += cube.Volume()
	}
	return n
}

func (r *Reactor) Reboot(ops ...Op) int {
	for _, op := range ops {
		r.do(op)
	}
	n := 0
	for cube := range r.m {
		n += cube.Volume()
	}
	return n
}
