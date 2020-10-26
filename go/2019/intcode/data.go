package intcode

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
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
