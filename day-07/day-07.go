package main

import (
	"math"
	"sort"
	"strings"
	. "utils"
)

func puzzle1(input string) int {
	return solve(input, func(distance int) int { return distance })
}

func puzzle2(input string) int {
	return solve(input, func(distance int) int { return (distance * (distance + 1)) / 2 })
}

func solve(input string, calculateFuel func(distance int) int) int {
	crabs := MapToInts(strings.Split(input, ","))
	sort.Ints(crabs)

	var results []int
	for i := crabs[0]; i <= crabs[len(crabs)-1]; i++ {
		result := 0
		for _, crab := range crabs {
			result += calculateFuel(int(math.Abs(float64(crab - i))))
		}
		results = append(results, result)
	}

	return MinInt(results)
}
