package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var (
	numberRegExp = regexp.MustCompile("[0-9]+")
)

type Race struct {
	time     int
	distance int
}

var races []Race

func main() {
	file := readFile("2023/06/input.txt")
	allNumbers := arrayToInt(numberRegExp.FindAllString(string(file), -1))

	for i := 0; i < len(allNumbers)/2; i++ {
		race := Race{
			time:     allNumbers[i],
			distance: allNumbers[i+len(allNumbers)/2],
		}
		races = append(races, race)
	}

	total := 0

	for _, race := range races {
		numberOfWins := 0

		for hold := 0; hold <= race.time; hold++ {
			if hold*(race.time-hold) > race.distance {
				numberOfWins++
			}
		}

		if total == 0 {
			total = numberOfWins
		} else {
			total *= numberOfWins
		}
	}

	fmt.Println(total)
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

func readFile(filepath string) []byte {
	file, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
