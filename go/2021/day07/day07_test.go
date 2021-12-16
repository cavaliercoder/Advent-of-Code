package day07

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"

	"aoc2021"
)

func mustOpenFixture(name string) []int {
	f := aoc2021.MustOpenFixture(name)
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	parts := strings.Split(strings.TrimSpace(string(b)), ",")
	a := make([]int, 0, len(parts))
	for _, s := range parts {
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(fmt.Sprintf("bad x-position: %s", s))
		}
		a = append(a, int(n))
	}
	return a
}

func TestPart1(t *testing.T) {
	aoc2021.AssertInt(
		t,
		37,
		AlignCrabmarines(16, 1, 2, 0, 4, 2, 7, 1, 2, 14),
		"bad fuel count",
	)
	positions := mustOpenFixture("day07")
	aoc2021.AssertInt(
		t,
		336721,
		AlignCrabmarines(positions...),
		"bad fuel count",
	)
}

func TestPart2(t *testing.T) {
	aoc2021.AssertInt(
		t,
		168,
		AlignCrabmarinesProper(16, 1, 2, 0, 4, 2, 7, 1, 2, 14),
		"bad fuel count",
	)
	positions := mustOpenFixture("day07")
	aoc2021.AssertInt(
		t,
		91638945,
		AlignCrabmarinesProper(positions...),
		"bad fuel count",
	)
}
