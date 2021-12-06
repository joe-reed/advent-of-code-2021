package main

import (
	"strings"
	. "utils"
)

func puzzle1(input string) int {
	return solve(input, 80)
}

func puzzle2(input string) int {
	return solve(input, 256)
}

func solve(input string, days int) (result int) {
	fish := MapToInts(strings.Split(input, ","))

	counts := make(map[int]int)

	for _, f := range fish {
		counts[f]++
	}

	for i := 0; i < days; i++ {
		newCounts := make(map[int]int)

		for j := 8; j >= 0; j-- {
			if j == 0 {
				newCounts[8] += counts[0]
				newCounts[6] += counts[0]
				continue
			}

			newCounts[j-1] = counts[j]
		}

		counts = newCounts
	}

	for i := 0; i <= 8; i++ {
		result += counts[i]
	}

	return
}
