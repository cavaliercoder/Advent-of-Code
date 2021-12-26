/*
Bootstrap Go files for a new year of Advent of Code.
*/
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fatalf("usage: %s YEAR\n", filepath.Base(os.Args[0]))
	}
	year := os.Args[1]
	if _, err := strconv.Atoi(year); err != nil {
		fatalf("Invalid year: %s", os.Args[1])
	}
	if err := os.MkdirAll(year, 0755); err != nil {
		fatal(err)
	}
	for i := 1; i <= 25; i++ {
		day := fmt.Sprintf("day%02d", i)
		dir := filepath.Join(year, day)
		if err := os.MkdirAll(dir, 0755); err != nil {
			fatal(err)
		}
		if err := createGoFile(filepath.Join(dir, day+".go"), day); err != nil {
			fatal(err)
		}
		if err := createGoFile(filepath.Join(dir, day+"_test.go"), day); err != nil {
			fatal(err)
		}
	}
}

func createGoFile(name, packageName string) error {
	_, err := os.Stat(name)
	if err == nil {
		return nil
	}
	if !os.IsNotExist(err) {
		return err
	}
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = fmt.Fprintf(f, "package %s\n\n", packageName)
	return err
}

func fatal(a ...interface{}) {
	fmt.Fprint(os.Stderr, a...)
	os.Exit(1)
}

func fatalf(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(1)
}
