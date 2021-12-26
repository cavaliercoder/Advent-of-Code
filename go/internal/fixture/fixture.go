package fixture

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"aoc/internal/geo"
)

func Open(t *testing.T, year, day int) *os.File {
	path, err := filepath.Abs(filepath.Join(
		"..", "..", "..",
		"inputs",
		fmt.Sprintf("%d", year),
		fmt.Sprintf("day%02d.txt", day),
	))
	if err != nil {
		t.Fatal(err)
	}
	f, err := os.Open(path)
	if err != nil {
		t.Fatal()
	}
	return f
}

func Bytes(t *testing.T, year, day int) []byte {
	f := Open(t, year, day)
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func String(t *testing.T, year, day int) string {
	return string(Bytes(t, year, day))
}

func scan(t *testing.T, year, day int, fn func([]byte) error) {
	f := Open(t, year, day)
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		if err := fn(s.Bytes()); err != nil {
			t.Fatal(err)
		}
	}
	if err := s.Err(); err != nil {
		t.Fatal(err)
	}
}

func ScanBytes(t *testing.T, year, day int, fn func([]byte) error) {
	scan(t, year, day, func(b []byte) error {
		p := make([]byte, len(b))
		copy(p, b)
		return fn(p)
	})
}

func ScanStrings(t *testing.T, year, day int, fn func(string) error) {
	scan(t, year, day, func(b []byte) error {
		return fn(string(b))
	})
}

func Strings(t *testing.T, year, day int) []string {
	a := make([]string, 0, 64)
	ScanStrings(t, year, day, func(s string) error {
		a = append(a, s)
		return nil
	})
	return a
}

// Ints reads on integer per line.
func Ints(t *testing.T, year, day int) []int {
	a := make([]int, 0, 64)
	ScanStrings(t, year, day, func(s string) error {
		n, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		a = append(a, n)
		return nil
	})
	return a
}

func Grid(t *testing.T, year, day int) *geo.Grid {
	f := Open(t, year, day)
	defer f.Close()
	g, err := geo.ReadGrid(f)
	if err != nil {
		t.Fatal(err)
	}
	return g
}
