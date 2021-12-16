package day06

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
			panic(fmt.Sprintf("bad fish: %s", s))
		}
		a = append(a, int(n))
	}
	return a
}

func TestPart1(t *testing.T) {
	fishes := mustOpenFixture("day06")
	aoc2021.AssertInt(t, 354564, GenerateFish(80, fishes), "bad fish count")
}

func TestPart2(t *testing.T) {
	fishes := mustOpenFixture("day06")
	aoc2021.AssertInt(
		t,
		1609058859115,
		GenerateFish(256, fishes),
		"bad fish count",
	)
}
