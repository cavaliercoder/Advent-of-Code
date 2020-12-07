package day02

import (
	"bufio"
	"bytes"
	"testing"

	. "aoc"
)

func ParseIds(name string) [][]byte {
	f := MustOpenFixture("day02")
	defer f.Close()

	ids := make([][]byte, 0)
	s := bufio.NewScanner(f)
	for s.Scan() {
		b := make([]byte, len(s.Bytes()))
		copy(b, s.Bytes())
		ids = append(ids, b)
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	return ids
}

func TestPartOne(t *testing.T) {
	tests := []struct {
		input  [][]byte
		expect int
	}{
		{
			input: [][]byte{
				[]byte("abcdef"),
				[]byte("bababc"),
				[]byte("abbcde"),
				[]byte("abcccd"),
				[]byte("aabcdd"),
				[]byte("abcdee"),
				[]byte("ababab"),
			},
			expect: 12,
		},
		{
			input:  ParseIds("fixture1"),
			expect: 8296,
		},
	}
	for i, test := range tests {
		actual := ChecksumIDList(test.input)
		if actual != test.expect {
			t.Errorf("expected: %v, got: %v, in test: %v", test.expect, actual, i+1)
		}
	}
}

func TestPartTwo(t *testing.T) {
	tests := []struct {
		input  [][]byte
		expect []byte
	}{
		{
			input: [][]byte{
				[]byte("abcde"),
				[]byte("fghij"),
				[]byte("klmno"),
				[]byte("pqrst"),
				[]byte("fguij"),
				[]byte("axcye"),
				[]byte("wvxyz"),
			},
			expect: []byte("fgij"),
		},
		{
			input:  ParseIds("fixture1"),
			expect: []byte("pazvmqbftrbeosiecxlghkwud"),
		},
	}
	for i, test := range tests {
		actual := FindPrototypeIDs(test.input)
		if !bytes.Equal(actual, test.expect) {
			t.Errorf("expected: %s, got: %s, in test: %v", test.expect, actual, i+1)
		}
	}
}
