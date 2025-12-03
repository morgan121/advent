package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"runtime"
	"slices"
	"strconv"
)

func main() {
	file, part := setup()

	defer file.Close()

	re := regexp.MustCompile(`[0-9]`)
	scanner := bufio.NewScanner(file)

	var iterations int

	switch part {
	case "1":
		iterations = 2
	case "2":
		iterations = 12
	}

	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		bank := re.FindAllString(line, -1)

		start := 0
		for i := iterations - 1; i >= 0; i-- {
			val, ind := maxIndex(bank[start : len(bank)-i])
			start += ind + 1
			total += val * int(math.Pow(10, float64(i)))
		}
	}

	fmt.Println(total)
}

func maxIndex(s []string) (int, int) {
	v := slices.Max(s)
	i := slices.Index(s, v)
	return toInt(v), i
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}

	return n
}

func setup() (*os.File, string) {
	if len(os.Args) < 3 {
		log.Fatal("Usage (from advent root): go run 2025/01/main.go <part> (1 or 2) <mode> (real or test)")
	}

	_, filename, _, _ := runtime.Caller(0)
	part := os.Args[1]
	mode := os.Args[2]

	re := regexp.MustCompile(`[0-9]+`)
	paths := re.FindAllString(filename, -1)

	file, err := os.Open(fmt.Sprintf("%s/%s/%s.txt", paths[0], paths[1], mode))
	if err != nil {
		log.Fatal(err)
	}

	return file, part
}
