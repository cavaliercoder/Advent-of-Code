package intcode

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	ErrByteOverflow = errors.New("byte overflow")
)

type IntReader interface {
	ReadInt() (v int, err error)
}

type IntWriter interface {
	WriteInt(v int) (err error)
}

type IntReadWriter interface {
	IntReader
	IntWriter
}

type IntBuffer interface {
	IntReadWriter

	Len() int
}

type intReader struct {
	r io.Reader
}

func NewIntReader(r io.Reader) IntReader { return &intReader{r: r} }

func (r *intReader) ReadInt() (v int, err error) {
	var p [1]byte
	_, err = r.r.Read(p[:])
	if err != nil {
		return
	}
	v = int(p[0])
	return
}

type intWriter struct {
	w io.Writer
}

func NewIntWriter(w io.Writer) IntWriter { return &intWriter{w: w} }

func (w *intWriter) WriteInt(v int) (err error) {
	if v < 0 || v > 255 {
		fmt.Fprintf(w.w, "%d\n", v)
		return
	}
	_, err = w.w.Write([]byte{byte(v)})
	return
}

type reader struct {
	r IntReader
}

func NewReader(r IntReader) io.Reader { return &reader{r: r} }

func (r *reader) Read(p []byte) (n int, err error) {
	var v int
	for i := 0; i < len(p); i++ {
		v, err = r.r.ReadInt()
		if err != nil {
			return
		}
		if v < 0 || v > 255 {
			return n, ErrByteOverflow
		}
		p[i] = byte(v)
		n++
	}
	return
}

type writer struct {
	w IntWriter
}

func NewWriter(w IntWriter) io.Writer { return &writer{w: w} }

func (w *writer) Write(p []byte) (n int, err error) {
	for i := 0; i < len(p); i++ {
		err = w.w.WriteInt(int(p[i]))
		if err != nil {
			return
		}
		n++
	}
	return
}

type intBuffer struct {
	buf []int
}

// NewIntBuffer initializes and returns a new IntBuffer.
func NewIntBuffer(buf []int) IntBuffer {
	if buf == nil {
		buf = make([]int, 0, 4096)
	}
	return &intBuffer{
		buf: buf,
	}
}

func (c *intBuffer) ReadInt() (v int, err error) {
	if len(c.buf) == 0 {
		return 0, io.EOF
	}
	v = c.buf[0]
	c.buf = c.buf[1:]
	return
}

func (c *intBuffer) WriteInt(v int) (err error) {
	if len(c.buf) >= 4096 {
		return io.ErrShortWrite
	}
	c.buf = append(c.buf, v)
	return
}

func (c *intBuffer) Len() int {
	return len(c.buf)
}

func Run(data Data, args ...int) (v int, err error) {
	vm := NewWithIO(data, NewIntBuffer(args), NewIntBuffer(nil))
	return vm.Run()
}

func RunASCII(data Data, args ...string) (v int, err error) {
	var stdin io.Reader = os.Stdin
	if len(args) > 0 {
		b := &bytes.Buffer{}
		for _, arg := range args {
			fmt.Fprintf(b, "%s\n", arg)
		}
		stdin = b
	}
	vm := NewWithIO(data, NewIntReader(stdin), NewIntWriter(os.Stdout))
	return vm.Run()
}
