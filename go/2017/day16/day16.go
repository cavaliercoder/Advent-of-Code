package day16

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type DanceFloor struct {
	b []byte       // array of programs
	m map[byte]int // map programs to their offset in b
}

// NewDanceFloor returns a DanceFloor of the given size, with all programs
// initialized in sequential order.
func NewDanceFloor(n int) *DanceFloor {
	df := &DanceFloor{
		b: make([]byte, n),
		m: make(map[byte]int, n),
	}
	var c byte
	for i := 0; i < n; i++ {
		c = 'a' + byte(i)
		df.b[i] = c
		df.m[c] = i
	}
	return df
}

// Do executes all the given dance moves, in order, n times.
//
// The function serves as my solution to Part One and Two. For Part Two, a
// simple optimization identifies the iteration at which the state of the dance
// floor loops back to its initial state. Subsequent loops can be skipped and
// the final state can be returned from cache.
//
// This solution runs in O(n) linear time where n is the number of times the
// dance moves must be run before the dance floor returns to its initial state.
func (d *DanceFloor) Do(moves []DanceMove, n int) {
	var loop int
	cache := make([][]byte, 0)
	for i := 0; i < n; i++ {
		cache = append(cache, make([]byte, len(d.b)))
		copy(cache[i], d.b)
		for m := 0; m < len(moves); m++ {
			moves[m].Do(d)
		}
		if bytes.Equal(d.b, cache[0]) {
			loop = i + 1
			break
		}
	}
	if loop == 0 {
		return
	}
	copy(d.b, cache[n%loop])
}

// A DanceMove transforms the arrangement of the dance floor.
type DanceMove interface {
	Do(*DanceFloor)
}

// ParseDanceMoves reads dance moves from the comma-separated input.
func ParseDanceMoves(s string) []DanceMove {
	tkns := strings.Split(strings.Trim(s, "\n"), ",")
	moves := make([]DanceMove, len(tkns))
	for i := 0; i < len(tkns); i++ {
		tkn := tkns[i]
		switch tkn[0] {
		case 's':
			v, err := strconv.Atoi(tkn[1:])
			if err != nil {
				panic(fmt.Sprintf("invalid spin token: %s", tkn))
			}
			moves[i] = &spin{size: v}

		case 'x':
			v := strings.Split(tkn[1:], "/")
			a, _ := strconv.Atoi(v[0])
			b, _ := strconv.Atoi(v[1])
			moves[i] = &exchange{a: a, b: b}

		case 'p':
			v := strings.Split(tkn[1:], "/")
			moves[i] = &partner{a: v[0][0], b: v[1][0]}

		default:
			panic(fmt.Sprintf("invalid token: %s", tkn))
		}
	}
	return moves
}

// spin moves programs from the end of the dance floor, to the beginning.
type spin struct {
	size int
}

func (s *spin) Do(d *DanceFloor) {
	x := len(d.b) - s.size
	d.b = append(d.b[x:], d.b[:x]...)
	for i := 0; i < len(d.b); i++ {
		d.m[d.b[i]] = i
	}
}

// exchange swaps the programs at positions a and b.
type exchange struct {
	a, b int
}

func (e *exchange) Do(d *DanceFloor) {
	d.b[e.a], d.b[e.b] = d.b[e.b], d.b[e.a]
	d.m[d.b[e.a]] = e.a
	d.m[d.b[e.b]] = e.b
}

// partner swaps the programs with the names a and b.
type partner struct {
	a, b byte
}

func (p *partner) Do(d *DanceFloor) {
	e := &exchange{a: d.m[p.a], b: d.m[p.b]}
	e.Do(d)
}
