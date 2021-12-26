package day09

import (
	"aoc/internal/geo"
)

func IsLowPoint(g *geo.Grid, i int) bool {
	p := g.Pos(i)
	v := g.Data[i]
	for _, adj := range geo.MapPos(p.Add, geo.PosUDLR...) {
		if neighbor, ok := g.MaybeGet(adj); ok {
			if neighbor <= v {
				return false
			}
		}
	}
	return true
}

func SumRisk(g *geo.Grid) int {
	risk := 0
	for i, v := range g.Data {
		if IsLowPoint(g, i) {
			risk += int(v) + 1
		}
	}
	return risk
}

func GetBasinSize(g *geo.Grid, i int) int {
	if !IsLowPoint(g, i) {
		return 0
	}
	p := g.Pos(i)
	size := 0
	seen := make(map[geo.Pos]struct{})
	stack := make([]geo.Pos, 0, 32)
	stack = append(stack, p)
	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if _, ok := seen[p]; ok {
			continue
		}
		seen[p] = struct{}{}
		if b := g.GetWithDefault(p, 9); b == 9 {
			continue
		}
		size++
		stack = append(
			stack,
			p.Add(geo.PosUp),
			p.Add(geo.PosRight),
			p.Add(geo.PosDown),
			p.Add(geo.PosLeft),
		)
	}
	return size
}

func GetBasinSizes(g *geo.Grid) []int {
	a := make([]int, 0, 64)
	for i := range g.Data {
		if IsLowPoint(g, i) {
			a = append(a, GetBasinSize(g, i))
		}
	}
	return a
}
