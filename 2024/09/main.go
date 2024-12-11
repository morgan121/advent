package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

type Block struct {
	value    int
	startIdx int
	length   int
}

var (
	blockLengths []int
	gapLengths   []int
	blocks       []Block
	gaps         []Block
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: program <parts> <mode>")
	}

	part := os.Args[1]
	mode := os.Args[2]

	file, _ := os.Open(fmt.Sprintf("2024/09/%s.txt", mode))
	defer file.Close()

	blockLengths, gapLengths = parse(file)

	switch part {
	case "1":
		calculate(condense(translate(blockLengths, gapLengths)))
	case "2":
		calculate(condenseByBlock(translate(blockLengths, gapLengths))) // 8317688930597 is too high
	}
}

func calculate(input []int) {
	total := 0

	for i, value := range input {
		if value != -1 {
			total += value * i
		}
	}

	fmt.Println(total)
}

func translate(blockLengths []int, gapLengths []int) []int {
	var disc []int

	blocks = make([]Block, 0)
	gaps = make([]Block, 0)

	initialBlock := makeBlock(0, blockLengths[0])
	disc = append(disc, initialBlock...)
	blocks = append(blocks, Block{length: len(initialBlock), value: 0, startIdx: 0})

	for i := 1; i < len(blockLengths); i++ {
		gap := makeBlock(-1, gapLengths[i-1])
		block := makeBlock(i, blockLengths[i])

		gaps = append(gaps, Block{value: -1, startIdx: len(disc), length: gapLengths[i-1]})
		disc = append(disc, gap...)

		blocks = append(blocks, Block{value: i, startIdx: len(disc), length: blockLengths[i]})
		disc = append(disc, block...)
	}

	return disc
}

func condense(disc []int) []int {
	firstGapIdx := slices.Index(disc, -1)
	lastBlockIdx := lastNonZeroIndex(disc)

	if firstGapIdx > lastBlockIdx {
		return disc
	} else {
		disc[firstGapIdx] = disc[lastBlockIdx]
		disc[lastBlockIdx] = -1
		condense(disc)
	}

	return disc
}

func condenseByBlock(disc []int) []int {
	for i := len(blocks) - 1; i >= 0; i-- {
		block := blocks[i]
		firstAvailableGapIdx := slices.IndexFunc(gaps, func(gap Block) bool {
			return gap.length >= block.length
		})

		if firstAvailableGapIdx == -1 {
			continue
		} else {
			gap := gaps[firstAvailableGapIdx]

			for g := gap.startIdx; g < gap.startIdx+block.length; g++ {
				disc[g] = block.value
			}

			for b := block.startIdx; b < block.startIdx+block.length; b++ {
				disc[b] = -1
			}

			gap.startIdx += block.length
			gap.length -= block.length
			gaps[firstAvailableGapIdx] = gap
		}
	}

	return disc
}

func lastNonZeroIndex(input []int) int {
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] != -1 {
			return i
		}
	}

	return -1
}

func makeBlock(value int, length int) []int {
	block := make([]int, 0)
	for i := 0; i < length; i++ {
		block = append(block, value)
	}

	return block
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

func parse(file *os.File) ([]int, []int) {
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`\d`)

	var blockLengths []int
	var gapLengths []int

	for scanner.Scan() {
		line := scanner.Text()
		numbers := arrayToInt(re.FindAllString(line, -1))

		for i, char := range numbers {
			if i%2 == 0 {
				blockLengths = append(blockLengths, char)
			} else {
				gapLengths = append(gapLengths, char)
			}
		}
	}

	return blockLengths, gapLengths
}
