package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var part1Total int
var part2Total int

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage (from advent root): go run 2025/01/main.go <parts> (1 or 2) <mode> (real or test)")
	}

	parts := os.Args[1]
	mode := os.Args[2]

	file, err := os.Open(fmt.Sprintf("2025/01/%s.txt", mode))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	re := regexp.MustCompile(`[0-9]+`)
	scanner := bufio.NewScanner(file)

	dialValue := 50

	for scanner.Scan() {
		line := scanner.Text()
		direction := line[0:1]
		distance := toInt(re.FindAllString(line, -1)[0])

		part2Total += distance / 100
		dialValue = rotate(
			dialValue,
			direction,
			distance%100,
		)
	}

	switch parts {
	case "1":
		fmt.Println(part1Total)
	case "2":
		fmt.Println(part2Total)
	}
}

func rotate(dialValue int, direction string, distance int) int {
	if direction == "R" {
		startVal := dialValue
		if dialValue == 100 {
			startVal = 0
		}
		return recalculate(startVal + distance)
	}

	if direction == "L" {
		startVal := dialValue
		if dialValue == 0 {
			startVal = 100
		}
		return recalculate(startVal - distance)
	}

	return dialValue
}

func recalculate(number int) int {
	if number%100 == 0 {
		part1Total++
		part2Total++
		return number
	}
	if number < 0 {
		part2Total++
		return number + 100
	} else if number > 100 {
		part2Total++
		return number - 100
	}

	return number
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}

	return n
}
