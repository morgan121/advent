package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Robot struct {
	xPos, yPos           int
	xVelocity, yVelocity int
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: program <parts> <mode>")
	}

	part := os.Args[1]
	mode := os.Args[2]

	values := map[string][]int{
		"test": {11, 7},
		"real": {101, 103},
	}[mode]
	xMax, yMax := values[0], values[1]

	file, _ := os.Open(fmt.Sprintf("2024/14/%s.txt", mode))
	defer file.Close()

	robots := parse(file)

	switch part {
	case "1":
	case "2":
	}

	fmt.Println(xMax, yMax, robots)
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

func parse(file *os.File) []Robot {
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`\d+`)

	robots := make([]Robot, 0)
	robot := Robot{}

	for scanner.Scan() {
		line := scanner.Text()
		values := arrayToInt(re.FindAllString(line, -1))

		robot.xPos = values[0]
		robot.yPos = values[1]
		robot.xVelocity = values[2]
		robot.yVelocity = values[3]

		robots = append(robots, robot)
	}

	return robots
}
