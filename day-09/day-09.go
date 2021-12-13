package main

import (
	"sort"
	"strings"
	. "utils"
)

func puzzle1(input []string) (result int) {
	floor := parseInput(input)

	for _, p := range getLowPoints(floor) {
		result += floor[p.y][p.x] + 1
	}

	return
}

func puzzle2(input []string) (result int) {
	floor := parseInput(input)

	lowPoints := getLowPoints(floor)

	var basinSizes []int

	for _, l := range lowPoints {
		basin := Basin{l}
		var previousBasin Basin

		for len(basin) != len(previousBasin) {
			var newBasin Basin
			newBasin = append(newBasin, basin...)

			for _, p := range basin.remove(previousBasin) {
				if p.x != 0 {
					newBasin.maybeAddPoint(Point{p.x - 1, p.y}, floor)
				}
				if p.y != 0 {
					newBasin.maybeAddPoint(Point{p.x, p.y - 1}, floor)
				}
				if p.x != len(floor[0])-1 {
					newBasin.maybeAddPoint(Point{p.x + 1, p.y}, floor)
				}
				if p.y != len(floor)-1 {
					newBasin.maybeAddPoint(Point{p.x, p.y + 1}, floor)
				}
			}

			previousBasin = basin
			basin = newBasin
		}

		basinSizes = append(basinSizes, len(basin))
	}

	sort.Slice(basinSizes, func(a, b int) bool {
		return basinSizes[b] < basinSizes[a]
	})

	return basinSizes[0] * basinSizes[1] * basinSizes[2]
}

func parseInput(input []string) [][]int {
	floor := make([][]int, len(input))
	for i, line := range input {
		floor[i] = MapToInts(strings.Split(line, ""))
	}
	return floor
}

func getLowPoints(floor [][]int) []Point {
	var lowPoints []Point
	for i := range floor {
		for j := range floor[i] {
			isLowPoint := true
			if i != 0 {
				isLowPoint = isLowPoint && floor[i-1][j] > floor[i][j]
			}
			if j != 0 {
				isLowPoint = isLowPoint && floor[i][j-1] > floor[i][j]
			}
			if i != len(floor)-1 {
				isLowPoint = isLowPoint && floor[i+1][j] > floor[i][j]
			}
			if j != len(floor[i])-1 {
				isLowPoint = isLowPoint && floor[i][j+1] > floor[i][j]
			}

			if isLowPoint {
				lowPoints = append(lowPoints, Point{j, i})
			}
		}
	}
	return lowPoints
}

type Point struct {
	x, y int
}

type Basin []Point

func (b *Basin) maybeAddPoint(point Point, floor [][]int) {
	if floor[point.y][point.x] != 9 && !b.contains(point) {
		*b = append(*b, point)
	}
}

func (a Basin) remove(b Basin) (result Basin) {
	for _, i := range a {
		if !b.contains(i) {
			result = append(result, i)
		}
	}
	return
}

func (a Basin) contains(p Point) bool {
	for _, i := range a {
		if i == p {
			return true
		}
	}
	return false
}
