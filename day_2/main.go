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
	f, err := os.Open("/tmp/aoc/input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

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

	var id int32 = 0
	var sum int32 = 0
	re := regexp.MustCompile(`(\d+) (\w+)`)
	for scanner.Scan() {
		id++
		game := strings.Split(strings.Split(scanner.Text(), ": ")[1], "; ")
		for i, draw := range game {
			matches := re.FindAllStringSubmatch(draw, -1)
			if isValid(matches) == false {
				break
			} else if i == len(game)-1 {
				sum += id
			}
		}
	}
	fmt.Printf("Sum of valid games: %v\n", sum)

}
