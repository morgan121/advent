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

var xMax, yMax int

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
		for _, direction := range directions {
			grid, startPoint = step(grid, direction, startPoint)
		}

		oPositions := make([]Point, 0)
		for point, value := range grid {
			if value == "O" {
				oPositions = append(oPositions, point)
			}
		}

		fmt.Println(calculate(oPositions))
	case "2":
	}
}

func calculate(oPositions []Point) int {
	total := 0
	for _, point := range oPositions {
		total += point.x + point.y*100
	}

	return total
}

func step(grid Grid, direction string, startPoint Point) (Grid, Point) {
	pointInFront := getPointInFront(startPoint, direction)
	if grid[pointInFront] == "#" {
		return grid, startPoint
	}

	switch grid[pointInFront] {
	case ".":
		grid[pointInFront] = "@"
		grid[startPoint] = "."
		startPoint = pointInFront
	case "O":
		grid, startPoint = pushObstacles(grid, startPoint, pointInFront, direction)
	}

	return grid, startPoint
}

func getPointInFront(startPoint Point, direction string) Point {
	switch direction {
	case "^":
		return Point{x: startPoint.x, y: startPoint.y - 1}
	case "v":
		return Point{x: startPoint.x, y: startPoint.y + 1}
	case "<":
		return Point{x: startPoint.x - 1, y: startPoint.y}
	case ">":
		return Point{x: startPoint.x + 1, y: startPoint.y}
	}

	return startPoint
}

func pushObstacles(grid Grid, startPoint, pointInFront Point, direction string) (Grid, Point) {
	switch direction {
	case "^":
		for y := pointInFront.y - 1; y >= 0; y-- {
			nextPoint := Point{x: startPoint.x, y: y}
			if grid[nextPoint] == "#" {
				return grid, startPoint
			}
			if grid[nextPoint] == "." {
				grid[nextPoint] = "O"
				grid[pointInFront] = "@"
				grid[startPoint] = "."
				return grid, pointInFront
			}
		}
	case "v":
		for y := pointInFront.y + 1; y <= yMax; y++ {
			nextPoint := Point{x: startPoint.x, y: y}
			if grid[nextPoint] == "#" {
				return grid, startPoint
			}
			if grid[nextPoint] == "." {
				grid[nextPoint] = "O"
				grid[pointInFront] = "@"
				grid[startPoint] = "."
				return grid, pointInFront
			}
		}
	case "<":
		for x := pointInFront.x - 1; x >= 0; x-- {
			nextPoint := Point{x: x, y: startPoint.y}
			if grid[nextPoint] == "#" {
				return grid, startPoint
			}
			if grid[nextPoint] == "." {
				grid[nextPoint] = "O"
				grid[pointInFront] = "@"
				grid[startPoint] = "."
				return grid, pointInFront
			}
		}
	case ">":
		for x := pointInFront.x + 1; x <= xMax; x++ {
			nextPoint := Point{x: x, y: startPoint.y}
			if grid[nextPoint] == "#" {
				return grid, startPoint
			}
			if grid[nextPoint] == "." {
				grid[nextPoint] = "O"
				grid[pointInFront] = "@"
				grid[startPoint] = "."
				return grid, pointInFront
			}
		}
	}

	return grid, startPoint
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
		xMax = len(line) - 1

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
			yVal++
		} else if section2 {
			for i := 0; i < len(line); i++ {
				directions = append(directions, string(line[i]))
			}
		}
	}

	yMax = yVal

	return grid, directions, startPoint
}
