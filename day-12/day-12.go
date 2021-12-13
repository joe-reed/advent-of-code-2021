package main

import (
	"strings"
)

func puzzle1(input []string) int {
	return countPaths("start", getPaths(input), make(map[string]int), func(start string, counts map[string]int) bool {
		return counts[start] == 1
	})
}

func puzzle2(input []string) (count int) {
	return countPaths("start", getPaths(input), make(map[string]int), func(start string, counts map[string]int) bool {
		haveUsedDoubleVisit := false
		for _, count := range counts {
			if count == 2 {
				haveUsedDoubleVisit = true
			}
		}
		return counts[start] == 2 || (counts[start] == 1 && haveUsedDoubleVisit)
	})
}

func countPaths(start string, paths []Path, counts map[string]int, stopCondition func(start string, counts map[string]int) bool) (count int) {
	if start == "end" {
		return 1
	}

	if stopCondition(start, counts) {
		return 0
	}

	newCounts := make(map[string]int)
	for k, v := range counts {
		newCounts[k] = v
	}

	if strings.ToLower(start) == start {
		newCounts[start]++
	}

	for _, path := range paths {
		if path.a == start && path.b != "start" {
			count += countPaths(path.b, paths, newCounts, stopCondition)
		}
		if path.b == start && path.a != "start" {
			count += countPaths(path.a, paths, newCounts, stopCondition)
		}
	}
	return
}

func getPaths(input []string) (paths []Path) {
	for _, line := range input {
		split := strings.Split(line, "-")
		paths = append(paths, Path{split[0], split[1]})
	}
	return
}

type Path struct {
	a, b string
}
