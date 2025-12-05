package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime"
	"sort"
	"strconv"
)

type Range struct {
	start, end int
}

func main() {
	file, part := setup()

	defer file.Close()

	re := regexp.MustCompile(`[0-9]+`)
	scanner := bufio.NewScanner(file)

	freshDates := make([]Range, 0)
	dates := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		values := re.FindAllString(line, -1)
		if len(values) == 2 {
			if fullyBetween(Range{toInt(values[0]), toInt(values[1])}, freshDates) {
				continue
			}
			freshDates = append(freshDates, Range{toInt(values[0]), toInt(values[1])})
		} else if len(values) == 1 {
			dates = append(dates, toInt(values[0]))
		}
	}

	freshDates = consolidate(freshDates)

	total := 0

	switch part {
	case "1":
		for _, date := range dates {
			fresh, _ := between(date, freshDates)
			if fresh {
				total++
			}
		}
		fmt.Println(total)
	case "2":
		for _, fd := range freshDates {
			total += fd.end - fd.start + 1
		}
		fmt.Println(total)
	}

}

func consolidate(ranges []Range) []Range {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})
	incorporatedIndexes := make([]int, 0)
	for i, fd := range ranges {
		between, index := between(fd.start, ranges)
		if between && index != i {
			incorporatedIndexes = append(incorporatedIndexes, i)
			ranges[index].end = max(ranges[index].end, fd.end)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(incorporatedIndexes)))
	for _, i := range incorporatedIndexes {
		ranges = append(ranges[:i], ranges[i+1:]...)
	}

	return ranges
}

func fullyBetween(r Range, rs []Range) bool {
	for _, rr := range rs {
		if r.start >= rr.start && r.end <= rr.end {
			return true
		}
	}
	return false
}

func between(r int, rs []Range) (bool, int) {
	for i, rr := range rs {
		if r >= rr.start && r <= rr.end {
			return true, i
		}
	}
	return false, -1
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
