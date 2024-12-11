package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Point struct {
	x, y int
}

type Trail struct {
	start Point
	end   Point
}

type Grid map[Point]int

var (
	grid            = make(Grid)
	trailHeads      = make(map[Point]int)
	completedTrails = make(map[Trail]bool)
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: program <parts> <mode>")
	}

	part := os.Args[1]
	mode := os.Args[2]

	file, _ := os.Open(fmt.Sprintf("2024/10/%s.txt", mode))
	defer file.Close()

	parse(file)

	for point := range trailHeads {
		walkTrails(point, point)
	}

	switch part {
	case "1":
		fmt.Println(len(completedTrails))
	case "2":
		total := 0
		for _, val := range trailHeads {
			total += val
		}
		fmt.Println(total)
	}
}

func walkTrails(startPoint Point, originPoint Point) {
	value := grid[startPoint]

	if value == 9 {
		completedTrails[Trail{start: originPoint, end: startPoint}] = true
		trailHeads[originPoint]++
		return
	}

	if grid[Point{x: startPoint.x, y: startPoint.y - 1}] == value+1 {
		walkTrails(Point{x: startPoint.x, y: startPoint.y - 1}, originPoint)
	}

	if grid[Point{x: startPoint.x, y: startPoint.y + 1}] == value+1 {
		walkTrails(Point{x: startPoint.x, y: startPoint.y + 1}, originPoint)
	}

	if grid[Point{x: startPoint.x - 1, y: startPoint.y}] == value+1 {
		walkTrails(Point{x: startPoint.x - 1, y: startPoint.y}, originPoint)
	}

	if grid[Point{x: startPoint.x + 1, y: startPoint.y}] == value+1 {
		walkTrails(Point{x: startPoint.x + 1, y: startPoint.y}, originPoint)
	}
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}

	return n
}

func parse(file *os.File) {
	scanner := bufio.NewScanner(file)

	yVal := 0

	for scanner.Scan() {
		line := scanner.Text()

		for i := 0; i < len(line); i++ {
			point := Point{x: i, y: yVal}
			grid[point] = toInt(string(line[i]))

			if grid[point] == 0 {
				trailHeads[point] = 0
			}
		}

		yVal++
	}
}
