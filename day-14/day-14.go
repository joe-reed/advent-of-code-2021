package main

import (
	"strings"
)

func puzzle1(input string) int {
	return solve(input, 10)
}

func puzzle2(input string) int {
	return solve(input, 40)
}

func solve(input string, steps int) int {
	initialPairs, rules := parseInput(input)

	pairCounts := make(map[Pair]int)
	for _, p := range initialPairs {
		pairCounts[p]++
	}

	for i := 0; i < steps; i++ {
		newPairCounts := make(map[Pair]int)
		for p, c := range pairCounts {
			for _, r := range rules {
				if p != r.pair {
					continue
				}
				newPairCounts[Pair{p.a, r.insert}] += c
				newPairCounts[Pair{r.insert, p.b}] += c
			}
		}
		pairCounts = newPairCounts
	}

	counts := make(map[string]int)
	for p, c := range pairCounts {
		counts[p.a] += c
	}

	finalElement := initialPairs[len(initialPairs)-1].b
	counts[finalElement] += 1

	min := counts[finalElement]
	max := 0
	for _, count := range counts {
		if count < min {
			min = count
		}

		if count > max {
			max = count
		}
	}

	return max - min
}

func parseInput(input string) (initialPairs []Pair, rules []Rule) {
	template, rulesString := strings.Split(input, "\n\n")[0], strings.Split(input, "\n\n")[1]

	templateElements := strings.Split(template, "")
	for i := range templateElements {
		if i == len(templateElements)-1 {
			break
		}

		initialPairs = append(initialPairs, Pair{templateElements[i], templateElements[i+1]})
	}

	for _, ruleString := range strings.Split(rulesString, "\n") {
		split := strings.Split(ruleString, " -> ")
		pairString, insert := split[0], split[1]
		rules = append(rules, Rule{Pair{string(pairString[0]), string(pairString[1])}, insert})
	}

	return
}

type Pair struct {
	a, b string
}

type Rule struct {
	pair   Pair
	insert string
}
