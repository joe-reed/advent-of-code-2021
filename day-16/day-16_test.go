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
			"8A004A801A8002F478",
			16,
		},
		{
			"620080001611562C8802118E34",
			12,
		},
		{
			"C0015000016115A2E0802F182340",
			23,
		},
		{
			"A0016C880162017C3686B18A3D4780",
			31,
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
			"C200B40A82",
			3,
		},
		{
			"04005AC33890",
			54,
		},
		{
			"880086C3E88112",
			7,
		},
		{
			"CE00C43D881120",
			9,
		},
		{
			"D8005AC2A8F0",
			1,
		},
		{
			"F600BC2D8F",
			0,
		},
		{
			"9C005AC2F8F0",
			0,
		},
		{
			"9C0141080250320F1802104A08",
			1,
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
