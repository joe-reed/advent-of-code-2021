package main

import (
	"math"
	"sort"
	"strings"
	. "utils"
)

func puzzle1(input []string) (result int) {
	for _, entry := range input {
		_, output := parseEntry(entry)

		for _, v := range output {
			if IntInArray(len(v), []int{2, 4, 3, 7}) {
				result++
			}
		}
	}
	return
}

func puzzle2(input []string) (result int) {
	for _, entry := range input {
		patterns, output := parseEntry(entry)

		one, seven, four, eight := patterns[0], patterns[1], patterns[2], patterns[9]
		nine := find(patterns, func(p string) bool { return contains(p, seven) && contains(p, four) && p != eight })
		three := find(patterns, func(p string) bool { return contains(p, seven) && contains(nine, p) && p != seven })
		five := find(patterns, func(p string) bool { return !contains(p, one) && contains(nine, p) })
		two := find(patterns, func(p string) bool { return len(p) == 5 && p != three && p != five })
		six := find(patterns, func(p string) bool { return len(p) == 6 && contains(p, five) && p != nine })
		zero := find(patterns, func(p string) bool { return len(p) == 6 && p != six && p != nine })

		numbers := map[string]int{
			SortString(zero):  0,
			SortString(one):   1,
			SortString(two):   2,
			SortString(three): 3,
			SortString(four):  4,
			SortString(five):  5,
			SortString(six):   6,
			SortString(seven): 7,
			SortString(eight): 8,
			SortString(nine):  9,
		}

		for i, v := range output {
			result += numbers[SortString(v)] * int(math.Pow(10, 3-float64(i)))
		}
	}
	return
}

func parseEntry(entry string) ([]string, []string) {
	split := strings.Split(entry, " | ")
	patterns := strings.Fields(split[0])
	sort.Slice(patterns, func(i, j int) bool {
		return len(patterns[i]) < len(patterns[j])
	})
	return patterns, strings.Fields(split[1])
}

func find(patterns []string, condition func(p string) bool) string {
	for _, p := range patterns {
		if condition(p) {
			return p
		}
	}
	return ""
}

func contains(a, b string) bool {
	for _, x := range b {
		if !strings.Contains(a, string(x)) {
			return false
		}
	}
	return true
}
