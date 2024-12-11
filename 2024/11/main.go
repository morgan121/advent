package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

var (
	stones = make([]int, 0)
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: program <parts> <mode>")
	}

	part := os.Args[1]
	mode := os.Args[2]

	file, _ := os.Open(fmt.Sprintf("2024/11/%s.txt", mode))
	defer file.Close()

	parse(file)

	switch part {
	case "1":
		for i := 0; i < 25; i++ {
			blink()
		}
		fmt.Println(len(stones))
	case "2":
	}
}

func blink() {
	for i := 0; i < len(stones); i++ {
		convertedStone := convertStone(stones[i])
		stones = slices.Delete(stones, i, i+1)
		stones = slices.Insert(stones, i, convertedStone...)

		if len(convertedStone) > 1 {
			i++
		}
	}
}

func convertStone(stone int) []int {
	if stone == 0 {
		return []int{1}
	} else if len(strconv.Itoa(stone))%2 == 0 {
		stringStone := strconv.Itoa(stone)
		return []int{
			toInt(stringStone[:len(stringStone)/2]),
			toInt(stringStone[len(stringStone)/2:]),
		}
	}

	return []int{stone * 2024}
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

func parse(file *os.File) {
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`\d+`)

	for scanner.Scan() {
		line := scanner.Text()
		numbers := arrayToInt(re.FindAllString(line, -1))
		stones = append(stones, numbers...)
	}
}
