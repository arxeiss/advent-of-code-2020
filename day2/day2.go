package day2

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type passwordEntry struct {
	password       string
	requiredLetter string
	minAmount      int
	maxAmount      int
}

func Day2(part int) (err error) {
	result := 0

	content, err := ioutil.ReadFile("day2/input.txt")
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
	entries, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	for _, entry := range entries {
		cnt := strings.Count(entry.password, entry.requiredLetter)
		if entry.minAmount <= cnt && cnt <= entry.maxAmount {
			result++
		}
	}
	return result, err
}

func Part2(input string) (result int, err error) {
	entries, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	for _, entry := range entries {
		if len(entry.password) < entry.maxAmount {
			continue
		}
		matchFirst := entry.password[entry.minAmount-1] == entry.requiredLetter[0]
		matchSecond := entry.password[entry.maxAmount-1] == entry.requiredLetter[0]
		if matchFirst != matchSecond {
			result++
		}
	}
	return result, err
}

func parseInput(input string) ([]passwordEntry, error) {
	patternRegex := regexp.MustCompile(`^(\d+)\-(\d+) ([a-z]): ([a-z]+)$`)
	var err error

	stringItems := strings.Split(input, "\n")
	entries := []passwordEntry{}
	for _, strItem := range stringItems {
		if strItem == "" {
			continue
		}
		matches := patternRegex.FindStringSubmatch(strItem)
		entry := passwordEntry{
			password:       matches[4],
			requiredLetter: matches[3],
		}
		entry.minAmount, err = strconv.Atoi(matches[1])
		if err != nil {
			return nil, fmt.Errorf("Cannot convert '%s' to number: %v", matches[1], err)
		}
		entry.maxAmount, err = strconv.Atoi(matches[2])
		if err != nil {
			return nil, fmt.Errorf("Cannot convert '%s' to number: %v", matches[2], err)
		}
		entries = append(entries, entry)
	}

	return entries, nil
}
