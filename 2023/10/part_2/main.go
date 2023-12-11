package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strings"
)

type Point struct {
	x, y int
}

type Border struct {
	xmin, ymin, xmax, ymax int
}

var (
	lineBreakRegExp = regexp.MustCompile(`\r?\n`)
	grid            = make(map[Point]string)
	mainPipe        []Point
	initialPoint    Point
	border          Border
)

func main() {
	file := readFile("2023/10/input.txt")
	lines := lineBreakRegExp.Split(string(file), -1)

	border = Border{
		xmin: 0,
		ymin: 0,
		xmax: len(lines[0]),
		ymax: len(lines),
	}

	// fill grid
	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			point := Point{
				x: x,
				y: y,
			}

			if line[x:x+1] == "S" {
				initialPoint = point
				mainPipe = append(mainPipe, initialPoint)
			}

			grid[point] = line[x : x+1]
		}
	}

	getMainPipe(initialPoint)

	totalInside := 0
	for point := range grid {
		if inMainPipe(point) {
			continue
		} else if point.x == border.xmin || point.x == border.xmax || point.y == border.ymin || point.y == border.ymax {
			continue
		} else {
			var testPoint Point
			var charsToLeft []string

			for x := 0; x < point.x; x++ {
				testPoint = Point{
					x: x,
					y: point.y,
				}
				if inMainPipe(testPoint) && strings.Contains("S7FJL|", character(testPoint)) {
					charsToLeft = append(charsToLeft, character(testPoint))
				}
			}

			charString := strings.Join(charsToLeft, "")
			ok := true
			for ok == true {
				if strings.Index(charString, "L7") > -1 {
					charString = remove(charString, strings.Index(charString, "L7"))
				} else {
					ok = false
				}
			}

			ok = true
			for ok == true {
				if strings.Index(charString, "FJ") > -1 {
					charString = remove(charString, strings.Index(charString, "FJ"))
				} else {
					ok = false
				}
			}

			if len(charString)%2 == 1 {
				totalInside++
			}
		}
	}

	fmt.Println(totalInside)
}

func remove(s string, i int) string {
	return strings.Join([]string{s[:i], s[i+1:]}, "")
}

func inMainPipe(p Point) bool {
	return slices.Contains(mainPipe, p)
}

