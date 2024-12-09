package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
)

var (
	permutations = make(map[int][][]string)
	total        = 0
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: program <parts> <mode>")
	}

	// part := os.Args[1]
	mode := os.Args[2]

	file, _ := os.Open(fmt.Sprintf("2024/07/%s.txt", mode))
	defer file.Close()

	scanner := bufio.NewScanner(file)

	re := regexp.MustCompile(`[0-9]+`)

	// var options map[int][][]string

	for scanner.Scan() {
		line := scanner.Text()
		elements := arrayToInt(re.FindAllString(line, -1))
		solution := elements[0]
		numbers := elements[1:]

		if len(numbers) == 1 {
			if numbers[0] == solution {
				total++
			}
			continue
		}

		if permutations[len(numbers)-1] == nil {
			permutations[len(numbers)-1] = permute([]string{"+", "*"}, len(numbers)-1)
		}

		for _, p := range permutations[len(numbers)-1] {
			tmp := numbers[0]
			for i := 0; i < len(p); i++ {
				if p[i] == "+" {
					tmp += numbers[i+1]
				} else {
					tmp *= numbers[i+1]
				}
			}

			if tmp == solution {
				total += solution
				break
			}
		}
	}

	fmt.Println(total)
}

func permute(choices []string, spots int) [][]string {
	res := make([][]string, 0)

	var first, last []string
	for i := 0; i < spots; i++ {
		first = append(first, choices[0])
		last = append(last, choices[1])
	}

	res = append(res, first)

	var backTrack func([]string)
	backTrack = func(variation []string) {
		if reflect.DeepEqual(variation, last) {
			return
		}

		temp := make([]string, 0)
		temp = append(temp, variation...)

		lastStarIdx := getLastIndex(variation, "*")
		lastPlusIdx := getLastIndex(variation, "+")

		if lastStarIdx == -1 {
			// e.g. +++ -> ++*
			temp[lastPlusIdx] = "*"
		} else if lastPlusIdx < lastStarIdx {
			// e.g. +++* -> ++*+
			for i := lastPlusIdx; i < len(temp); i++ {
				temp[i] = "+"
			}
			temp[lastPlusIdx] = "*"
		} else if lastPlusIdx > lastStarIdx {
			// e.g. +*++ -> +*+*
			temp[lastPlusIdx] = "*"
		}

		res = append(res, temp)
		backTrack(temp)
	}

	backTrack(first)
	return res
}

func getLastIndex(input []string, value string) int {
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] == value {
			return i
		}
	}

	return -1
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
