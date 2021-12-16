package day04

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"testing"

	"aoc2021"
)

func parseCalls(r *bufio.Reader) ([]int, error) {
	calls := make([]int, 0, 64)
	line, err := r.ReadString('\n')
	if err != nil {
		return nil, err
	}
	line = strings.TrimSuffix(line, "\n")
	parts := strings.Split(line, ",")
	for _, s := range parts {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("bad call: %s: %v", s, err)
		}
		calls = append(calls, n)
	}
	return calls, nil
}

func parseBoard(r *bufio.Reader) (*Board, error) {
	board := &Board{}
	for i := 0; i < 25; {
		line, err := r.ReadString('\n')
		if err != nil {
			return nil, err
		}
		if line == "\n" {
			continue
		}
		if len(line) != 15 {
			return nil, fmt.Errorf("bad line: %s", line)
		}
		for j := 0; j < 5; j++ {
			s := line[j*3 : (j*3)+2]
			if s[0] == ' ' {
				s = s[1:]
			}
			n, err := strconv.Atoi(s)
			if err != nil {
				return nil, fmt.Errorf("bad cell: %s, in line: %s", s, line)
			}
			board.data[i] = n
			i++
		}
	}
	return board, nil
}

func mustOpenFixture(name string) (calls []int, boards []*Board) {
	f, err := aoc2021.OpenFixture(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	calls, err = parseCalls(r)
	if err != nil {
		panic(err)
	}
	boards = make([]*Board, 0, 64)
	for {
		board, err := parseBoard(r)
		if err == io.EOF {
			return
		}
		if err != nil {
			panic(err)
		}
		boards = append(boards, board)
	}
}

func TestPart1(t *testing.T) {
	calls, boards := mustOpenFixture("day04")
	for _, call := range calls {
		for _, board := range boards {
			if board.Call(call) {
				aoc2021.AssertInt(t, 64084, call*board.Score(), "bad score")
				return
			}
		}
	}
	t.Errorf("no bingo")
}

func TestPart2(t *testing.T) {
	var lastBoard *Board
	var lastCall int
	calls, boards := mustOpenFixture("day04")
	for _, call := range calls {
		for i, board := range boards {
			if board == nil {
				continue
			}
			if board.Call(call) {
				lastBoard = board
				lastCall = call
				boards[i] = nil
			}
		}
	}
	aoc2021.AssertInt(t, 12833, lastCall*lastBoard.Score(), "bad score")
}
