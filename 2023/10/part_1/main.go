package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type Path struct {
	prevCharacter      string
	prevCharacterIndex int
	nextDirection      int
	length             int
}

var (
	numberRegExp    = regexp.MustCompile("-?[0-9]+")
	lineBreakRegExp = regexp.MustCompile(`\r?\n`)
	availableSteps  = make(map[string]map[int]int)
	paths           = make(map[int][]Path)
	allLines        string
)

func main() {
	file := readFile("2023/10/input.txt")
	lines := lineBreakRegExp.Split(string(file), -1)
	lineLength := len(lines[0])

	allLines = strings.Join(lines, "")

	// 4 directions: N E S W - offsets from the starting point
	directions := []int{-lineLength, 1, lineLength, -1}

	// depnding on the letter + previous step, only certain available steps
	availableSteps["S"] = make(map[int]int)
	availableSteps["S"][-1] = -1
	availableSteps["S"][1] = 1
	availableSteps["S"][-lineLength] = -lineLength
	availableSteps["S"][lineLength] = lineLength
	availableSteps["S"][-lineLength] = 1
	availableSteps["F"] = make(map[int]int)
	availableSteps["F"][-1] = lineLength
	availableSteps["F"][-lineLength] = 1
	availableSteps["L"] = make(map[int]int)
	availableSteps["L"][-1] = -lineLength
	availableSteps["L"][lineLength] = 1
	availableSteps["J"] = make(map[int]int)
	availableSteps["J"][1] = -lineLength
	availableSteps["J"][lineLength] = -1
	availableSteps["7"] = make(map[int]int)
	availableSteps["7"][1] = lineLength
	availableSteps["7"][-lineLength] = -1
	availableSteps["|"] = make(map[int]int)
	availableSteps["|"][-lineLength] = -lineLength
	availableSteps["|"][lineLength] = lineLength
	availableSteps["-"] = make(map[int]int)
	availableSteps["-"][-1] = -1
	availableSteps["-"][1] = 1

	initialIndex := strings.Index(allLines, "S")
	for _, direction := range directions {
		currentCharacter := allLines[initialIndex+direction : (initialIndex+direction)+1]
		_, ok := availableSteps[currentCharacter][direction]

		if ok {
			path := Path{
				prevCharacter:      "S",
				prevCharacterIndex: initialIndex,
				nextDirection:      direction,
				length:             0,
			}
			paths[len(paths)] = append(paths[len(paths)], path)
		}
	}

	for i, path := range paths {

		ok := true
		for ok != false {
			nextPath := getNextPathStep(path[len(path)-1])

			if nextPath.prevCharacter == "." || nextPath.prevCharacter == "S" {
				ok = false
			} else {
				path = append(path, nextPath)
				paths[i] = path
			}
		}
	}

	fmt.Println(len(paths[0]) / len(paths))
}

func getNextPathStep(prevPath Path) Path {
	nextDirection := availableSteps[prevPath.prevCharacter][prevPath.nextDirection]
	nextIndex := prevPath.prevCharacterIndex + nextDirection
	nextCharacter := allLines[nextIndex : nextIndex+1]

	path := Path{
		prevCharacter:      nextCharacter,
		prevCharacterIndex: nextIndex,
		nextDirection:      nextDirection,
		length:             prevPath.length + 1,
	}

	return path
}

func readFile(filepath string) []byte {
	file, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
