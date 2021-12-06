package main

import (
	"fmt"
	"testing"
	. "utils"
)

func TestPuzzle1(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			"./test-input-1.txt",
			5934,
		},
	}

	for _, test := range tests {
		a := puzzle1(LoadFile(test.input)[0])
		if a != test.expected {
			t.Errorf("expected: %d, actual: %d", test.expected, a)
		}
	}
}

func TestSolvePuzzle1(t *testing.T) {
	fmt.Println("Puzzle 1:", puzzle1(LoadFile("./input.txt")[0]))
}

func TestPuzzle2(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			"./test-input-1.txt",
			26984457539,
		},
	}

	for _, test := range tests {
		a := puzzle2(LoadFile(test.input)[0])
		if a != test.expected {
			t.Errorf("expected: %d, actual: %d", test.expected, a)
		}
	}
}

func TestSolvePuzzle2(t *testing.T) {
	fmt.Println("Puzzle 2:", puzzle2(LoadFile("./input.txt")[0]))
}
