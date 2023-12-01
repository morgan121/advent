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
	file, err := os.Open("2023/01/part_1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	re := regexp.MustCompile("[0-9]+")
	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		line := fix(scanner.Text())
		number_array := re.FindAllString(line, -1)
		number_string := strings.Join(number_array, "")
		two_digits := number_string[0:1] + number_string[len(number_string)-1:]

		n, err := strconv.Atoi(two_digits)
		if err != nil {
			fmt.Println(err)
		} else {
			total = total + n
		}
	}
	fmt.Println(total)
}

func fix(line string) string {
	line = strings.Replace(line, "one", "o1e", -1)
	line = strings.Replace(line, "two", "t2o", -1)
	line = strings.Replace(line, "three", "t3e", -1)
	line = strings.Replace(line, "four", "f4r", -1)
	line = strings.Replace(line, "five", "f5e", -1)
	line = strings.Replace(line, "six", "s6i", -1)
	line = strings.Replace(line, "seven", "s7n", -1)
	line = strings.Replace(line, "eight", "e8t", -1)
	line = strings.Replace(line, "nine", "n9e", -1)

	return line
}
