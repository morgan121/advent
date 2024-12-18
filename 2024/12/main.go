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
type VisitedGrid map[Point]bool

var (
	points  = make([]Point, 0)
	grid    = make(Grid)
	visited = make(VisitedGrid)
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: program <parts> <mode>")
	}

	part := os.Args[1]
	mode := os.Args[2]

	file, _ := os.Open(fmt.Sprintf("2024/12/%s.txt", mode))
	defer file.Close()

	parse(file)

	switch part {
	case "1":
		total := 0
		for _, point := range points {
			total += handlePointPart1(point)
		}
		fmt.Println(total)
	case "2":
		total := 0
		for _, point := range points {
			total += handlePointPart2(point)
		}
		fmt.Println(total)
	}
}

func handlePointPart1(point Point) int {
	if visited[point] {
		return 0
	}

	region := make([]Point, 0)
	region = append(region, point)
	region = findRegion(point, region)

	area := len(region)
	perimeter := calcPerimeter(region)

	return area * perimeter
}

func handlePointPart2(point Point) int {
	if visited[point] {
		return 0
	}

	region := make([]Point, 0)
	region = append(region, point)
	region = findRegion(point, region)

	area := len(region)
	sides := calcSides(region)

	return area * sides
}

func calcPerimeter(region []Point) int {
	perimeter := 0

	for _, point := range region {
		maxPerimeter := 4
		potentialNeighbours := getNeighbours(point, "nsew")

		for _, neighbour := range potentialNeighbours {
			if grid[neighbour] == grid[point] {
				maxPerimeter--
			}
		}

		perimeter += maxPerimeter
	}

	return perimeter
}

func calcSides(region []Point) int {
	corners := 0

	directions := []string{"nw", "se", "sw", "ne"}

	for _, point := range region {
		for _, direction := range directions {
			neighbours := getNeighbours(point, direction)

			if grid[neighbours[0]] == grid[point] && grid[neighbours[1]] == grid[point] {
				if grid[getDiagonal(point, direction)] != grid[point] {
					corners++
				}
			} else if grid[neighbours[0]] != grid[point] && grid[neighbours[1]] != grid[point] {
				corners++
			}
		}
	}

	return corners
}

func getDiagonal(point Point, direction string) Point {
	switch direction {
	case "nw":
		return Point{x: point.x - 1, y: point.y - 1}
	case "se":
		return Point{x: point.x + 1, y: point.y + 1}
	case "sw":
		return Point{x: point.x - 1, y: point.y + 1}
	case "ne":
		return Point{x: point.x + 1, y: point.y - 1}
	}

	return Point{}
}

func findRegion(startPoint Point, region []Point) []Point {
	visited[startPoint] = true

	neighbours := getNeighbours(startPoint, "nsew")

	for _, neighbour := range neighbours {
		if grid[neighbour] == grid[startPoint] && !visited[neighbour] {
			region = append(region, neighbour)
			region = findRegion(neighbour, region)
		}
	}

	return region
}

func getNeighbours(point Point, mode string) []Point {
	neighbours := make([]Point, 0)

	switch mode {
	case "nsew":
		neighbours = append(neighbours, Point{x: point.x, y: point.y - 1}) // North
		neighbours = append(neighbours, Point{x: point.x, y: point.y + 1}) // South
		neighbours = append(neighbours, Point{x: point.x + 1, y: point.y}) // East
		neighbours = append(neighbours, Point{x: point.x - 1, y: point.y}) // West
	case "ne":
		neighbours = append(neighbours, Point{x: point.x, y: point.y - 1}) // North
		neighbours = append(neighbours, Point{x: point.x + 1, y: point.y}) // East
	case "se":
		neighbours = append(neighbours, Point{x: point.x, y: point.y + 1}) // South
		neighbours = append(neighbours, Point{x: point.x + 1, y: point.y}) // East
	case "sw":
		neighbours = append(neighbours, Point{x: point.x, y: point.y + 1}) // South
		neighbours = append(neighbours, Point{x: point.x - 1, y: point.y}) // West
	case "nw":
		neighbours = append(neighbours, Point{x: point.x, y: point.y - 1}) // North
		neighbours = append(neighbours, Point{x: point.x - 1, y: point.y}) // West
	}

	return neighbours
}

func parse(file *os.File) (Grid, VisitedGrid) {
	scanner := bufio.NewScanner(file)

	yVal := 0

	for scanner.Scan() {
		line := scanner.Text()

		for i := 0; i < len(line); i++ {
			point := Point{x: i, y: yVal}
			grid[point] = string(line[i])
			visited[point] = false
			points = append(points, point)
		}

		yVal++
	}

	return grid, visited
}
