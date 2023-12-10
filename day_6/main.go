package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	content, err := os.ReadFile("/tmp/aoc/input.txt")
	check(err)
	times := strings.Fields(strings.Split(string(content), "\n")[0])[1:]
	distances := strings.Fields(strings.Split(string(content), "\n")[1])[1:]
	races := make([][]int, len(times))
	for i := 0; i < len(races); i++ {
		t, err1 := strconv.Atoi(times[i])
		check(err1)
		d, err2 := strconv.Atoi(distances[i])
		check(err2)
		races[i] = []int{t, d}
	}
	product := 0
	for _, r := range races {
		t := r[0]
		bestDist := r[1]
		numWaysToBeat := 0
		for x := 1; x < t; x++ {
			dist := (t - x) * x
			if dist > bestDist {
				numWaysToBeat++
			}
		}
		if product == 0 {
			product = numWaysToBeat
		} else {
			product = product * numWaysToBeat
		}
	}
	fmt.Printf("Part 1 - num ways to beat: %d\n", product)
}
