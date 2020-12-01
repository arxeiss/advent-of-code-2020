package day1

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func Day1(part int) (err error) {
	result := 0

	content, err := ioutil.ReadFile("day1/input.txt")
	if err != nil {
		return err
	}

	if part == 1 {
		_, _, result, err = Part1(string(content))
	} else {
		part = 2
		_, _, _, result, err = Part2(string(content))
	}
	if err != nil {
		return err
	}

	fmt.Printf("Done, result of part %d is %d \n", part, result)
	return nil
}

func Part1(input string) (a, b, result int, err error) {
	intItems, err := getIntSlice(input)
	for lower := 0; lower < len(intItems); lower++ {
		for upper := len(intItems) - 1; upper > lower; upper-- {
			a = intItems[lower]
			b = intItems[upper]
			if a+b == 2020 && a != 0 && b != 0 {
				return a, b, a * b, nil
			}
			// If it is lower, we have to go to next lower, because now the upper is going to be less and less
			if a+b < 2020 {
				break
			}
		}
	}

	return 0, 0, 0, nil
}

func Part2(input string) (a, b, c, result int, err error) {
	intItems, err := getIntSlice(input)

	for lower := 0; lower < len(intItems); lower++ {
		for middle := lower + 1; middle < len(intItems); middle++ {
			for upper := middle + 1; upper < len(intItems); upper++ {
				a = intItems[lower]
				b = intItems[middle]
				c = intItems[upper]
				if a+b+c == 2020 && a != 0 && b != 0 && c != 0 {
					return a, b, c, a * b * c, nil
				}
			}
		}
	}
	return 0, 0, 0, 0, nil
}

func getIntSlice(input string) ([]int, error) {
	stringItems := strings.Split(input, "\n")
	intItems := []int{}
	for _, strItem := range stringItems {
		if strItem == "" {
			continue
		}
		intItem, err := strconv.Atoi(strItem)
		if err != nil {
			return nil, fmt.Errorf("invalid '%s', %v", strItem, err)
		}
		if intItem > 2020 {
			continue
		}
		intItems = append(intItems, intItem)
	}
	sort.Ints(intItems)

	return intItems, nil
}
