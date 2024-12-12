package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: program <parts> <mode>")
	}

	part := os.Args[1]
	mode := os.Args[2]

	file, _ := os.Open(fmt.Sprintf("2024/11/%s.txt", mode))
	defer file.Close()

	initialStones := parse(file)

	blinkCount := map[string]int{
		"1": 25,
		"2": 75,
	}[part]

	total := simulateStones(initialStones, blinkCount)
	fmt.Println(total)
}

func simulateStones(initialStones []int, blinks int) int {
	// Map to track the count of each stone
	stoneCounts := make(map[int]int)
	for _, stone := range initialStones {
		stoneCounts[stone]++
	}

	for i := 0; i < blinks; i++ {
		nextStoneCounts := make(map[int]int)
		for stone, count := range stoneCounts {
			if stone == 0 {
				nextStoneCounts[1] += count
			} else {
				strStone := strconv.Itoa(stone)
				if len(strStone)%2 == 0 {
					halfLen := len(strStone) / 2
					left, right := toInt(strStone[:halfLen]), toInt(strStone[halfLen:])
					nextStoneCounts[left] += count
					nextStoneCounts[right] += count
				} else {
					nextStoneCounts[stone*2024] += count
				}
			}
		}
		stoneCounts = nextStoneCounts
	}

	// Calculate the total number of stones
	totalStones := 0
	for _, count := range stoneCounts {
		totalStones += count
	}
	return totalStones
}

// func processStones(blinkCount int) int {
// 	total := 0

// 	for _, stone := range stones {
// 		if stoneCount[stone] == 0 {
// 			transformStone(stone, stone, 0, blinkCount)
// 		}
// 		total += stoneCount[stone]
// 	}

// 	return total
// }

// func transformStone(stone int, initialStone int, iterations int, maxIterations int) {
// 	if iterations == maxIterations {
// 		return
// 	}

// 	if _, exists := converter[stone]; !exists {
// 		converter[stone] = convertStone(stone)
// 	}

// 	newStones := converter[stone]
// 	if len(newStones) == 2 {
// 		stoneCount[initialStone]++
// 	}

// 	for _, stone := range newStones {
// 		transformStone(stone, initialStone, iterations+1, maxIterations)
// 	}
// }

// func convertStone(stone int) []int {
// 	if stone == 0 {
// 		return []int{1}
// 	}

// 	strStone := strconv.Itoa(stone)
// 	if len(strStone)%2 == 0 {
// 		halfLen := len(strStone) / 2
// 		return []int{
// 			toInt(strStone[:halfLen]),
// 			toInt(strStone[halfLen:]),
// 		}
// 	}

// 	return []int{stone * 2024}
// }

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

func parse(file *os.File) []int {
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`\d+`)

	var stones []int

	for scanner.Scan() {
		line := scanner.Text()
		numbers := arrayToInt(re.FindAllString(line, -1))
		stones = append(stones, numbers...)
	}

	return stones
}
