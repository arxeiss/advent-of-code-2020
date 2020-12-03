package day3

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Day3(part int) (err error) {
	result := 0

	content, err := ioutil.ReadFile("day3/input.txt")
	if err != nil {
		return err
	}

	if part == 1 {
		result, err = Part1(string(content))
	} else {
		part = 2
		result, err = Part2(string(content))
	}
	if err != nil {
		return err
	}

	fmt.Printf("Done, result of part %d is %d \n", part, result)
	return nil
}

func Part1(input string) (result int, err error) {
	forrest, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	result = treesInSlope(forrest, 3, 1)

	return result, err
}

func Part2(input string) (result int, err error) {
	forrest, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	result = treesInSlope(forrest, 1, 1) *
		treesInSlope(forrest, 3, 1) *
		treesInSlope(forrest, 5, 1) *
		treesInSlope(forrest, 7, 1) *
		treesInSlope(forrest, 1, 2)
	return result, err
}

func parseInput(input string) ([][]bool, error) {
	forrest := [][]bool{}
	inputMap := strings.Split(input, "\n")
	for _, mapLine := range inputMap {
		if mapLine == "" {
			continue
		}
		forrestLine := make([]bool, len(mapLine))
		for i := 0; i < len(mapLine); i++ {
			forrestLine[i] = mapLine[i] == '#'
		}

		forrest = append(forrest, forrestLine)
	}

	return forrest, nil
}

func treesInSlope(forrest [][]bool, right, down int) (result int) {
	for r, c := 0, 0; r < len(forrest); r, c = r+down, (c+right)%len(forrest[0]) {
		if forrest[r][c] {
			result++
		}
	}
	return
}
