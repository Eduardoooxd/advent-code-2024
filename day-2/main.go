package main

// report per line
// list number called levels

/*

six reports each containing five levels

7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9


*/

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	reports := utils.ReadInput("input.txt")

	nSafeReports := safeReports(reports)
	fmt.Printf("Part 1 (Safe Reports): %d\n", nSafeReports)
}

type OrderDirection int

const (
	Unspecified OrderDirection = iota
	Increasing
	Decreasing
)

func safeReports(reports []string) int {
	nSafeReports := 0

	for _, report := range reports {
		levels := getLevels(report)
		nLevels := len(levels)

		var direction OrderDirection = Unspecified
		isSafeReport := true

		for i, _ := range levels {
			if (i + 1) >= nLevels {
				// last number
				continue
			}

			currentLevel := levels[i]
			nextLevel := levels[i+1]

			// If numbers are equal, report is unsafe
			if currentLevel == nextLevel {
				isSafeReport = false
				break
			}

			if direction == Unspecified {
				if nextLevel > currentLevel {
					direction = Increasing
					// I want to ignore the case where both are equal, because I can't specify
				} else {
					direction = Decreasing
				}
			}

			if !validTransition(currentLevel, nextLevel, direction) {
				isSafeReport = false
				break
			}
		}

		if isSafeReport {
			nSafeReports++
		}
	}

	return nSafeReports
}

func validTransition(src int, nxt int, dir OrderDirection) bool {
	if dir == Unspecified {
		return false
	}

	diff := nxt - src
	if dir == Increasing {
		return diff >= 1 && diff <= 3
	}
	if dir == Decreasing {
		return diff <= -1 && diff >= -3
	}

	return false
}

func getLevels(report string) []int {
	levelsParsed := strings.Split(report, " ")
	nLevels := len(levelsParsed)

	levelsList := make([]int, 0, nLevels)
	for _, levelParsed := range levelsParsed {
		level, err := strconv.Atoi(levelParsed)
		if err != nil {
			fmt.Errorf("invalid line format: %q", levelParsed)
			continue
		}

		levelsList = append(levelsList, level)
	}

	return levelsList
}
