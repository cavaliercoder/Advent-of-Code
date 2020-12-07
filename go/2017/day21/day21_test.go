package day21

import (
	"bufio"
	"bytes"
	"testing"

	. "aoc"
)

func TestPartOneAndTwo(t *testing.T) {
	tests := map[int]int{
		5:  125,
		18: 1782917,
	}

	rules := make(EnhancementRuleList, 0)
	s := bufio.NewScanner(bytes.NewReader(MustReadFixture("day21")))
	for s.Scan() {
		//rules = append(rules, ParseEnhancementRule(s.Text()))
		rules.Add(ParseEnhancementRule(s.Text()))
	}

	for input, expect := range tests {
		actual := Generate(input, rules)
		if actual != expect {
			t.Errorf("expected %d, got %d for %d", expect, actual, input)
		}
	}
}
