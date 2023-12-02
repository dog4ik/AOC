package day2

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/dog4ik/AOC.git/advent"
)

func parseGameId(game string) int {
	gamePart := strings.Split(game, ":")[0]
	stringNumber := strings.Split(gamePart, " ")[1]
	gameNumber, err := strconv.Atoi(stringNumber)
	if err != nil {
		log.Fatalf("gameId is not reachable")
	}
	return gameNumber
}

func trimSplit(str string, splitChar string) []string {
	return strings.Split(strings.Trim(str, " "), splitChar)
}

func colorToNumber(str string) int {
	if str == "red" {
		return 0
	}

	if str == "green" {
		return 1
	}

	if str == "blue" {
		return 2
	}
	log.Fatalf("unknown color %v", str)
	return -1
}

func parseRoundPart(str string) (int, int) {
	splittedRoundPart := strings.Split(str, " ")
	amount, err := strconv.Atoi(splittedRoundPart[0])
	if err != nil {
		log.Fatal("could not parse gamePart amount")
	}
	colorIdx := colorToNumber(splittedRoundPart[1])
	return amount, colorIdx
}

func parseGameRGB(game string) [3]int {
	maxRgbAmount := [3]int{0, 0, 0}
	payloadPart := strings.Trim(strings.Split(game, ":")[1], " ")
	rounds := strings.Split(payloadPart, ";")
	for _, round := range rounds {
		round = strings.Trim(round, " ")
		roundParts := trimSplit(round, ",")
		for _, roundPart := range roundParts {
			roundPart = strings.Trim(roundPart, " ")
			amount, color := parseRoundPart(roundPart)
			if maxRgbAmount[color] < amount {
				maxRgbAmount[color] = amount
			}
		}
	}
	return maxRgbAmount
}

func Day2() int {
	gameConfig := [3]int{12, 13, 14}
	lines := advent.ReadLines(2, true)
	result := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		gameId := parseGameId(line)
		gameItems := parseGameRGB(line)
		isFair := true
		for i, configValue := range gameConfig {
			if configValue < gameItems[i] {
				isFair = false
			}
		}
		fmt.Printf("game %v %v -> %v\n", gameId, gameItems, isFair)
		if isFair {
			result += gameId
		}
	}
	return result
}
