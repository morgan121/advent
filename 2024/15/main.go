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

	grid, wideGrid, directions, startPoint, wideStartPoint := parse(file)

	switch part {
	case "1":
		for _, direction := range directions {
			grid, startPoint = step(grid, direction, startPoint)
		}

		obstacles := make([]Point, 0)
		for point, value := range grid {
			if value == "O" {
				obstacles = append(obstacles, point)
			}
		}

		fmt.Println(calculate(obstacles))
	case "2":
		for _, direction := range directions {
			wideGrid, wideStartPoint = step(wideGrid, direction, wideStartPoint)
		}

		boxes := make([]Point, 0)
		for point, value := range grid {
			if value == "[" {
				boxes = append(boxes, point)
			}
		}

		fmt.Println(calculate(boxes))
		fmt.Println(wideStartPoint)
	}
}

func calculate(positions []Point) int {
	total := 0
	for _, point := range positions {
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
	case "[", "]":
		grid, startPoint = pushBoxes(grid, startPoint, pointInFront, direction)
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

func pushBoxes(grid Grid, startPoint Point, pointInFront Point, direction string) (Grid, Point) {
	switch direction {
	case "^":
		for y := pointInFront.y - 1; y >= 0; y-- {
			nextPoint := Point{x: startPoint.x, y: y}
			if grid[nextPoint] == "#" {
				return grid, startPoint
			}
			if grid[nextPoint] == "." {
				for i := y; i < pointInFront.y; i++ {
					grid[Point{x: startPoint.x, y: i}] = grid[Point{x: startPoint.x, y: i + 1}]
				}
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
				for i := y; i > pointInFront.y; i-- {
					grid[Point{x: startPoint.x, y: i}] = grid[Point{x: startPoint.x, y: i - 1}]
				}
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
				for i := x; i < pointInFront.x; i++ {
					grid[Point{x: i, y: startPoint.y}] = grid[Point{x: i + 1, y: startPoint.y}]
				}
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
				for i := x; i > pointInFront.x; i-- {
					grid[Point{x: i, y: startPoint.y}] = grid[Point{x: i - 1, y: startPoint.y}]
				}
				grid[pointInFront] = "@"
				grid[startPoint] = "."
				return grid, pointInFront
			}
		}
	}

	return grid, startPoint
}

func pushBox(grid Grid, startPoint Point, direction string) (Grid, Point) {
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
	case "[", "]":
		grid, startPoint = pushBoxes(grid, startPoint, pointInFront, direction)
	}

	return grid, startPoint
}

func canPushBox(grid Grid, leftPoint Point, rightPoint Point, direction string) bool {
	switch direction {
	case "^":
		aboveLeft := Point{x: leftPoint.x, y: leftPoint.y - 1}
		aboveRight := Point{x: rightPoint.x, y: rightPoint.y - 1}

		if grid[aboveLeft] == "#" || grid[aboveRight] == "#" {
			return false
		} else if grid[aboveLeft] == "." || grid[aboveRight] == "." {
			return true
		} else {
			canPush := true
			if grid[aboveLeft] == "[" {
				canPush = canPushBox(grid, aboveLeft, aboveRight, direction) && canPush
			}
			if grid[aboveLeft] == "]" {
				canPush = canPushBox(grid, Point{x: aboveLeft.x - 1, y: aboveLeft.y}, aboveLeft, direction) && canPush
			}
			if grid[aboveRight] == "[" {
				canPush = canPushBox(grid, aboveRight, Point{x: aboveLeft.x + 1, y: aboveLeft.y}, direction) && canPush
			}

			return canPush
		}
	case "v":
		belowLeft := Point{x: leftPoint.x, y: leftPoint.y + 1}
		belowRight := Point{x: rightPoint.x, y: rightPoint.y + 1}

		if grid[belowLeft] == "#" || grid[belowRight] == "#" {
			return false
		} else if grid[belowLeft] == "." || grid[belowRight] == "." {
			return true
		} else {
			canPush := true
			if grid[belowLeft] == "[" {
				canPush = canPushBox(grid, belowLeft, belowRight, direction) && canPush
			}
			if grid[belowLeft] == "]" {
				canPush = canPushBox(grid, Point{x: belowLeft.x - 1, y: belowLeft.y}, belowLeft, direction) && canPush
			}
			if grid[belowRight] == "[" {
				canPush = canPushBox(grid, belowRight, Point{x: belowLeft.x + 1, y: belowLeft.y}, direction) && canPush
			}

			return canPush
		}
	}

	return false
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

func parse(file *os.File) (Grid, Grid, Directions, Point, Point) {
	scanner := bufio.NewScanner(file)

	yVal := 0
	grid := make(Grid)
	wideGrid := make(Grid)
	directions := make(Directions, 0)
	startPoint := Point{}
	wideStartPoint := Point{}

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
				switch string(line[i]) {
				case "#", ".":
					wideGrid[Point{x: 2 * i, y: point.y}] = string(line[i])
					wideGrid[Point{x: 2*i + 1, y: point.y}] = string(line[i])
				case "O":
					wideGrid[Point{x: 2 * i, y: point.y}] = "["
					wideGrid[Point{x: 2*i + 1, y: point.y}] = "]"
				case "@":
					wideGrid[Point{x: 2 * i, y: point.y}] = "@"
					wideGrid[Point{x: 2*i + 1, y: point.y}] = "."
				}

				if grid[point] == "@" {
					startPoint = point
					wideStartPoint = Point{x: 2 * i, y: point.y}
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

	return grid, wideGrid, directions, startPoint, wideStartPoint
}
