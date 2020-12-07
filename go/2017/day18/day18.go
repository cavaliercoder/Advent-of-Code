package day18

import (
	"bufio"
	"strconv"
	"strings"
)

// Instruction is a single assembly instruction and its parameters.
type Instruction [3]string

// ParseInstructions returns an array of Instructions given a line separated
// list of instructions as input.
func ParseInstructions(s string) []Instruction {
	v := make([]Instruction, 0)
	scanner := bufio.NewScanner(strings.NewReader(s))
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
	snd       int
	rcv       []int
}

// NewProcess returns a process for the given program with the register p set to
// pid.
func NewProcess(pid int, program string) *Process {
	return &Process{
		registers: RegisterFile{'p': pid},
		program:   ParseInstructions(program),
	}
}

// Run mutates the state of a Process by stepping though the associated program.
// Run yeilds whenever snd is called, or if rcv is called but no values are
// available in the receive buffer. The name of the last instruction is returned
// whenever control is yielded.
func (p *Process) Run() string {
	reg := p.registers
	for ; p.pointer < len(p.program); p.pointer++ {
		ins := p.program[p.pointer]
		switch ins[0] {
		case "snd":
			p.snd = reg.Get(ins[1])
			p.pointer++
			return "snd"

		case "set":
			reg[ins[1][0]] = reg.Get(ins[2])

		case "add":
			reg[ins[1][0]] += reg.Get(ins[2])

		case "mul":
			reg[ins[1][0]] *= reg.Get(ins[2])

		case "mod":
			reg[ins[1][0]] %= reg.Get(ins[2])

		case "rcv":
			r := ins[1][0]
			if r != 0 {
				if len(p.rcv) == 0 {
					return "rcv"
				}
				reg[r] = p.rcv[0]
				p.rcv = p.rcv[1:]
			}

		case "jgz":
			v := reg.Get(ins[1])
			if v > 0 {
				v = reg.Get(ins[2])
				p.pointer += v - 1
			}
		}
	}
	panic("program finished early")
}

// LastSnd is my solution to Part One and returns the value of the last call to
// snd before the program is terminated by any call to rcv.
func LastSnd(program string) int {
	pid0 := NewProcess(0, program)
	for pid0.Run() != "rcv" {
	}
	return pid0.snd
}

// Duet is my solution to Part Two and returns the number of times the second
// instance of the given program sends messages to the first program, before the
// programs terminate in a deadlock.
func Duet(program string) int {
	count := 0
	pid0 := NewProcess(0, program)
	pid1 := NewProcess(1, program)
	var state0, state1 string
	for {
		state0 = pid0.Run()
		if state0 == "snd" {
			pid1.rcv = append(pid1.rcv, pid0.snd)
		}

		state1 = pid1.Run()
		if state1 == "snd" {
			count++
			pid0.rcv = append(pid0.rcv, pid1.snd)
		}

		if state0 == "rcv" && state1 == "rcv" {
			return count
		}
	}
}
