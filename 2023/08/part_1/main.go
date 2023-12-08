package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
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

		location := line[0:3]

		instruction := Instruction{
			left:  line[7:10],
			right: line[12:15],
		}

		instructions[location] = instruction
	}

	nextLocation := "AAA"
	for i := 0; i <= len(pattern); i++ {
		if i == len(pattern) {
			i = 0
		}

		direction := pattern[i]
		nextLocation = step(nextLocation, direction)

		if nextLocation == "ZZZ" {
			break
		}
	}

	fmt.Println(stepCount)
}

func step(currentLocation, direction string) string {
	switch direction {
	case "L":
		stepCount++
		return instructions[currentLocation].left
	case "R":
		stepCount++
		return instructions[currentLocation].right
	}

	return "Uh-oh"
}

func readFile(filepath string) []byte {
	file, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
