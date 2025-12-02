package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage (from advent root): go run 2025/01/main.go <parts> (1 or 2) <mode> (real or test)")
	}

	part := os.Args[1]
	mode := os.Args[2]

	file, err := os.Open(fmt.Sprintf("2025/02/%s.txt", mode))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var invalidSum int

	for scanner.Scan() {
		line := scanner.Text()
		rawRanges := strings.Split(line, ",")

		for _, rawRange := range rawRanges {
			startVal := toInt(strings.Split(rawRange, "-")[0])
			endVal := toInt(strings.Split(rawRange, "-")[1])

			for value := startVal; value <= endVal; value++ {
				if part == "1" {
					if isInvalidPart1(strconv.Itoa(value)) {
						invalidSum += value
					}
				}

				if part == "2" {
					if isInvalidPart2(strconv.Itoa(value)) {
						invalidSum += value
					}
				}
			}
		}
	}

	fmt.Println(invalidSum)
}

func isInvalidPart1(value string) bool {
	n := len(value)
	if n%2 != 0 {
		return false
	}
	half := n / 2
	return value[:half] == value[half:]
}

func isInvalidPart2(value string) bool {
	n := len(value)
	if n <= 1 {
		return false
	}
	ss := (value + value)[1 : 2*n-1]
	return strings.Contains(ss, value)
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}

	return n
}
