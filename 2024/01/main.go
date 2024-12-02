package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	parts := os.Args[1]
	mode := os.Args[2]

	file, err := os.Open(fmt.Sprintf("2024/01/%s.txt", mode))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	re := regexp.MustCompile("[0-9]+")
	scanner := bufio.NewScanner(file)

	var col_1 []int
	var col_2 []int

	for scanner.Scan() {
		line := scanner.Text()
		number_pair := re.FindAllString(line, -1)
		col_1 = append(col_1, toInt(number_pair[0]))
		col_2 = append(col_2, toInt(number_pair[1]))
	}

	switch parts {
	case "1":
		part1(col_1, col_2)
	case "2":
		part2(col_1, col_2)
	case "both":
		part1(col_1, col_2)
		part2(col_1, col_2)
	}
}

func part1(col_1 []int, col_2 []int) {
	total := 0

	sortSlice(col_1)
	sortSlice(col_2)

	for i := 0; i < len(col_1); i++ {
		total += absInt(col_2[i], col_1[i])
	}

	fmt.Println(total)
}

func part2(col_1 []int, col_2 []int) {
	total := 0
	freq := make(map[int]int)

	for i := 0; i < len(col_2); i++ {
		freq[col_2[i]]++
	}

	for i := 0; i < len(col_1); i++ {
		total += col_1[i] * freq[col_1[i]]
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

func sortSlice(s []int) []int {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	return s
}
