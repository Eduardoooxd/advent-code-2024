package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input := readInput("input.txt")

	var leftValues []int
	var rightValues []int

	for _, line := range input {
		splitValue := strings.Split(line, "   ")
		leftInt, _ := strconv.Atoi(splitValue[0])
		rightInt, _ := strconv.Atoi(splitValue[1])

		leftValues = append(leftValues, leftInt)
		rightValues = append(rightValues, rightInt)
	}

	slices.Sort(leftValues)
	slices.Sort(rightValues)

	sum := firstChallenge(leftValues, rightValues)
	fmt.Println(sum)

	similarity := secondChallenge(leftValues, rightValues)
	fmt.Println(similarity)
}

func countElements(toFind int, values []int) int {
	count := 0

	for index, _ := range values {
		if values[index] == toFind {
			count++
		}
	}

	return count
}

func secondChallenge(leftValues []int, rightValues []int) int {
	similarity := 0
	cache := make(map[int]int)

	for index, _ := range leftValues {
		leftInt := leftValues[index]
		if cache[leftInt] == 0 {
			numberElements := countElements(leftInt, rightValues)
			similarity += leftInt * numberElements

			cache[leftInt] = numberElements
		} else {
			numberElements := cache[leftInt]
			similarity += leftInt * numberElements
		}
	}

	return similarity
}

func firstChallenge(leftValues []int, rightValues []int) int {
	sum := 0

	for index, _ := range leftValues {
		leftInt := leftValues[index]
		rightInt := rightValues[index]

		if leftInt > rightInt {
			sum += leftInt - rightInt
		} else {
			sum += rightInt - leftInt
		}
	}
	return sum
}

func readInput(filename string) []string {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
