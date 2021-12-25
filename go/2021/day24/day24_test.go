package day24

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"testing"

	. "aoc2021"
)

func mustOpenFixture(name string) *ALU {
	f := MustOpenFixture(name)
	defer f.Close()
	data, err := readProgram(f)
	if err != nil {
		panic(err)
	}
	return NewALU(data)
}

func readProgram(r io.Reader) (program []Instruction, err error) {
	program = make([]Instruction, 0, 64)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s := scanner.Text()
		if strings.HasPrefix(s, "#") {
			continue
		}
		var ins Instruction
		ins, err = parseInstruction(scanner.Text())
		if err != nil {
			return
		}
		program = append(program, ins)
	}
	if err = scanner.Err(); err != nil {
		return
	}
	return
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
	c := mustOpenFixture("day24")
	AssertBool(t, true, c.Run(99911993949684), "bad model number")
	AssertBool(t, true, Reimp(99911993949684), "bad model number")
}

func TestPart2(t *testing.T) {
	c := mustOpenFixture("day24")
	// see constraints in Reimp
	AssertBool(t, true, c.Run(uint64(62911941716111)), "bad model number")
	AssertBool(t, true, Reimp(62911941716111), "bad model number")
}
