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
	if len(os.Args) < 3 {
		log.Fatal("Usage: program <parts> <mode>")
	}

	parts := os.Args[1]
	mode := os.Args[2]

	file, err := os.Open(fmt.Sprintf("2024/02/%s.txt", mode))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	re := regexp.MustCompile("[0-9]+")
	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		numbers := arrayToInt(re.FindAllString(line, -1))

		switch parts {
		case "1":
			if safe(numbers) {
				total++
			}
		case "2":
			if safe(numbers) {
				total++
			} else {
				for i := 0; i < len(numbers); i++ {
					if safe(excludeElement(numbers, i)) {
						total++
						break
					}
				}
			}
		}
	}

	fmt.Println(total)
}

func safe(numbers []int) bool {
	if duplicates(numbers) {
		return false
	}

	if sort.SliceIsSorted(numbers, func(i, j int) bool { return numbers[i] < numbers[j] }) || sort.SliceIsSorted(numbers, func(i, j int) bool { return numbers[i] > numbers[j] }) {
		if maxDistance(numbers) <= 3 {
			return true
		}
	}

	return false
}

func duplicates(numbers []int) bool {
	freq := make(map[int]int)

	for _, value := range numbers {
		if freq[value] >= 1 {
			return true
		}
		freq[value]++
	}

	return false
}

func maxDistance(numbers []int) int {
	max := 0

	for i := 1; i < len(numbers); i++ {
		if absInt(numbers[i], numbers[i-1]) > max {
			max = absInt(numbers[i], numbers[i-1])
		}
	}

	return max
}

func excludeElement(s []int, i int) []int {
	remainingNumbers := make([]int, 0)

	remainingNumbers = append(remainingNumbers, s[:i]...)
	remainingNumbers = append(remainingNumbers, s[i+1:]...)

	return remainingNumbers
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

func absInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
