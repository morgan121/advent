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

type VisitedPoint struct {
	x, y      int
	direction string
}

type Grid map[Point]string

var (
	grid                 = make(Grid)
	startPoint           Point
	direction            string
	maxX                 int
	maxY                 int
	visited              []Point
	visitedWithDirection = make(map[VisitedPoint]bool)
	numberOfLoops        = 0
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: program <parts> <mode>")
	}

	part := os.Args[1]
	mode := os.Args[2]

	file, _ := os.Open(fmt.Sprintf("2024/06/%s.txt", mode))
	defer file.Close()

	grid, startPoint, maxX, maxY = parse(file)

	direction = grid[startPoint]

	switch part {
	case "1":
		traverse(startPoint)
		fmt.Println(len(unique(visited)))
	case "2":
		traverse(startPoint)

		for _, p := range unique(visited) {
			visitedWithDirection = make(map[VisitedPoint]bool)
			direction = "^"

			if grid[p] == "." {
				grid[p] = "#"
				traverse(startPoint)
				grid[p] = "."
			}
		}

		fmt.Println(numberOfLoops) // 1117 is too low
	}

}

func unique(points []Point) []Point {
	var unique []Point

	for _, v := range points {
		skip := false
		for _, u := range unique {
			if v == u {
				skip = true
				break
			}
		}
		if !skip {
			unique = append(unique, v)
		}
	}

	return unique
}

func traverse(startPoint Point) {
	if visitedWithDirection[VisitedPoint{x: startPoint.x, y: startPoint.y, direction: direction}] {
		numberOfLoops++
		return
	}

	switch direction {
	case "^":
		if startPoint.y < 0 {
			return
		} else if grid[Point{x: startPoint.x, y: startPoint.y - 1}] == "#" {
			direction = ">"
		} else {
			visited = append(visited, startPoint)
			visitedWithDirection[VisitedPoint{x: startPoint.x, y: startPoint.y, direction: direction}] = true
			startPoint = Point{x: startPoint.x, y: startPoint.y - 1}
		}
	case ">":
		if startPoint.x == maxX {
			return
		} else if grid[Point{x: startPoint.x + 1, y: startPoint.y}] == "#" {
			direction = "v"
		} else {
			visited = append(visited, startPoint)
			visitedWithDirection[VisitedPoint{x: startPoint.x, y: startPoint.y, direction: direction}] = true
			startPoint = Point{x: startPoint.x + 1, y: startPoint.y}
		}
	case "v":
		if startPoint.y == maxY {
			return
		} else if grid[Point{x: startPoint.x, y: startPoint.y + 1}] == "#" {
			direction = "<"
		} else {
			visited = append(visited, startPoint)
			visitedWithDirection[VisitedPoint{x: startPoint.x, y: startPoint.y, direction: direction}] = true
			startPoint = Point{x: startPoint.x, y: startPoint.y + 1}
		}
	case "<":
		if startPoint.x < 0 {
			return
		}
		if grid[Point{x: startPoint.x - 1, y: startPoint.y}] == "#" {
			direction = "^"
		} else {
			visited = append(visited, startPoint)
			visitedWithDirection[VisitedPoint{x: startPoint.x, y: startPoint.y, direction: direction}] = true
			startPoint = Point{x: startPoint.x - 1, y: startPoint.y}
		}
	default:
		return
	}

	traverse(startPoint)
}

func parse(file *os.File) (Grid, Point, int, int) {
	scanner := bufio.NewScanner(file)

	xVal := 0
	yVal := 0

	var startPoint Point

	for scanner.Scan() {
		line := scanner.Text()
		xVal = len(line) - 1

		for i := 0; i < len(line); i++ {
			point := Point{x: i, y: yVal}
			grid[point] = string(line[i])

			if string(line[i]) == "^" {
				startPoint = Point{x: point.x, y: point.y}
			}
		}

		yVal++
	}

	return grid, startPoint, xVal, yVal
}
