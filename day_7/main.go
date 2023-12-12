package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type HandType uint8

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	Cards string
	Type  HandType
	Bid   int
}

type Pair struct {
	Key   rune
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func getHandType(cardCounts PairList) HandType {
	unique := len(cardCounts)
	if unique == 1 {
		return FiveOfAKind
	}
	if unique == 2 {
		if cardCounts[0].Value == 4 || cardCounts[1].Value == 4 {
			return FourOfAKind
		} else {
			return FullHouse
		}
	}
	threes := 0
	pairs := 0
	for _, p := range cardCounts {
		if p.Value == 3 {
			threes += 1
			continue
		}
		if p.Value == 2 {
			pairs += 1
			continue
		}
	}
	if threes > 0 {
		return ThreeOfAKind
	}
	if pairs > 1 {
		return TwoPair
	}
	if pairs > 0 {
		return OnePair
	}
	return HighCard
}

func compareLabels(compare rune, with rune, jokersAreWild bool) int {
	labels := map[rune]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'J': 11,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
	}
	if jokersAreWild {
		labels['J'] = 1
	}
	return labels[compare] - labels[with]
}

func solvePartOne(lines []string) (int, time.Duration) {
	startTime := time.Now()
	hands := make([]Hand, 0, len(lines))
	for _, ln := range lines {
		cards := strings.Split(ln, " ")[0]
		bid, err := strconv.Atoi(strings.Split(ln, " ")[1])
		check(err)
		m := make(map[rune]int)
		for _, card := range cards {
			m[card] = m[card] + 1
		}
		cardCounts := make(PairList, 0, len(m))
		for card, cnt := range m {
			cardCounts = append(cardCounts, Pair{card, cnt})
		}
		sort.Sort(sort.Reverse(cardCounts))
		var t HandType = getHandType(cardCounts)
		hands = append(hands, Hand{cards, t, bid})
	}
	slices.SortFunc(hands, func(a, b Hand) int {
		if a.Type != b.Type {
			return int(a.Type) - int(b.Type)
		}
		for i := 0; i < 5; i++ {
			if a.Cards[i] != b.Cards[i] {
				return compareLabels(rune(a.Cards[i]), rune(b.Cards[i]), false)
			}
		}
		return 0
	})
	winnings := 0
	for i := 0; i < len(hands); i++ {
		winnings += (i + 1) * hands[i].Bid
	}
	return winnings, time.Since(startTime)
}

func solvePartTwo(lines []string) (int, time.Duration) {
	startTime := time.Now()
	hands := make([]Hand, 0, len(lines))
	for _, ln := range lines {
		cards := strings.Split(ln, " ")[0]
		bid, err := strconv.Atoi(strings.Split(ln, " ")[1])
		check(err)
		m := make(map[rune]int)
		for _, card := range cards {
			m[card] = m[card] + 1
		}
		var jokers int
		if m['J'] > 0 && len(m) > 1 {
			jokers = m['J']
			delete(m, 'J')
		}

		cardCounts := make(PairList, 0, len(m))
		for card, cnt := range m {
			cardCounts = append(cardCounts, Pair{card, cnt})
		}
		sort.Sort(sort.Reverse(cardCounts))

		// add jokers to highest card count
		if jokers > 0 {
			cardCounts[0].Value += jokers
		}
		var t HandType = getHandType(cardCounts)

		hands = append(hands, Hand{cards, t, bid})
	}
	slices.SortFunc(hands, func(a, b Hand) int {
		if a.Type != b.Type {
			return int(a.Type) - int(b.Type)
		}
		for i := 0; i < 5; i++ {
			if a.Cards[i] != b.Cards[i] {
				return compareLabels(rune(a.Cards[i]), rune(b.Cards[i]), true)
			}
		}
		return 0
	})
	winnings := 0
	for i := 0; i < len(hands); i++ {
		winnings += (i + 1) * hands[i].Bid
	}
	return winnings, time.Since(startTime)
}

func main() {
	content, err := os.ReadFile("/tmp/aoc/input.txt")
	check(err)
	lines := strings.Split(string(content), "\n")
	ansPartOne, time := solvePartOne(lines)
	fmt.Printf("Part 1 - total winnings: %d 	Duration: %s\n", ansPartOne, time)
	ansPartTwo, time := solvePartTwo(lines)
	fmt.Printf("Part 2 - total winnings: %d 	Duration: %s\n", ansPartTwo, time)
}
