package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var (
	numberRegExp    = regexp.MustCompile("[0-9]+")
	characterRegexp = regexp.MustCompile(`[^a-zA-Z\d\s.:]+`)
	lineBreakRegExp = regexp.MustCompile(`\r?\n`)
	characterHash   = make(map[int]int)
)

func main() {
	file, err := os.ReadFile("2023/03/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	total := 0

	lines := lineBreakRegExp.Split(string(file), -1)

	for lineIndex, line := range lines {
		numberPositions := numberRegExp.FindAllStringIndex(line, -1)

		for _, numberIndexPair := range numberPositions {
			aboveIndex := max(lineIndex-1, 0)
			belowIndex := min(lineIndex+1, len(lines)-1)

			firstIndex := max(numberIndexPair[0]-1, 0)
			lastIndex := min(numberIndexPair[1]+1, len(line)-1)

			above := append(characterRegexp.FindStringIndex(lines[aboveIndex][firstIndex:lastIndex]), -1)[0]
			equal := append(characterRegexp.FindStringIndex(lines[lineIndex][firstIndex:lastIndex]), -1)[0]
			below := append(characterRegexp.FindStringIndex(lines[belowIndex][firstIndex:lastIndex]), -1)[0]

			if above >= 0 || equal >= 0 || below >= 0 {
				total += getNumberAtPositions(lineIndex, numberIndexPair, lines)
			}
		}
	}

	fmt.Println(total)
}

func getNumberAtPositions(lineIndex int, numberIndexPair []int, lines []string) int {
	return toInt(lines[lineIndex][numberIndexPair[0]:numberIndexPair[1]])
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}

	return n
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
