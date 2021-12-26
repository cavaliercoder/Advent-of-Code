package day23

import (
	"testing"

	"aoc/internal/assert"
)

var (
	ExampleFixture1 = NewState(2, "BACDBCDA")
	ExampleFixture2 = NewState(4, "BDDACCBDBBACDACA")

	TargetState1 = NewState(2, "AABBCCDD")
	TargetState2 = NewState(4, "AAAABBBBCCCCDDDD")

	TestFixture1 = NewState(2, "DCDCABAB")
	TestFixture2 = NewState(4, "DDDCDCBCABABAACB")
)

func TestExample(t *testing.T) {
	assert.Int(
		t,
		12521,
		ExampleFixture1.Organize(TargetState1),
		"bad energy cost",
	)
	assert.Int(
		t,
		44169,
		ExampleFixture2.Organize(TargetState2),
		"bad energy cost",
	)
}

func TestPart1(t *testing.T) {
	assert.Int(
		t,
		16489,
		TestFixture1.Organize(TargetState1),
		"bad energy cost",
	)
}

func TestPart2(t *testing.T) {
	assert.Int(
		t,
		43413,
		TestFixture2.Organize(TargetState2),
		"bad energy cost",
	)
}
