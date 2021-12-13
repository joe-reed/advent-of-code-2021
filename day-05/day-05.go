package main

import (
	"math"
	"strings"
	. "utils"
)

func puzzle1(input []string) int {
	lines, grid := parseInput(input)

	lines = filterOutDiagonals(lines)

	return getDangerousCount(lines, grid)
}

func puzzle2(input []string) int {
	return getDangerousCount(parseInput(input))
}

func parseInput(input []string) ([]Line, [][]int) {
	var lines []Line
	for _, v := range input {
		pointStrings := strings.Split(v, " -> ")

		startStrings := strings.Split(pointStrings[0], ",")
		endStrings := strings.Split(pointStrings[1], ",")

		start := Point{ConvertToInt(startStrings[0]), ConvertToInt(startStrings[1])}
		end := Point{ConvertToInt(endStrings[0]), ConvertToInt(endStrings[1])}

		lines = append(lines, Line{start, end, nil})
	}

	grid := createGrid(lines)
	return lines, grid
}

func getDangerousCount(lines []Line, grid [][]int) int {
	for _, line := range lines {
		for {
			point := line.advance()

			grid[point.y][point.x]++

			if line.isAtEnd() {
				break
			}
		}
	}

	dangerousCount := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] >= 2 {
				dangerousCount++
			}
		}
	}

	return dangerousCount
}

func createGrid(lines []Line) [][]int {
	maxX, maxY := 0, 0
	for _, line := range lines {
		if line.start.x > maxX {
			maxX = line.start.x
		}
		if line.end.x > maxX {
			maxX = line.end.x
		}
		if line.start.y > maxY {
			maxY = line.start.y
		}
		if line.end.y > maxY {
			maxY = line.end.x
		}
	}

	grid := make([][]int, maxY+1)
	for i := range grid {
		grid[i] = make([]int, maxX+1)
	}

	return grid
}

func filterOutDiagonals(lines []Line) (result []Line) {
	for _, line := range lines {
		if line.start.x == line.end.x {
			result = append(result, line)
			continue
		}

		if line.start.y == line.end.y {
			result = append(result, line)
			continue
		}
	}

	return
}

type Line struct {
	start, end Point
	curr       *Point
}

type Point struct {
	x, y int
}

func (line *Line) advance() Point {
	if line.curr == nil {
		line.curr = &line.start
		return line.start
	}

	xDistance := line.end.x - line.start.x
	yDistance := line.end.y - line.start.y

	xStep := 0
	if xDistance != 0 {
		xStep = xDistance / int(math.Abs(float64(xDistance)))
	}

	yStep := 0
	if yDistance != 0 {
		yStep = yDistance / int(math.Abs(float64(yDistance)))
	}

	line.curr = &Point{line.curr.x + xStep, line.curr.y + yStep}

	return *line.curr
}

func (line *Line) isAtEnd() bool {
	return *line.curr == line.end
}
