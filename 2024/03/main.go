package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: program <parts> <mode>")
	}

	parts := os.Args[1]
	mode := os.Args[2]

	file, err := os.Open(fmt.Sprintf("2024/03/%s.txt", mode))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var re *regexp.Regexp
	var validInstructions []string

	switch parts {
	case "1":
		re = regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	case "2":
		re = regexp.MustCompile(`do\(\)|don't\(\)|mul\(\d{1,3},\d{1,3}\)`)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		validInstructions = append(validInstructions, re.FindAllString(line, -1)...)
	}

	fmt.Println(calculate(validInstructions))
}

func calculate(instructions []string) int {
	re := regexp.MustCompile("[0-9]+")
	lineTotal := 0

	enabled := true

	for _, e := range instructions {
		if e == "do()" {
			enabled = true
		} else if e == "don't()" {
			enabled = false
		} else {
			if enabled {
				numbers := re.FindAllString(e, -1)
				lineTotal += toInt(numbers[0]) * toInt(numbers[1])
			}
		}
	}

	return lineTotal
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}

	return n
}
