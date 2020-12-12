package day10

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func Day10(part int) (err error) {
	result := 0

	content, err := ioutil.ReadFile("day10/input.txt")
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
	adapters, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	differencesCnt := map[int]int{1: 0, 2: 0, 3: 0}
	sort.Ints(adapters)
	prev := 0
	for i := 0; i < len(adapters); i++ {
		diff := adapters[i] - prev
		prev = adapters[i]
		differencesCnt[diff]++
	}

	return differencesCnt[1] * (differencesCnt[3] + 1), err
}

func Part2(input string) (result int, err error) {
	adapters, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	adapters = append(adapters, 0)
	sort.Ints(adapters)

	countedPaths := map[int]int{adapters[len(adapters)-1]: 1}
	// The last one is already added to the map, start from the second-to-last
	for i := len(adapters) - 2; i >= 0; i-- {
		waysFromIth := 0

		current := adapters[i]
		for w := 1; w <= 3; w++ {
			if ways, exist := countedPaths[current+w]; exist {
				waysFromIth += ways
			}
		}
		if waysFromIth > 0 {
			countedPaths[adapters[i]] = waysFromIth
		}
	}

	if ways, exist := countedPaths[0]; exist {
		return ways, nil
	}
	return 0, fmt.Errorf("Bad input, there is no way...")
}

func parseInput(input string) (adapters []int, err error) {
	inputLines := strings.Split(input, "\n")
	adapters = make([]int, 0)

	for _, line := range inputLines {
		if line == "" {
			continue
		}
		number, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		adapters = append(adapters, number)
	}

	return adapters, nil
}
