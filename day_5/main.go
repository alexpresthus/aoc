package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func stringsToIntegers(s []string) ([]int, error) {
	ints := make([]int, 0, len(s))
	for _, line := range s {
		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		ints = append(ints, n)
	}
	return ints, nil
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func solvePartOne() int {
	f, err := os.Open("/tmp/aoc/input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	i := -1
	sources := make([]int, 0, 20)
	destinations := make([]int, 0, 20)
	for scanner.Scan() {
		i++
		ln := scanner.Text()
		// handle first line of sources (seeds) - handle as already mapped to destinations
		if i == 0 {
			for _, seedStr := range strings.Fields(ln)[1:] {
				seedInt, err := strconv.Atoi(seedStr)
				check(err)
				destinations = append(destinations, seedInt)
			}
			continue
		}
		// handle end of map
		if len(ln) == 0 {
			// destinations -> sources for next map
			// sources that have not been mapped to dest keeps on as sources with the same values
			for _, dest := range destinations {
				sources = append(sources, dest)
			}
			destinations = destinations[:0]
			continue
		}
		// skip titles (non digit)
		if unicode.IsDigit(rune(ln[0])) != true {
			continue
		}
		mapValues, err := stringsToIntegers(strings.Fields(ln))
		check(err)
		remainingSources := make([]int, 0, len(sources))
		for _, src := range sources {
			if src >= mapValues[1] && src < mapValues[1]+mapValues[2] {
				destinations = append(destinations, mapValues[0]+src-mapValues[1])
			} else {
				remainingSources = append(remainingSources, src)
			}
		}
		sources = remainingSources
	}
	// final remaining sources -> destinations
	for _, src := range sources {
		destinations = append(destinations, src)
	}
	return slices.Min(destinations)
}

func solvePartTwo() int {
	f, err := os.Open("/tmp/aoc/input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	i := -1
	sourceRanges := make([][]int, 0, 20)
	destinationRanges := make([][]int, 0, 20)
	for scanner.Scan() {
		i++
		ln := scanner.Text()
		// handle first line of sources (seeds) - handle as already mapped to destinations
		if i == 0 {
			seedRanges := strings.Fields(ln)[1:]
			for x := 0; x < len(seedRanges); x += 2 {
				seedInt, err1 := strconv.Atoi(seedRanges[x])
				check(err1)
				rangeInt, err2 := strconv.Atoi(seedRanges[x+1])
				check(err2)
				destinationRanges = append(destinationRanges, []int{seedInt, rangeInt})
			}
			continue
		}
		// handle end of map
		if len(ln) == 0 {
			// destinations -> sources for next map
			// sources that have not been mapped to dest keeps on as sources with the same values
			for _, destRange := range destinationRanges {
				sourceRanges = append(sourceRanges, destRange)
			}
			destinationRanges = destinationRanges[:0]
			continue
		}
		// skip titles (non digit)
		if unicode.IsDigit(rune(ln[0])) != true {
			continue
		}
		mapValues, err := stringsToIntegers(strings.Fields(ln))
		check(err)
		remainingSourceRanges := make([][]int, 0, len(sourceRanges))
		for _, srcRange := range sourceRanges {
			src := srcRange[0]
			lng := srcRange[1]

			if src >= mapValues[1] && src < mapValues[1]+mapValues[2] {
				// source start is within current destination range
				if src+lng <= mapValues[1]+mapValues[2] {
					// entire source range is within the current destination range
					destinationRanges = append(destinationRanges, []int{mapValues[0] + src - mapValues[1], lng})
				} else {
					// part of source range is within current destination range - add part to dest and rest to remaining
					overflow := (src + lng) - (mapValues[1] + mapValues[2])
					destinationRanges = append(destinationRanges, []int{mapValues[0] + src - mapValues[1], lng - overflow})
					remainingSourceRanges = append(remainingSourceRanges, []int{src + lng - overflow, overflow})
				}
			} else if src+lng >= mapValues[1] && src+lng < mapValues[1]+mapValues[2] {
				// source ends within current dest range, e.g. some numbers of the sourcerange are in range of destination range
				srcStart := mapValues[1]
				srcLng := lng - (srcStart - src)
				if srcLng > mapValues[2] {
					overflow := srcLng - mapValues[2]
					srcLng = mapValues[2]
					remainingSourceRanges = append(remainingSourceRanges, []int{srcStart + srcLng, overflow})
				}
				destinationRanges = append(destinationRanges, []int{mapValues[0] + srcStart - mapValues[1], srcLng})
				remainingSourceRanges = append(remainingSourceRanges, []int{src, srcStart - src})
			} else if src < mapValues[1] && src+lng >= mapValues[1]+mapValues[2] {
				// not start or end, but overlaps
				srcStart := mapValues[1]
				srcLng := mapValues[2]
				destinationRanges = append(destinationRanges, []int{mapValues[0] + srcStart - mapValues[1], srcLng})
				remainingSourceRanges = append(remainingSourceRanges, []int{src, srcStart - src}, []int{srcStart + srcLng, src + lng - (srcStart + srcLng)})
			} else {
				// not in range
				remainingSourceRanges = append(remainingSourceRanges, srcRange)
			}
		}
		sourceRanges = remainingSourceRanges
	}
	// final remaining sources -> destinations
	for _, srcRange := range sourceRanges {
		destinationRanges = append(destinationRanges, srcRange)
	}
	startingDestinations := make([]int, 0, len(destinationRanges))
	for _, destRange := range destinationRanges {
		startingDestinations = append(startingDestinations, destRange[0])
	}
	return slices.Min(startingDestinations)
}

func main() {
	fmt.Printf("Part 1 - lowest location number: %d\n", solvePartOne())
	fmt.Printf("Part 2 - lowest location number: %d\n", solvePartTwo())
}
