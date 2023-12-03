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
	starRegExp      = regexp.MustCompile(`[*]+`)
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
		starPositions := starRegExp.FindAllStringIndex(line, -1)

		for _, starIndexPair := range starPositions {
			var previousLine [][]int
			var thisLine [][]int
			var nextLine [][]int
			var arrayOfNumbersAroundThisStar []int

			startLeftBufferIndex := max(starIndexPair[0]-1, 0)
			startRightBufferIndex := min(starIndexPair[1]+1, len(line)-1)

			if lineIndex > 0 {
				previousLine = numberRegExp.FindAllStringIndex(lines[lineIndex-1][startLeftBufferIndex:startRightBufferIndex], -1)
			}

			thisLine = numberRegExp.FindAllStringIndex(lines[lineIndex][startLeftBufferIndex:startRightBufferIndex], -1)

			if lineIndex < len(lines)-1 {
				nextLine = numberRegExp.FindAllStringIndex(lines[lineIndex+1][startLeftBufferIndex:startRightBufferIndex], -1)
			}

			for _, numberPos := range previousLine {
				arrayOfNumbersAroundThisStar = append(arrayOfNumbersAroundThisStar, parseNumberFromLine(lines[lineIndex-1], numberPos[0]+startLeftBufferIndex))
			}
			for _, numberPos := range thisLine {
				arrayOfNumbersAroundThisStar = append(arrayOfNumbersAroundThisStar, parseNumberFromLine(lines[lineIndex], numberPos[0]+startLeftBufferIndex))
			}
			for _, numberPos := range nextLine {
				arrayOfNumbersAroundThisStar = append(arrayOfNumbersAroundThisStar, parseNumberFromLine(lines[lineIndex+1], numberPos[0]+startLeftBufferIndex))
			}

			if len(arrayOfNumbersAroundThisStar) == 2 {
				total += ratio(arrayOfNumbersAroundThisStar)
			}
		}
	}

	fmt.Println(total)
}

func isDigit(s string) bool {
	return s == "0" || s == "1" || s == "2" || s == "3" || s == "4" || s == "5" || s == "6" || s == "7" || s == "8" || s == "9"
}

func parseNumberFromLine(line string, numberIndex int) int {
	numberString := ""
	lineLength := len(line)

	numberString = line[numberIndex : numberIndex+1]

	// post-append numbers
	for x := numberIndex + 1; x < lineLength; x++ {
		nextNumber := line[x : x+1]

		if isDigit(nextNumber) {
			numberString = numberString + nextNumber
		} else {
			break
		}
	}

	// pre-append numbers
	for x := numberIndex - 1; x >= 0; x-- {
		prevNumber := line[x : x+1]

		if isDigit(prevNumber) {
			numberString = prevNumber + numberString
		} else {
			break
		}
	}

	return toInt(numberString)
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

func ratio(x []int) int {
	return x[0] * x[1]
}
