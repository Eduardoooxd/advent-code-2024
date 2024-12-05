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
