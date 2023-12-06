package day5

import (
	"log"
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
	seedsRanges := make([]int, len(seeds))
	for i := range seedsRanges {
		number, err := strconv.Atoi(seeds[i])
		if err != nil {
			log.Fatalf("failed to parse seeds numbers")
		}
		seedsRanges[i] = number
	}
  return seedsRanges
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

// TODO: filter ranges to speed it up
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
	minimal := 9999999999999

	for i := 0; i < len(seeds); i += 2 {
		start := seeds[i]
		amount := seeds[i+1]
		for seed := start; seed < start+amount; seed++ {
      res := seed
			for _, sector := range ranges {
				// fmt.Printf("new sector: %v\n", sectorIdx)
				for _, rng := range sector {
					// fmt.Printf("not in range(%v): %v\n", rng, res)
					if inRange(rng, res) {
						res = passLayer(rng, res)
						// fmt.Printf("in range(%v): %v\n", rng, res)
						break
					}
				}
			}
      if res < minimal {
        minimal = res
      }
		}
	}
	return minimal
}
