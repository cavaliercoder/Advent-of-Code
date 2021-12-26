package day15

import (
	"aoc/internal/geo"
)

func GrowGrid(g *geo.Grid) *geo.Grid {
	g2 := geo.NewGrid(g.Width*5, g.Height*5)
	for y := 0; y < g2.Height; y++ {
		for x := 0; x < g2.Width; x++ {
			v := int(g.MustGet(geo.Pos{X: x % g.Width, Y: y % g.Height}))
			v += x/g.Width + y/g.Height
			v %= 9
			if v == 0 {
				v = 9
			}
			g2.Set(geo.Pos{X: x, Y: y}, byte(v))
		}
	}
	return g2
}

func ShortestPath(g *geo.Grid, a, b int) int {
	costs := make([]int, g.Len())
	seen := make([]byte, g.Len())
	q := make([]int, 1, g.Len())
	q[0] = a
	for len(q) > 0 {
		i := q[0]
		q = q[1:]
		if seen[i] > 0 {
			continue
		}
		seen[i] = 1
		p := g.Pos(i)
		for _, np := range g.UDLR(p) {
			ni := g.Index(np)
			tc := costs[i] + int(g.Data[ni])
			nc := costs[ni]
			if nc == 0 || tc < nc {
				costs[ni] = tc
				seen[ni] = 0
				q = append(q, ni)
			}
		}
	}
	return costs[b]
}
