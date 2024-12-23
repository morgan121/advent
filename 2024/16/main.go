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

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: program <parts> <mode>")
	}

	part := os.Args[1]
	mode := os.Args[2]

	file, _ := os.Open(fmt.Sprintf("2024/16/%s.txt", mode))
	defer file.Close()

	grid, startPoint, endPoint := parse(file)

	switch part {
	case "1":
		fmt.Println(startPoint, grid[startPoint], endPoint, grid[endPoint])
	case "2":
	}
}

func parse(file *os.File) (Grid, Point, Point) {
	scanner := bufio.NewScanner(file)

	yVal := 0
	grid := make(Grid)
	startPoint := Point{}
	endPoint := Point{}

	for scanner.Scan() {
		line := scanner.Text()

		for i := 0; i < len(line); i++ {
			point := Point{x: i, y: yVal}
			grid[point] = string(line[i])

			if grid[point] == "S" {
				startPoint = point
			}
			if grid[point] == "E" {
				endPoint = point
			}
		}

		yVal++
	}

	return grid, startPoint, endPoint
}
