package main

import (
	"regexp"
	"strings"
	. "utils"
)

func puzzle1(input string) int {
	paper, folds := parseInput(input)

	return len(paper.fold(folds[0]))
}

func puzzle2(input string) (result string) {
	paper, folds := parseInput(input)

	for _, f := range folds {
		paper = paper.fold(f)
	}

	maxX, maxY := 0, 0
	for _, dot := range paper {
		if dot.x > maxX {
			maxX = dot.x
		}
		if dot.y > maxY {
			maxY = dot.y
		}
	}

	for j := 0; j <= maxY; j++ {
		for i := 0; i <= maxX; i++ {
			foundDot := false
			for _, dot := range paper {
				if dot.x == i && dot.y == j {
					result += "#"
					foundDot = true
					break
				}
			}
			if !foundDot {
				result += "."
			}
		}
		result += "\n"
	}

	return
}

func parseInput(input string) (paper Paper, folds []Fold) {
	paperString, instructionString := strings.Split(input, "\n\n")[0], strings.Split(input, "\n\n")[1]

	for _, p := range strings.Split(paperString, "\n") {
		x, y := strings.Split(p, ",")[0], strings.Split(p, ",")[1]
		paper = append(paper, Dot{ConvertToInt(x), ConvertToInt(y)})
	}

	for _, i := range strings.Split(instructionString, "\n") {
		re := regexp.MustCompile(`fold along (.{1})=(\d+)`)
		matches := re.FindStringSubmatch(i)
		folds = append(folds, Fold{matches[1], ConvertToInt(matches[2])})
	}

	return
}

type Paper []Dot

type Dot struct {
	x, y int
}

type Fold struct {
	direction string
	location  int
}

func (p Paper) fold(f Fold) (folded Paper) {
	if f.direction == "x" {
		for _, dot := range p {
			if dot.x > f.location {
				folded = append(folded, Dot{dot.x - 2*(dot.x-f.location), dot.y})
			} else {
				folded = append(folded, dot)
			}
		}
	} else {
		for _, dot := range p {
			if dot.y > f.location {
				folded = append(folded, Dot{dot.x, dot.y - 2*(dot.y-f.location)})
			} else {
				folded = append(folded, dot)
			}
		}
	}

	return folded.dedupe()
}

func (p Paper) dedupe() (deduped Paper) {
	for _, d := range p {
		if !deduped.contains(d) {
			deduped = append(deduped, d)
		}
	}
	return
}

func (p Paper) contains(dot Dot) bool {
	for _, d := range p {
		if d == dot {
			return true
		}
	}
	return false
}
