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
	grid          = make(Grid)
	startPoint    Point
	direction     string
	maxX          int
	maxY          int
	visited       = make(map[Point]bool)
	hashLocations []Point
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
		fmt.Println(len(visited))
	case "2":
		total := 0
		for _, hash := range hashLocations {
			if startFromBottomLeft(hash) {
				total++
			}
			if startFromBottomRight(hash) {
				total++
			}
			if startFromTopLeft(hash) {
				total++
			}
			if startFromTopRight(hash) {
				total++
			}
		}
		fmt.Println(total) // 1117 is too low
	}

}

/*
.#...
....#
X.... <-- startPoint
...O.
*/
func startFromBottomLeft(bottomLeft Point) bool {
	if bottomLeft.x == maxX || bottomLeft.y == 0 {
		return false
	}

	topLeft := findNextHash(Point{x: bottomLeft.x + 1, y: bottomLeft.y}, "^")
	if topLeft.x > maxX || topLeft.y < 0 {
		return false
	}

	topRight := findNextHash(Point{x: topLeft.x, y: topLeft.y + 1}, ">")
	if topRight.x > maxX || topRight.y > maxY {
		return false
	}

	nextHash := findNextHash(Point{x: topRight.x - 1, y: topRight.y}, "v")

	neededBottomRight := Point{x: topRight.x - 1, y: bottomLeft.y + 1}
	if neededBottomRight.y > maxY {
		return false
	}

	if neededBottomRight.y < nextHash.y {
		fmt.Println("BL: ", neededBottomRight)
	}
	return neededBottomRight.y < nextHash.y
}

/*
.#...
....O
#....
...X. <-- startPoint
*/
func startFromBottomRight(bottomRight Point) bool {
	if bottomRight.x == 0 || bottomRight.y == 0 {
		return false
	}

	bottomLeft := findNextHash(Point{x: bottomRight.x, y: bottomRight.y - 1}, "<")
	if bottomLeft.x < 0 {
		return false
	}

	topLeft := findNextHash(Point{x: bottomLeft.x + 1, y: bottomLeft.y}, "^")
	if topLeft.y < 0 {
		return false
	}

	nextHash := findNextHash(Point{x: topLeft.x, y: topLeft.y + 1}, ">")

	neededTopRight := Point{x: bottomRight.x + 1, y: topLeft.y + 1}
	if neededTopRight.x > maxX {
		return false
	}

	if neededTopRight.x < nextHash.x {
		fmt.Println("BR: ", neededTopRight)
	}
	return neededTopRight.x < nextHash.x
}

/*
.X... <-- startPoint
....#
O....
...#.
*/
func startFromTopLeft(topLeft Point) bool {
	if topLeft.x == maxX || topLeft.y == maxY {
		return false
	}

	topRight := findNextHash(Point{x: topLeft.x, y: topLeft.y + 1}, ">")
	if topRight.x > maxX || topRight.y > maxY {
		return false
	}

	bottomRight := findNextHash(Point{x: topRight.x - 1, y: topRight.y}, "v")
	if bottomRight.y > maxY {
		return false
	}

	nextHash := findNextHash(Point{x: bottomRight.x, y: bottomRight.y - 1}, "<")

	neededBottomLeft := Point{x: topLeft.x - 1, y: bottomRight.y - 1}
	if neededBottomLeft.x < 0 {
		return false
	}

	if neededBottomLeft.x > nextHash.x {
		fmt.Println("TL: ", neededBottomLeft)
	}
	return neededBottomLeft.x > nextHash.x
}

/*
.O...
....X <-- startPoint
#....
...#.
*/
func startFromTopRight(topRight Point) bool {
	if startPoint.x == 0 || startPoint.y == maxY {
		return false
	}

	bottomRight := findNextHash(Point{x: topRight.x - 1, y: topRight.y}, "v")
	if bottomRight.y > maxY {
		return false
	}

	bottomLeft := findNextHash(Point{x: bottomRight.x, y: bottomRight.y - 1}, "<")
	if bottomLeft.x < 0 {
		return false
	}

	nextHash := findNextHash(Point{x: bottomLeft.x + 1, y: bottomLeft.y}, "^")

	neededTopLeft := Point{x: bottomLeft.x + 1, y: topRight.y - 1}
	if neededTopLeft.y < 0 {
		return false
	}

	if neededTopLeft.y > nextHash.y {
		fmt.Println("TR: ", neededTopLeft)
	}
	return neededTopLeft.y > nextHash.y
}

func findNextHash(startPoint Point, direction string) Point {
	switch direction {
	case "^":
		if startPoint.y < 0 {
			return startPoint
		} else if grid[Point{x: startPoint.x, y: startPoint.y - 1}] == "#" {
			return Point{x: startPoint.x, y: startPoint.y - 1}
		} else {
			startPoint = Point{x: startPoint.x, y: startPoint.y - 1}
		}
	case ">":
		if startPoint.x > maxX {
			return startPoint
		} else if grid[Point{x: startPoint.x + 1, y: startPoint.y}] == "#" {
			return Point{x: startPoint.x + 1, y: startPoint.y}
		} else {
			startPoint = Point{x: startPoint.x + 1, y: startPoint.y}
		}
	case "v":
		if startPoint.y > maxY {
			return startPoint
		} else if grid[Point{x: startPoint.x, y: startPoint.y + 1}] == "#" {
			return Point{x: startPoint.x, y: startPoint.y + 1}
		} else {
			startPoint = Point{x: startPoint.x, y: startPoint.y + 1}
		}
	case "<":
		if startPoint.x < 0 {
			return startPoint
		} else if grid[Point{x: startPoint.x - 1, y: startPoint.y}] == "#" {
			return Point{x: startPoint.x - 1, y: startPoint.y}
		} else {
			startPoint = Point{x: startPoint.x - 1, y: startPoint.y}
		}
	}

	return findNextHash(startPoint, direction)
}

func traverse(startPoint Point) {
	switch direction {
	case "^":
		if startPoint.y < 0 {
			return
		} else if grid[Point{x: startPoint.x, y: startPoint.y - 1}] == "#" {
			direction = ">"
		} else {
			visited[startPoint] = true
			startPoint = Point{x: startPoint.x, y: startPoint.y - 1}
		}
	case ">":
		if startPoint.x == maxX {
			return
		} else if grid[Point{x: startPoint.x + 1, y: startPoint.y}] == "#" {
			direction = "v"
		} else {
			visited[startPoint] = true
			startPoint = Point{x: startPoint.x + 1, y: startPoint.y}
		}
	case "v":
		if startPoint.y == maxY {
			return
		} else if grid[Point{x: startPoint.x, y: startPoint.y + 1}] == "#" {
			direction = "<"
		} else {
			visited[startPoint] = true
			startPoint = Point{x: startPoint.x, y: startPoint.y + 1}
		}
	case "<":
		if startPoint.x < 0 {
			return
		}
		if grid[Point{x: startPoint.x - 1, y: startPoint.y}] == "#" {
			direction = "^"
		} else {
			visited[startPoint] = true
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
				startPoint = point
			}

			if string(line[i]) == "#" {
				hashLocations = append(hashLocations, point)
			}
		}

		yVal++
	}

	return grid, startPoint, xVal, yVal
}
