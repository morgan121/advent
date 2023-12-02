package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file := readFile("2023/02/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	red_re := regexp.MustCompile(" [0-9]+ red")
	green_re := regexp.MustCompile(" [0-9]+ green")
	blue_re := regexp.MustCompile(" [0-9]+ blue")
	total := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")

		all_red := extractNumberArray(red_re.FindAllString(line[1], -1))
		all_green := extractNumberArray(green_re.FindAllString(line[1], -1))
		all_blue := extractNumberArray(blue_re.FindAllString(line[1], -1))
		sort.Ints(all_red)
		sort.Ints(all_green)
		sort.Ints(all_blue)

		max_red := all_red[len(all_red)-1]
		max_green := all_green[len(all_green)-1]
		max_blue := all_blue[len(all_blue)-1]

		total += max_red * max_green * max_blue
	}

	fmt.Println(total)
}

func extractNumbers(s string) []int {
	number_re := regexp.MustCompile("[0-9]+")

	var number_array []int
	for _, e := range number_re.FindAllString(s, -1) {
		number_array = append(number_array, toInt(e))
	}

	return number_array
}

func extractNumberArray(s []string) []int {
	number_re := regexp.MustCompile("[0-9]+")

	var number_array []int
	for _, e := range s {
		number_array = append(number_array, toInt(number_re.FindAllString(e, -1)[0]))
	}

	return number_array
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}

	return n
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func readFile(filepath string) *os.File {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
