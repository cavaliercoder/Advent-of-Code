package day01

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func fixture(name string) string {
	f, err := os.Open(fmt.Sprintf("../../../inputs/2015/%s.txt", name))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func TestPartOne(t *testing.T) {
	tests := map[string]int{
		"(())":           0,
		"(((":            3,
		"(()(()(":        3,
		"))(((((":        3,
		"())":            -1,
		"))(":            -1,
		")))":            -3,
		")())())":        -3,
		fixture("day01"): 138,
	}
	for input, expect := range tests {
		actual := Elevate([]byte(input))
		if actual != expect {
			t.Errorf("expected %d, got %d", expect, actual)
		}
	}
}

func TestPartTwo(t *testing.T) {
	tests := map[string]int{
		")":                      1,
		"()())":                  5,
		string(fixture("day01")): 1771,
	}
	for input, expect := range tests {
		actual := GoToBasement([]byte(input))
		if actual != expect {
			t.Errorf("expected %d, got %d", expect, actual)
		}
	}
}
