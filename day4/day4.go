package day4

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/dog4ik/AOC.git/advent"
)

func normalizeNumbers(str string) []string {
	trimmed := strings.Trim(str, " ")
	normalized := strings.ReplaceAll(trimmed, "  ", " ")
	return strings.Split(normalized, " ")
}

func parseCard(card string) (int, []int, []int) {
	split := strings.Split(card, ":")
  fmt.Printf("split is:%v\n", split[0])
	cardName := strings.Split(split[0], " ")
	cardNumber, err := strconv.Atoi(cardName[len(cardName)-1])
	if err != nil {
		log.Fatalf("could not figure game number: %v\n", split[0])
	}

	numbersParts := strings.Split(split[1], "|")
	winningPart := normalizeNumbers(numbersParts[0])
	winningNumbers := make([]int, len(winningPart))
	havePart := normalizeNumbers(numbersParts[1])
	haveNumbers := make([]int, len(havePart))
	for i, part := range winningPart {
		number, err := strconv.Atoi(part)
		if err != nil {
			log.Fatalf("cant convert winning number: (%v)\n", winningPart)
		}
		winningNumbers[i] = number
	}

	for i, part := range havePart {
		number, err := strconv.Atoi(part)
		if err != nil {
			log.Fatalf("cant convert have number: (%v)\n", havePart)
		}
		haveNumbers[i] = number
	}
	return cardNumber, winningNumbers, haveNumbers
}

func Day4() int {
	lines := advent.ReadLines(4, false)
	result := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		points := 0
		cardNumber, winningNums, haveNums := parseCard(line)
		for _, haveNum := range haveNums {
			if slices.Contains(winningNums, haveNum) {
				if points == 0 {
					points = 1
				} else {
					points = points * 2
				}
			}
		}
		result += points
		fmt.Printf("%v: %v | %v\n", cardNumber, winningNums, haveNums)
	}

	return result
}
