package day25

import (
	. "aoc2021"
)

func stepEast(g *Grid) (*Grid, bool) {
	moved := false
	g2 := NewGrid(g.Width, g.Height).Reset('.')
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			p := Pos{X: x, Y: y}
			g2.Set(p, g.MustGet(p))
			if g.MustGet(p) != '>' {
				continue
			}
			p2 := Pos{X: (x + 1) % g.Width, Y: y}
			if g.MustGet(p2) == '.' {
				g2.Set(p, '.')
				g2.Set(p2, '>')
				x++
				moved = true
			}
		}
	}
	return g2, moved
}

func stepSouth(g *Grid) (*Grid, bool) {
	moved := false
	g2 := NewGrid(g.Width, g.Height).Reset('.')
	for x := 0; x < g.Width; x++ {
		for y := 0; y < g.Height; y++ {
			p := Pos{X: x, Y: y}
			g2.Set(p, g.MustGet(p))
			if g.MustGet(p) != 'v' {
				continue
			}
			p2 := Pos{X: x, Y: (y + 1) % g.Height}
			if g.MustGet(p2) == '.' {
				g2.Set(p, '.')
				g2.Set(p2, 'v')
				y++
				moved = true
			}
		}
	}
	return g2, moved
}

func step(g *Grid) (*Grid, bool) {
	var movedEast, movedSouth bool
	g, movedEast = stepEast(g)
	g, movedSouth = stepSouth(g)
	return g, movedEast || movedSouth
}

func MaxSteps(g *Grid) int {
	var moved bool
	for i := 1; ; i++ {
		g, moved = step(g)
		if !moved {
			return i
		}
	}
}
