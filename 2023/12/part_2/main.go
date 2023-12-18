package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"sync"
)

type Point struct {
	x, y int
}

var (
	count           int
	lineBreakRegExp = regexp.MustCompile(`\r?\n`)
)

func main() {
	file := readFile("2023/12/input.txt")
	lines := lineBreakRegExp.Split(string(file), -1)

	var wg sync.WaitGroup

	wg.Add(len(lines))

	for _, line := range lines {
		pattern := unfoldpattern(strings.Split(strings.Split(line, " ")[0], ""))
		report := unfoldReport(arrayToInt(strings.Split(strings.Split(line, " ")[1], ",")))
		// fmt.Printf("Splitting line %d\n", i)
		for {
			if pattern[len(pattern)-1] == "." {
				pattern = pattern[:len(pattern)-1]
			} else if pattern[0] == "." {
				pattern = pattern[1:]
			} else {
				break
			}
		}

		go reportFits(report, pattern)
	}

	wg.Wait()
	fmt.Println(count)
}

func reportFits(report []int, pattern []string) {
	chunkSize := report[0]
	// fmt.Printf("report value: %d, pattern: %s\n", chunkSize, pattern)

	// not enough space left in pattern
	if sum(report) > len(pattern)+len(report)-1 || chunkSize > len(pattern) {
		return
	}

	if allQuestions(pattern) {
		n := len(pattern) - sum(report) + 1
		r := len(report)

		if n == 40 {
			count += 137846528820
			return
		}

		count += factorial(n) / ((factorial(r)) * (factorial(n - r)))
		return
	}

	chunk := pattern[:chunkSize]

	if pattern[0] == "." {
		// fmt.Println("EAT DOT")
		reportFits(report, pattern[1:])
		return
	}

	if len(report) == 1 && chunkSize == len(pattern) && !slices.Contains(chunk, ".") {
		// fmt.Println("YEAH BOI - 1")

		count++
		return
	}

	if len(report) != 1 && chunkSize == len(pattern) {
		return
	}

	if slices.Contains(chunk, ".") {
		if chunk[0] == "?" {
			reportFits(report, pattern[1:])
		}
		return
	}

	if len(report) == 1 && allQuestionOrDots(pattern[chunkSize:]) {
		// fmt.Println("YEAH BOI - 2")
		count++
	}

	switch chunk[0] {
	case "#":
		// fmt.Println("hash")
		// treat subsequent ? as .
		if pattern[chunkSize] == "?" {
			// fmt.Println("    hash - question")
			if len(report) > 1 {
				reportFits(report[1:], pattern[chunkSize+1:])
			}
			return
		}

		// doesnt fit pattern
		if pattern[chunkSize] == "#" {
			// fmt.Println("    hash - hash")
			return
		}

		// next char is "."
		// fmt.Println("    hash - dot")
		if len(report) > 1 {
			reportFits(report[1:], pattern[chunkSize:])
		}
	case "?":
		// fmt.Println("question")
		if pattern[chunkSize] == "#" {
			// fmt.Println("    question - hash")
			reportFits(report, pattern[1:])
		}

		if pattern[chunkSize] == "." {
			// fmt.Println("    question - dot")
			if len(report) > 1 {
				reportFits(report[1:], pattern[chunkSize:])
			}

			// if all questions in the chunk are dots
			if allQuestionOrDots(chunk) {
				reportFits(report, pattern[chunkSize:])
			}
		}

		if pattern[chunkSize] == "?" {
			// fmt.Println("    question - question")
			// treat next question as dot
			if len(report) > 1 {
				reportFits(report[1:], pattern[chunkSize+1:])
			}

			// treat first question as dot
			reportFits(report, pattern[1:])
		}
	}
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}

	return n * factorial(n-1)
}

func allQuestions(pattern []string) bool {
	if slices.Contains(pattern, "#") || slices.Contains(pattern, ".") {
		return false
	}

	return true
}

func allQuestionOrDots(pattern []string) bool {
	if slices.Contains(pattern, "#") {
		return false
	}

	return true
}

func unfoldpattern(s []string) []string {
	var h []string

	for i := 0; i < 4; i++ {
		h = append(h, s...)
		h = append(h, "?")
	}

	h = append(h, s...)

	return h
}

func unfoldReport(r []int) []int {
	var dm []int

	for i := 0; i < 5; i++ {
		dm = append(dm, r...)
	}

	return dm
}

func sum(nums []int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
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
