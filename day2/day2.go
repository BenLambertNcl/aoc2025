package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func isValidId(id string, exactlyTwiceMultiplier bool) bool {
	testChars := ""
	strLen := len(id)

	for i := 0; i < strLen/2; i++ {
		testChars += string(id[i])
		lenTest := len(testChars)

		// If the lengths of these don't divide exactly, they won't repeat so ID is valid
		stringRatio := strLen % lenTest
		if stringRatio > 0 {
			continue
		}

		multiplier := 2
		if !exactlyTwiceMultiplier {
			multiplier = strLen / lenTest
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
	start := time.Now()
	part1 := part1(lines)
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Day 1:", part1, ", Time:", elapsed)

	start = time.Now()
	part2 := part2(lines)
	t = time.Now()
	elapsed = t.Sub(start)
	fmt.Println("Day 2:", part2, ", Time:", elapsed)
}

func readInstructions() []string {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(data), "\n")
}

func parseIntOrError(num string) int {
	parsed, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(parsed)
}
