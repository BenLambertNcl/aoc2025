package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type positions struct {
	x, y int
}

func checkAdjacentPositions(grid [][]string, x, y int) int {
	gridHeight := len(grid)
	gridWidth := len(grid[0])

	count := 0

	allPositions := []positions{
		{x - 1, y - 1}, // left upper diag
		{x, y - 1},     // upper
		{x + 1, y - 1}, // right upper diag
		{x - 1, y},     // left
		{x + 1, y},     // right
		{x - 1, y + 1}, // right lower diag
		{x, y + 1},     // lower
		{x + 1, y + 1}, // right lower diag
	}

	for _, pos := range allPositions {
		if pos.x >= 0 && pos.x < gridWidth && pos.y >= 0 && pos.y < gridHeight {
			if grid[pos.y][pos.x] == "@" {
				count++
			}
		}
	}
	return count
}

func forEachGridPosition(grid [][]string, fn func(x, y int)) {
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			fn(x, y)
		}
	}
}

func part1(grid [][]string) int {
	total := 0
	forEachGridPosition(grid, func(x, y int) {
		if grid[y][x] == "@" {
			count := checkAdjacentPositions(grid, x, y)
			if count < 4 {
				total++
			}
		}
	})

	return total
}

func part2(grid [][]string) int {
	total := 0
	var positionsToReplace []positions
	lastLoopTotal := -1

	// Run until the previous look found no valid positions
	for lastLoopTotal != 0 {
		lastLoopTotal = 0
		forEachGridPosition(grid, func(x, y int) {
			if grid[y][x] == "@" {
				count := checkAdjacentPositions(grid, x, y)
				if count < 4 {
					lastLoopTotal++
					positionsToReplace = append(positionsToReplace, positions{x, y})
				}
			}
		})

		for _, pos := range positionsToReplace {
			grid[pos.y][pos.x] = "."
		}

		total += lastLoopTotal
	}

	return total
}

func main() {
	grid := readInstructionsAsGrid()
	fmt.Println("Day 1:", part1(grid))
	fmt.Println("Day 2:", part2(grid))
}

func readInstructions() []string {
	data, _ := os.ReadFile("input.txt")
	return strings.Split(string(data), "\n")
}

func readInstructionsAsGrid() [][]string {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(data), "\r\n")
	grid := make([][]string, len(lines))
	for i, line := range lines {
		grid[i] = strings.Split(line, "")
	}
	return grid
}

func parseIntOrError(num string) int {
	parsed, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(parsed)
}
