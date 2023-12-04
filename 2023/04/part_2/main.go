package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	numberRegExp    = regexp.MustCompile("[0-9]+")
	lineBreakRegExp = regexp.MustCompile(`\r?\n`)
)

type cardLineInfoType struct {
	cardCount   int
	lineMatches int
}

func main() {
	file := readFile("2023/04/input.txt")
	lines := lineBreakRegExp.Split(string(file), -1)

	total := 0

	cardLineInfoMap := make(map[int]cardLineInfoType)

	for lineIndex, line := range lines {
		var cardLineInfo cardLineInfoType

		numbers := strings.Split(line, ":")
		winningNumbers := numberRegExp.FindAllString(strings.Split(numbers[1], "|")[0], -1)
		uncoveredNumbers := numberRegExp.FindAllString(strings.Split(numbers[1], "|")[1], -1)
		matches := intersection(winningNumbers, uncoveredNumbers)

		cardLineInfo.cardCount = 1
		cardLineInfo.lineMatches = len(matches)
		cardLineInfoMap[lineIndex] = cardLineInfo
	}

	keys := make([]int, 0, len(cardLineInfoMap))
	for k := range cardLineInfoMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, key := range keys {
		var cardLineInfo cardLineInfoType

		info := cardLineInfoMap[key]
		count := info.cardCount

		for times := 0; times < count; times++ {
			matches := info.lineMatches

			for i := 1; i <= matches; i++ {
				cardLineInfo.cardCount = cardLineInfoMap[key+i].cardCount + 1
				cardLineInfo.lineMatches = cardLineInfoMap[key+i].lineMatches
				cardLineInfoMap[key+i] = cardLineInfo
			}
		}
	}

	for _, info := range cardLineInfoMap {
		total += info.cardCount
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

func readFile(filepath string) []byte {
	file, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
