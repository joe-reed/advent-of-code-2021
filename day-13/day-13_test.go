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

func TestSolvePuzzle2(t *testing.T) {
	fmt.Println("Test:")
	testFile, _ := ioutil.ReadFile("./test-input-1.txt")
	puzzle2(string(testFile))

	fmt.Println("Puzzle:")
	file, _ := ioutil.ReadFile("./input.txt")
	puzzle2(string(file))
}
