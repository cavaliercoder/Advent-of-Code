package day16

import (
	"bytes"
	"fmt"
	"testing"

	"aoc/internal/assert"
	"aoc/internal/fixture"
)

func openFixture(t *testing.T) []byte {
	b := fixture.Bytes(t, 2021, 16)
	b = bytes.TrimSuffix(b, []byte("\n"))
	return b
}

func TestPart1(t *testing.T) {
	b := openFixture(t)
	n := 0
	tokens, err := Lex(b)
	if err != nil {
		t.Error(err)
		return
	}
	for _, p := range tokens {
		n += p.Version
	}
	assert.Int(t, 989, n, "bad version sum")
}

func TestPart2(t *testing.T) {
	tests := []struct {
		Fixture []byte
		Expect  int
	}{
		{
			Fixture: []byte("C200B40A82"),
			Expect:  3,
		},
		{
			Fixture: []byte("04005AC33890"),
			Expect:  54,
		},
		{
			Fixture: []byte("880086C3E88112"),
			Expect:  7,
		},
		{
			Fixture: []byte("CE00C43D881120"),
			Expect:  9,
		},
		{
			Fixture: []byte("D8005AC2A8F0"),
			Expect:  1,
		},
		{
			Fixture: []byte("F600BC2D8F"),
			Expect:  0,
		},
		{
			Fixture: []byte("9C005AC2F8F0"),
			Expect:  0,
		},
		{
			Fixture: []byte("9C0141080250320F1802104A08"),
			Expect:  1,
		},
		{
			Fixture: openFixture(t),
			Expect:  7936430475134,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("%02d", i+1), func(t *testing.T) {
			expr, err := Parse(test.Fixture)
			if err != nil {
				t.Error(err)
				return
			}
			v := expr.Eval()
			t.Log(expr, "=", v)
			assert.Int(t, test.Expect, v, "bad expression result")
		})
	}
}
