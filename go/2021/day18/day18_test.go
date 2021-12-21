package day18

import (
	"bufio"
	"fmt"
	"testing"

	. "aoc2021"
)

func mustOpenFixture(name string) []*Number {
	f := MustOpenFixture(name)
	defer f.Close()
	a := make([]*Number, 0, 64)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Bytes()
		a = append(a, Parse(s))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return a
}

func TestParser(t *testing.T) {
	tests := []string{
		"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
		"[[[5,[2,8]],4],[5,[[9,9],0]]]",
		"[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
		"[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]",
		"[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
		"[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
		"[[[[5,4],[7,7]],8],[[8,3],8]]",
		"[[9,3],[[9,9],[6,[4,9]]]]",
		"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
		"[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]",
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("%02d", i+1), func(t *testing.T) {
			n := Parse([]byte(test))
			AssertString(t, test, n.String(), "bad parser output")
		})
	}
}

func TestExplode(t *testing.T) {
	tests := map[string]string{
		"[[[[[9,8],1],2],3],4]":                 "[[[[0,9],2],3],4]",
		"[7,[6,[5,[4,[3,2]]]]]":                 "[7,[6,[5,[7,0]]]]",
		"[[6,[5,[4,[3,2]]]],1]":                 "[[6,[5,[7,0]]],3]",
		"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]": "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
		"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]":     "[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
	}
	for input, expect := range tests {
		n := Parse([]byte(input))
		n.explode(0, nil, nil)
		AssertString(t, expect, n.String(), "bad reduction")
	}
}

func TestSplit(t *testing.T) {
	tests := map[string]string{
		"[[[[0,7],4],[15,[0,13]]],[1,1]]":    "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]",
		"[[[[0,7],4],[[7,8],[0,13]]],[1,1]]": "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]",
	}
	for input, expect := range tests {
		n := Parse([]byte(input))
		n.split()
		AssertString(t, expect, n.String(), "bad split")
	}
}

func TestAdd(t *testing.T) {
	a := Parse([]byte("[[[[4,3],4],4],[7,[[8,4],9]]]"))
	b := Parse([]byte("[1,1]"))
	v := Add(a, b)
	AssertString(
		t,
		"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		v.String(),
		"bad addition",
	)
}

func TestMaxMagnitude(t *testing.T) {
	A := []*Number{
		Parse([]byte("[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]")),
		Parse([]byte("[[[5,[2,8]],4],[5,[[9,9],0]]]")),
		Parse([]byte("[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]")),
		Parse([]byte("[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]")),
		Parse([]byte("[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]")),
		Parse([]byte("[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]")),
		Parse([]byte("[[[[5,4],[7,7]],8],[[8,3],8]]")),
		Parse([]byte("[[9,3],[[9,9],[6,[4,9]]]]")),
		Parse([]byte("[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]")),
		Parse([]byte("[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]")),
	}
	AssertInt(t, 3993, MaxMagnitude(A...), "bad maximum magnitude")
}

func TestPart1(t *testing.T) {
	a := mustOpenFixture("day18")
	AssertInt(t, 3524, Add(a...).M(), "bad magnitude")
}

func TestPart2(t *testing.T) {
	a := mustOpenFixture("day18")
	AssertInt(t, 4656, MaxMagnitude(a...), "bad maximum magnitude")
}
