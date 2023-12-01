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
	file, err := os.Open("2023/01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	re := regexp.MustCompile("[0-9]+")
	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
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
