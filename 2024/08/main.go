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

var (
	antennae  = make(map[string][]Point)
	grid      = make(Grid)
	antinodes = make(map[Point]bool)
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: program <parts> <mode>")
	}

	part := os.Args[1]
	mode := os.Args[2]

	file, _ := os.Open(fmt.Sprintf("2024/08/%s.txt", mode))
	defer file.Close()

	grid = parse(file)

	switch part {
	case "1":
		calculatePart1()
	case "2":
		calculatePart2()
	}

	fmt.Println(len(antinodes))
}

func calculatePart1() {
	for _, value := range antennae {
		for i := 0; i < len(value); i++ {
			startPoint := value[i]

			for j := 0; j < len(value); j++ {
				if i == j {
					continue
				}
				otherPoint := value[j]

				potentialAntinode := getAntinode(startPoint, otherPoint)

				if grid[potentialAntinode] != "" {
					antinodes[potentialAntinode] = true
				}
			}
		}
	}
}

func calculatePart2() {
	for _, value := range antennae {
		for i := 0; i < len(value); i++ {
			startPoint := value[i]
			antinodes[startPoint] = true

			for j := 0; j < len(value); j++ {
				if i == j {
					continue
				}
				otherPoint := value[j]

				stillInsideGrid := true
				tmpStartPoint := copyOf(startPoint)
				tmpOtherPoint := copyOf(otherPoint)
				for ok := true; ok; ok = stillInsideGrid {
					potentialAntinode := getAntinode(tmpStartPoint, tmpOtherPoint)

					if grid[potentialAntinode] != "" {
						antinodes[potentialAntinode] = true
						tmpStartPoint = copyOf(tmpOtherPoint)
						tmpOtherPoint = copyOf(potentialAntinode)
					} else {
						stillInsideGrid = false
					}
				}
			}
		}
	}
}

func copyOf(p Point) Point {
	return Point{x: p.x, y: p.y}
}

func getAntinode(p1 Point, p2 Point) Point {
	xDiff := p1.x - p2.x
	yDiff := p1.y - p2.y

	potentialAntinode := Point{x: p2.x - xDiff, y: p2.y - yDiff}

	return potentialAntinode
}

func parse(file *os.File) Grid {
	scanner := bufio.NewScanner(file)

	yVal := 0

	for scanner.Scan() {
		line := scanner.Text()

		for i := 0; i < len(line); i++ {
			point := Point{x: i, y: yVal}
			grid[point] = string(line[i])

			if grid[point] != "." {
				antennae[grid[point]] = append(antennae[grid[point]], point)
			}
		}

		yVal++
	}

	return grid
}
