package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	x, y int
}

type Grid map[Point]string
type Directions []string

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: program <parts> <mode>")
	}

	part := os.Args[1]
	mode := os.Args[2]

	file, _ := os.Open(fmt.Sprintf("2024/15/%s.txt", mode))
	defer file.Close()

	grid, directions, startPoint := parse(file)

	switch part {
	case "1":
	case "2":
	}

	fmt.Println(grid, directions, startPoint)
}

func parse(file *os.File) (Grid, Directions, Point) {
	scanner := bufio.NewScanner(file)

	yVal := 0
	grid := make(Grid)
	directions := make(Directions, 0)
	startPoint := Point{}

	section1 := true
	section2 := false

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			section1 = false
			section2 = true
			continue
		}

		if section1 {
			for i := 0; i < len(line); i++ {
				point := Point{x: i, y: yVal}
				grid[point] = string(line[i])

				if grid[point] == "@" {
					startPoint = point
				}
			}
		} else if section2 {
			directions = append(directions, line)
		}

		yVal++
	}

	return grid, directions, startPoint
}
