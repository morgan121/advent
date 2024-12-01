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
	mode := os.Args[1]
	file, err := os.Open(fmt.Sprintf("2024/01/%s.txt", mode))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	re := regexp.MustCompile("[0-9]+")
	scanner := bufio.NewScanner(file)
	total := 0

	var first_column []int
	var second_column []int

	for scanner.Scan() {
		line := scanner.Text()
		number_pair := re.FindAllString(line, -1)
		first_column = append(first_column, toInt(number_pair[0]))
		second_column = append(second_column, toInt(number_pair[1]))
	}

	sortSlice(first_column)
	sortSlice(second_column)

	for i := 0; i < len(first_column); i++ {
		total += absInt(second_column[i] - first_column[i])
	}

	fmt.Println(total)
}

func absInt(x int) int {
	return absDiffInt(x, 0)
}

func absDiffInt(x, y int) int {
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
