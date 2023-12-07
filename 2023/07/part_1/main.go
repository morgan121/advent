package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
)

// handType will be a point value:
//
//	7 = 5 of a kind
//	6 = 4 of a kind
//	5 = full house
//	4 = 3 of a kind
//	3 = 2 pair
//	2 = 1 pair
//	1 = high card
type Hand struct {
	cards    []string
	bid      int
	handType int
}

var (
	numberRegExp    = regexp.MustCompile("[0-9]+")
	lineBreakRegExp = regexp.MustCompile(`\r?\n`)
	hands           []Hand
)

func main() {
	file := readFile("2023/07/input.txt")
	lines := lineBreakRegExp.Split(string(file), -1)

	totalHands := len(lines)

	totalWinnings := 0

	for _, line := range lines {
		cards := strings.Split(strings.Split(line, " ")[0], "")
		hand := Hand{
			cards:    cards,
			bid:      toInt(strings.Split(line, " ")[1]),
			handType: handType(cards),
		}

		hands = append(hands, hand)
	}

	sort.Slice(hands, sortHands)

	for i, hand := range hands {
		totalWinnings += hand.bid * (totalHands - i)
	}

	fmt.Println(totalWinnings)
}

func sortHands(i, j int) bool {
	if hands[i].handType > hands[j].handType {
		return true
	}

	if hands[i].handType == hands[j].handType {
		if hands[i].cards[0] == hands[j].cards[0] {
			if hands[i].cards[1] == hands[j].cards[1] {
				if hands[i].cards[2] == hands[j].cards[2] {
					if hands[i].cards[3] == hands[j].cards[3] {
						return higherCard(hands[i].cards[4], hands[j].cards[4])
					}
					return higherCard(hands[i].cards[3], hands[j].cards[3])
				}
				return higherCard(hands[i].cards[2], hands[j].cards[2])
			}
			return higherCard(hands[i].cards[1], hands[j].cards[1])
		}
		return higherCard(hands[i].cards[0], hands[j].cards[0])
	}

	return false
}

func higherCard(x, y string) bool {
	cardValues := make(map[string]int)
	cardValues["A"] = 13
	cardValues["K"] = 12
	cardValues["Q"] = 11
	cardValues["J"] = 10
	cardValues["T"] = 9
	cardValues["9"] = 8
	cardValues["8"] = 7
	cardValues["7"] = 6
	cardValues["6"] = 5
	cardValues["5"] = 4
	cardValues["4"] = 3
	cardValues["3"] = 2
	cardValues["2"] = 1

	return cardValues[x] > cardValues[y]
}

func handType(cards []string) int {
	counts := countOfCards(cards)

	// 5 of a kind
	if sortedCountValues(counts)[0] == 5 {
		return 7
	}

	// 4 of a kind
	if sortedCountValues(counts)[0] == 4 {
		return 6
	}

	// full house OR 3 of a kind
	if sortedCountValues(counts)[0] == 3 {
		if len(counts) == 2 {
			return 5
		}

		return 4
	}

	// 2 pair OR 1 pair
	if sortedCountValues(counts)[0] == 2 {
		if len(counts) == 3 {
			return 3
		}

		return 2
	}

	// highest card
	return 1
}

func countOfCards(slice []string) map[string]int {
	dict := make(map[string]int)
	for _, v := range slice {
		dict[v] = dict[v] + 1
	}
	return dict
}

func sortedCountValues(dict map[string]int) []int {
	values := make([]int, 0, len(dict))

	for _, value := range dict {
		values = append(values, value)
	}

	slices.Sort(values)
	slices.Reverse(values)

	return values
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}

	return n
}

func readFile(filepath string) []byte {
	file, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
