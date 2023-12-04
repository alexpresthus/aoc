package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	isValid := func(matches [][]string) bool {
		for _, match := range matches {
			clr := strings.Fields(match[0])[1]
			num, err := strconv.Atoi(match[1])
			check(err)
			if clr == "red" && num > 12 {
				return false
			} else if clr == "green" && num > 13 {
				return false
			} else if clr == "blue" && num > 14 {
				return false
			}
		}
		return true
	}

	getDrawCounts := func(matches [][]string) map[string]int {
		cnts := make(map[string]int)
		for _, match := range matches {
			clr := strings.Fields(match[0])[1]
			num, err := strconv.Atoi(match[1])
			check(err)
			cnts[clr] = num
		}
		return cnts
	}

	// Part 1
	f1, err1 := os.Open("/tmp/aoc/input.txt")
	check(err1)

	scanner1 := bufio.NewScanner(f1)
	var id int32 = 0
	var sum int32 = 0
	re := regexp.MustCompile(`(\d+) (\w+)`)
	for scanner1.Scan() {
		id++
		game := strings.Split(strings.Split(scanner1.Text(), ": ")[1], "; ")
		for i, draw := range game {
			matches := re.FindAllStringSubmatch(draw, -1)
			if isValid(matches) == false {
				break
			} else if i == len(game)-1 {
				sum += id
			}
		}
	}
	fmt.Printf("Part 1 Sum of valid games: %v\n", sum)
	f1.Close()

	// Part 2
	f2, err2 := os.Open("/tmp/aoc/input.txt")
	check(err2)

	scanner2 := bufio.NewScanner(f2)
	var sum2 int = 0
	for scanner2.Scan() {
		gamesLeast := make(map[string]int)
		for _, draw := range strings.Split(strings.Split(scanner2.Text(), ": ")[1], "; ") {
			matches := re.FindAllStringSubmatch(draw, -1)
			counts := getDrawCounts(matches)
			for clr, num := range counts {
				curClrLst, ex := gamesLeast[clr]
				if ex == false || num > curClrLst {
					gamesLeast[clr] = num
				}
			}
		}
		pw := 1
		for _, num := range gamesLeast {
			pw = pw * num
		}
		sum2 += pw
	}
	fmt.Printf("Part 2 Sum of power of sets: %v\n", sum2)
	f2.Close()
}
