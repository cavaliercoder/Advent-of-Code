package day12

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"testing"

	. "aoc2021"
)

func parseGraph(r io.Reader) (map[string][]string, error) {
	m := make(map[string][]string)
	addEdge := func(a, b string) {
		e, ok := m[a]
		if !ok {
			e = make([]string, 0, 1)
		}
		m[a] = append(e, b)
	}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s := scanner.Text()
		edges := strings.Split(s, "-")
		if len(edges) != 2 {
			return nil, fmt.Errorf("bad edge: %s", s)
		}
		addEdge(edges[0], edges[1])
		addEdge(edges[1], edges[0])
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return m, nil
}

func mustOpenFixture(name string) map[string][]string {
	f := MustOpenFixture(name)
	defer f.Close()
	m, err := parseGraph(f)
	if err != nil {
		panic(err)
	}
	return m
}

func TestPart1(t *testing.T) {
	m := mustOpenFixture("day12")
	AssertInt(t, 5076, CountPaths1(m), "bad path count")
}
func TestPart2(t *testing.T) {
	m := mustOpenFixture("day12")
	AssertInt(t, 145643, CountPaths2(m), "bad path count")
}
