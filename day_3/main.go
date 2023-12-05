package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func isSymbol(r rune) bool {
	if unicode.IsDigit(r) || unicode.IsControl(r) || r == '.' {
		return false
	}
	return true
}

func isGear(r rune) bool {
	return r == '*'
}

// Returns an []int of the indexes adjacent to the provided i...j indexes in a 11 x ? matrix (including line feed)
func someCheckAtAdjancentIndex(i int, j int, matrix *[]rune, check func(r rune) bool) (int, bool) {
	w := 141
	h := len(*matrix) / w

	// check side and diagonal indexes if not edges
	if i%w != 0 {
		// left
		if check((*matrix)[i-1]) {
			return i - 1, true
		}
		// left over
		if i/w > 0 {
			if check((*matrix)[i-1-w]) {
				return i - 1 - w, true
			}
		}
		// left under
		if i/w < h-1 {
			if check((*matrix)[i-1+w]) {
				return i - 1 + w, true
			}
		}
	}
	if j%w != w-1 {
		// right
		if check((*matrix)[j+1]) {
			return j + 1, true
		}
		// right over
		if j/w > 0 {
			if check((*matrix)[j+1-w]) {
				return j + 1 - w, true
			}
		}
		// right under
		if j/w < h-1 {
			if check((*matrix)[j+1+w]) {
				return j + 1 + w, true
			}
		}
	}

	// over and under indexes
	for x := i; x <= j; x++ {
		// over
		if x/w > 0 {
			if check((*matrix)[x-w]) {
				return x - w, true
			}
		}
		// under
		if x/w < h-1 {
			if check((*matrix)[x+w]) {
				return x + w, true
			}
		}
	}

	return -1, false
}

func main() {
	content, err := os.ReadFile("/tmp/aoc/input.txt")
	check(err)
	runes := bytes.Runes(content)
	var sum1 int = 0
	var curNumRunes []rune = make([]rune, 0, 5)
	gearPartsMap := make(map[int][]int) // idx: []int
	for i, r := range runes {
		// add digits to current num
		if unicode.IsDigit(r) {
			curNumRunes = append(curNumRunes, r)
			continue
		}
		// not a digit - check current num for adjacent symbols
		if len(curNumRunes) > 0 {
			if _, isAdjacentSymbol := someCheckAtAdjancentIndex(i-len(curNumRunes), i-1, &runes, isSymbol); isAdjacentSymbol != false {
				num, err := strconv.Atoi(string(curNumRunes))
				check(err)
				sum1 += num
			}
			if gearIdx, isAdjacentGear := someCheckAtAdjancentIndex(i-len(curNumRunes), i-1, &runes, isGear); isAdjacentGear != false {
				num, err := strconv.Atoi(string(curNumRunes))
				check(err)
				gearNums, ex := gearPartsMap[gearIdx]
				if ex != true {
					gearPartsMap[gearIdx] = []int{num}
				} else {
					gearPartsMap[gearIdx] = append(gearNums, num)
				}
			}
		}
		curNumRunes = curNumRunes[:0]
	}
	var sumGearRatios int = 0
	for _, nums := range gearPartsMap {
		if len(nums) == 2 {
			sumGearRatios += nums[0] * nums[1]
		}
	}

	fmt.Printf("Part 1 sum: %d\n", sum1)
	fmt.Printf("Part 2 sum of gear ratios %v\n", sumGearRatios)
}
