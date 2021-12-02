package main

import (
	"strings"
	. "utils"
)

func puzzle1(input []string) int {
	depth := 0
	position := 0
	for _, v := range input {
		instruction := strings.Split(v, " ")
		x := ConvertToInt(instruction[1])

		switch instruction[0] {
		case "up":
			depth -= x
			continue
		case "down":
			depth += x
			continue
		case "forward":
			position += x
			continue
		}
	}
	return depth * position
}

func puzzle2(input []string) int {
	aim := 0
	depth := 0
	position := 0
	for _, v := range input {
		instruction := strings.Split(v, " ")
		x := ConvertToInt(instruction[1])

		switch instruction[0] {
		case "up":
			aim -= x
			continue
		case "down":
			aim += x
			continue
		case "forward":
			position += x
			depth += aim * x
			continue
		}
	}
	return depth * position
}
