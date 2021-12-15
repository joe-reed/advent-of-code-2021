package main

import (
	"math"
	"strings"
	. "utils"
)

func puzzle1(input string) int {
	return lengthOfShortestPath(parseInput(input))
}

func puzzle2(input string) int {
	cave := parseInput(input)

	newCave := [][]int{}
	for i := 0; i < 5; i++ {
		for y := range cave {
			newRow := []int{}
			for j := 0; j < 5; j++ {
				for x := range cave[y] {
					value := cave[y][x] + i + j
					for value > 9 {
						value -= 9
					}
					newRow = append(newRow, value)
				}
			}
			newCave = append(newCave, newRow)
		}
	}

	return lengthOfShortestPath(newCave)
}

func lengthOfShortestPath(cave [][]int) int {
	distances := make(map[Point]int)
	for i := range cave {
		for j := range cave[i] {
			distances[Point{j, i}] = math.MaxInt
		}
	}

	distances[Point{0, 0}] = 0

	queue := []Point{{0, 0}}
	end := Point{len(cave[0]) - 1, len(cave) - 1}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		left := math.MaxInt
		up := math.MaxInt
		right := math.MaxInt
		down := math.MaxInt

		if current.x > 0 {
			leftPoint := Point{current.x - 1, current.y}
			left = cave[leftPoint.y][leftPoint.x] + distances[current]
			if left < distances[leftPoint] {
				distances[leftPoint] = left
				queue = append(queue, leftPoint)
			}
		}
		if current.y > 0 {
			upPoint := Point{current.x, current.y - 1}
			up = cave[upPoint.y][upPoint.x] + distances[current]
			if up < distances[upPoint] {
				distances[upPoint] = up
				queue = append(queue, upPoint)
			}
		}
		if current.x < len(cave[0])-1 {
			rightPoint := Point{current.x + 1, current.y}
			right = cave[rightPoint.y][rightPoint.x] + distances[current]
			if right < distances[rightPoint] {
				distances[rightPoint] = right
				queue = append(queue, rightPoint)
			}
		}
		if current.y < len(cave)-1 {
			downPoint := Point{current.x, current.y + 1}
			down = cave[downPoint.y][downPoint.x] + distances[current]
			if down < distances[downPoint] {
				distances[downPoint] = down
				queue = append(queue, downPoint)
			}
		}
	}

	return distances[end]
}

func parseInput(input string) (result [][]int) {
	for _, line := range strings.Split(input, "\n") {
		resultLine := []int{}
		for _, v := range strings.Split(line, "") {
			resultLine = append(resultLine, ConvertToInt(v))
		}
		result = append(result, resultLine)
	}
	return
}

type Point struct {
	x, y int
}
