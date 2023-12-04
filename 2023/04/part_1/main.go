package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	numberRegExp = regexp.MustCompile("[0-9]+")
)

func main() {
	file := readFile("2023/04/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")
		winningNumbers := numberRegExp.FindAllString(strings.Split(line[1], "|")[0], -1)
		uncoveredNumbers := numberRegExp.FindAllString(strings.Split(line[1], "|")[1], -1)
		matches := intersection(winningNumbers, uncoveredNumbers)

		if len(matches) > 0 {
			total += int(math.Pow(2, float64((len(matches) - 1))))
		}
	}

	fmt.Println(total)
}

func intersection(first, second []string) []string {
	out := []string{}
	bucket := map[string]bool{}

	for _, i := range first {
		for _, j := range second {
			if i == j && !bucket[i] {
				out = append(out, i)
				bucket[i] = true
			}
		}
	}

	return out
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

func readFile(filepath string) *os.File {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
