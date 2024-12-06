package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strconv"
)

type Rule []int

var (
	rules             []Rule
	orders            [][]int
	validOrderTotal   int
	invalidOrderTotal int
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: program <parts> <mode>")
	}

	part := os.Args[1]
	mode := os.Args[2]

	file, _ := os.Open(fmt.Sprintf("2024/05/%s.txt", mode))
	defer file.Close()

	rules, orders = parse(file)

	for _, order := range orders {
		originalOrder := append([]int(nil), order...)

		sort.SliceStable(order, func(i, j int) bool {
			for _, rule := range rules {
				if rule[0] == order[i] && rule[1] == order[j] {
					return true
				}
			}

			return false
		})

		if reflect.DeepEqual(order, originalOrder) {
			validOrderTotal += order[len(order)/2]
		} else {
			invalidOrderTotal += order[len(order)/2]
		}
	}

	switch part {
	case "1":
		fmt.Println(validOrderTotal)
	case "2":
		fmt.Println(invalidOrderTotal)
	}
}

func arrayToInt(s []string) []int {
	var intArray []int

	for _, e := range s {
		intArray = append(intArray, toInt(e))
	}

	return intArray
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}

	return n
}

func parse(file *os.File) ([]Rule, [][]int) {
	reRules := regexp.MustCompile(`\d{1,2}\|\d{1,2}`)
	reorder := regexp.MustCompile(`\d{1,2},\d{1,2}`)
	reDigit := regexp.MustCompile(`\d{1,2}`)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if reRules.MatchString(line) {
			startEnd := reDigit.FindAllString(line, -1)
			rules = append(rules, arrayToInt(startEnd))
		} else if reorder.MatchString(line) {
			orders = append(orders, arrayToInt(reDigit.FindAllString(line, -1)))
		}
	}

	return rules, orders
}
