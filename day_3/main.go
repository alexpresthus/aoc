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

// Returns an []int of the indexes adjacent to the provided i...j indexes in a 11 x ? matrix (including line feed)
func isSymbolAtAdjancentIndex(i int, j int, matrix *[]rune) bool {
	w := 141
	h := len(*matrix) / w

	// check side and diagonal indexes if not edges
	if i%w != 0 {
		// left
		if isSymbol((*matrix)[i-1]) {
			return true
		}
		// left over
		if i/w > 0 {
			if isSymbol((*matrix)[i-1-w]) {
				return true
			}
		}
		// left under
		if i/w < h-1 {
			if isSymbol((*matrix)[i-1+w]) {
				return true
			}
		}
	}
	if j%w != w-1 {
		// right
		if isSymbol((*matrix)[j+1]) {
			return true
		}
		// right over
		if j/w > 0 {
			if isSymbol((*matrix)[j+1-w]) {
				return true
			}
		}
		// right under
		if j/w < h-1 {
			if isSymbol((*matrix)[j+1+w]) {
				return true
			}
		}
	}

	// over and under indexes
	for x := i; x <= j; x++ {
		// over
		if x/w > 0 {
			if isSymbol((*matrix)[x-w]) {
				return true
			}
		}
		// under
		if x/w < h-1 {
			if isSymbol((*matrix)[x+w]) {
				return true
			}
		}
	}

	return false
}

func main() {
	content, err := os.ReadFile("/tmp/aoc/input.txt")
	check(err)
	runes := bytes.Runes(content)
	var sum int = 0
	var curNumRunes []rune = make([]rune, 0, 5)
	for i, r := range runes {
		// add digits to current num
		if unicode.IsDigit(r) {
			curNumRunes = append(curNumRunes, r)
			continue
		}
		// not a digit - check current num for adjacent symbols
		if len(curNumRunes) > 0 {
			if isAdjacentSymbol := isSymbolAtAdjancentIndex(i-len(curNumRunes), i-1, &runes); isAdjacentSymbol != false {
				num, err := strconv.Atoi(string(curNumRunes))
				check(err)
				sum += num
			}

		}
		curNumRunes = curNumRunes[:0]
	}
	fmt.Printf("Part 1 sum: %d\n", sum)
}
