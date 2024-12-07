package main

import (
	"aoc/utils"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines := utils.ReadInput("input.txt")
	pairs, err := parseInput(lines)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing input: %v\n", err)
		os.Exit(1)
	}

	leftValues := make([]int, len(pairs))
	rightValues := make([]int, len(pairs))

	for i, pair := range pairs {
		leftValues[i] = pair.Left
		rightValues[i] = pair.Right
	}

	sortedLeft := slices.Clone(leftValues)
	sortedRight := slices.Clone(rightValues)
	slices.Sort(sortedLeft)
	slices.Sort(sortedRight)

	distance := calculateDistance(sortedLeft, sortedRight)
	similarity := calculateSimilarity(leftValues, rightValues)

	fmt.Printf("Part 1 (Total Distance): %d\n", distance)
	fmt.Printf("Part 2 (Similarity Score): %d\n", similarity)
}

func calculateSimilarity(leftValues []int, rightValues []int) int {
	similarity := 0

	rightFreq := make(map[int]int)
	for _, v := range rightValues {
		rightFreq[v]++
	}

	for _, v := range leftValues {
		similarity += v * rightFreq[v]
	}

	return similarity
}

// Calculate absolute difference
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func calculateDistance(leftValues, rightValues []int) int {
	sum := 0
	for i := range leftValues {
		sum += abs(leftValues[i] - rightValues[i])
	}
	return sum
}

type ParseError struct {
	line string
	err  error
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("failed to parse line %q: %v", e.line, e.err)
}

func parseInput(lines []string) ([]utils.Pair[int], error) {
	// capacity = 3rd argument
	pairs := make([]utils.Pair[int], 0, len(lines))

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		parts := strings.Split(line, "   ")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line format: %q", line)
		}

		Left, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, &ParseError{line, err}
		}

		Right, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, &ParseError{line, err}
		}

		pairs = append(pairs, utils.Pair[int]{Left, Right})
	}

	return pairs, nil
}
