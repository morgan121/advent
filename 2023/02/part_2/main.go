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
	file, err := os.Open("2023/02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	red_re := regexp.MustCompile(" [0-9]+ red")
	green_re := regexp.MustCompile(" [0-9]+ green")
	blue_re := regexp.MustCompile(" [0-9]+ blue")
	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")
		game_results := strings.Split(line[1], ";")

		max_red := 0
		max_green := 0
		max_blue := 0

		for i := 0; i < len(game_results); i++ {
			// add 0 to games where that colour did not get pulled out
			red_n := extractNumbers(append(red_re.FindAllString(game_results[i], -1), "0")[0])
			green_n := extractNumbers(append(green_re.FindAllString(game_results[i], -1), "0")[0])
			blue_n := extractNumbers(append(blue_re.FindAllString(game_results[i], -1), "0")[0])

			max_red = max(max_red, red_n)
			max_green = max(max_green, green_n)
			max_blue = max(max_blue, blue_n)
		}

		total += max_red * max_green * max_blue
	}

	fmt.Println(total)
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

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
