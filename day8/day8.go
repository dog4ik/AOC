package day8

import (
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

func Day8() int {
	lines := advent.ReadLines(8, false)
	var navigations []string
	roads := make(map[string][2]string, len(lines)-3)
	for i, line := range lines {
		if i == 0 {
			navigations = parseNavigation(line)
			continue
		}
		if line == "" {
			continue
		}
		key, value := parsePoint(line)
		roads[key] = value
	}

	currentLocation := "AAA"
	navigationIdx := 0
	steps := 0
	for currentLocation != "ZZZ" {
		direction := navigations[navigationIdx]
		if direction == "L" {
			currentLocation = roads[currentLocation][0]
		}
		if direction == "R" {
			currentLocation = roads[currentLocation][1]
		}
		if len(navigations)-1 == navigationIdx {
			navigationIdx = 0
		} else {
			navigationIdx += 1
		}
		steps += 1
	}

	return steps
}
