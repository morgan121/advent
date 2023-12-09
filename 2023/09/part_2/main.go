package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var (
	numberRegExp    = regexp.MustCompile("-?[0-9]+")
	lineBreakRegExp = regexp.MustCompile(`\r?\n`)
)

func main() {
	file := readFile("2023/09/input.txt")
	lines := lineBreakRegExp.Split(string(file), -1)

	total := 0
	for _, line := range lines {
		var sequence [][]int
		sequence = append(sequence, arrayToInt(numberRegExp.FindAllString(line, -1)))

		finished := false
		for finished != true {
			nextLayer := nextLayer(sequence[len(sequence)-1])
			sequence = append(sequence, nextLayer)

			if sum(nextLayer) == 0 {
				finished = true
			}
		}

		for i := len(sequence) - 2; i >= 0; i-- {
			prevElement := sequence[i][0] - sequence[i+1][0]
			sequence[i] = append([]int{prevElement}, sequence[i]...)

			if i == 0 {
				total += prevElement
			}
		}
	}

	fmt.Println(total)
}

func sum(nums []int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

func nextLayer(seq []int) []int {
	var nextLayer []int

	for i := 1; i < len(seq); i++ {
		nextLayer = append(nextLayer, seq[i]-seq[i-1])
	}

	return nextLayer
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
