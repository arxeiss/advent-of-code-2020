package day16

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var (
	ruleRegex   = regexp.MustCompile(`([a-z ]+): (\d+)-(\d+) or (\d+)-(\d+)`)
	assignRegex = regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)
)

type Rule struct {
	LowerFrom, LowerTo int
	UpperFrom, UpperTo int
}

func (r *Rule) IsValid(num int) bool {
	return (r.LowerFrom <= num && num <= r.LowerTo) || (r.UpperFrom <= num && num <= r.UpperTo)
}

func Day16(part int) (err error) {
	result := 0

	content, err := ioutil.ReadFile("day16/input.txt")
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
	rules, _, nearbyTickets, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	for _, ticket := range nearbyTickets {
		for _, inputNumber := range ticket {
			foundValid := false
			for _, rule := range rules {
				if rule.IsValid(inputNumber) {
					foundValid = true
					break
				}
			}
			if !foundValid {
				result += inputNumber
				break
			}
		}
	}

	return result, nil
}

func Part2(input string) (result int, err error) {
	rules, myTicket, nearbyTickets, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	validTickets := [][]int{}
ticketLoop:
	for _, ticket := range nearbyTickets {
		for _, inputNumber := range ticket {
			foundValid := false
			for _, rule := range rules {
				if rule.IsValid(inputNumber) {
					foundValid = true
					break
				}
			}
			if !foundValid {
				continue ticketLoop
			}
		}
		validTickets = append(validTickets, ticket)
	}

	finalRuleForColumn := map[string]int{}
	matchedColumn := map[int]bool{}

	for len(matchedColumn) < len(myTicket) {
		rulesForColumn := map[string][]int{}
		for ruleName, rule := range rules {
		columnLoop:
			for col := 0; col < len(myTicket); col++ {
				if _, matched := matchedColumn[col]; matched {
					continue
				}
				for _, ticket := range validTickets {
					if !rule.IsValid(ticket[col]) {
						continue columnLoop
					}
				}
				rulesForColumn[ruleName] = append(rulesForColumn[ruleName], col)
			}
		}
		for ruleName, columns := range rulesForColumn {
			if len(columns) == 1 {
				finalRuleForColumn[ruleName] = columns[0]
				matchedColumn[columns[0]] = true
			}
		}
	}

	result = 1
	for ruleName, colIndex := range finalRuleForColumn {
		if strings.HasPrefix(ruleName, "departure") {
			result *= myTicket[colIndex]
		}
	}

	return result, nil
}

func parseInput(input string) (rules map[string]*Rule, myTicket []int, nearbyTickets [][]int, err error) {
	inputLines := strings.Split(input, "\n")
	rules = make(map[string]*Rule)

	nextLineIsMy := false
	for _, line := range inputLines {
		if line == "" || strings.HasPrefix(line, "nearby") {
			continue
		}
		if matches := ruleRegex.FindStringSubmatch(line); matches != nil {
			rule := &Rule{}
			if rule.LowerFrom, err = strconv.Atoi(matches[2]); err != nil {
				return
			}
			if rule.LowerTo, err = strconv.Atoi(matches[3]); err != nil {
				return
			}
			if rule.UpperFrom, err = strconv.Atoi(matches[4]); err != nil {
				return
			}
			if rule.UpperTo, err = strconv.Atoi(matches[5]); err != nil {
				return
			}
			rules[matches[1]] = rule
		} else if strings.HasPrefix(line, "your ticket") {
			nextLineIsMy = true
		} else if nextLineIsMy {
			myTicket, err = parseIntLine(line)
			nextLineIsMy = false
		} else {
			var nearbyTicket []int
			nearbyTicket, err = parseIntLine(line)
			nearbyTickets = append(nearbyTickets, nearbyTicket)
		}
		if err != nil {
			return
		}
	}

	return
}
func parseIntLine(input string) (numbers []int, err error) {
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
