package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Direction int

const (
	LEFT Direction = iota
	RIGHT
)

func getRotationAmount(line string) (direction Direction, amount int64) {
	if strings.HasPrefix(line, "L") {
		direction = LEFT
		num := strings.Trim(strings.Split(line, "L")[1], "\r\n ")
		rotations, err := strconv.ParseInt(num, 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		amount = rotations
		return
	}

	if strings.HasPrefix(line, "R") {
		direction = RIGHT
		num := strings.Trim(strings.Split(line, "R")[1], "\r\n ")
		rotations, err := strconv.ParseInt(num, 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		amount = rotations
		return
	}

	log.Fatalf("Unknown prefix...: %q", line)
	return
}

func part1(lines []string) int {
	var position int64 = 50
	var max int64 = 100

	zeroCount := 0

	for _, element := range lines {
		if element == "" {
			continue
		}

		direction, amount := getRotationAmount(element)

		adjustedAmount := amount % max

		switch direction {
		case LEFT:
			position -= adjustedAmount
		case RIGHT:
			position += adjustedAmount
		}

		if position == 0 {
			zeroCount++
		} else if position >= max {
			position = position % max
		} else if position < 0 {
			position += max
		}
	}

	return zeroCount
}

func part2(lines []string) int {
	var position int64 = 50
	var max int64 = 100

	zeroCount := 0

	for _, element := range lines {
		if element == "" {
			continue
		}

		direction, amount := getRotationAmount(element)

		adjustedAmount := amount % max
		zeroCount += int(amount / max)

		prevPosition := position

		switch direction {
		case LEFT:
			position -= adjustedAmount
		case RIGHT:
			position += adjustedAmount
		}

		if position == 0 {
			zeroCount++
		} else if position >= max {
			if prevPosition != 0 {
				zeroCount++
			}
			position = position % max
		} else if position < 0 {
			if prevPosition != 0 {
				zeroCount++
			}
			position += max
		}
	}

	return zeroCount
}

func main() {
	lines := readInstructions()
	fmt.Println("Day 1:", part1(lines))
	fmt.Println("Day 2:", part2(lines))
}

func readInstructions() []string {
	data, _ := os.ReadFile("day2.txt")
	return strings.Split(string(data), "\n")
}
