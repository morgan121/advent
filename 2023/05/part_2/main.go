package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var (
	numberRegExp    = regexp.MustCompile("[0-9]+")
	linebreakRegExp = regexp.MustCompile(`\r?\n`)
	chunkRegExp     = regexp.MustCompile("seeds: |seed-to-soil map:|soil-to-fertilizer map:|fertilizer-to-water map:|water-to-light map:|light-to-temperature map:|temperature-to-humidity map:|humidity-to-location map:")
)

type PuzzleMap struct {
	destinationStart int
	sourceStart      int
	rangeLength      int
}

type SeedRange struct {
	start  int
	length int
}

func main() {
	file := readFile("2023/05/input.txt")
	chunks := chunkRegExp.Split(string(file), -1)

	output := 0

	var seeds []SeedRange
	var seedToSoil []PuzzleMap
	var soilToFertilizer []PuzzleMap
	var fertilizerToWater []PuzzleMap
	var waterToLight []PuzzleMap
	var lightToTemp []PuzzleMap
	var tempToHumidity []PuzzleMap
	var humidityToLocation []PuzzleMap

	for i, chunk := range chunks {
		allNumbers := arrayToInt(numberRegExp.FindAllString(chunk, -1))

		switch i {
		case 1:
			for n := 0; n < len(allNumbers); n += 2 {
				seed := SeedRange{
					start:  allNumbers[n],
					length: allNumbers[n+1],
				}
				seeds = append(seeds, seed)
			}
		case 2:
			seedToSoil = parseChunk(allNumbers)
		case 3:
			soilToFertilizer = parseChunk(allNumbers)
		case 4:
			fertilizerToWater = parseChunk(allNumbers)
		case 5:
			waterToLight = parseChunk(allNumbers)
		case 6:
			lightToTemp = parseChunk(allNumbers)
		case 7:
			tempToHumidity = parseChunk(allNumbers)
		case 8:
			humidityToLocation = parseChunk(allNumbers)
		default:
		}
	}

	for _, seed := range seeds {
		for j := 0; j < seed.length; j++ {
			soil := sourceToDestination(seedToSoil, seed.start+j)
			fertilizer := sourceToDestination(soilToFertilizer, soil)
			water := sourceToDestination(fertilizerToWater, fertilizer)
			light := sourceToDestination(waterToLight, water)
			temp := sourceToDestination(lightToTemp, light)
			humidity := sourceToDestination(tempToHumidity, temp)
			location := sourceToDestination(humidityToLocation, humidity)

			if output == 0 {
				output = location
			} else {
				output = min(output, location)
			}
		}
	}

	fmt.Println(output)
}

func parseChunk(numbers []int) []PuzzleMap {
	var puzzleMap []PuzzleMap

	for e := 0; e < len(numbers); e += 3 {
		row := PuzzleMap{
			destinationStart: numbers[e],
			sourceStart:      numbers[e+1],
			rangeLength:      numbers[e+2],
		}
		puzzleMap = append(puzzleMap, row)
	}

	return puzzleMap
}

func sourceToDestination(mapper []PuzzleMap, sourceValue int) int {
	for _, info := range mapper {
		start := info.sourceStart
		end := info.sourceStart + info.rangeLength

		if sourceValue >= start && sourceValue < end {
			return (sourceValue - info.sourceStart) + info.destinationStart
		}
	}

	return sourceValue
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

func readFile(filepath string) []byte {
	file, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
