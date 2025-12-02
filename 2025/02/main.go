package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage (from advent root): go run 2025/01/main.go <parts> (1 or 2) <mode> (real or test)")
	}

	part := os.Args[1]
	mode := os.Args[2]

	file, err := os.Open(fmt.Sprintf("2025/02/%s.txt", mode))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var invalidSum int

	for scanner.Scan() {
		line := scanner.Text()
		ranges := strings.Split(line, ",")

		for _, r := range ranges {
			start := toInt(strings.Split(r, "-")[0])
			end := toInt(strings.Split(r, "-")[1])

			for v := start; v <= end; v++ {
				if isInvalid(strconv.Itoa(v), part) {
					invalidSum += v
				}
			}
		}
	}

	fmt.Println(invalidSum)
}

func isInvalid(v string, part string) bool {
	n := len(v)

	switch part {
	case "1":
		if n%2 != 0 {
			return false
		}

		half := n / 2
		return v[:half] == v[half:]
	case "2":
		ss := (v + v)[1 : 2*n-1]
		return strings.Contains(ss, v)
	}

	return false
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}

	return n
}
