package intcode

import (
	"bytes"
	"errors"
	"io"
	"strings"
)

var (
	ErrByteOverflow = errors.New("byte overflow")
)

type BytesIO interface {
	io.Reader
	io.Writer

	ReadLine() (s string, err error)
	WriteString(s string) (n int64, err error)
}

type bytesIO struct {
	vm VM
}

func NewBytesIO(vm VM) BytesIO {
	return &bytesIO{vm: vm}
}

func (c *bytesIO) Read(p []byte) (n int, err error) {
	var v int
	for n = 0; n < len(p); n++ {
		v, err = c.vm.IOPop(true)
		if err != nil {
			return
		}
		if v < 0 || v > 255 {
			return n, ErrByteOverflow
		}
		p[n] = byte(v)
	}
	return
}

func (c *bytesIO) ReadLine() (s string, err error) {
	b := &bytes.Buffer{}
	p := make([]byte, 1)
	for {
		_, err = c.Read(p)
		if err != nil {
			return
		}
		b.Write(p)
		if p[0] == '\n' {
			s = b.String()
			// fmt.Printf("readline() -> [%s]", s)
			return
		}
	}
}

func (c *bytesIO) Write(p []byte) (n int, err error) {
	for n = 0; n < len(p); n++ {
		err = c.vm.IOPush(int(p[n]), true)
		if err != nil {
			return
		}
	}
	return
}

func (c *bytesIO) WriteString(s string) (n int64, err error) {
	return io.Copy(c, strings.NewReader(s))
}
