package day24

import "fmt"

type Operator byte

const (
	OpINP Operator = iota
	OpADD
	OpMUL
	OpDIV
	OpMOD
	OpEQL
)

func (op Operator) IsIndirect() bool { return op&0x80 != 0 }

func (op Operator) String() string {
	switch op & 0x7F {
	case OpINP:
		return "inp"
	case OpADD:
		return "add"
	case OpMUL:
		return "mul"
	case OpDIV:
		return "div"
	case OpMOD:
		return "mod"
	case OpEQL:
		return "eql"
	default:
		return fmt.Sprintf("unknown(0x%02X)", byte(op))
	}
}

type Instruction struct {
	Op   Operator
	A, B int
}

func (ins Instruction) String() string {
	if ins.Op.IsIndirect() {
		return fmt.Sprintf("%s %c %c", ins.Op, 'w'+byte(ins.A), 'w'+byte(ins.B))
	}
	return fmt.Sprintf("%s %c %d", ins.Op, 'w'+byte(ins.A), ins.B)
}

type ALU struct {
	data      []Instruction
	registers [4]int
}

func NewALU(data []Instruction) *ALU { return &ALU{data: data} }

func (c *ALU) Run(input uint64) bool {
	i := 0
	mag := uint64(10000000000000)
	c.registers = [4]int{}
	for p := 0; p < len(c.data); p++ {
		ins := c.data[p]
		a, b := c.registers[ins.A], ins.B
		if ins.Op.IsIndirect() {
			b = c.registers[b]
		}
		op := ins.Op & 0x7F
		switch op {
		case OpINP:
			i++
			n := (input / mag) % 10
			if n == 0 {
				return false
			}
			if n > 9 {
				panic(fmt.Sprintf("bad input (%d %d %d)", input, mag, n))
			}
			mag /= 10
			c.registers[ins.A] = int(n)
		case OpADD:
			c.registers[ins.A] = a + b
		case OpMUL:
			c.registers[ins.A] = a * b
		case OpDIV:
			c.registers[ins.A] = a / b
		case OpMOD:
			c.registers[ins.A] = a % b
		case OpEQL:
			if a == b {
				c.registers[ins.A] = 1
			} else {
				c.registers[ins.A] = 0
			}
		}
	}
	return c.registers[3] == 0
}

// Reimp is a native reimplementation of the input program.
// The program basically converts the input to base-26 with some unexpected
// offsets and precendence. Following along with the program, model numbers have
// the following constraints:
//
//  0: Must be in [6, 9]
//  1: Must be in [2, 9]
//  2: Must be 9
//  3: Must be 1
//  4: Must be 1
//  5: Must be 9
//  6: Must be in [4, 9]
//  7: Must be in [1, 3]
//  8: Must be [7] + 6
//  9: Must be in [1, 4]
// 10: Must be [9] + 5
// 11: Must be [6] - 3]
// 12: Must be [1] - 1
// 13: Must be [0] - 5
func Reimp(model uint64) bool {
	if model%1000000 == 0 {
		fmt.Println(model)
	}
	var z int
	var m [14]int
	mag := uint64(10000000000000)
	for i := 0; i < len(m); i++ {
		m[i] = int((model / mag) % 10)
		mag /= 10
		if m[i] <= 0 || m[i] > 9 {
			return false
		}
	}
	z = m[0] + 6
	z = z*26 + m[1] + 6
	z = z*26 + m[2] + 3
	if m[3] != z%26-11 {
		return false
	}
	z = (z / 26)
	z = z*26 + m[4] + 9
	if m[5] != z%26-1 {
		return false
	}
	z /= 26
	z = z*26 + m[6] + 13
	z = z*26 + m[7] + 6
	if m[8] != z%26 {
		return false
	}
	z /= 26
	z = z*26 + m[9] + 10
	if m[10] != z%26-5 {
		return false
	}
	z /= 26
	if m[11] != z%26-16 {
		return false
	}
	z /= 26
	if m[12] != z%26-7 {
		return false
	}
	z /= 26
	return m[13] == z%26-11
}
