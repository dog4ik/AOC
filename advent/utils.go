package advent

import (
	"fmt"
	"log"
	"os"
	"strings"
)

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
