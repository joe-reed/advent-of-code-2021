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
			10,
		},
		{
			"./test-input-2.txt",
			19,
		},
		{
			"./test-input-3.txt",
			226,
		},
	}

	for _, test := range tests {
		a := puzzle1(LoadFile(test.input))
		if a != test.expected {
			t.Errorf("expected: %d, actual: %d", test.expected, a)
		}
	}
}

func TestSolvePuzzle1(t *testing.T) {
	fmt.Println("Puzzle 1:", puzzle1(LoadFile("./input.txt")))
}

func TestPuzzle2(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			"./test-input-1.txt",
			36,
		},
		{
			"./test-input-2.txt",
			103,
		},
		{
			"./test-input-3.txt",
			3509,
		},
	}

	for _, test := range tests {
		a := puzzle2(LoadFile(test.input))
		if a != test.expected {
			t.Errorf("expected: %d, actual: %d", test.expected, a)
		}
	}
}

func TestSolvePuzzle2(t *testing.T) {
	fmt.Println("Puzzle 2:", puzzle2(LoadFile("./input.txt")))
}
