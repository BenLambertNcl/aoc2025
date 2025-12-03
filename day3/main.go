package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func findLargest(line string, from, to int) (int, int) {
	max := -1
	index := -1
	for i := from; i < len(line)-to; i++ {
		num, _ := strconv.Atoi(string(line[i]))
		if num > max {
			max = num
			index = i
		}
	}
	return max, index
}

func findJoltage(line string, numValues int) string {
	lastIndex := -1
	joltage := ""
	for i := 1; i <= numValues; i++ {
		value, index := findLargest(line, lastIndex+1, numValues-i)
		joltage += strconv.Itoa(value)
		lastIndex = index
	}

	return joltage
}

func part1(lines []string) int {
	total := 0
	for _, line := range lines {
		output := findJoltage(strings.TrimSpace(line), 2)
		joltage, _ := strconv.Atoi(output)
		total += joltage
	}
	return total
}

func part2(lines []string) int {
	total := 0
	for _, line := range lines {
		output := findJoltage(strings.TrimSpace(line), 12)
		joltage, _ := strconv.Atoi(output)
		total += joltage
	}
	return total
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

func parseIntOrError(num string) int {
	parsed, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(parsed)
}
