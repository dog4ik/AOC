package day10

import (
	"fmt"
	"log"

	"github.com/dog4ik/AOC.git/advent"
)

func inPath(path [][2]int, x int, y int) bool {
	for _, p := range path {
		if p[0] == x && p[1] == y {
			return true
		}
	}
	return false
}

func filterMatrix(matrix [][]string, path [][2]int) {
	for i := range matrix {
		for j := range matrix[i] {
			inPath := inPath(path, j, i)
			if inPath || matrix[i][j] == "." {
				continue
			} else {
				matrix[i][j] = " "
			}
		}
	}
}

func canWalk(dir [2]int, from string, to string) bool {
	if to == "." {
		return false
	}
	x, y := dir[0], dir[1]
	if x == 0 && y == -1 {
		return (from == "|" || from == "L" || from == "J" || from == "S") && (to == "|" || to == "7" || to == "F" || to == "S")
	}
	if x == 1 && y == 0 {
		return (from == "-" || from == "L" || from == "F" || from == "S") && (to == "-" || to == "7" || to == "J" || to == "S")
	}
	if x == 0 && y == 1 {
		return (from == "|" || from == "7" || from == "F" || from == "S") && (to == "|" || to == "L" || to == "J" || to == "S")
	}
	if x == -1 && y == 0 {
		return (from == "-" || from == "J" || from == "7" || from == "S") && (to == "-" || to == "L" || to == "F" || to == "S")
	}
	log.Fatalf("all cases should be covered at this point")
	return false
}

func getDirs() [4][2]int {
	return [4][2]int{
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, 0},
	}
}

func indexExists(arr [][]string, i int, j int) bool {
	xLen := len(arr[0])
	yLen := len(arr) - 1
	return i >= 0 && i < xLen && j >= 0 && j < yLen
}

func findStart(matrix [][]string) [2]int {
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == "S" {
				return [2]int{j, i}
			}
		}
	}
	log.Fatalf("could not find start position")
	return [2]int{-1, -1}
}

func createSeenMap(matrix [][]string) [][]bool {
	seen := make([][]bool, len(matrix))
	for i := range seen {
		seen[i] = make([]bool, len(matrix[0]))
	}
	return seen
}

func isSeen(point [2]int, seen [][]bool) bool {
	return seen[point[1]][point[0]]
}

func selectDirection(current [2]int, matrix [][]string, seen [][]bool) [2]int {
	dirs := getDirs()
	for _, dir := range dirs {
		next := [2]int{current[0] + dir[0], current[1] + dir[1]}
		if !indexExists(matrix, next[1], next[0]) {
			continue
		}
		nextChar := matrix[next[1]][next[0]]
		currChar := matrix[current[1]][current[0]]
		isSeen := isSeen(next, seen)
		canWalk := canWalk(dir, currChar, nextChar)
		if !isSeen && canWalk {
			return next
		}
		if nextChar == "S" && isSeen && canWalk {
			return next
		}
	}
	log.Fatalf("could not find the direction: %v", current)
	return [2]int{-1, -1}
}

func walk(current [2]int, matrix [][]string, seen [][]bool) [][2]int {
	var path [][2]int
	for !isSeen(current, seen) {
		seen[current[1]][current[0]] = true
		path = append(path, current)
		current = selectDirection(current, matrix, seen)
		fmt.Printf("went to %v (%v)\n", current, matrix[current[1]][current[0]])
		newCurrentChar := matrix[current[1]][current[0]]
		if newCurrentChar == "S" {
			break
		}
	}
	return path
}

func Day10() int {
	matrix := advent.ReadIntoMatrix(10, false)
	start := findStart(matrix)
	fmt.Printf("start position: %v\n", start)
	seen := createSeenMap(matrix)
	path := walk(start, matrix, seen)
	path = append(path, start)
	path = path[1:]
	filterMatrix(matrix, path)
  advent.PrintMatrix(matrix)

	return len(path) / 2
}
