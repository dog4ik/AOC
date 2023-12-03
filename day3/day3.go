package day3

import (
	"log"
	"strconv"
	"strings"

	"github.com/dog4ik/AOC.git/advent"
)

func isSymbol(str string) bool {
	return str == "*"
}

func parseNumber(str string) *int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return nil
	}
	return &num
}

func indexExists(arr [][]string, i int, j int) bool {
	xLen := len(arr[0])
	yLen := len(arr) - 1
	return i >= 0 && i < xLen && j >= 0 && j < yLen
}

func intoMatrix(lines []string) [][]string {
	var collector [][]string
	for _, line := range lines {
		chars := strings.Split(line, "")
		collector = append(collector, chars)
	}
	return collector
}

func isNumber(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func exploreNumber(row []string, idx int) int {
	firstIdx := idx
	lastIdx := idx
	for firstIdx > 0 {
		_, err := strconv.Atoi(row[firstIdx-1])
		if err == nil {
			firstIdx -= 1
		} else {
			break
		}
	}
	for lastIdx < len(row) {
		_, err := strconv.Atoi(row[lastIdx])
		if err == nil {
			lastIdx += 1
		} else {
			break
		}
	}
	slice := strings.Join(row[firstIdx:lastIdx], "")
	num, err := strconv.Atoi(slice)
	if err != nil {
		log.Fatalf("this should never happen! Value:(%v)\n", slice)
	}
	for i := firstIdx; i < lastIdx; i++ {
		row[i] = "."
	}
	return num
}

func scanNeighbors(matrix [][]string, row int, col int) []int {
	var numbers []int = []int{}
	for i := row - 1; i < row+2; i++ {
		for j := col - 1; j < col+2; j++ {
			if indexExists(matrix, i, j) {
				char := matrix[i][j]
				if isNumber(char) == true {
					number := exploreNumber(matrix[i], j)
					numbers = append(numbers, number)
				}
			}
		}
	}
	return numbers
}

func Day3() int {
	input := advent.ReadLines(3, false)
	matrix := intoMatrix(input)
	result := 0
	for i, row := range matrix {
		for j, col := range row {
			if isSymbol(col) {
				nums := scanNeighbors(matrix, i, j)
				for i := range nums {
					if len(nums) == 2 {
						result += nums[i] * nums[i+1]
						break
					}
				}
			}
		}
	}

	return result
}
