package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	numberRegExp    = regexp.MustCompile("[0-9]+")
	lineBreakRegExp = regexp.MustCompile(`\r?\n`)
)

func main() {
	file := readFile("2023/06/input.txt")
	lines := lineBreakRegExp.Split(string(file), -1)
	time := toInt(strings.Join(numberRegExp.FindAllString(lines[0], -1), ""))
	distance := toInt(strings.Join(numberRegExp.FindAllString(lines[1], -1), ""))

	numberOfWins := 0

	for hold := 0; hold <= time; hold++ {
		calcedDistance := hold * (time - hold)

		if calcedDistance > distance {
			numberOfWins++
		}
	}

	fmt.Println(numberOfWins)
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
