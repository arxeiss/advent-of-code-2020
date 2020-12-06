package day6

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type GroupAnswers struct {
	answers     map[byte]int
	respondents int
}

func Day6(part int) (err error) {
	result := 0

	content, err := ioutil.ReadFile("day6/input.txt")
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
	groupsAnswers, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	for _, group := range groupsAnswers {
		result += len(group.answers)
	}

	return result, err
}

func Part2(input string) (result int, err error) {
	groupsAnswers, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	for _, group := range groupsAnswers {
		_ = group
		for _, yesAnswers := range group.answers {
			if yesAnswers == group.respondents {
				result++
			}
		}
	}

	return result, err
}

func parseInput(input string) ([]*GroupAnswers, error) {
	inputLines := strings.Split(input, "\n")
	groupsAnswers := []*GroupAnswers{}
	currentGroup := &GroupAnswers{
		answers:     map[byte]int{},
		respondents: 0,
	}

	for _, line := range inputLines {
		if line == "" {
			if currentGroup.respondents > 0 {
				groupsAnswers = append(groupsAnswers, currentGroup)
			}
			currentGroup = &GroupAnswers{
				answers:     map[byte]int{},
				respondents: 0,
			}
			continue
		}
		for i := 0; i < len(line); i++ {
			if val, exists := currentGroup.answers[line[i]]; exists {
				currentGroup.answers[line[i]] = val + 1
			} else {
				currentGroup.answers[line[i]] = 1
			}
		}
		currentGroup.respondents++
	}
	// If there is no new line at the end
	if currentGroup.respondents > 0 {
		groupsAnswers = append(groupsAnswers, currentGroup)
	}

	return groupsAnswers, nil
}
