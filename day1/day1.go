package day1

import (
	"strconv"
	"strings"

	"github.com/dog4ik/AOC.git/advent"
)

func getNumbers(str string) (int, int) {
	numbers := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	var first *int = nil
	var last *int = nil
	firstIdx := 9999
	lastIdx := -1

	for i, num := range numbers {
		firstNumIdx := strings.Index(str, num)
		lastNumIdx := strings.LastIndex(str, num)
		number := i + 1
		if firstNumIdx != -1 {
			if firstIdx > firstNumIdx {
				firstIdx = firstNumIdx
				first = &number
			}
		}

		if lastNumIdx != -1 {
			if lastIdx < lastNumIdx {
				lastIdx = lastNumIdx
				last = &number
			}
		}
	}

	chars := strings.Split(str, "")

	for i, char := range chars {
		number, err := strconv.Atoi(char)
		if err != nil {
			continue
		} else {
			if firstIdx > i {
				firstIdx = i
				first = &number
			}
			if lastIdx < i {
				lastIdx = i
				last = &number
			}
		}
	}
	return *first, *last
}

func Day1() int {
  lines := advent.ReadLines(1, false)

	sum := 0

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		first, last := getNumbers(line)
		res := first*10 + last
		sum += res
	}
	return sum
}
