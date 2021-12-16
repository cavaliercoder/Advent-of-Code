package day05

import "fmt"

type Point struct {
	X, Y int
}

func (p Point) String() string    { return fmt.Sprintf("%d,%d", p.X, p.Y) }
func (p Point) Add(v Point) Point { return Point{X: p.X + v.X, Y: p.Y + v.Y} }

func (p Point) Orientation(v Point) Point {
	o := Point{}
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

type Vent struct {
	A, B Point
}

func (c *Vent) String() string   { return fmt.Sprintf("%v -> %v", c.A, c.B) }
func (c *Vent) IsDiagonal() bool { return c.A.X != c.B.X && c.A.Y != c.B.Y }

func CountIntersects(countDiagonals bool, vents ...*Vent) int {
	count := 0
	m := make(map[Point]int)
	mark := func(p Point) {
		n := m[p]
		if n == 1 {
			count++
		}
		m[p] = n + 1
	}
	for _, vent := range vents {
		if vent.IsDiagonal() && !countDiagonals {
			continue
		}
		p, o := vent.A, vent.A.Orientation(vent.B)
		mark(p)
		for {
			p = p.Add(o)
			mark(p)
			if p == vent.B {
				break
			}
		}
	}
	return count
}
