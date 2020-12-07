package aoc

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func Fixture(name string) string {
	return filepath.Join("..", "..", "..", "inputs", "2017", name+".txt")
}

func MustOpenFixture(name string) *os.File {
	f, err := os.Open(Fixture(name))
	if err != nil {
		panic(err)
	}
	return f
}

func MustReadFixture(name string) []byte {
	f := MustOpenFixture(name)
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return b
}
