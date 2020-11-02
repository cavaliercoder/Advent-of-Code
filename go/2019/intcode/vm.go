package intcode

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	ErrIllegal = errors.New("illegal")
	ErrHalted  = errors.New("halted")
)

// Opcode is the ID of a function that the VM can perform.
type Opcode int

const (
	OpcodeIllegal Opcode = iota
	OpcodeAdd
	OpcodeMultiply
	OpcodeInput
	OpcodeOutput
	OpcodeJCC
	OpcodeJNC
	OpcodeLessThan
	OpcodeEqual
	OpcodeRelBase
	OpcodeHalt = 99
)

type opcodeSignature struct {
	name   string
	params []string
}

var opcodeSignatures = map[Opcode]opcodeSignature{
	OpcodeAdd:      {"add", []string{"a", "b", "addr"}},
	OpcodeMultiply: {"multiple", []string{"a", "b", "addr"}},
	OpcodeInput:    {"input", []string{"addr"}},
	OpcodeOutput:   {"output", []string{"addr"}},
	OpcodeJCC:      {"jump-if-true", []string{"v", "addr"}},
	OpcodeJNC:      {"jump-if-false", []string{"v", "addr"}},
	OpcodeLessThan: {"less-than", []string{"a", "b", "addr"}},
	OpcodeEqual:    {"equal", []string{"a", "b", "addr"}},
	OpcodeRelBase:  {"adjust-rel-base", []string{"offset"}},
	OpcodeHalt:     {"halt", []string{}},
}

type opcodeHandler func() error

// AddressMode describes how an Address should be dereferenced by the VM.
type AddressMode int

const (
	// AddressModePosition reads the memory address of a value.
	AddressModePosition AddressMode = iota

	// AddressModeImmediate reads a value directly.
	AddressModeImmediate

	// AddressModeRelative reads the memory address of a value relative to the
	// value of the Relative Base Register.
	AddressModeRelative
)

// Address points to a location in VM mempory.
type Address int

// Instruction is an integer encoding of an opcode and the address mode of each
// of its parameters.
type Instruction int

// Opcode returns the Opcode of the instruction.
func (c Instruction) Opcode() Opcode {
	return Opcode(c) % 100
}

// Modes returns the address mode of each paramter in the instruction.
func (c Instruction) Modes() [4]AddressMode {
	modes := [4]AddressMode{0, 0, 0, 0}
	v := AddressMode(c) / 100
	for i := 0; i < len(modes); i++ {
		modes[i] = v % 10
		v /= 10
	}
	return modes
}

// A VM runs IntCode programs.
type VM interface {
	// Step will fetch, decode and execute a single instruction
	Step() error

	// Run until the program halts
	Run() (v int, err error)

	// MemGet reads a value from VM memory
	MemGet(addr Address, mode AddressMode) int

	// MemSet writes a value to VM memory
	MemSet(addr Address, v int)
}

type vm struct {
	data        Data
	handlers    map[Opcode]opcodeHandler
	stdin       IntReader
	stdout      IntWriter
	exitCode    int
	opcode      Opcode
	modes       [4]AddressMode
	regPC       Address
	regRelBase  Address
	traceWriter io.Writer
}

// New returns a new VM.
func New(data Data) VM {
	return NewWithIO(data, NewIntBuffer(nil), NewIntBuffer(nil))
}

// NewWithIO returns a new VM with the given IO devices.
func NewWithIO(data Data, stdin IntReader, stdout IntWriter) VM {
	v := &vm{
		data:   data,
		stdin:  stdin,
		stdout: stdout,
	}
	v.handlers = map[Opcode]opcodeHandler{
		OpcodeIllegal:  v.opIllegal,
		OpcodeAdd:      v.opAdd,
		OpcodeMultiply: v.opMultiply,
		OpcodeInput:    v.opInput,
		OpcodeOutput:   v.opOutput,
		OpcodeJCC:      v.opJumpIfTrue,
		OpcodeJNC:      v.opJumpIfNotTrue,
		OpcodeLessThan: v.opLessThan,
		OpcodeEqual:    v.opEqual,
		OpcodeRelBase:  v.opRelBase,
		OpcodeHalt:     v.opHalt,
	}
	if os.Getenv("LOGLEVEL") == "DEBUG" {
		v.traceWriter = os.Stderr
	}
	return v
}

func (c *vm) tracef(format string, a ...interface{}) {
	if c.traceWriter == nil {
		return
	}
	fmt.Fprintf(c.traceWriter, format, a...)
}

func (c *vm) traceStep() {
	if c.traceWriter == nil {
		return
	}
	sig := opcodeSignatures[c.opcode]
	fmt.Fprintf(c.traceWriter, "[%04d] %s(", c.regPC, sig.name)
	for i := 0; i < len(sig.params); i++ {
		if i > 0 {
			fmt.Fprintf(c.traceWriter, ", ")
		}
		v := c.data[c.regPC+Address(1+i)]
		switch c.modes[i] {
		case AddressModePosition:
			fmt.Fprintf(c.traceWriter, "%s=%d", sig.params[i], v)
		case AddressModeImmediate:
			fmt.Fprintf(c.traceWriter, "%s=#%d", sig.params[i], v)
		case AddressModeRelative:
			fmt.Fprintf(c.traceWriter, "%s=%d+$%d", sig.params[i], v, c.regRelBase)
		}
	}
	fmt.Fprintf(c.traceWriter, ")  // ")
	for i := 0; i < len(sig.params)+1; i++ {
		if i > 0 {
			fmt.Fprintf(c.traceWriter, ", ")
		}
		fmt.Fprintf(c.traceWriter, "%d", c.data[c.regPC+Address(i)])
	}
	fmt.Fprintf(c.traceWriter, "\n")
}

