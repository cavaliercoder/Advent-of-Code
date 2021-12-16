package day07

import (
	"sort"
)

type fuelFunc func(x int, positions []int) (fuel int)

func AlignCrabmarines(positions ...int) (fuel int) {
	f := func(x int, positions []int) (fuel int) {
		for _, p := range positions {
			n := p - x
			if n < 0 {
				n = -n
			}
			fuel += n
		}
		return
	}
	return alignCrabmarines(f, positions...)
}

func AlignCrabmarinesProper(positions ...int) (fuel int) {
	f := func(x int, positions []int) (fuel int) {
		for _, p := range positions {
			n := p - x
			if n < 0 {
				n = -n
			}
			fuel += n * (n + 1) / 2
		}
		return
	}
	return alignCrabmarines(f, positions...)
}

func alignCrabmarines(f fuelFunc, positions ...int) (fuel int) {
	sort.Ints(positions)
	l, r := positions[0], positions[len(positions)-1]
	lCost := f(l, positions)
	rCost := f(r, positions)
	for {
		if l == r || l == r-1 {
			if lCost < rCost {
				return lCost
			} else {
				return rCost
			}
		}
		m := (r + l) / 2
		if lCost < rCost {
			r = m
			rCost = f(r, positions)
		} else {
			l = m
			lCost = f(l, positions)
		}
	}
}
