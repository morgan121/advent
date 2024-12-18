package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Button struct {
	x, y int
}

type Prize struct {
	x, y int
}

type Game struct {
	buttonA Button
	buttonB Button
	prize   Prize
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: program <parts> <mode>")
	}

	part := os.Args[1]
	mode := os.Args[2]

	file, _ := os.Open(fmt.Sprintf("2024/13/%s.txt", mode))
	defer file.Close()

	games := parse(file)

	total := 0

	switch part {
	case "1":
		for _, game := range games {
			total += calculatePart1(game.buttonA.x, game.buttonA.y, game.buttonB.x, game.buttonB.y, game.prize.x, game.prize.y)
		}
	case "2":
		for _, game := range games {
			total += calculatePart2(game.buttonA.x, game.buttonA.y, game.buttonB.x, game.buttonB.y, game.prize.x, game.prize.y)
		}
	}

	fmt.Println(total)
}

func calculatePart1(a int, b int, c int, d int, x int, y int) int {
	A, B := calculate(a, b, c, d, x, y)

	if A > 100 || B > 100 {
		return 0
	}

	return A*3 + B
}

func calculatePart2(a int, b int, c int, d int, x int, y int) int {
	A, B := calculate(a, b, c, d, x+10000000000000, y+10000000000000)

	return A*3 + B
}

func calculate(a int, b int, c int, d int, x int, y int) (int, int) {
	A := ((d * x) + (-c * y)) / (a*d - b*c)
	B := ((-b * x) + (a * y)) / (a*d - b*c)

	if A*a+B*c != x || A*b+B*d != y {
		return 0, 0
	}

	return A, B
}

func arrayToInt(s []string) []int {
	var intArray []int

	for _, e := range s {
		intArray = append(intArray, toInt(e))
	}

	return intArray
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}

	return n
}

func parse(file *os.File) []Game {
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`\d+`)

	step := 0
	games := make([]Game, 0)
	game := Game{}

	for scanner.Scan() {
		line := scanner.Text()
		coords := arrayToInt(re.FindAllString(line, -1))

		if len(coords) == 0 {
			step = 0
		} else {
			switch step {
			case 0:
				game.buttonA.x = coords[0]
				game.buttonA.y = coords[1]
			case 1:
				game.buttonB.x = coords[0]
				game.buttonB.y = coords[1]
			case 2:
				game.prize.x = coords[0]
				game.prize.y = coords[1]
				games = append(games, game)
			}

			step++
		}
	}

	return games
}
