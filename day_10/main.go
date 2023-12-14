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

type PipeNode struct {
	Pipe     Pipe
	i        int
	j        int
	Prev     *PipeNode
	Next     *PipeNode
	Distance int
}

func (node PipeNode) print() {
	fmt.Printf("PipeNode { Pipe: %s, i: %d, j: %d, Prev: &'%v', Next: &'%v', Distance: %d }\n", string(node.Pipe), node.i, node.j, string(node.Prev.Pipe), string(node.Next.Pipe), node.Distance)
}

type Pipe = rune

const (
	Start             Pipe = 'S'
	Vertical          Pipe = '|'
	Horizontal        Pipe = '-'
	BottomLeftCorner  Pipe = 'L'
	BottomRightCorner Pipe = 'J'
	TopLeftCorner     Pipe = 'F'
	TopRightCorner    Pipe = '7'
	Ground            rune = '.'
)

type PipeLoop struct {
	Head *PipeNode
}

func (loop PipeLoop) len() int {
	check := loop.Head.Next
	len := 1
	for check != loop.Head {
		check = check.Next
		len++
	}
	return len
}

func (loop PipeLoop) print() {
	fmt.Print("    ")
	loop.Head.print()
	check := loop.Head.Next
	for check != loop.Head {
		fmt.Print(" -> ")
		check.print()
		check = check.Next
	}
}

func (loop PipeLoop) longestDistance() int {
	check := loop.Head.Next
	for true {
		if check.Next.Distance <= check.Distance {
			break
		}
		check = check.Next
	}
	return check.Distance
}

func addPipeToLoop(r rune, loop *PipeLoop) {
}
func isConnectedNorth(r rune) bool {
	switch r {
	case Vertical:
		return true
	case BottomLeftCorner:
		return true
	case BottomRightCorner:
		return true
	default:
		return false
	}
}
func isConnectedSouth(r rune) bool {
	switch r {
	case Vertical:
		return true
	case TopLeftCorner:
		return true
	case TopRightCorner:
		return true
	default:
		return false
	}
}
func isConnectedEast(r rune) bool {
	switch r {
	case Horizontal:
		return true
	case BottomLeftCorner:
		return true
	case TopLeftCorner:
		return true
	default:
		return false
	}
}
func isConnectedWest(r rune) bool {
	switch r {
	case Horizontal:
		return true
	case BottomRightCorner:
		return true
	case TopRightCorner:
		return true
	default:
		return false
	}
}

type Direction uint8

const (
	North Direction = iota
	South
	East
	West
	NoDirection
)

type Edge uint8

const (
	Top Edge = iota
	Bottom
	Left
	Right
	NoEdge
)

// assumes node and node.Prev is adjacent vertically or horizontally
func getNotDirection(node *PipeNode, direction int) Direction {
	if node.Prev == nil && node.Next == nil {
		return NoDirection
	}
	var check *PipeNode
	if direction > 0 {
		check = node.Prev
	} else {
		check = node.Next
	}

	if check.j == node.j {
		if check.i > node.i {
			return South
		} else {
			return North
		}
	}
	if check.j > node.j {
		return East
	} else {
		return West
	}
}

func getEdge(node *PipeNode, h int, w int) Edge {
	if node.i == 0 {
		return Top
	}
	if node.i == h-1 {
		return Bottom
	}
	if node.j == 0 {
		return Left
	}
	if node.j == w-1 {
		return Right
	}
	return NoEdge
}

func areInSamePosition(a *PipeNode, b *PipeNode) bool {
	if a.i == b.i && a.j == b.j {
		return true
	}
	return false
}

