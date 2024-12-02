package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: program <parts> <mode>")
	}

	parts := os.Args[1]
	mode := os.Args[2]

	file, err := os.Open(fmt.Sprintf("2024/01/%s.txt", mode))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	re := regexp.MustCompile("[0-9]+")
	scanner := bufio.NewScanner(file)

	var left []int
	var right []int

	for scanner.Scan() {
		line := scanner.Text()
		number_pair := re.FindAllString(line, -1)
		left = append(left, toInt(number_pair[0]))
		right = append(right, toInt(number_pair[1]))
	}

	switch parts {
	case "1":
		part1(left, right)
	case "2":
		part2(left, right)
	case "both":
		part1(left, right)
		part2(left, right)
	}
}

func part1(left []int, right []int) {
	total := 0

	slices.Sort(left)
	slices.Sort(right)

	for i := 0; i < len(left); i++ {
		total += absInt(left[i], right[i])
	}

	fmt.Println(total)
}

func part2(left []int, right []int) {
	total := 0
	freq := make(map[int]int)

	for _, value := range right {
		freq[value]++
	}

	for i := 0; i < len(left); i++ {
		total += left[i] * freq[left[i]]
	}

	fmt.Println(total)
}

func absInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}

	return n
}
