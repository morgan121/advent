package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strings"
)

var (
	lineBreakRegExp = regexp.MustCompile(`\r?\n`)
	pattern         []string
	instructions    = make(map[string]Instruction)
	stepCount       = 0
)

type Instruction struct {
	left  string
	right string
}

func main() {
	file := readFile("2023/08/input.txt")
	lines := lineBreakRegExp.Split(string(file), -1)

	for i, line := range lines {
		if i == 0 {
			pattern = strings.Split(line, "")
			continue
		}
		if i == 1 {
			continue
		}

		instructions[line[0:3]] = Instruction{
			left:  line[7:10],
			right: line[12:15],
		}
	}

	var nextLocations []string
	for key := range instructions {
		if key[2:3] == "A" {
			nextLocations = append(nextLocations, key)
		}
	}

	pathTotals := make([]int, len(nextLocations))
	for i := 0; i <= len(pattern); i++ {
		stepCount++

		if i == len(pattern) {
			i = 0
		}

		direction := pattern[i]

		// step for each path
		for j := 0; j < len(nextLocations); j++ {
			nextLocations[j] = step(nextLocations[j], direction)

			if nextLocations[j][2:3] == "Z" {
				pathTotals[j] = stepCount
			}
		}

		if slices.Min(pathTotals) > 0 {
			break
		}
	}

	fmt.Println(LCM(pathTotals[0], pathTotals[1], pathTotals...))
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func step(currentLocation, direction string) string {
	if direction == "L" {
		return instructions[currentLocation].left
	}

	return instructions[currentLocation].right
}

func readFile(filepath string) []byte {
	file, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