func findConnectedPipes(head *PipeNode, m *[][]rune, direction int) []*PipeNode {
	count := 1
	if head.Prev == nil && head.Next == nil {
		count = 2
	}
	nodes := make([]*PipeNode, 0, count)
	// look for pipes until we have enough
	not := getNotDirection(head, direction)
	edge := getEdge(head, len(*m), len((*m)[0]))
	// check if south rune is connected in north
	if (head.Pipe == Start || isConnectedSouth(head.Pipe)) && edge != Bottom && not != South && isConnectedNorth((*m)[head.i+1][head.j]) {
		nodes = append(nodes, &PipeNode{Pipe: (*m)[head.i+1][head.j], i: head.i + 1, j: head.j, Distance: head.Distance + 1})
	}
	// check if north rune is connected in south if not already found all
	if (head.Pipe == Start || isConnectedNorth(head.Pipe)) && edge != Top && not != North && len(nodes) < count && isConnectedSouth((*m)[head.i-1][head.j]) {
		node := &PipeNode{Pipe: (*m)[head.i-1][head.j], i: head.i - 1, j: head.j, Distance: head.Distance + 1}
		nodes = append(nodes, node)
	}
	// check if east rune is connected in west if not already found all
	if (head.Pipe == Start || isConnectedEast(head.Pipe)) && edge != Right && not != East && len(nodes) < count && isConnectedWest((*m)[head.i][head.j+1]) {
		node := &PipeNode{Pipe: (*m)[head.i][head.j+1], i: head.i, j: head.j + 1, Distance: head.Distance + 1}
		nodes = append(nodes, node)
	}
	// check if west rune is connected in east if not already found all
	if (head.Pipe == Start || isConnectedWest(head.Pipe)) && edge != Left && not != West && len(nodes) < count && isConnectedEast((*m)[head.i][head.j-1]) {
		node := &PipeNode{Pipe: (*m)[head.i][head.j-1], i: head.i, j: head.j - 1, Distance: head.Distance + 1}
		nodes = append(nodes, node)
	}
	for i, n := range nodes {
		if i == 0 {
			if direction > 0 {
				n.Prev = head
				head.Next = n
			} else {
				n.Next = head
				head.Prev = n
			}
		} else {
			if direction > 0 {
				n.Next = head
				head.Prev = n
			} else {
				n.Prev = head
				head.Next = n
			}
		}
	}
	return nodes
}

func buildLoop(start *PipeNode, m *[][]rune) *PipeLoop {
	loop := PipeLoop{Head: start}

	currentPipeNodes := findConnectedPipes(start, m, 1)

	for true {
		if areInSamePosition(currentPipeNodes[0], currentPipeNodes[1]) {
			// close the loop, keeping only one of the final nodes
			currentPipeNodes[0].Next = currentPipeNodes[1].Next
			currentPipeNodes[1].Next.Prev = currentPipeNodes[0]
			break
		}
		nextOne := findConnectedPipes(currentPipeNodes[0], m, 1)[0]
		nextTwo := findConnectedPipes(currentPipeNodes[1], m, -1)[0]
		currentPipeNodes = []*PipeNode{nextOne, nextTwo}
	}

	return &loop
}

func solvePartOne(lines []string) (int, time.Duration) {
	startTime := time.Now()

	m := make([][]rune, len(lines))
	start := PipeNode{Pipe: Pipe(Start), Distance: 0}
	for i, ln := range lines {
		m[i] = make([]rune, len(ln))
		for j, r := range ln {
			if r == Start {
				// populate start node
				start.i = i
				start.j = j
			}
			m[i][j] = r
		}
	}

	loop := buildLoop(&start, &m)
	//loop.print()

	sum := loop.longestDistance()

	return sum, time.Since(startTime)
}

func solvePartTwo(lines []string) (int, time.Duration) {
	startTime := time.Now()
	sum := 0

	return sum, time.Since(startTime)
}

func main() {
	content, err := os.ReadFile("/tmp/aoc/input.txt")
	check(err)
	lines := strings.Split(string(content), "\n")
	ansPartOne, time := solvePartOne(lines)
	fmt.Printf("Part 1 - shortest distance: %d    Duration: %s\n", ansPartOne, time)
	// ansPartTwo, time := solvePartTwo(lines)
	// fmt.Printf("Part 2 - steps: %d    Duration: %s\n", ansPartTwo, time)
}
