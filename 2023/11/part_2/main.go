package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strings"
)

type Point struct {
	x, y int
}

var (
	lineBreakRegExp = regexp.MustCompile(`\r?\n`)
	galaxies        = make(map[int]Point)
)

func main() {
	file := readFile("2023/11/input.txt")
	lines := lineBreakRegExp.Split(string(file), -1)

	fillGalaxy(lines)

	totalDistance := 0
	for start, startGalaxy := range galaxies {
		for end, endGalaxy := range galaxies {
			if start >= end {
				continue
			}
			totalDistance += Abs(endGalaxy.y-startGalaxy.y) + Abs(endGalaxy.x-startGalaxy.x)
		}
	}

	fmt.Println(totalDistance)
}

func Abs(i int) int {
	if i < 0 {
		i = i * -1
	}

	return i
}

func fillGalaxy(lines []string) {
	var emptyRows []int
	var emptyCols []int
	var filledCols []int

	for y, line := range lines {
		if strings.Count(line, ".") == len(line) {
			emptyRows = append(emptyRows, y)
		}
		for x := 0; x < len(line); x++ {
			if line[x:x+1] != "." {
				filledCols = append(filledCols, x)
			}
		}
	}

	for c := 0; c < len(lines); c++ {
		if !slices.Contains(filledCols, c) {
			emptyCols = append(emptyCols, c)
		}
	}

	addedToY := 0
	for y, line := range lines {
		if slices.Contains(emptyRows, y) {
			addedToY++
		}
		expandedY := y + addedToY*(1000000-1)

		for x := 0; x < len(line); x++ {
			if line[x:x+1] != "." {
				newEmptyCols := make([]int, len(emptyCols))
				copy(newEmptyCols, emptyCols)

				if slices.Index(emptyCols, x) == -1 {
					newEmptyCols = append(newEmptyCols, x)
				}
				slices.Sort(newEmptyCols)

				expandedX := x + max(slices.Index(newEmptyCols, x), 0)*(1000000-1)
				point := Point{
					x: expandedX,
					y: expandedY,
				}

				galaxies[len(galaxies)] = point
			}
		}
	}
}

func readFile(filepath string) []byte {
	file, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
