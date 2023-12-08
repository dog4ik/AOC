package day8

import (
	"math"
	"strings"

	"github.com/dog4ik/AOC.git/advent"
)

func parsePoint(line string) (string, [2]string) {
	split := strings.Split(line, " = ")
	key := split[0]
	values := split[1]
	left := values[1:4]
	right := values[6:9]
	return key, [2]string{left, right}

}

func parseNavigation(line string) []string {
	return strings.Split(line, "")
}

func notDone(trackingRoads []string) bool {
	for _, road := range trackingRoads {
		lastChar := string(road[2])
		if string(lastChar) != "Z" {
			return true
		}
	}
	return false
}

func findLCM(numbers []int) int {
	lcmResult := numbers[0]
  for i := range numbers {
		lcmResult = lcm(lcmResult, numbers[i])
	}
	return lcmResult
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return int(math.Abs(float64(a)))
}

func Day8() int {
	lines := advent.ReadLines(8, false)
	var navigations []string
	roads := make(map[string][2]string, len(lines)-3)
	var trackingRoads []string
	for i, line := range lines {
		if i == 0 {
			navigations = parseNavigation(line)
			continue
		}
		if line == "" {
			continue
		}
		key, value := parsePoint(line)
		if string(key[2]) == "A" {
			trackingRoads = append(trackingRoads, key)
		}
		roads[key] = value
	}

	navigationIdx := 0
	steps := 0
	sycles := make([]int, len(trackingRoads))

	for notDone(trackingRoads) {
		direction := navigations[navigationIdx]
		for i := range trackingRoads {

			if string(trackingRoads[i][2]) == "Z" {
				if sycles[i] == 0 {
					sycles[i] = steps
				}
			}
			if direction == "L" {
				trackingRoads[i] = roads[trackingRoads[i]][0]
			}
			if direction == "R" {
				trackingRoads[i] = roads[trackingRoads[i]][1]
			}
		}

		if len(navigations)-1 == navigationIdx {
			navigationIdx = 0
		} else {
			navigationIdx += 1
		}

		shouldBreak := true
		for _, sycle := range sycles {
			if sycle == 0 {
				shouldBreak = false
			}
		}

		if shouldBreak {
			break
		}
		steps += 1
	}

	lcd := findLCM(sycles)

	return lcd
}
