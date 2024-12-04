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
	total = 0
	grid  = make(Grid)
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: program <parts> <mode>")
	}

	part := os.Args[1]
	mode := os.Args[2]

	file, err := os.Open(fmt.Sprintf("2024/04/%s.txt", mode))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	xVal := 0
	yVal := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		xVal = len(line)

		for i := 0; i < len(line); i++ {
			grid[Point{x: i, y: yVal}] = string(line[i])
		}

		yVal++
	}

	switch part {
	case "1":
		part1(grid, xVal, yVal)
	case "2":
		part2(grid, xVal, yVal)
	}

}

func part1(grid Grid, maxX int, maxY int) {
	for p, v := range grid {
		if v == "X" {
			total += checkNorth(p) +
				checkNorthEast(p, maxX, maxY) +
				checkEast(p, maxX) +
				checkSouthEast(p, maxX, maxY) +
				checkSouth(p, maxY) +
				checkSouthWest(p, maxY) +
				checkWest(p) +
				checkNorthWest(p)
		}
	}

	fmt.Println(total)
}

func part2(grid Grid, maxX int, maxY int) {
	for p, v := range grid {
		if v == "A" {
			total += checkMiddle(p, maxX, maxY)
		}
	}

	fmt.Println(total)
}

func checkNorth(p Point) int {
	if p.y < 3 {
		return 0
	}

	if grid[Point{x: p.x, y: p.y - 1}] != "M" {
		return 0
	}

	if grid[Point{x: p.x, y: p.y - 2}] != "A" {
		return 0
	}

	if grid[Point{x: p.x, y: p.y - 3}] != "S" {
		return 0
	}

	return 1
}

func checkNorthEast(p Point, maxX, maxY int) int {
	if p.y < 3 || p.x > maxX-3 {
		return 0
	}

	if grid[Point{x: p.x + 1, y: p.y - 1}] != "M" {
		return 0
	}

	if grid[Point{x: p.x + 2, y: p.y - 2}] != "A" {
		return 0
	}

	if grid[Point{x: p.x + 3, y: p.y - 3}] != "S" {
		return 0
	}

	return 1
}

func checkEast(p Point, maxX int) int {
	if p.x > maxX-3 {
		return 0
	}

	if grid[Point{x: p.x + 1, y: p.y}] != "M" {
		return 0
	}

	if grid[Point{x: p.x + 2, y: p.y}] != "A" {
		return 0
	}

	if grid[Point{x: p.x + 3, y: p.y}] != "S" {
		return 0
	}

	return 1
}

func checkSouthEast(p Point, maxX, maxY int) int {
	if p.y > maxY-3 || p.x > maxX-3 {
		return 0
	}

	if grid[Point{x: p.x + 1, y: p.y + 1}] != "M" {
		return 0
	}

	if grid[Point{x: p.x + 2, y: p.y + 2}] != "A" {
		return 0
	}

	if grid[Point{x: p.x + 3, y: p.y + 3}] != "S" {
		return 0
	}

	return 1
}

func checkSouth(p Point, maxY int) int {
	if p.y > maxY-3 {
		return 0
	}

	if grid[Point{x: p.x, y: p.y + 1}] != "M" {
		return 0
	}

	if grid[Point{x: p.x, y: p.y + 2}] != "A" {
		return 0
	}

	if grid[Point{x: p.x, y: p.y + 3}] != "S" {
		return 0
	}

	return 1
}

func checkSouthWest(p Point, maxY int) int {
	if p.y > maxY-3 || p.x < 3 {
		return 0
	}

	if grid[Point{x: p.x - 1, y: p.y + 1}] != "M" {
		return 0
	}

	if grid[Point{x: p.x - 2, y: p.y + 2}] != "A" {
		return 0
	}

	if grid[Point{x: p.x - 3, y: p.y + 3}] != "S" {
		return 0
	}

	return 1
}

func checkWest(p Point) int {
	if p.x < 3 {
		return 0
	}

	if grid[Point{x: p.x - 1, y: p.y}] != "M" {
		return 0
	}

	if grid[Point{x: p.x - 2, y: p.y}] != "A" {
		return 0
	}

	if grid[Point{x: p.x - 3, y: p.y}] != "S" {
		return 0
	}

	return 1
}

func checkNorthWest(p Point) int {
	if p.y < 3 || p.x < 3 {
		return 0
	}

	if grid[Point{x: p.x - 1, y: p.y - 1}] != "M" {
		return 0
	}

	if grid[Point{x: p.x - 2, y: p.y - 2}] != "A" {
		return 0
	}

	if grid[Point{x: p.x - 3, y: p.y - 3}] != "S" {
		return 0
	}

	return 1
}

func checkMiddle(p Point, maxX int, maxY int) int {
	if p.y < 1 || p.y > maxY-1 || p.x < 1 || p.x > maxX-1 {
		return 0
	}

	topLeft := grid[Point{x: p.x - 1, y: p.y - 1}]
	topRight := grid[Point{x: p.x + 1, y: p.y - 1}]
	bottomLeft := grid[Point{x: p.x - 1, y: p.y + 1}]
	bottomRight := grid[Point{x: p.x + 1, y: p.y + 1}]

	if topLeft == "M" && bottomRight == "S" && topRight == "M" && bottomLeft == "S" {
		return 1
	}

	if topLeft == "M" && bottomRight == "S" && topRight == "S" && bottomLeft == "M" {
		return 1
	}

	if topLeft == "S" && bottomRight == "M" && topRight == "M" && bottomLeft == "S" {
		return 1
	}

	if topLeft == "S" && bottomRight == "M" && topRight == "S" && bottomLeft == "M" {
		return 1
	}

	return 0
}
