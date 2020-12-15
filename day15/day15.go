package day15

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Instruction struct {
	Mask    string
	Address int
	Value   int
}

func Day15(part int) (err error) {
	result := 0

	content, err := ioutil.ReadFile("day15/input.txt")
	if err != nil {
		return err
	}

	if part == 1 {
		result, err = Solve(string(content), 2020)
	} else {
		part = 2
		result, err = Solve(string(content), 30000000)
	}
	if err != nil {
		return err
	}

	fmt.Printf("Done, result of part %d is %d \n", part, result)
	return nil
}

func Solve(input string, nth int) (result int, err error) {
	numbers, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	mem := map[int]int{}
	lastSaid := -1
	for i, v := range numbers {
		mem[v] = i + 1
		lastSaid = v
	}
	delete(mem, lastSaid)
	for i := len(numbers) + 1; i <= nth; i++ {
		if ith, exists := mem[lastSaid]; exists {
			mem[lastSaid] = i - 1
			lastSaid = i - ith - 1
		} else {
			mem[lastSaid] = i - 1
			lastSaid = 0
		}
	}

	return lastSaid, nil
}

func parseInput(input string) (numbers []int, err error) {
	inputNumbers := strings.Split(strings.TrimSpace(input), ",")

	for _, number := range inputNumbers {
		if number == "" {
			continue
		}
		num, err := strconv.Atoi(number)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}

	return
}
