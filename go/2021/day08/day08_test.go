package day08

import (
	"fmt"
	"strings"
	"testing"

	"aoc/internal/assert"
	"aoc/internal/fixture"
)

type Sample struct {
	Patterns [10]Display
	Output   [4]Display
}

func parseFixture(s string) (*Sample, error) {
	c := &Sample{}
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

func openFixture(t *testing.T) []*Sample {
	a := make([]*Sample, 0, 64)
	fixture.ScanStrings(t, 2021, 8, func(s string) error {
		v, err := parseFixture(s)
		if err != nil {
			return err
		}
		a = append(a, v)
		return nil
	})
	return a
}

func TestPart1(t *testing.T) {
	count := 0
	tests := openFixture(t)
	for _, test := range tests {
		for _, output := range test.Output {
			switch output.SegmentCount() {
			case 2, 4, 3, 7: // segment counts of 1, 4, 7, 8
				count++
			}
		}
	}
	assert.Int(t, 284, count, "bad value count")
}

func TestPart2(t *testing.T) {
	sum := 0
	tests := openFixture(t)
	for _, test := range tests {
		cipher := Rewire(test.Patterns)
		sum += Decode(cipher, test.Output[:]...)
	}
	assert.Int(t, 973499, sum, "bad display sum")
}
