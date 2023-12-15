package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

var (
	hashRegExp      = regexp.MustCompile("#+")
	questionRegExp  = regexp.MustCompile("[?]")
	lineBreakRegExp = regexp.MustCompile(`\r?\n`)
)

func main() {
	file := readFile("2023/12/input.txt")
	lines := lineBreakRegExp.Split(string(file), -1)

	for _, line := range lines {
		pattern := unfoldpattern(strings.Split(strings.Split(line, " ")[0], ""))
		report := unfoldReport(arrayToInt(strings.Split(strings.Split(line, " ")[1], ",")))
		total := 0

		fullReport := make([]int, len(pattern))
		var tempReport []int

		// create the initial report layout (of the correct size)
		for i, count := range report {
			var nextChunk []int

			for j := 0; j < count; j++ {
				nextChunk = append(nextChunk, j+1)
			}

			tempReport = append(tempReport, nextChunk...)

			if i < len(report)-1 {
				tempReport = append(tempReport, 0)
			}
		}

		copy(fullReport, tempReport)

		if reportMatchesPattern(pattern, fullReport) {
			total++
		}

		// for {
		// 	// if report fits totalCombos++
		// 	// report = get next report
		// 	// break if nil report
		// }

		fmt.Println(fullReport)
		fmt.Println(total)
	}
}

func nextReportLayout(report []int) []int {
	var nextReport []int
	if report[len(report)-1] == 0 {
		// I give up
	}

	return nextReport
}

func reportMatchesPattern(pattern []string, report []int) bool {
	// . = 0
	// # = 1
	// ? could be either
	for index := range pattern {
		if report[index] == 0 && pattern[index] == "#" || report[index] > 0 && pattern[index] == "." {
			return false
		}
	}

	return true
}

func moveInt(array []int, srcIndex int, dstIndex int) []int {
	value := array[srcIndex]
	return insertInt(removeInt(array, srcIndex), value, dstIndex)
}

func insertInt(array []int, value int, index int) []int {
	return append(array[:index], append([]int{value}, array[index:]...)...)
}

func removeInt(array []int, index int) []int {
	return append(array[:index], array[index+1:]...)
}

func unfoldpattern(s []string) []string {
	var h []string

	// for i := 0; i < 4; i++ {
	// 	h = append(h, s...)
	// 	h = append(h, "?")
	// }

	h = append(h, s...)

	return h
}

func unfoldReport(r []int) []int {
	var dm []int

	for i := 0; i < 1; i++ {
		dm = append(dm, r...)
	}

	return dm
}

func arrayToInt(s []string) []int {
	var intArray []int

	for _, e := range s {
		intArray = append(intArray, toInt(e))
	}

	return intArray
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}

	return n
}

func readFile(filepath string) []byte {
	file, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
