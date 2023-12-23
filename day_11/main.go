package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Galaxy struct {
	id int
	y  int
	x  int
}

func solve(content string, expansionRate int) (int, time.Duration) {
	startTime := time.Now()
	galaxies := make([]Galaxy, 0)
	shortestPathsSum := 0
	// populate galaxies
	id := 1
	for y, row := range strings.Split(content, "\n") {
		for x, r := range row {
			if r == '#' {
				galaxies = append(galaxies, Galaxy{id, y, x})
				id++
			}
		}
	}

	// expand y positions
	expandY := 0
	for i, g := range galaxies {
		galaxies[i].y += expandY
		if i == len(galaxies)-1 {
			continue
		}
		nextG := galaxies[i+1]
		if nextG.y-g.y > 1 {
			expandY += (expansionRate - 1) * (nextG.y - g.y - 1)
		}
	}
	// sort galaxies by x
	slices.SortFunc(galaxies, func(a, b Galaxy) int { return a.x - b.x })
	// expand x positions
	expandX := 0
	for i, g := range galaxies {
		galaxies[i].x += expandX
		if i == len(galaxies)-1 {
			continue
		}
		nextG := galaxies[i+1]
		if nextG.x-g.x > 1 {
			expandX += (expansionRate - 1) * (nextG.x - g.x - 1)
		}
	}
	slices.SortFunc(galaxies, func(a, b Galaxy) int { return a.id - b.id })

	// find shortest paths for each pair
	for i, startG := range galaxies {
		for _, endG := range galaxies[i:] {
			stepsX := Abs(endG.x - startG.x)
			stepsY := Abs(endG.y - startG.y)
			shortestPathsSum += stepsX + stepsY
		}
	}
	return shortestPathsSum, time.Since(startTime)
}

func main() {
	f, err := os.ReadFile("/tmp/aoc/input.txt")
	content := string(f)
	check(err)
	ansPartOne, durationPartOne := solve(content, 2)
	fmt.Printf("Part 1: %d	Duration: %v\n", ansPartOne, durationPartOne)
	ansPartTwo, durationPartTwo := solve(content, 1000000)
	fmt.Printf("Part 2: %d	Duration: %v\n", ansPartTwo, durationPartTwo)
}
