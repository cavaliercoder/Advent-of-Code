package day13

import (
	"fmt"
	"io"

	"aoc/internal/geo"
)

func Fold(p geo.Pos, grid map[geo.Pos]struct{}) map[geo.Pos]struct{} {
	if p.X != 0 && p.Y == 0 {
		return foldX(p.X, grid)
	}
	if p.Y != 0 && p.X == 0 {
		return foldY(p.Y, grid)
	}
	panic(fmt.Sprintf("bad fold: %v", p))
}

func foldX(x int, grid map[geo.Pos]struct{}) map[geo.Pos]struct{} {
	result := make(map[geo.Pos]struct{}, len(grid))
	for p := range grid {
		if p.X <= x {
			result[p] = struct{}{}
			continue
		}
		delta := p.X - x
		result[geo.Pos{X: x - delta, Y: p.Y}] = struct{}{}
	}
	return result
}

func foldY(y int, grid map[geo.Pos]struct{}) map[geo.Pos]struct{} {
	result := make(map[geo.Pos]struct{}, len(grid))
	for p := range grid {
		if p.Y <= y {
			result[p] = struct{}{}
			continue
		}
		delta := p.Y - y
		result[geo.Pos{X: p.X, Y: y - delta}] = struct{}{}
	}
	return result
}

func PrintGrid(w io.Writer, grid map[geo.Pos]struct{}) {
	maxX, maxY := 0, 0
	for p := range grid {
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}
	g := geo.NewGrid(maxX+1, maxY+1)
	g.Reset('.')
	for p := range grid {
		g.Set(p, '#')
	}
	g.Format(w)
}
