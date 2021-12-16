package day09

import (
	. "aoc2021"
)

func IsLowPoint(g *Grid, i int) bool {
	v := g.Data[i]
	p := g.Pos(i)
	for _, adj := range []Pos{PosUp, PosRight, PosDown, PosLeft} {
		if neighbor, ok := g.Get(p.Add(adj)); ok {
			if neighbor <= v {
				return false
			}
		}
	}
	return true
}

func SumRisk(g *Grid) int {
	risk := 0
	for i, v := range g.Data {
		if IsLowPoint(g, i) {
			risk += int(v) + 1
		}
	}
	return risk
}

func GetBasinSize(g *Grid, i int) int {
	if !IsLowPoint(g, i) {
		return 0
	}
	p := g.Pos(i)
	size := 0
	seen := make(map[Pos]struct{})
	stack := make([]Pos, 0, 32)
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
			p.Add(PosUp),
			p.Add(PosRight),
			p.Add(PosDown),
			p.Add(PosLeft),
		)
	}
	return size
}

func GetBasinSizes(g *Grid) []int {
	a := make([]int, 0, 64)
	for i := range g.Data {
		if IsLowPoint(g, i) {
			a = append(a, GetBasinSize(g, i))
		}
	}
	return a
}
