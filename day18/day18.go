package day18

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Day18(part int) (err error) {
	result := 0

	content, err := ioutil.ReadFile("day18/input.txt")
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
	equations, err := parseInput(input)

	for _, e := range equations {
		p := NewParser(e, SameMulAsAdd)
		ast, err := p.Run()
		if err != nil {
			return 0, err
		}
		result += ast.Eval()
	}

	return result, nil
}

func Part2(input string) (result int, err error) {
	equations, err := parseInput(input)

	for _, e := range equations {
		p := NewParser(e, FlipAddAndMul)
		ast, err := p.Run()
		if err != nil {
			return 0, err
		}
		result += ast.Eval()
	}

	return result, nil
}

func parseInput(input string) (equations []Expression, err error) {
	inputLines := strings.Split(input, "\n")

	for _, line := range inputLines {
		if line == "" {
			continue
		}
		expr, err := Tokenize(line)
		if err != nil {
			return nil, err
		}
		equations = append(equations, expr)
	}

	return
}
