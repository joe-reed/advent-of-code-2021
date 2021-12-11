package main

import (
	"strings"
	. "utils"
)

func puzzle1(input []string) (flashCount int) {
	grid := parseInput(input)
	var flashers []Point

	for i := 0; i < 100; i++ {
		grid, flashers = handleStep(grid)
		flashCount += len(flashers)
	}

	return
}

func puzzle2(input []string) int {
	grid := parseInput(input)
	var flashers []Point

	i := 0
	for true {
		i++
		grid, flashers = handleStep(grid)
		if len(flashers) == len(grid)*len(grid[0]) {
			return i
		}
	}

	return 0
}

func parseInput(input []string) [][]int {
	grid := make([][]int, len(input))
	for i, line := range input {
		grid[i] = MapToInts(strings.Split(line, ""))
	}
	return grid
}

func handleStep(grid [][]int) ([][]int, []Point) {
	newGrid := make([][]int, len(grid))
	copy(newGrid, grid)

	for i, line := range newGrid {
		for j := range line {
			newGrid[i][j]++
		}
	}

	var flashers []Point

	for i, line := range newGrid {
		for j := range line {
			if shouldFlash(i, j, newGrid, flashers) {
				handleFlash(&newGrid, i, j, &flashers)
			}
		}
	}

	for _, flasher := range flashers {
		newGrid[flasher.y][flasher.x] = 0
	}

	return newGrid, flashers
}

func handleFlash(grid *[][]int, i, j int, flashers *[]Point) {
	*flashers = append(*flashers, Point{j, i})

	if i != 0 {
		incrementOctopus(i-1, j, grid, flashers)

		if j != 0 {
			incrementOctopus(i-1, j-1, grid, flashers)
		}

		if j != len((*grid)[0])-1 {
			incrementOctopus(i-1, j+1, grid, flashers)
		}
	}

	if i != len(*grid)-1 {
		incrementOctopus(i+1, j, grid, flashers)

		if j != 0 {
			incrementOctopus(i+1, j-1, grid, flashers)
		}
		if j != len((*grid)[0])-1 {
			incrementOctopus(i+1, j+1, grid, flashers)
		}
	}

	if j != len((*grid)[0])-1 {
		incrementOctopus(i, j+1, grid, flashers)
	}

	if j != 0 {
		incrementOctopus(i, j-1, grid, flashers)
	}
}

func incrementOctopus(i, j int, grid *[][]int, flashers *[]Point) {
	(*grid)[i][j]++
	if shouldFlash(i, j, *grid, *flashers) {
		handleFlash(grid, i, j, flashers)
	}
}

func shouldFlash(i, j int, grid [][]int, flashers []Point) bool {
	return grid[i][j] >= 10 && !containsPoint(flashers, Point{j, i})
}

func containsPoint(points []Point, p Point) bool {
	for _, point := range points {
		if point == p {
			return true
		}
	}
	return false
}

type Point struct {
	x int
	y int
}
