package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage (from advent root): go run 2025/01/main.go <part> (1 or 2) <mode> (real or test)")
	}

	part := os.Args[1]
	mode := os.Args[2]

	file, err := os.Open(fmt.Sprintf("2025/03/%s.txt", mode))
	if err != nil {
		log.Fatal(err)
	}
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
