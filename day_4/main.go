package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.Open("/tmp/aoc/input.txt")
	check(err)
	defer f.Close()

	var sum int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		nums := strings.Split(strings.Split(scanner.Text(), ": ")[1], " | ")
		winningNums, playingNums := strings.Fields(nums[0]), strings.Fields(nums[1])
		var cardScore int
		for _, p := range playingNums {
			for _, w := range winningNums {
				if p == w {
					if cardScore > 0 {
						cardScore = cardScore * 2
					} else {
						cardScore = 1
					}
				}
			}
		}
		sum += cardScore
	}
	fmt.Printf("Part 1 sum: %d\n", sum)
}
