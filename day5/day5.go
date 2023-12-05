package day5

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/dog4ik/AOC.git/advent"
)

type Range struct {
	dest int
	src  int
	len  int
}

func parseSeeds(line string) []int {
	seedsPart := strings.Split(line, ": ")[1]
	seeds := strings.Split(seedsPart, " ")
	seedsNums := make([]int, len(seeds))
	for i := range seedsNums {
		number, err := strconv.Atoi(seeds[i])
		if err != nil {
			log.Fatalf("failed to parse seeds numbers")
		}
		seedsNums[i] = number
	}
	return seedsNums
}

func parseRange(line string) Range {
	ranges := strings.Split(line, " ")
	rangeNums := make([]int, len(ranges))
	for i := range rangeNums {
		number, err := strconv.Atoi(ranges[i])
		if err != nil {
			log.Fatalf("failed to parse range numbers")
		}
		rangeNums[i] = number
	}
	return Range{dest: rangeNums[0], src: rangeNums[1], len: rangeNums[2]}
}

func inRange(r Range, source int) bool {
	return source < r.src+r.len && source >= r.src
}

func passLayer(r Range, source int) int {
	return r.dest + source - r.src
}

func Day5() int {
	// var seedToSoil map[string]int
	lines := advent.ReadLines(5, false)
	ranges := make([][]Range, 7)
	var seeds []int
	rangeIdx := -1
	for i, line := range lines {
		if i == 0 {
			seeds = parseSeeds(line)
			continue
		}
		if line == "" {
			rangeIdx += 1
			continue
		}
		if lines[i-1] == "" {
			continue
		}
		ranges[rangeIdx] = append(ranges[rangeIdx], parseRange(line))
	}
	for seedIdx := range seeds {
		for sectorIdx, sector := range ranges {
			fmt.Printf("new sector: %v\n", sectorIdx)
			for _, rng := range sector {
				fmt.Printf("not in range(%v): %v\n", rng, seeds[seedIdx])
				if inRange(rng, seeds[seedIdx]) {
					seeds[seedIdx] = passLayer(rng, seeds[seedIdx])
					fmt.Printf("in range(%v): %v\n", rng, seeds[seedIdx])
          break
				}
			}
		}
	}
	fmt.Printf("final seeds: %v\n", seeds)
	return slices.Min(seeds)
}
