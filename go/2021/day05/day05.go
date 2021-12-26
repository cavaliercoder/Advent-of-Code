package day05

import (
	"fmt"

	"aoc/internal/geo"
)

type Vent struct {
	A, B geo.Pos
}

func (c *Vent) String() string   { return fmt.Sprintf("%v -> %v", c.A, c.B) }
func (c *Vent) IsDiagonal() bool { return c.A.X != c.B.X && c.A.Y != c.B.Y }

func CountIntersects(countDiagonals bool, vents ...*Vent) int {
	count := 0
	m := make(map[geo.Pos]int)
	mark := func(p geo.Pos) {
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
		p, o := vent.A, vent.A.Orient(vent.B)
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
