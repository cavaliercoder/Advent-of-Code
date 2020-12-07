/*
This example returns the number of valid passphrases for a given input. A
passphrase is considered valid if it contains only words where no two words are
anagrams or equal.

Passphrases are read - one per line - from standard input. The number of valid
passphrases is printed to standard output.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	count := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if isValidPassphrase(scanner.Text()) {
			count++
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%d\n", count)
}

// RuneSlice implements sort.Interface for a slice of runes.
type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// isValidPassphrase returns true if the given passphrase contains only words
// where no two words are anagrams or equal.
//
// The passphrase is assumed to be any unicode string that includes space
// separated words.
//
// Time complexity is O(n log n) for each passphrase, given that each word in
// the phrase is sorted in logarithmic time.
func isValidPassphrase(s string) bool {
	words := strings.Split(s, " ")
	seen := make(map[string]int, len(words))
	for _, word := range words {
		runes := RuneSlice(word)
		sort.Sort(runes)
		word = string(runes)
		if _, ok := seen[word]; ok {
			return false
		}
		seen[word] = 1
	}
	return true
}
