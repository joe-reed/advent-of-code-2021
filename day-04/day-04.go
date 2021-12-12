package main

import (
	"errors"
	"strings"
	. "utils"
)

func puzzle1(file string) int {
	numbers, boards := parseInput(file)

	for _, number := range numbers {
		for _, board := range boards {
			if board.contains(number) {
				board.mark(number)
			}

			if board.isWinner() {
				return ConvertToInt(number) * board.getUnmarkedNumberTotal()
			}
		}
	}

	return 0
}

func puzzle2(file string) int {
	numbers, boards := parseInput(file)

	loserCandidates := make([]Board, len(boards))
	copy(loserCandidates, boards)

	for _, number := range numbers {
		newLoserCandidates := []Board{}

		for _, board := range loserCandidates {
			if board.contains(number) {
				board.mark(number)
			}

			if !board.isWinner() {
				newLoserCandidates = append(newLoserCandidates, board)
			}
		}

		if len(newLoserCandidates) == 0 {
			return ConvertToInt(number) * loserCandidates[0].getUnmarkedNumberTotal()
		}

		loserCandidates = newLoserCandidates
	}

	return 0
}

func parseInput(file string) ([]string, []Board) {
	blocks := strings.Split(string(file), "\n\n")
	numbersString, boardStrings := blocks[0], blocks[1:]
	numbers := strings.Split(numbersString, ",")

	var boards []Board
	for _, boardString := range boardStrings {
		boards = append(boards, mapStringToBoard(boardString))
	}

	return numbers, boards
}

func mapStringToBoard(boardString string) Board {
	rowStrings := strings.Split(boardString, "\n")

	var rows [][]string
	for _, rowString := range rowStrings {
		rows = append(rows, strings.Fields(rowString))
	}

	marks := make([][]bool, len(rows))
	for i := range marks {
		marks[i] = make([]bool, len(rows))
	}

	return Board{rows, marks}
}

type Board struct {
	rows  [][]string
	marks [][]bool
}

func (b *Board) contains(n string) bool {
	_, _, err := b.findLocation(n)

	return err == nil
}

func (b *Board) mark(n string) {
	x, y, _ := b.findLocation(n)

	b.marks[y][x] = true
}

func (b *Board) isWinner() bool {
	return b.hasCompleteRow() || b.hasCompleteColumn()
}

func (b *Board) hasCompleteRow() bool {
	return hasRowOfTrues(b.marks)
}

func (b *Board) hasCompleteColumn() bool {
	return hasRowOfTrues(transpose(b.marks))
}

func (b *Board) getUnmarkedNumberTotal() int {
	var unmarkedNumbers []string

	for i, row := range b.rows {
		for j, number := range row {
			if !b.marks[i][j] {
				unmarkedNumbers = append(unmarkedNumbers, number)
			}
		}
	}

	result := 0
	for _, number := range unmarkedNumbers {
		result += ConvertToInt(number)
	}
	return result
}

func (b *Board) findLocation(n string) (int, int, error) {
	for y, row := range b.rows {
		for x, number := range row {
			if n == number {
				return x, y, nil
			}
		}
	}

	return -1, -1, errors.New("number not found")
}

func hasRowOfTrues(grid [][]bool) bool {
	for _, row := range grid {
		isRowComplete := true

		for _, value := range row {
			isRowComplete = isRowComplete && value
		}

		if isRowComplete {
			return true
		}
	}

	return false
}

func transpose(slice [][]bool) [][]bool {
	result := make([][]bool, len(slice[0]))
	for i := range result {
		result[i] = make([]bool, len(slice))
	}

	for i := 0; i < len(slice[0]); i++ {
		for j := 0; j < len(slice); j++ {
			result[i][j] = slice[j][i]
		}
	}

	return result
}
