package intcode

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var (
	ErrIllegal  = errors.New("illegal")
	ErrHalted   = errors.New("halted")
	ErrNoInput  = errors.New("no input")
	ErrNoOutput = errors.New("no output")
)

// Data is a slice of program and data memory for the VM.
type Data []int

// DecodeData reads Data from a file and decodes it from a comma-seperated
// string format.
func DecodeData(r io.Reader) (data Data, err error) {
	var b []byte
	var v int
	b, err = ioutil.ReadAll(r)
	if err != nil {
		return
	}
	parts := bytes.Split(b, []byte{','})
	data = make([]int, len(parts))
	for i := 0; i < len(parts); i++ {
		v, err = strconv.Atoi(strings.TrimRight(string(parts[i]), "\n"))
		if err != nil {
			return
		}
		data[i] = v
	}
	return
}

// OpenData reads all data from a file.
func OpenData(name string) (data Data, err error) {
	var f *os.File
	f, err = os.Open(name)
	if err != nil {
		return
	}
	defer f.Close()
	return DecodeData(f)
}

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

// A VirtualMachine runs IntCode programs.
type VirtualMachine interface {
	// Step will fetch, decode and execute a single instruction
	Step() error

	// Run until the program halts
	Run() error

	// MemGet reads a value from VM memory
	MemGet(addr Address, mode AddressMode) int

	// MemSet writes a value to VM memory
	MemSet(addr Address, v int)

	// IOPush sends a value to the VM input buffer
	IOPush(v int, block bool) (err error)

	// IOPop reads a value from the VM output buffer
	IOPop(block bool) (v int, err error)
}

type virtualMachine struct {
	data        Data
	handlers    []opcodeHandler
	stdin       []int
	stdout      []int
	opcode      Opcode
	modes       [4]AddressMode
	regPC       Address
	regRelBase  Address
	traceWriter io.Writer
}

// New returns a new VirtualMachine.
func New(data Data) VirtualMachine {
	v := &virtualMachine{
		data:   data,
		stdin:  make([]int, 0, 64),
		stdout: make([]int, 0, 64),
	}
	v.handlers = []opcodeHandler{
		v.opIllegal,
		v.opAdd,
		v.opMultiply,
		v.opInput,
		v.opOutput,
		v.opJumpIfTrue,
		v.opJumpIfNotTrue,
		v.opLessThan,
		v.opEqual,
		v.opRelBase,
	}
	if os.Getenv("LOGLEVEL") == "DEBUG" {
		v.traceWriter = os.Stderr
	}
	return v
}

func (c *virtualMachine) tracef(format string, a ...interface{}) {
	if c.traceWriter == nil {
		return
	}
	fmt.Fprintf(c.traceWriter, format, a...)
}

func (c *virtualMachine) traceStep() {
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

func (c *virtualMachine) Step() error {
	// fetch and decode
	v := Instruction(c.MemGet(c.regPC, AddressModeImmediate))
	c.opcode = v.Opcode()
	c.modes = v.Modes()
	c.traceStep()

	// dispatch
	if c.opcode == OpcodeHalt {
		return ErrHalted
	}
	return c.handlers[c.opcode]()
}

func (c *virtualMachine) Run() error {
	for {
		if err := c.Step(); err != nil {
			if err == ErrHalted {
				return nil
			}
			return err
		}
	}
}

func (c *virtualMachine) MemGet(addr Address, mode AddressMode) (v int) {
	eaddr := c.effectiveAddr(addr, mode)
	if int(eaddr) >= len(c.data) {
		v = 0
	} else {
		v = c.data[eaddr]
	}
	c.tracef("  memget(%04d, %d) -> %v\n", addr, mode, v)
	return
}

func (c *virtualMachine) MemSet(addr Address, v int) {
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

func (c *virtualMachine) IOPush(v int, block bool) (err error) {
	c.tracef("  iopush(%d)\n", v)
	c.stdin = append(c.stdin, v)
	for block && len(c.stdin) > 0 {
		err = c.Step()
		if err != nil {
			return
		}
	}
	return nil
}

func (c *virtualMachine) IOPop(block bool) (v int, err error) {
	for block && len(c.stdout) == 0 {
		err = c.Step()
		if err != nil {
			c.tracef("  iopop() -> ERROR\n")
			return
		}
	}
	if len(c.stdout) == 0 {
		c.tracef("  iopop() -> ERROR\n")
		return 0, ErrNoOutput
	}
	v = c.stdout[0]
	c.stdout = c.stdout[1:]
	c.tracef("  iopop() -> %d\n", v)
	return
}

func (c *virtualMachine) effectiveAddr(addr Address, mode AddressMode) Address {
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

func (c *virtualMachine) opIllegal() error {
	return ErrIllegal
}

func (c *virtualMachine) opAdd() error {
	a := c.MemGet(c.regPC+1, c.modes[0])
	b := c.MemGet(c.regPC+2, c.modes[1])
	addr := c.effectiveAddr(c.regPC+3, c.modes[2])
	v := a + b
	c.MemSet(addr, v)
	c.regPC += 4
	return nil
}

func (c *virtualMachine) opMultiply() error {
	a := c.MemGet(c.regPC+1, c.modes[0])
	b := c.MemGet(c.regPC+2, c.modes[1])
	addr := c.effectiveAddr(c.regPC+3, c.modes[2])
	v := a * b
	c.MemSet(addr, v)
	c.regPC += 4
	return nil
}

func (c *virtualMachine) opInput() error {
	addr := c.effectiveAddr(c.regPC+1, c.modes[0])
	if len(c.stdin) == 0 {
		return ErrNoInput
	}
	v := c.stdin[0]
	c.stdin = c.stdin[1:]
	c.MemSet(addr, v)
	c.regPC += 2
	return nil
}

func (c *virtualMachine) opOutput() error {
	v := c.MemGet(c.regPC+1, c.modes[0])
	c.stdout = append(c.stdout, v)
	c.regPC += 2
	return nil
}

func (c *virtualMachine) opJumpIfTrue() error {
	v := c.MemGet(c.regPC+1, c.modes[0])
	addr := Address(c.MemGet(c.regPC+2, c.modes[1]))
	if v != 0 {
		c.regPC = addr
	} else {
		c.regPC += 3
	}
	return nil
}

func (c *virtualMachine) opJumpIfNotTrue() error {
	v := c.MemGet(c.regPC+1, c.modes[0])
	addr := Address(c.MemGet(c.regPC+2, c.modes[1]))
	if v == 0 {
		c.regPC = addr
	} else {
		c.regPC += 3
	}
	return nil
}

func (c *virtualMachine) opLessThan() error {
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

func (c *virtualMachine) opEqual() error {
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

func (c *virtualMachine) opRelBase() error {
	c.regRelBase += Address(c.MemGet(c.regPC+1, c.modes[0]))
	c.regPC += 2
	return nil
}
