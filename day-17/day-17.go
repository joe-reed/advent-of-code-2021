package main

import (
	"regexp"
	. "utils"
)

func puzzle1(input string) int {
	maxHeight, _ := solve(input)
	return maxHeight
}

func puzzle2(input string) int {
	_, options := solve(input)
	return options
}

func solve(input string) (int, int) {
	area := area(input)

	maxHeights := []int{}
	for x := 0; x < 500; x++ {
		for y := -500; y < 500; y++ {
			isSuccessful := false
			heights := []int{}

			point := Point{0, 0, x, y}
			for i := 0; i < 500; i++ {
				heights = append(heights, point.y)
				point.advance()

				if area.contains(point) {
					isSuccessful = true
					break
				}
			}

			if isSuccessful {
				maxHeight := 0
				for _, height := range heights {
					if height > maxHeight {
						maxHeight = height
					}
				}
				maxHeights = append(maxHeights, maxHeight)
			}
		}
	}

	maxHeight := 0
	for _, height := range maxHeights {
		if height > maxHeight {
			maxHeight = height
		}
	}

	return maxHeight, len(maxHeights)
}

func area(input string) Area {
	re := regexp.MustCompile(`target area: x=(-?\d+)..(-?\d+), y=(-?\d+)..(-?\d+)`)
	matches := re.FindStringSubmatch(input)
	return Area{ConvertToInt(matches[1]), ConvertToInt(matches[2]), ConvertToInt(matches[3]), ConvertToInt(matches[4])}
}

type Area struct {
	xMin, xMax, yMin, yMax int
}

type Point struct {
	x, y, vX, vY int
}

func (p *Point) advance() {
	(*p).x += (*p).vX
	(*p).y += (*p).vY

	if (*p).vX > 0 {
		(*p).vX -= 1
	} else if (*p).vX < 0 {
		(*p).vX += 1
	}

	(*p).vY -= 1
}

func (a Area) contains(p Point) bool {
	return a.xMin <= p.x && p.x <= a.xMax && a.yMin <= p.y && p.y <= a.yMax
}
