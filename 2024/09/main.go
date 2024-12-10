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
	blockLength []int
	gapLength   []int
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: program <parts> <mode>")
	}

	part := os.Args[1]
	mode := os.Args[2]

	file, _ := os.Open(fmt.Sprintf("2024/09/%s.txt", mode))
	defer file.Close()

	blockLength, gapLength = parse(file)

	switch part {
	case "1":
		calculatePart1()
	case "2":
		calculatePart2()
	}

}

func calculatePart1() {
	disc := condense(translate(blockLength, gapLength))
	total := 0

	for i, value := range disc {
		if value != -1 {
			total += value * i
		}
	}

	fmt.Println(total)
}

func calculatePart2() {
}

func condense(disc []int) []int {
	processing := true
	for ok := true; ok; ok = processing {
		firstGapIdx := slices.Index(disc, -1)
		lastBlockIdx := lastNonZeroIndex(disc)

		if firstGapIdx > lastBlockIdx {
			processing = false
		} else {
			lastBlockValue := disc[lastBlockIdx]
			disc[firstGapIdx] = lastBlockValue
			disc[lastBlockIdx] = -1
		}
	}

	return disc
}

func translate(blockLength []int, gapLength []int) []int {
	var translated []int

	initialBlock := makeBlock(0, blockLength[0])
	translated = append(translated, initialBlock...)
	for i := 1; i < len(blockLength); i++ {
		gap := makeBlock(-1, gapLength[i-1])
		block := makeBlock(i, blockLength[i])
		translated = append(translated, gap...)
		translated = append(translated, block...)
	}

	return translated
}

func lastNonZeroIndex(input []int) int {
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] != -1 {
			return i
		}
	}

	return -1
}

func makeBlock(value int, length int) []int {
	block := make([]int, 0)
	for i := 0; i < length; i++ {
		block = append(block, value)
	}

	return block
}

func parse(file *os.File) ([]int, []int) {
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`\d`)

	var blockLength []int
	var gapLength []int

	for scanner.Scan() {
		line := scanner.Text()
		numbers := arrayToInt(re.FindAllString(line, -1))

		for i, char := range numbers {
			if i%2 == 0 {
				blockLength = append(blockLength, char)
			} else {
				gapLength = append(gapLength, char)
			}
		}
	}

	return blockLength, gapLength
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
