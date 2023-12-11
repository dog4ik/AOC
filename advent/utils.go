package advent

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func PrintMatrix(matrix [][]string) {
	for i := range matrix {
		for j := range matrix[i] {
			fmt.Printf("%v", matrix[i][j])
		}
		fmt.Printf("\n")
	}
}

func intoMatrix(lines []string) [][]string {
	var collector [][]string
	for _, line := range lines {
		chars := strings.Split(line, "")
		collector = append(collector, chars)
	}
	return collector
}

func ReadIntoMatrix(day int, isTest bool) [][]string {
	lines := ReadLines(day, isTest)
	return intoMatrix(lines)
}

func ReadLines(day int, isTest bool) []string {
	fileName := "input"
	if isTest {
		fileName = "test"
	}

	path := fmt.Sprintf("/home/dog4ik/personal/go-aoc/inputs/day%d/%v.txt", day, fileName)
	text, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read file %v", path)
	}
	return strings.Split(string(text), "\n")
}
