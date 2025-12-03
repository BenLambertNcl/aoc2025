package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func findJoltage(line string) string {
	max := -1
	maxIndex := -1
	for i := 0; i < len(line)-1; i++ {
		num, _ := strconv.Atoi(string(line[i]))
		if num > max {
			max = num
			maxIndex = i
		}

		if max == 9 {
			break
		}
	}

	nextMax := -1
	for i := maxIndex + 1; i < len(line); i++ {
		num, _ := strconv.Atoi(string(line[i]))
		if num > nextMax {
			nextMax = num
		}
	}

	return fmt.Sprintf("%d%d", max, nextMax)
}

func findLargestWithEndBuffer(line string, start, buffer int) (int, int) {
	max := -1
	index := -1
	for i := start; i < len(line)-buffer; i++ {
		num, _ := strconv.Atoi(string(line[i]))
		if num > max {
			max = num
			index = i
		}
	}
	return max, index
}

func findLargerJoltage(line string) string {
	lastIndex := -1
	joltage := ""
	for i := 1; i <= 12; i++ {
		value, index := findLargestWithEndBuffer(line, lastIndex+1, 12-i)
		joltage += strconv.Itoa(value)
		lastIndex = index
	}

	return joltage
}

func part1(lines []string) int {
	total := 0
	for _, line := range lines {
		output := findJoltage(strings.TrimSpace(line))
		joltage, _ := strconv.Atoi(output)
		total += joltage
	}
	return total
}

func part2(lines []string) int {
	total := 0
	for _, line := range lines {
		output := findLargerJoltage(strings.TrimSpace(line))
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
