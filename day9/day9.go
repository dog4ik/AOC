package day9

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/dog4ik/AOC.git/advent"
)

func parseRow(line string) []int {
	split := strings.Split(line, " ")
	numbers := make([]int, len(split))
	for i, chunk := range split {
		num, err := strconv.Atoi(chunk)
		if err != nil {
			log.Fatalf("could not parse row %v", chunk)
		}
		numbers[i] = num
	}
	return numbers
}

func isAllZero(values []int) bool {
	for _, num := range values {
		if num != 0 {
			return false
		}
	}
	return true
}

func walk(values []int) ([]int, int) {
	if isAllZero(values) {
		return values, 0
	}
	newValues := make([]int, len(values)-1)
	for i := range values {
		if i != len(values)-1 {
			newValues[i] = values[i+1] - values[i]
		}
	}
	fmt.Printf("new numbers:%v\n", newValues)
	values, reminder := walk(newValues)
	last := newValues[len(newValues)-1]

	return values, last + reminder
}

func calculateValue(values []int) int {
	_, value := walk(values)
	return value + values[len(values)-1]
}

func Day9() int {
	lines := advent.ReadLines(9, false)
	res := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		numbers := parseRow(line)
		res += calculateValue(numbers)
		fmt.Printf("result:%v\n", res)
	}
	return res
}
