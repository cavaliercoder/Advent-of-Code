package day08

import (
	"fmt"
	"strconv"
	"strings"
)

type RegisterFile map[string]int

type Operator int

const (
	NotEqual Operator = iota
	LesserThan
	LesserOrEqual
	Equal
	GreaterOrEqual
	GreaterThan
)

type Instruction struct {
	Register   string
	Offset     int
	Predictate struct {
		Register string
		Operator Operator
		Value    int
	}
}

// ParseInstruction parses a CPU instruction from the given string of the form:
// "[register] [inc|dec] if [register] [operator] [value]"
func ParseInstruction(s string) (*Instruction, error) {
	tkns := strings.Split(s, " ")
	if len(tkns) != 7 {
		return nil, fmt.Errorf("invalid token count: %d in '%s'", len(tkns), s)
	}
	ins := &Instruction{
		Register: tkns[0],
	}

	i, err := strconv.ParseInt(tkns[2], 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid offset: %s: %v", tkns[2], err)
	}
	ins.Offset = int(i)

	switch tkns[1] {
	case "dec":
		ins.Offset = 0 - ins.Offset
	case "inc":
		// do nothing
	default:
		return nil, fmt.Errorf("invalid offset direction: %s", tkns[1])
	}

	if tkns[3] != "if" {
		return nil, fmt.Errorf("invalid token: %s", tkns[3])
	}

	ins.Predictate.Register = tkns[4]

	switch tkns[5] {
	case "!=":
		ins.Predictate.Operator = NotEqual
	case "<":
		ins.Predictate.Operator = LesserThan
	case "<=":
		ins.Predictate.Operator = LesserOrEqual
	case "==":
		ins.Predictate.Operator = Equal
	case ">=":
		ins.Predictate.Operator = GreaterOrEqual
	case ">":
		ins.Predictate.Operator = GreaterThan
	default:
		return nil, fmt.Errorf("invalid operator: %s", tkns[5])
	}

	i, err = strconv.ParseInt(tkns[6], 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid integer value: %s: %v", tkns[6], err)
	}
	ins.Predictate.Value = int(i)

	return ins, nil
}

// Mutate executes an instruction to mutate the state of the RegisterFile if the
// instruction's condition is evaluated as true.
func (r RegisterFile) Mutate(ins *Instruction) {
	v := r[ins.Register]
	x := r[ins.Predictate.Register]

	var ok bool
	switch ins.Predictate.Operator {
	case NotEqual:
		ok = x != ins.Predictate.Value
	case LesserThan:
		ok = x < ins.Predictate.Value
	case LesserOrEqual:
		ok = x <= ins.Predictate.Value
	case Equal:
		ok = x == ins.Predictate.Value
	case GreaterOrEqual:
		ok = x >= ins.Predictate.Value
	case GreaterThan:
		ok = x > ins.Predictate.Value
	}
	if ok {
		r[ins.Register] = v + ins.Offset
	}
}

// Max returns the value of the register with the highest value.
func (r RegisterFile) Max() int {
	var m int = ^(1 << 32)
	for _, v := range r {
		if v > m {
			m = v
		}
	}
	return m
}
