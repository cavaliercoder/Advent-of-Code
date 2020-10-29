package main

import (
	. "aoc/2019/common"
	"strings"
	"testing"
)

func TestPart1Example1(t *testing.T) {
	example := `#########
#b.A.@.a#
#########
`
	g, _ := ReadGrid(strings.NewReader(example))
	AssertInt(t, 8, ShortestPath(g), "incorrect shortest path")
}

func TestPart1Example2(t *testing.T) {
	example := `########################
#f.D.E.e.C.b.A.@.a.B.c.#
######################.#
#d.....................#
########################
`
	g, _ := ReadGrid(strings.NewReader(example))
	AssertInt(t, 86, ShortestPath(g), "incorrect shortest path")
}

func TestPart1Example3(t *testing.T) {
	example := `########################
#...............b.C.D.f#
#.######################
#.....@.a.B.c.d.A.e.F.g#
########################
`
	g, _ := ReadGrid(strings.NewReader(example))
	AssertInt(t, 132, ShortestPath(g), "incorrect shortest path")
}

func TestPart1Example4(t *testing.T) {
	example := `#################
#i.G..c...e..H.p#
########.########
#j.A..b...f..D.o#
########@########
#k.E..a...g..B.n#
########.########
#l.F..d...h..C.m#
#################
`
	g, _ := ReadGrid(strings.NewReader(example))
	AssertInt(t, 136, ShortestPath(g), "incorrect shortest path")
}

func TestPart1Example5(t *testing.T) {
	example := `########################
#@..............ac.GI.b#
###d#e#f################
###A#B#C################
###g#h#i################
########################
`
	g, _ := ReadGrid(strings.NewReader(example))
	AssertInt(t, 81, ShortestPath(g), "incorrect shortest path")
}

func TestPart1(t *testing.T) {
	g, _ := OpenGrid(Fixture("day18"))
	AssertInt(t, 5858, ShortestPath(g), "incorrect shortest path")
}

func TestPart2Example1(t *testing.T) {
	example := `#######
#a.#Cd#
##@#@##
#######
##@#@##
#cB#Ab#
#######
`
	g, _ := ReadGrid(strings.NewReader(example))
	AssertInt(t, 8, ShortestPath(g), "incorrect shortest path")
}

func TestPart2Example2(t *testing.T) {
	example := `###############
#d.ABC.#.....a#
######@#@######
###############
######@#@######
#b.....#.....c#
###############
`
	g, _ := ReadGrid(strings.NewReader(example))
	AssertInt(t, 24, ShortestPath(g), "incorrect shortest path")
}

func TestPart2Example3(t *testing.T) {
	example := `#############
#DcBa.#.GhKl#
#.###@#@#I###
#e#d#####j#k#
###C#@#@###J#
#fEbA.#.FgHi#
#############
`
	g, _ := ReadGrid(strings.NewReader(example))
	AssertInt(t, 32, ShortestPath(g), "incorrect shortest path")
}

func TestPart2(t *testing.T) {
	g, _ := OpenGrid(Fixture("day18"))
	Split(g)
	// g.Print(os.Stdout)
	AssertInt(t, 0, ShortestPath(g), "incorrect shortest path") // < 2224
}
