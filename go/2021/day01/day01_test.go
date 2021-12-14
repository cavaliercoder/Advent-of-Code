package day01

import (
	"bufio"
	"strconv"
	"testing"

	"aoc2021"
)

func OpenFixture(name string) []int {
	f, err := aoc2021.OpenFixture(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	a := make([]int, 0, 64)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		a = append(a, n)
	}
	return a
}

func TestPart1(t *testing.T) {
	a := OpenFixture("day01")
	aoc2021.AssertInt(t, 1462, CountIncreases(a), "Bad")
}
func TestPart2(t *testing.T) {
	a := OpenFixture("day01")
	aoc2021.AssertInt(t, 1497, CountIncreasesSliding(a), "Bad")
}
