package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Node string
type Children struct {
	L Node
	R Node
}

func solvePartOne(lines []string) (int, time.Duration) {
	startTime := time.Now()
	instructions := lines[0]
	var steps int
	var current Node
	replacer := strings.NewReplacer("=", "", "(", "", ")", "", ",", "")
	m := make(map[Node]Children)
	for _, ln := range lines[2:] {
		fields := strings.Fields(replacer.Replace(ln))
		head := Node(fields[0])
		left := Node(fields[1])
		right := Node(fields[2])
		m[head] = Children{left, right}
		if head == "AAA" {
			current = head
		}
	}

	for true {
		for _, direction := range instructions {
			if current == "ZZZ" {
				break
			}
			if direction == 'L' {
				current = m[current].L
				steps += 1
			} else {
				current = m[current].R
				steps += 1
			}
		}
		if current == "ZZZ" {
			break
		}
	}
	return steps, time.Since(startTime)
}

func main() {
	content, err := os.ReadFile("/tmp/aoc/input.txt")
	check(err)
	lines := strings.Split(string(content), "\n")
	ansPartOne, time := solvePartOne(lines)
	fmt.Printf("Part 1 - steps: %d    Duration: %s\n", ansPartOne, time)
}
