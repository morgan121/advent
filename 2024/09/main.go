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
	length   int
	startIdx int
}

var (
	blockLength []int
	gapLength   []int
	blocks      []Block
	gaps        []Block
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: program <parts> <mode>")
	}

	part := os.Args[1]
	mode := os.Args[2]

	file, _ := os.Open(fmt.Sprintf("2024/09/%s.txt", mode))
	defer file.Close()

	blockLength, gapLength = parse(file)

	switch part {
	case "1":
		calculatePart1()
	case "2":
		calculatePart2() // 8317688930597 is too high
	}

}

func calculatePart1() {
	disc := condense(translate(blockLength, gapLength))
	total := 0

	for i, value := range disc {
		if value != -1 {
			total += value * i
		}
	}

	fmt.Println(total)
}

func calculatePart2() {
	disc := condenseByBlock(translate(blockLength, gapLength))
	total := 0

	for i, value := range disc {
		if value != -1 {
			total += value * i
		}
	}

	fmt.Println(total)
}

func translate(blockLength []int, gapLength []int) []int {
	var translated []int

	blocks = make([]Block, 0)
	gaps = make([]Block, 0)

	initialBlock := makeBlock(0, blockLength[0])
	translated = append(translated, initialBlock...)
	blocks = append(blocks, Block{length: len(initialBlock), value: 0, startIdx: 0})
	for i := 1; i < len(blockLength); i++ {
		gap := makeBlock(-1, gapLength[i-1])
		block := makeBlock(i, blockLength[i])
		gaps = append(gaps, Block{length: gapLength[i-1], value: -1, startIdx: len(translated)})
		translated = append(translated, gap...)
		blocks = append(blocks, Block{length: blockLength[i], value: i, startIdx: len(translated)})
		translated = append(translated, block...)
	}

	return translated
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

func lastNonZeroIndex(input []int) int {
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] != -1 {
			return i
		}
	}

	return -1
}

func condenseByBlock(disc []int) []int {
	for i := len(blocks) - 1; i >= 0; i-- {

		block := blocks[i]
		firstAppropriateGapIdx := slices.IndexFunc(gaps, func(gap Block) bool {
			return gap.length >= block.length
		})

		if firstAppropriateGapIdx == -1 {
			continue
		} else {
			gap := gaps[firstAppropriateGapIdx]

			for g := gap.startIdx; g < gap.startIdx+block.length; g++ {
				disc[g] = block.value
			}

			for b := block.startIdx; b < block.startIdx+block.length; b++ {
				disc[b] = -1
			}

			gap.startIdx += block.length
			gap.length -= block.length
			gaps[firstAppropriateGapIdx] = gap
		}
	}

	return disc
}

func makeBlock(value int, length int) []int {
	block := make([]int, 0)
	for i := 0; i < length; i++ {
		block = append(block, value)
	}

	return block
}

func parse(file *os.File) ([]int, []int) {
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`\d`)

	var blockLength []int
	var gapLength []int

	for scanner.Scan() {
		line := scanner.Text()
		numbers := arrayToInt(re.FindAllString(line, -1))

		for i, char := range numbers {
			if i%2 == 0 {
				blockLength = append(blockLength, char)
			} else {
				gapLength = append(gapLength, char)
			}
		}
	}

	return blockLength, gapLength
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
