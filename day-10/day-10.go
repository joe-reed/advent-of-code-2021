package main

import (
	"sort"
	"strings"
)

func puzzle1(input []string) (result int) {
	points := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	for _, line := range input {
		reduced := reduce(line)

		c := findCorruption(reduced)

		if c != -1 {
			result += points[string(reduced[c])]
		}
	}

	return
}

func puzzle2(input []string) int {
	points := map[string]int{
		"(": 1,
		"[": 2,
		"{": 3,
		"<": 4,
	}

	var incomplete []string
	for _, line := range input {
		reduced := reduce(line)

		if findCorruption(reduced) == -1 {
			incomplete = append(incomplete, reduced)
		}
	}

	scores := make([]int, len(incomplete))
	for i, line := range incomplete {
		for j := range line {
			scores[i] *= 5
			scores[i] += points[string(line[len(line)-1-j])]
		}
	}

	sort.Ints(scores)

	return scores[len(scores)/2]
}

func reduce(line string) (r string) {
	r = line
	for strings.Contains(r, "()") || strings.Contains(r, "[]") || strings.Contains(r, "{}") || strings.Contains(r, "<>") {
		r = strings.ReplaceAll(r, "()", "")
		r = strings.ReplaceAll(r, "[]", "")
		r = strings.ReplaceAll(r, "{}", "")
		r = strings.ReplaceAll(r, "<>", "")
	}
	return
}

func findCorruption(reduced string) int {
	i := 0
	for i != len(reduced) && strings.Contains("([{<", string(reduced[i])) {
		i++
	}

	if i == len(reduced) {
		return -1
	}

	return i
}
