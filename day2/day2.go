package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isValidId(id string, exactlyTwiceMultiplier bool) bool {
	testChars := ""
	for i := 0; i < len(id); i++ {
		testChars += string(id[i])
		stringRatio := len(id) % len(testChars)
		multiplier := 2
		if !exactlyTwiceMultiplier {
			multiplier = len(id) / len(testChars)
		}

		// If the lengths of these don't divide exactly, they won't repeat so ID is valid
		if stringRatio > 0 {
			continue
		}

		if multiplier == 1 {
			return true
		}

		repeatedString := strings.Repeat(testChars, multiplier)

		if repeatedString == id {
			return false
		}
	}

	return true
}

func part1(lines []string) int {
	// Only a single line in this part:
	input := strings.TrimSpace(lines[0])
	parts := strings.Split(input, ",")

	count := 0

	for _, part := range parts {
		numRange := strings.Split(part, "-")
		lower := numRange[0]
		upper := numRange[1]

		for i := parseIntOrError(lower); i <= parseIntOrError(upper); i++ {
			if !isValidId(fmt.Sprint(i), true) {
				count += i
			}
		}
	}
	return count
}

func part2(lines []string) int {
	// Only a single line in this part:
	input := strings.TrimSpace(lines[0])
	parts := strings.Split(input, ",")

	count := 0

	for _, part := range parts {
		numRange := strings.Split(part, "-")
		lower := numRange[0]
		upper := numRange[1]

		for i := parseIntOrError(lower); i <= parseIntOrError(upper); i++ {
			if !isValidId(fmt.Sprint(i), false) {
				count += i
			}
		}
	}
	return count
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
