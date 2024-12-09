package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var (
	total = 0
	match = false
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: program <parts> <mode>")
	}

	part := os.Args[1]
	mode := os.Args[2]

	file, _ := os.Open(fmt.Sprintf("2024/07/%s.txt", mode))
	defer file.Close()

	scanner := bufio.NewScanner(file)

	re := regexp.MustCompile(`[0-9]+`)

	for scanner.Scan() {
		line := scanner.Text()
		elements := re.FindAllString(line, -1)
		solution := toInt(elements[0])
		numbers := arrayToInt(elements[1:])

		match = false
		calculate(numbers, solution, part)

		if match {
			total += solution
		}
	}

	fmt.Println(total)
}

func calculate(input []int, solution int, part string) {
	if len(input) == 1 {
		if input[0] == solution {
			match = true
		}
		return
	}

	var operations []string
	if part == "1" {
		operations = []string{"+", "*"}
	} else {
		operations = []string{"+", "*", "||"}
	}

	for _, o := range operations {
		switch o {
		case "+":
			tmp := input[0] + input[1]
			calculate(append([]int{tmp}, input[2:]...), solution, part)
		case "*":
			tmp := input[0] * input[1]
			calculate(append([]int{tmp}, input[2:]...), solution, part)
		case "||":
			tmp := toInt(fmt.Sprintf("%d%d", input[0], input[1]))
			calculate(append([]int{tmp}, input[2:]...), solution, part)
		}
	}
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}

	return n
}

func arrayToInt(s []string) []int {
	var intArray []int

	for _, e := range s {
		intArray = append(intArray, toInt(e))
	}

	return intArray
}
