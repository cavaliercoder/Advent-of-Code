package day23

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

// Instruction is a single assembly instruction and its parameters.
type Instruction [3]string

// ParseInstructions returns an array of Instructions given a line separated
// list of instructions as input.
func ParseInstructions(r io.Reader) []Instruction {
	v := make([]Instruction, 0)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		tkns := strings.Split(scanner.Text(), " ")
		ins := Instruction{}
		copy(ins[:], tkns)
		v = append(v, ins)
	}
	return v
}

// RegisterFile is a mapping of registers and their values, indexed by name.
type RegisterFile map[byte]int

// Get parses the given string as an integer or the name of a register. If the
// name of register is given, the value of that register is returned.
func (r RegisterFile) Get(s string) int {
	if s[0] >= 'a' && s[0] <= 'z' {
		return r[s[0]]
	}
	v, _ := strconv.Atoi(s)
	return v
}

// A Process tracks the state of a running program.
type Process struct {
	registers RegisterFile
	program   []Instruction
	pointer   int
	mul       int
}

// NewProcess returns a process for the given program.
func NewProcess(r io.Reader) *Process {
	return &Process{
		registers: make(RegisterFile, 16),
		program:   ParseInstructions(r),
	}
}

// Run mutates the state of a Process by stepping though the associated program.
// This function is my solution to Part One, as it tracks the number of calls to
// `mul` in the Process struct.
func (p *Process) Run() {
	reg := p.registers
	for p.pointer = 0; p.pointer < len(p.program); p.pointer++ {
		ins := p.program[p.pointer]
		switch ins[0] {
		case "set":
			reg[ins[1][0]] = reg.Get(ins[2])

		case "sub":
			reg[ins[1][0]] -= reg.Get(ins[2])

		case "mul":
			reg[ins[1][0]] *= reg.Get(ins[2])
			p.mul++

		case "jnz":
			v := reg.Get(ins[1])
			if v != 0 {
				v = reg.Get(ins[2])
				p.pointer += v - 1
			}
		}
	}
}

// Optimized is my solution to Part Two. It represents a decompiled (by hand)
// version of the input program, in native Go.
//
// In the original program, two inner loops were used to find all divisors
// greater than one - effectively testing if a number is a prime number in
// O(n²). This version replaces the inner loops with a single loop that tests
// for prime membership in O(n).
func Optimized(a int) int {
	i, n, v := 81, 81, 0
	if a != 0 {
		i, n = 108100, 125100
	}
	for ; i <= n; i += 17 {
		for d := 2; d < i/2; d++ {
			if i%d == 0 {
				v++
				break
			}
		}
	}
	return v
}
