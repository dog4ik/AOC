package day6

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dog4ik/AOC.git/advent"
)

func parseLine(line string) []int {
	results := strings.Split(line, ":")[1]
	resultsArr := strings.Split(results, " ")
	var numbers = []int{}

	for _, result := range resultsArr {
		num, err := strconv.Atoi(result)
		if err != nil {
			continue
		}
		numbers = append(numbers, num)
	}
	return numbers
}

func calculateWins(time int, distance int) int {
	wins := 0

	for i := 0; i < time; i++ {
		remainingTime := time - i
		distanceReached := remainingTime * i
		if distanceReached > distance {
			wins += 1
		}
	}

	return wins
}

func Day6() int {
	lines := advent.ReadLines(6, false)
	times := parseLine(lines[0])
	distances := parseLine(lines[1])
	result := 1
	fmt.Printf("times: %v\ndistances: %v\n", times, distances)
	for i := range times {
		wins := calculateWins(times[i], distances[i])
		fmt.Printf("game %v got %v wins\n", i, wins)
		result = result * wins
	}
	return result
}
