package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Use to find first num in part 1 (only unicode digits)
func firstDig(runes *[]rune) (*rune, error) {
	for _, r := range *runes {
		if unicode.IsDigit(r) {
			return &r, nil
		}
	}
	return nil, errors.New("empty")
}

// Used to check if rune array is a digit representation in part 2
func isDigitRepresentation(input *[]rune) (*rune, bool) {
	digitWords1 := map[rune][]rune{
		'o': []rune("one"),
		'e': []rune("eight"),
		'n': []rune("nine"),
	}
	digitWords2 := map[rune]map[rune][]rune{
		't': {
			'w': []rune("two"),
			'h': []rune("three"),
		},
		'f': {
			'o': []rune("four"),
			'i': []rune("five"),
		},
		's': {
			'i': []rune("six"),
			'e': []rune("seven"),
		},
	}
	lngth := len(*input)

	if input == nil || lngth == 0 || lngth > 5 {
		return nil, false
	}

	var expectedRunes, exists = digitWords1[(*input)[0]]
	if exists == false {
		_, existsSec := digitWords2[(*input)[0]]
		if existsSec == false {
			return nil, false
		} else if lngth < 2 {
			return nil, true
		}
		expectedRunes, exists = digitWords2[(*input)[0]][(*input)[1]]
		if exists == false {
			return nil, false
		}
	}

	if len(expectedRunes) > lngth {
		return nil, true
	}
	for i, r := range expectedRunes {
		if r != (*input)[i] {
			return nil, false
		}
	}

	wordRuneMap := map[string]rune{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}
	r := wordRuneMap[string((*input)[:len(expectedRunes)])]

	return &r, true
}

// Use to find first num in part 2 (unicode digits plus string rep "one", "two", etc.)
func allNums(runes *[]rune) *[]int32 {
	nums := make([]int32, 0, 10)
	curWord := make([]rune, 0, 5)
	for _, r := range *runes {
		curWord = append(curWord, r)
		wr, correctStart := isDigitRepresentation(&curWord)
		if correctStart == false {
			if len(curWord) > 1 {
				curWord = curWord[1:] // rm first el
			} else {
				curWord = curWord[:0] // clean
			}
		} else if wr != nil {
			nums = append(nums, int32(*wr-'0')) // add to nums
			curWord = curWord[1:]               // rm first el
		}
		if unicode.IsDigit(r) {
			nums = append(nums, int32(r-'0'))
			curWord = curWord[:0] // empty current word
		}
	}
	for len(curWord) > 2 {
		wr, correctStart := isDigitRepresentation(&curWord)
		if correctStart == false {
			if len(curWord) > 1 {
				curWord = curWord[1:] // rm first el
			} else {
				curWord = curWord[:0] // clean
			}
		} else if wr != nil {
			nums = append(nums, int32(*wr-'0')) // add to nums
			curWord = curWord[1:]               // rm first el
		} else {
			if len(curWord) > 1 {
				curWord = curWord[1:] // rm first el
			} else {
				curWord = curWord[:0] // clean
			}
		}
	}
	return &nums
}

func main() {
	// Part 1
	f1, err1 := os.Open("/tmp/aoc/input.txt")
	check(err1)

	scanner1 := bufio.NewScanner(f1)
	var sum1 int32 = 0
	for scanner1.Scan() {
		runes := []rune(scanner1.Text())
		d1, err := firstDig(&runes)
		if err != nil {
			// no number in line - go to next line
			continue
		}
		slices.Reverse(runes)
		d2, _ := firstDig(&runes)
		sum1 += (int32(*d1-'0') * 10) + int32(*d2-'0')
	}
	fmt.Printf("Part 1 sum: %d\n", sum1)
	check(scanner1.Err())
	f1.Close()

	// Part 2
	f2, err2 := os.Open("/tmp/aoc/input.txt")
	check(err2)

	scanner2 := bufio.NewScanner(f2)
	var sum2 int32 = 0
	for scanner2.Scan() {
		ln := scanner2.Text()
		runes := []rune(ln)

		nums := allNums(&runes)
		if nums == nil || len(*nums) == 0 {
			continue
		}
		d1 := (*nums)[0]
		d2 := (*nums)[len(*nums)-1]
		sum2 += (d1 * 10) + d2
	}
	fmt.Printf("Part 2 sum: %d\n", sum2)
	check(scanner2.Err())
	f2.Close()
}
