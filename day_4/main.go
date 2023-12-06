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

// matchArr holds number of matching numbers for each card at the corresponding index (arr[0] is the value for card 1 etc.)
func getTotalCards(matchArr []int) int {
	var tot int
	numCopiesMap := make(map[int]int)
	for i, v := range matchArr {
		// copies
		copies, hasCopies := numCopiesMap[i]
		// add current card's 1 original match
		tot++
		// add current card's copies
		if hasCopies != false {
			tot += copies
		}
		for x := i + 1; x < i+1+v; x++ {
			// x is the next card index
			// loop over copies + 1 and add 1 to the next cards' copies
			for j := 0; j <= copies; j++ {
				_, ex := numCopiesMap[x]
				if ex != false {
					numCopiesMap[x] += 1
				} else {
					numCopiesMap[x] = 1
				}
			}
		}
	}
	return tot
}

func main() {
	f, err := os.Open("/tmp/aoc/input.txt")
	check(err)
	defer f.Close()

	var sum int
	scanner := bufio.NewScanner(f)
	matchArr := make([]int, 0, 220) // part 2
	cardIdx := 0                    // part 2
	for scanner.Scan() {
		nums := strings.Split(strings.Split(scanner.Text(), ": ")[1], " | ")
		winningNums, playingNums := strings.Fields(nums[0]), strings.Fields(nums[1])
		var cardScore int
		matchArr = append(matchArr, 0) // part 2
		for _, p := range playingNums {
			for _, w := range winningNums {
				if p == w {
					if cardScore > 0 {
						cardScore = cardScore * 2
					} else {
						cardScore = 1
					}
					matchArr[cardIdx] += 1 // part 2
				}
			}
		}
		sum += cardScore
		cardIdx++ // part 2
	}
	totCards := getTotalCards(matchArr) // part 2
	fmt.Printf("Part 1 sum: %d\n", sum)
	fmt.Printf("Part 2 tot: %d\n", totCards)
}