func (c *vm) Step() error {
	// fetch
	v := Instruction(c.MemGet(c.regPC, AddressModeImmediate))

	// decode
	c.opcode = v.Opcode()
	c.modes = v.Modes()

	// execute
	c.traceStep()
	fn, ok := c.handlers[c.opcode]
	if !ok {
		fn = c.opIllegal
	}
	return fn()
}

func (c *vm) Run() (v int, err error) {
	for {
		if err = c.Step(); err != nil {
			if err == ErrHalted {
				return c.exitCode, nil
			}
			return
		}
	}
}

func (c *vm) MemGet(addr Address, mode AddressMode) (v int) {
	eaddr := c.effectiveAddr(addr, mode)
	if addr < 0 {
		panic("invalid address")
	}
	if int(eaddr) >= len(c.data) {
		v = 0
	} else {
		v = c.data[eaddr]
	}
	c.tracef("  memget(%04d, %d) -> %v\n", addr, mode, v)
	return
}

func (c *vm) MemSet(addr Address, v int) {
	if int(addr) >= len(c.data) {
		// grow mem
		sz := 64
		for sz < int(addr)+1 {
			sz *= 2
		}
		data := make(Data, sz)
		copy(data, c.data)
		c.data = data
	}

	c.data[addr] = v
	c.tracef("  memset(%04d, %d)\n", addr, v)
}

func (c *vm) effectiveAddr(addr Address, mode AddressMode) Address {
	if addr < 0 {
		panic("bad address")
	}
	switch mode {
	case AddressModeImmediate:
		return addr
	case AddressModePosition:
		return Address(c.MemGet(addr, AddressModeImmediate))
	case AddressModeRelative:
		return Address(c.MemGet(addr, AddressModeImmediate)) + c.regRelBase
	}
	panic(mode)
}

func (c *vm) opIllegal() error {
	return ErrIllegal
}

func (c *vm) opAdd() error {
	a := c.MemGet(c.regPC+1, c.modes[0])
	b := c.MemGet(c.regPC+2, c.modes[1])
	addr := c.effectiveAddr(c.regPC+3, c.modes[2])
	v := a + b
	c.MemSet(addr, v)
	c.regPC += 4
	return nil
}

func (c *vm) opMultiply() error {
	a := c.MemGet(c.regPC+1, c.modes[0])
	b := c.MemGet(c.regPC+2, c.modes[1])
	addr := c.effectiveAddr(c.regPC+3, c.modes[2])
	v := a * b
	c.MemSet(addr, v)
	c.regPC += 4
	return nil
}

func (c *vm) opInput() error {
	addr := c.effectiveAddr(c.regPC+1, c.modes[0])
	v, err := c.stdin.ReadInt()
	if err != nil {
		return err
	}
	c.MemSet(addr, v)
	c.regPC += 2
	return nil
}

func (c *vm) opOutput() error {
	v := c.MemGet(c.regPC+1, c.modes[0])
	c.exitCode = v
	if err := c.stdout.WriteInt(v); err != nil {
		return err
	}
	c.regPC += 2
	return nil
}

func (c *vm) opJumpIfTrue() error {
	v := c.MemGet(c.regPC+1, c.modes[0])
	addr := Address(c.MemGet(c.regPC+2, c.modes[1]))
	if v != 0 {
		c.regPC = addr
	} else {
		c.regPC += 3
	}
	return nil
}

func (c *vm) opJumpIfNotTrue() error {
	v := c.MemGet(c.regPC+1, c.modes[0])
	addr := Address(c.MemGet(c.regPC+2, c.modes[1]))
	if v == 0 {
		c.regPC = addr
	} else {
		c.regPC += 3
	}
	return nil
}

func (c *vm) opLessThan() error {
	a := c.MemGet(c.regPC+1, c.modes[0])
	b := c.MemGet(c.regPC+2, c.modes[1])
	addr := (c.effectiveAddr(c.regPC+3, c.modes[2]))
	v := 0
	if a < b {
		v = 1
	}
	c.MemSet(addr, v)
	c.regPC += 4
	return nil
}

func (c *vm) opEqual() error {
	a := c.MemGet(c.regPC+1, c.modes[0])
	b := c.MemGet(c.regPC+2, c.modes[1])
	addr := (c.effectiveAddr(c.regPC+3, c.modes[2]))
	v := 0
	if a == b {
		v = 1
	}
	c.MemSet(addr, v)
	c.regPC += 4
	return nil
}

func (c *vm) opRelBase() error {
	c.regRelBase += Address(c.MemGet(c.regPC+1, c.modes[0]))
	c.regPC += 2
	return nil
}

func (c *vm) opHalt() error {
	return ErrHalted
}
