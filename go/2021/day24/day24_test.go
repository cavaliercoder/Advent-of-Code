package day24

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"aoc/internal/assert"
	"aoc/internal/fixture"
)

func openFixture(t *testing.T) *ALU {
	data := make([]Instruction, 0, 64)
	fixture.ScanStrings(t, 2021, 24, func(s string) error {
		if strings.HasPrefix(s, "#") {
			return nil
		}
		var ins Instruction
		ins, err := parseInstruction(s)
		if err != nil {
			return err
		}
		data = append(data, ins)
		return nil
	})
	return NewALU(data)
}

func parseInstruction(s string) (ins Instruction, err error) {
	regID := func(b byte) int { return int(b - 'w') }
	if strings.HasPrefix(s, "inp ") {
		ins.Op = OpINP
		ins.A = regID(s[4])
		return
	}
	tokens := strings.Split(s, " ")
	if len(tokens) != 3 {
		err = fmt.Errorf("bad instruction: %s", s)
		return
	}
	op, a, b := tokens[0], tokens[1], tokens[2]
	switch op {
	case "add":
		ins.Op = OpADD
	case "mul":
		ins.Op = OpMUL
	case "div":
		ins.Op = OpDIV
	case "mod":
		ins.Op = OpMOD
	case "eql":
		ins.Op = OpEQL
	default:
		err = fmt.Errorf("bad operator: %s", s)
		return
	}
	ins.A = regID(a[0])
	if n, err := strconv.Atoi(b); err == nil {
		ins.B = n
	} else {
		ins.B = regID(b[0])
		ins.Op |= 0x80
	}
	return
}

func TestPart1(t *testing.T) {
	// see constraints in Reimp
	c := openFixture(t)
	assert.Bool(t, true, c.Run(99911993949684), "bad model number")
	assert.Bool(t, true, Reimp(99911993949684), "bad model number")
}

func TestPart2(t *testing.T) {
	c := openFixture(t)
	// see constraints in Reimp
	assert.Bool(t, true, c.Run(uint64(62911941716111)), "bad model number")
	assert.Bool(t, true, Reimp(62911941716111), "bad model number")
}
