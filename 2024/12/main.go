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
			total += handlePoint(point)
		}

		fmt.Println(total)
	case "2":
	}
}

func handlePoint(point Point) int {
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

func findRegion(startPoint Point, region []Point) []Point {
	visited[startPoint] = true

	neighbours := make([]Point, 0)
	neighbours = append(neighbours, Point{x: startPoint.x, y: startPoint.y - 1})
	neighbours = append(neighbours, Point{x: startPoint.x, y: startPoint.y + 1})
	neighbours = append(neighbours, Point{x: startPoint.x - 1, y: startPoint.y})
	neighbours = append(neighbours, Point{x: startPoint.x + 1, y: startPoint.y})

	for _, neighbour := range neighbours {
		if grid[neighbour] == grid[startPoint] && !visited[neighbour] {
			region = append(region, neighbour)
			region = findRegion(neighbour, region)
		}
	}

	return region
}

func calcPerimeter(region []Point) int {
	perimeter := 0

	for _, point := range region {
		maxPerimeter := 4

		potentialNeighbours := make([]Point, 0)
		potentialNeighbours = append(potentialNeighbours, Point{x: point.x, y: point.y - 1})
		potentialNeighbours = append(potentialNeighbours, Point{x: point.x, y: point.y + 1})
		potentialNeighbours = append(potentialNeighbours, Point{x: point.x - 1, y: point.y})
		potentialNeighbours = append(potentialNeighbours, Point{x: point.x + 1, y: point.y})

		for _, neighbour := range potentialNeighbours {
			if grid[neighbour] == grid[point] {
				maxPerimeter--
			}
		}

		perimeter += maxPerimeter
	}

	return perimeter
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
