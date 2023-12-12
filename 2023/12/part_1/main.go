package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/stat/combin"
)

type Point struct {
	x, y int
}

var (
	dotRegExp       = regexp.MustCompile(`\.+`)
	hashRegExp      = regexp.MustCompile("#+")
	questionRegExp  = regexp.MustCompile("[?]")
	lineBreakRegExp = regexp.MustCompile(`\r?\n`)
)

func main() {
	file := readFile("2023/12/input.txt")
	lines := lineBreakRegExp.Split(string(file), -1)

	characterMap := make(map[int]string)
	characterMap[0] = "."
	characterMap[1] = "#"

	totalCombos := 0

	for _, line := range lines {
		hotsprings := strings.Split(line, " ")[0]
		damagedMap := arrayToInt(strings.Split(strings.Split(line, " ")[1], ","))

		questionIndex := questionRegExp.FindAllStringIndex(hotsprings, -1)
		combinations := allArrangements(questionIndex)

		for _, combination := range combinations {
			for i, e := range combination {
				hotsprings = replace(hotsprings, questionIndex[i][0], characterMap[e])
			}

			if damagedReportFits(hotsprings, damagedMap) {
				totalCombos++
			}
		}
	}

	fmt.Println(totalCombos)
}

func damagedReportFits(hotsprings string, damage []int) bool {
	hashGroups := hashRegExp.FindAllStringIndex(hotsprings, -1)

	// there must be one hash group per damage report
	if len(hashGroups) != len(damage) {
		return false
	} else {
		for i, hashGroup := range hashGroups {
			if damage[i] != hashGroup[1]-hashGroup[0] {
				return false
			}
		}
	}

	return true
}

func replace(str string, index int, replacement string) string {
	return str[:index] + replacement + str[index+1:]
}

func allArrangements(questionMarks [][]int) [][]int {
	var lens []int

	for i := 0; i < len(questionMarks); i++ {
		lens = append(lens, 2)
	}

	combinations := combin.Cartesian(lens)

	return combinations
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

func arrayToString(i []int) []string {
	var stringArray []string

	for _, e := range i {
		stringArray = append(stringArray, strconv.Itoa(e))
	}

	return stringArray
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
