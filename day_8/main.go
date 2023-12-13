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
type Direction rune

const (
	Left  Direction = 'L'
	Right Direction = 'R'
)

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

func allNodesOnZ(nodes []Node) ([]Node, bool) {
	remaining := make([]Node, 0, len(nodes))
	all := true
	for _, n := range nodes {
		if n[2] != 'Z' {
			all = false
			remaining = append(remaining, n)
		}
	}
	return remaining, all
}

func gcd(a int, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func lcm(a int, b int) int {
	return (a * b) / gcd(a, b)
}

// expects len(x) > 2
func multiLcm(x []int) int {
	if len(x) < 3 {
		panic("Slice too short for multiLcm() - use lcm() instead")
	}
	curr := lcm(x[0], x[1])
	for _, k := range x[2:] {
		curr = lcm(curr, k)
	}
	return curr
}

func solvePartTwo(lines []string) (int, time.Duration) {
	startTime := time.Now()
	instructions := make([]Direction, len(lines[0]))
	for i, r := range lines[0] {
		instructions[i] = Direction(r)
	}
	var steps int
	var current []Node
	replacer := strings.NewReplacer("=", "", "(", "", ")", "", ",", "")
	m := make(map[Node]Children)
	for _, ln := range lines[2:] {
		fields := strings.Fields(replacer.Replace(ln))
		head := Node(fields[0])
		left := Node(fields[1])
		right := Node(fields[2])
		m[head] = Children{left, right}
		if head[2] == 'A' {
			current = append(current, head)
		}
	}

	next := func(node Node, dir Direction) Node {
		if dir == Left {
			return m[node].L
		} else {
			return m[node].R
		}
	}

	cycleSteps := make([]int, 0, len(current))

	done := false
	for !done {
		for _, dir := range instructions {
			remainingNodes, allDone := allNodesOnZ(current)
			for x := 0; x < len(current)-len(remainingNodes); x++ {
				cycleSteps = append(cycleSteps, steps)
			}
			if allDone {
				done = true
				break
			}
			current = remainingNodes
			for i, node := range current {
				current[i] = next(node, dir)
			}
			steps++
		}
	}

	totalSteps := multiLcm(cycleSteps)
	return totalSteps, time.Since(startTime)
}

func main() {
	content, err := os.ReadFile("/tmp/aoc/input.txt")
	check(err)
	lines := strings.Split(string(content), "\n")
	ansPartOne, time := solvePartOne(lines)
	fmt.Printf("Part 1 - steps: %d    Duration: %s\n", ansPartOne, time)
	ansPartTwo, time := solvePartTwo(lines)
	fmt.Printf("Part 2 - steps: %d    Duration: %s\n", ansPartTwo, time)
}
