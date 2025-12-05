package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime"
	"slices"
	"sort"
)

type Roll struct {
	line, index int
}

func main() {
	file, part := setup()

	defer file.Close()

	rolls := parse(file)

	switch part {
	case "1":
		removedIndexes := removeRolls(rolls)
		fmt.Println(len(removedIndexes))
	case "2":
		totalRemoved := 0
		for {
			removedIndexes := removeRolls(rolls)
			if len(removedIndexes) == 0 {
				break
			} else {
				totalRemoved += len(removedIndexes)
				for _, i := range removedIndexes {
					rolls = append(rolls[:i], rolls[i+1:]...)
				}
			}
		}
		fmt.Println(totalRemoved)
	}
}

func removeRolls(rs []Roll) []int {
	removedIndexes := []int{}

	for i, roll := range rs {
		if calcNeighbors(roll, rs) < 4 {
			removedIndexes = append(removedIndexes, i)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(removedIndexes)))
	return removedIndexes
}

func calcNeighbors(r Roll, rs []Roll) int {
	total := 0
	if slices.Contains(rs, Roll{line: r.line - 1, index: r.index - 1}) {
		total++
	}
	if slices.Contains(rs, Roll{line: r.line - 1, index: r.index}) {
		total++
	}
	if slices.Contains(rs, Roll{line: r.line - 1, index: r.index + 1}) {
		total++
	}

	if slices.Contains(rs, Roll{line: r.line, index: r.index - 1}) {
		total++
	}
	if slices.Contains(rs, Roll{line: r.line, index: r.index + 1}) {
		total++
	}

	if slices.Contains(rs, Roll{line: r.line + 1, index: r.index - 1}) {
		total++
	}
	if slices.Contains(rs, Roll{line: r.line + 1, index: r.index}) {
		total++
	}
	if slices.Contains(rs, Roll{line: r.line + 1, index: r.index + 1}) {
		total++
	}

	return total
}

func parse(file *os.File) []Roll {
	scanner := bufio.NewScanner(file)

	rolls := []Roll{}

	lineNo := 1
	for scanner.Scan() {
		line := scanner.Text()

		for i := 0; i < len(line); i++ {
			if string(line[i]) == "@" {
				rolls = append(rolls, Roll{line: lineNo, index: i})
			}
		}

		lineNo++
	}

	return rolls
}

func setup() (*os.File, string) {
	if len(os.Args) < 3 {
		log.Fatal("Usage (from advent root): go run 2025/01/main.go <part> (1 or 2) <mode> (real or test)")
	}

	_, filename, _, _ := runtime.Caller(0)
	part := os.Args[1]
	mode := os.Args[2]

	re := regexp.MustCompile(`[0-9]+`)
	paths := re.FindAllString(filename, -1)

	file, err := os.Open(fmt.Sprintf("%s/%s/%s.txt", paths[0], paths[1], mode))
	if err != nil {
		log.Fatal(err)
	}

	return file, part
}
