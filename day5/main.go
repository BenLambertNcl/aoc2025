package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(lines []string) int {
	return 0
}

func part2(lines []string) int {
	return 0
}

func main() {
	lines := readInstructions()
	fmt.Println("Day 1:", part1(lines))
	fmt.Println("Day 2:", part2(lines))
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
