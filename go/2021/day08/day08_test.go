package day08

import (
	"aoc2021"
	"bufio"
	"fmt"
	"strings"
	"testing"
)

type fixture struct {
	Patterns [10]Display
	Output   [4]Display
}

func parseFixture(s string) (*fixture, error) {
	c := &fixture{}
	parts := strings.Split(s, " ")
	if len(parts) != 15 {
		return nil, fmt.Errorf("bad fixture: %s", s)
	}
	for i := 0; i < 10; i++ {
		p, err := ParseDisplay([]byte(parts[i]))
		if err != nil {
			return nil, err
		}
		c.Patterns[i] = p
	}
	for i := 0; i < 4; i++ {
		p, err := ParseDisplay([]byte(parts[11+i]))
		if err != nil {
			return nil, err
		}
		c.Output[i] = p
	}
	return c, nil
}

func mustOpenFixture(name string) []*fixture {
	f, err := aoc2021.OpenFixture(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	a := make([]*fixture, 0, 64)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		v, err := parseFixture(s)
		if err != nil {
			panic(err)
		}
		a = append(a, v)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return a
}

func TestPart1(t *testing.T) {
	count := 0
	tests := mustOpenFixture("day08")
	for _, test := range tests {
		for _, output := range test.Output {
			switch output.SegmentCount() {
			case 2, 4, 3, 7: // segment counts of 1, 4, 7, 8
				count++
			}
		}
	}
	aoc2021.AssertInt(t, 284, count, "bad value count")
}

func TestPart2(t *testing.T) {
	sum := 0
	tests := mustOpenFixture("day08")
	for _, test := range tests {
		cipher := Rewire(test.Patterns)
		sum += Decode(cipher, test.Output[:]...)
	}
	aoc2021.AssertInt(t, 973499, sum, "bad display sum")
}
