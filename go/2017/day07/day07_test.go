package day07

import (
	"testing"

	. "aoc"
)

func getTestGraph() Graph {
	f := MustOpenFixture("day07")
	defer f.Close()
	G, err := parseGraph(f)
	if err != nil {
		panic(err)
	}
	return G
}

func TestBottomProgram(t *testing.T) {
	G := getTestGraph()
	expect := "rqwgj"
	actual := getBottomProgram(G)
	if actual.Name != expect {
		t.Errorf("expected %v, got %v", expect, actual.Name)
	}
}

func TestCorrectProgramWeight(t *testing.T) {
	G := getTestGraph()
	expect := 333
	actual := getCorrectProgramWeight(G)
	if actual != expect {
		t.Errorf("expected %v, got %v", expect, actual)
	}
}
