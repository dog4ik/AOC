package day7

import (
	"cmp"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/dog4ik/AOC.git/advent"
)

func parseHand(line string) ([]int, int) {
	split := strings.Split(line, " ")
	parsedBid, err := strconv.Atoi(split[1])
	strength := make([]int, 5)
	chars := strings.Split(split[0], "")
	for i, char := range chars {
		intCard, err := strconv.Atoi(char)
		if err == nil {
			strength[i] = intCard
			continue
		}
		if char == "T" {
			strength[i] = 10
		}
		if char == "J" {
			strength[i] = 1
		}
		if char == "Q" {
			strength[i] = 12
		}
		if char == "K" {
			strength[i] = 13
		}
		if char == "A" {
			strength[i] = 14
		}
	}
	if err != nil {
		log.Default().Fatalf("could not parse bid:(%v)", line)
	}
	return strength, parsedBid
}

func getHandRank(card []int) int {
	values := make(map[int]int)
	for _, v := range card {
		_, ok := values[v]
		if !ok {
			values[v] = 1
		} else {
			values[v] += 1
		}
	}

	_, haveJocker := values[1]

	if len(values) > 1 && haveJocker {
		maxVal := 0
		maxCard := -1
		for key, value := range values {
			// joker
			if key == 1 {
				continue
			}
			if value > maxVal {
				maxVal = value
				maxCard = key
			}
		}
		jockerAmount := values[1]
		delete(values, 1)
    values[maxCard] += jockerAmount
	}

	// five of kind
	if len(values) == 1 {
		return 7
	}

	if len(values) == 2 {
		for _, value := range values {
			// four of kind
			if value == 4 || value == 1 {
				return 6
			}
			// full house
			if value == 3 || value == 2 {
				return 5
			}
		}
	}

	if len(values) == 3 {
		// 3 of kind
		for _, value := range values {
			if value == 3 {
				return 4
			}
		}

		// 2 pair
		pair := false
		for _, value := range values {
			if value == 2 && !pair {
				pair = true
			} else {
				return 3
			}
		}
	}

	if len(values) == 4 {
		return 2
	}

	if len(values) == 5 {
		return 1
	}

	return 0
}

type Hand struct {
	bid  int
	hand []int
}

func Day7() int {
	lines := advent.ReadLines(7, false)
	hands := make([]Hand, len(lines)-1)
	for i, line := range lines {
		if line == "" {
			continue
		}
		hand, bid := parseHand(line)
		hands[i] = Hand{bid: bid, hand: hand}
	}

	slices.SortFunc(hands, func(a, b Hand) int {
		if n := cmp.Compare(getHandRank(a.hand), getHandRank(b.hand)); n != 0 {
			return n
		}
		for i := range a.hand {
			n := cmp.Compare(a.hand[i], b.hand[i])
			if n == 0 {
				continue
			} else {
				return n
			}

		}
		return 0
	})

	result := 0
	for i, hand := range hands {
		rank := i + 1
		result += hand.bid * rank
	}

	return result
}
