package day6

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/dog4ik/AOC.git/advent"
)

func parseLine(line string) int {
	results := strings.Split(line, ":")[1]
	resultsNumber := strings.ReplaceAll(results, " ", "")
	num, err := strconv.Atoi(resultsNumber)
	if err != nil {
		log.Fatalf("could not parse number (%v)", resultsNumber)
	}
	return num
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
	time := parseLine(lines[0])
	distance := parseLine(lines[1])
	result := 1
	fmt.Printf("times: %v\ndistances: %v\n", time, distance)
	wins := calculateWins(time, distance)
	fmt.Printf("game 0 got %v wins\n", wins)
	result = result * wins
	return result
}
