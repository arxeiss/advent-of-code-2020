package day9

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Day9(part int) (err error) {
	result := 0

	content, err := ioutil.ReadFile("day9/input.txt")
	if err != nil {
		return err
	}

	if part == 1 {
		result, err = Part1(string(content), 25)
	} else {
		part = 2
		result, err = Part2(string(content), 25)
	}
	if err != nil {
		return err
	}

	fmt.Printf("Done, result of part %d is %d \n", part, result)
	return nil
}

func Part1(input string, preamble int) (result int, err error) {
	xmasData, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	result = firstXMASViolation(xmasData, preamble)

	return result, err
}

func Part2(input string, preamble int) (result int, err error) {
	_, err = parseInput(input)
	if err != nil {
		return 0, err
	}

	xmasData, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	sumToSearch := firstXMASViolation(xmasData, preamble)
	result = findInvalidSet(xmasData, sumToSearch)

	return result, err
}

func parseInput(input string) (xmasData []int, err error) {
	inputLines := strings.Split(input, "\n")
	xmasData = make([]int, 0)

	for _, line := range inputLines {
		if line == "" {
			continue
		}
		number, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		xmasData = append(xmasData, number)
	}

	return xmasData, nil
}

func firstXMASViolation(xmasData []int, preamble int) (result int) {
	for i := 0; i < len(xmasData)-preamble; i++ {
		if !containsSum(xmasData[i:i+preamble], xmasData[i+preamble]) {
			return xmasData[i+preamble]
		}
	}
	return -1
}

func containsSum(xmasData []int, searchSum int) bool {
	for i := 0; i < len(xmasData); i++ {
		for j := i + 1; j < len(xmasData); j++ {
			if i == j {
				continue
			}
			if xmasData[i]+xmasData[j] == searchSum {
				return true
			}
		}
	}

	return false
}

func findInvalidSet(xmasData []int, sumToSearch int) (result int) {
	for i := 0; i < len(xmasData); i++ {
		sum := xmasData[i]
		min, max := sum, sum

		for j := i + 1; j < len(xmasData); j++ {
			if xmasData[j] < min {
				min = xmasData[j]
			}
			if xmasData[j] > max {
				max = xmasData[j]
			}
			sum += xmasData[j]
			if sum == sumToSearch {
				return min + max
			}
			if sum > sumToSearch {
				break
			}
		}
	}

	return 0
}
