package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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
	index_total := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")
		game_index := extractNumbers(line[0])
		draw_results := strings.Split(line[1], ";")

		valid_red := true
		valid_green := true
		valid_blue := true

		for i := 0; i < len(draw_results); i++ {
			red_n := extractNumbers(append(red_re.FindAllString(draw_results[i], -1), "0")[0])
			green_n := extractNumbers(append(green_re.FindAllString(draw_results[i], -1), "0")[0])
			blue_n := extractNumbers(append(blue_re.FindAllString(draw_results[i], -1), "0")[0])

			if red_n <= 12 {
				valid_red = valid_red && true
			} else {
				valid_red = false
			}

			if green_n <= 13 {
				valid_green = valid_green && true
			} else {
				valid_green = false
			}

			if blue_n <= 14 {
				valid_blue = valid_blue && true
			} else {
				valid_blue = false
			}
		}

		if valid_red && valid_green && valid_blue {
			index_total += game_index
		}
	}

	fmt.Println(index_total)
}

func extractNumbers(s string) int {
	number_re := regexp.MustCompile("[0-9]+")

	return toInt(number_re.FindAllString(s, -1)[0])
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}

	return n
}

func readFile(filepath string) *os.File {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
