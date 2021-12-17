package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestPuzzle1(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			"target area: x=20..30, y=-10..-5",
			45,
		},
	}

	for _, test := range tests {
		a := puzzle1(test.input)
		if a != test.expected {
			t.Errorf("expected: %d, actual: %d", test.expected, a)
		}
	}
}

func TestSolvePuzzle1(t *testing.T) {
	file, _ := ioutil.ReadFile("./input.txt")
	fmt.Println("Puzzle 1:", puzzle1(string(file)))
}

func TestPuzzle2(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			"target area: x=20..30, y=-10..-5",
			112,
		},
	}

	for _, test := range tests {
		a := puzzle2(test.input)
		if a != test.expected {
			t.Errorf("expected: %d, actual: %d", test.expected, a)
		}
	}
}

func TestSolvePuzzle2(t *testing.T) {
	file, _ := ioutil.ReadFile("./input.txt")
	fmt.Println("Puzzle 2:", puzzle2(string(file)))
}