func getMainPipe(initialPoint Point) {
	prevPoint := initialPoint
	ok := true
	for ok {
		var nextPoint Point

		switch character(prevPoint) {
		case "S":
			if strings.Contains("7F|", character(goNorth(prevPoint))) && prevPoint.y > border.ymin {
				nextPoint = goNorth(prevPoint)
				mainPipe = append(mainPipe, nextPoint)
				prevPoint = nextPoint
			} else if strings.Contains("7J-", character(goEast(prevPoint))) && prevPoint.x < border.xmax {
				nextPoint = goEast(prevPoint)
				mainPipe = append(mainPipe, nextPoint)
				prevPoint = nextPoint
			} else if strings.Contains("JL|", character(goSouth(prevPoint))) && prevPoint.y < border.ymax {
				nextPoint = goSouth(prevPoint)
				mainPipe = append(mainPipe, nextPoint)
				prevPoint = nextPoint
			} else if strings.Contains("FL-", character(goWest(prevPoint))) && prevPoint.x > border.xmin {
				nextPoint = goWest(prevPoint)
				mainPipe = append(mainPipe, nextPoint)
				prevPoint = nextPoint
			} else {
				ok = false
				break
			}
		case "L":
			if strings.Contains("7F|", character(goNorth(prevPoint))) && !slices.Contains(mainPipe, goNorth(prevPoint)) && prevPoint.y != border.ymin {
				nextPoint = goNorth(prevPoint)
				mainPipe = append(mainPipe, nextPoint)
				prevPoint = nextPoint
			} else if strings.Contains("7J-", character(goEast(prevPoint))) && !slices.Contains(mainPipe, goEast(prevPoint)) && prevPoint.x != border.xmax {
				nextPoint = goEast(prevPoint)
				mainPipe = append(mainPipe, nextPoint)
				prevPoint = nextPoint
			} else {
				ok = false
				break
			}
		case "F":
			if strings.Contains("JL|", character(goSouth(prevPoint))) && !slices.Contains(mainPipe, goSouth(prevPoint)) && prevPoint.y != border.ymax {
				nextPoint = goSouth(prevPoint)
				mainPipe = append(mainPipe, nextPoint)
				prevPoint = nextPoint
			} else if strings.Contains("7J-", character(goEast(prevPoint))) && !slices.Contains(mainPipe, goEast(prevPoint)) && prevPoint.x != border.xmax {
				nextPoint = goEast(prevPoint)
				mainPipe = append(mainPipe, nextPoint)
				prevPoint = nextPoint
			} else {
				ok = false
				break
			}
		case "J":
			if strings.Contains("7F|", character(goNorth(prevPoint))) && !slices.Contains(mainPipe, goNorth(prevPoint)) && prevPoint.y != border.ymin {
				nextPoint = goNorth(prevPoint)
				mainPipe = append(mainPipe, nextPoint)
				prevPoint = nextPoint
			} else if strings.Contains("FL-", character(goWest(prevPoint))) && !slices.Contains(mainPipe, goWest(prevPoint)) && prevPoint.x != border.xmin {
				nextPoint = goWest(prevPoint)
				mainPipe = append(mainPipe, nextPoint)
				prevPoint = nextPoint
			} else {
				ok = false
				break
			}
		case "7":
			if strings.Contains("JL|", character(goSouth(prevPoint))) && !slices.Contains(mainPipe, goSouth(prevPoint)) && prevPoint.y != border.ymax {
				nextPoint = goSouth(prevPoint)
				mainPipe = append(mainPipe, nextPoint)
				prevPoint = nextPoint
			} else if strings.Contains("FL-", character(goWest(prevPoint))) && !slices.Contains(mainPipe, goWest(prevPoint)) && prevPoint.x != border.xmin {
				nextPoint = goWest(prevPoint)
				mainPipe = append(mainPipe, nextPoint)
				prevPoint = nextPoint
			} else {
				ok = false
				break
			}
		case "|":
			if strings.Contains("JL|", character(goSouth(prevPoint))) && !slices.Contains(mainPipe, goSouth(prevPoint)) && prevPoint.y != border.ymax {
				nextPoint = goSouth(prevPoint)
				mainPipe = append(mainPipe, nextPoint)
				prevPoint = nextPoint
			} else if strings.Contains("7F|", character(goNorth(prevPoint))) && !slices.Contains(mainPipe, goNorth(prevPoint)) && prevPoint.y != border.ymin {
				nextPoint = goNorth(prevPoint)
				mainPipe = append(mainPipe, nextPoint)
				prevPoint = nextPoint
			} else {
				ok = false
				break
			}
		case "-":
			if strings.Contains("7J-", character(goEast(prevPoint))) && !slices.Contains(mainPipe, goEast(prevPoint)) && prevPoint.x != border.xmax {
				nextPoint = goEast(prevPoint)
				mainPipe = append(mainPipe, nextPoint)
				prevPoint = nextPoint
			} else if strings.Contains("FL-", character(goWest(prevPoint))) && !slices.Contains(mainPipe, goWest(prevPoint)) && prevPoint.x != border.xmin {
				nextPoint = goWest(prevPoint)
				mainPipe = append(mainPipe, nextPoint)
				prevPoint = nextPoint
			} else {
				ok = false
				break
			}
		case ".":
			ok = false
			break
		}
	}
}

func character(p Point) string {
	return grid[p]
}

func goNorth(p Point) Point {
	return Point{
		x: p.x,
		y: p.y - 1,
	}
}

func goEast(p Point) Point {
	return Point{
		x: p.x + 1,
		y: p.y,
	}
}

func goSouth(p Point) Point {
	return Point{
		x: p.x,
		y: p.y + 1,
	}
}

func goWest(p Point) Point {
	return Point{
		x: p.x - 1,
		y: p.y,
	}
}

func readFile(filepath string) []byte {
	file, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
