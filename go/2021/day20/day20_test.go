package day20

import (
	"bufio"
	"testing"

	. "aoc2021"
)

func mustReadFixture(name string) (algo []byte, state *State) {
	f := MustOpenFixture(name)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	if scanner.Scan() {
		b := scanner.Bytes()
		algo = make([]byte, len(b))
		for i, c := range b {
			switch c {
			case '.':
				continue
			case '#':
				algo[i] = 1
			default:
				panic(c)
			}
		}
	}
	y := 0
	state = NewState()
	scanner.Scan()
	for scanner.Scan() {
		for x, c := range scanner.Bytes() {
			switch c {
			case '.':
				continue
			case '#':
				state.Set(Pos{X: x, Y: y})
			default:
				panic(c)
			}
		}
		y--
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return
}

func TestPart1(t *testing.T) {
	algo, state := mustReadFixture("day20")
	for i := 0; i < 2; i++ {
		state = state.Step(algo)
	}
	t.Log(state)
	AssertInt(t, 5291, state.LitCount(), "bad lit count")
}

func TestPart2(t *testing.T) {
	algo, state := mustReadFixture("day20")
	for i := 0; i < 50; i++ {
		state = state.Step(algo)
	}
	t.Log(state)
	AssertInt(t, 16665, state.LitCount(), "bad lit count")
}
