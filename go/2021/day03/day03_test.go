package day03

import (
	"aoc2021"
	"bufio"
	"testing"
)

func mustOpenFixture(name string) []int {
	f, err := aoc2021.OpenFixture(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	a := make([]int, 0, 64)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		a = append(a, Parse(scanner.Bytes()))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return a
}

func TestPart1(t *testing.T) {
	a := mustOpenFixture("day03")
	aoc2021.AssertInt(
		t,
		3969000,
		PowerConsumptionRate(a...),
		"bad power consumption rate",
	)
}

func TestPart2(t *testing.T) {
	a := mustOpenFixture("day03")
	aoc2021.AssertInt(
		t,
		4267809,
		LifeSupportRating(a...),
		"bad life support rating",
	)
}
