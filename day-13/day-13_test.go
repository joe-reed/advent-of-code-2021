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
			"./test-input-1.txt",
			17,
		},
	}

	for _, test := range tests {
		file, _ := ioutil.ReadFile(test.input)
		a := puzzle1(string(file))
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
		expected string
	}{
		{
			"./test-input-1.txt",
			"#####\n#...#\n#...#\n#...#\n#####\n",
		},
	}

	for _, test := range tests {
		file, _ := ioutil.ReadFile(test.input)
		a := puzzle2(string(file))
		if a != test.expected {
			t.Errorf("\nexpected:\n%s\nactual:\n%s", test.expected, a)
		}
	}
}

func TestSolvePuzzle2(t *testing.T) {
	file, _ := ioutil.ReadFile("./input.txt")
	fmt.Println("Puzzle 2:")
	fmt.Print(puzzle2(string(file)))
}
