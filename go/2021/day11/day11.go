package day11

import "aoc/internal/geo"

func Step(g *geo.Grid) (flashes int) {
	stack := make([]int, 0, 64)
	for i, c := range g.Data {
		if c == 9 {
			stack = append(stack, i)
			continue
		}
		if c > 9 {
			panic("grid corrupted")
		}
		g.Data[i] = c + 1
	}
	for len(stack) > 0 {
		i := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if g.Data[i] == 0 {
			continue
		}
		flashes++
		g.Data[i] = 0
		for _, j := range g.Indexes(g.Adj(g.Pos(i))...) {
			switch g.Data[j] {
			case 0:
				continue
			case 9:
				stack = append(stack, j)
				continue
			default:
				g.Data[j]++
			}
		}
	}
	return
}

func GetSyncSteps(g *geo.Grid) (steps int) {
	isSynced := func(g *geo.Grid) bool {
		for _, c := range g.Data {
			if c != 0 {
				return false
			}
		}
		return true
	}
	for {
		if isSynced(g) {
			return
		}
		steps++
		Step(g)
	}
}
