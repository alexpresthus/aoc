package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Node struct {
	Value int
	L     *Node
	R     *Node
}

func allZeroes(seq []int) bool {
	for _, num := range seq {
		if num != 0 {
			return false
		}
	}
	return true
}

// rec get next number in sequence from extrapolation
func rec(seq []int) int {
	if allZeroes(seq) {
		return 0
	}
	nextSeqStart := make([]int, len(seq)-1)
	for i := range nextSeqStart {
		nextSeqStart[i] = seq[i+1] - seq[i]
	}
	return rec(nextSeqStart) + seq[len(seq)-1]
}

// rec get next number in sequence from extrapolation
func recPrev(seq []int) int {
	if allZeroes(seq) {
		return 0
	}
	nextSeqStart := make([]int, len(seq)-1)
	for i := range nextSeqStart {
		nextSeqStart[i] = seq[i+1] - seq[i]
	}
	return seq[0] - recPrev(nextSeqStart)
}

func solvePartOne(lines []string) (int, time.Duration) {
	startTime := time.Now()
	sum := 0
	for _, ln := range lines {
		strSeq := strings.Fields(ln)
		numSeq := make([]int, len(strSeq))
		for i, str := range strSeq {
			num, err := strconv.Atoi(str)
			check(err)
			numSeq[i] = num
		}
		sum += rec(numSeq)
	}

	return sum, time.Since(startTime)
}

func solvePartTwo(lines []string) (int, time.Duration) {
	startTime := time.Now()
	sum := 0
	for _, ln := range lines {
		strSeq := strings.Fields(ln)
		numSeq := make([]int, len(strSeq))
		for i, str := range strSeq {
			num, err := strconv.Atoi(str)
			check(err)
			numSeq[i] = num
		}
		sum += recPrev(numSeq)
	}

	return sum, time.Since(startTime)
}

func main() {
	content, err := os.ReadFile("/tmp/aoc/input.txt")
	check(err)
	lines := strings.Split(string(content), "\n")
	ansPartOne, time := solvePartOne(lines)
	fmt.Printf("Part 1 - sum: %d    Duration: %s\n", ansPartOne, time)
	ansPartTwo, time := solvePartTwo(lines)
	fmt.Printf("Part 2 - sum: %d    Duration: %s\n", ansPartTwo, time)
}
