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

		one := patterns[0]
		seven := patterns[1]
		four := patterns[2]
		eight := patterns[9]
		nine := getNine(patterns)
		three := getThree(patterns, nine)
		five := getFive(patterns, nine)
		two := getTwo(patterns, three, five)
		six := getSix(patterns, five, nine)
		zero := getZero(patterns, six, nine)

		numbers := make(map[string]int, 10)
		numbers[sortString(zero)] = 0
		numbers[sortString(one)] = 1
		numbers[sortString(two)] = 2
		numbers[sortString(three)] = 3
		numbers[sortString(four)] = 4
		numbers[sortString(five)] = 5
		numbers[sortString(six)] = 6
		numbers[sortString(seven)] = 7
		numbers[sortString(eight)] = 8
		numbers[sortString(nine)] = 9

		for i, v := range output {
			result += numbers[sortString(v)] * int(math.Pow(10, 3-float64(i)))
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

func getNine(patterns []string) string {
	return find(patterns, func(p string, i int) bool { return contains(p, patterns[1]) && contains(p, patterns[2]) && i != 9 })
}

func getThree(patterns []string, nine string) string {
	return find(patterns, func(p string, i int) bool { return contains(p, patterns[1]) && contains(nine, p) && i != 1 })
}

func getFive(patterns []string, nine string) string {
	return find(patterns, func(p string, i int) bool { return !contains(p, patterns[0]) && contains(nine, p) })
}

func getTwo(patterns []string, three string, five string) string {
	return find(patterns, func(p string, i int) bool { return len(p) == 5 && p != three && p != five })
}

func getSix(patterns []string, five string, nine string) string {
	return find(patterns, func(p string, i int) bool { return len(p) == 6 && contains(p, five) && p != nine })
}

func getZero(patterns []string, six string, nine string) string {
	return find(patterns, func(p string, i int) bool { return len(p) == 6 && p != six && p != nine })
}

func find(patterns []string, condition func(p string, i int) bool) string {
	for i, p := range patterns {
		if condition(p, i) {
			return p
		}
	}
	return ""
}

func contains(a, b string) bool {
	return len(remove(b, a)) == 0
}

func remove(a, b string) (result string) {
	for _, x := range a {
		if !strings.Contains(b, string(x)) {
			result += string(x)
		}
	}
	return
}

func sortString(s string) string {
	split := strings.Split(s, "")
	sort.Strings(split)
	return strings.Join(split, "")
}
