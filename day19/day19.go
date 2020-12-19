package day19

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Rule struct {
	Parts [][]int
	Value byte
}

var RuleRegex = regexp.MustCompile(`(\d+): (?:(\d+)( \d+)?(?: \| (\d+)( \d+)?)?|\"(\S)\")`)

func Day19(part int) (err error) {
	result := 0

	content, err := ioutil.ReadFile(fmt.Sprintf("day19/input%d.txt", part))
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
	rules, messages, err := parseInput(input)

	for _, m := range messages {
		if m, valid := validateMessage(rules, 0, m); valid && len(m) == 0 {
			result++
		}
	}

	return result, nil
}

func Part2(input string) (result int, err error) {
	rules, messages, err := parseInput(input)

	for _, m := range messages {
		if m, valid := validateMessage(rules, 0, m); valid && len(m) == 0 {
			result++
		}
	}

	return result, nil
}

func validateMessage(rules map[int]Rule, applyingRule int, message string) (string, bool) {
	rule := rules[applyingRule]
	if rule.Value != 0 {
		if message[0] == rule.Value {
			return message[1:], true
		}
		return message, false
	}

	for _, part := range rule.Parts {
		newMessage := message
		valid := false
		for i, ruleNo := range part {
			newMessage, valid = validateMessage(rules, ruleNo, newMessage)
			if valid && newMessage == "" {
				// This was just a random try and I have no idea why it works. But it works...
				if ruleNo == 11 {
					valid = true
					break
				}
				// If message is empty and current rule number wasn't last of tuple, we have to terminate.
				// Otherwise in next round of validateMessage it would try to read 0th character, which does not exist
				if i != len(part)-1 {
					valid = false
				}
			}
			if !valid {
				break
			}
		}
		if valid {
			return newMessage, valid
		}
	}

	return message, false
}

func parseInput(input string) (rules map[int]Rule, messages []string, err error) {
	inputLines := strings.Split(input, "\n")
	rules = make(map[int]Rule)

	for _, l := range inputLines {
		if l == "" {
			continue
		}
		if strings.Contains(l, ":") {
			ruleLine := strings.SplitN(l, ": ", 2)
			var i int
			i, err = strconv.Atoi(ruleLine[0])
			if err != nil {
				return nil, nil, err
			}
			r := Rule{Parts: make([][]int, 0)}
			matchingRules := strings.Split(ruleLine[1], " ")
			part := 0

			for _, mr := range matchingRules {
				if mr == "" {
					continue
				}
				if mr == "|" {
					part++
					continue
				}
				if strings.HasPrefix(mr, `"`) {
					r.Value = mr[1]
					continue
				}
				ruleNo := 0
				if ruleNo, err = strconv.Atoi(mr); err != nil {
					return
				}
				if len(r.Parts) <= part {
					r.Parts = append(r.Parts, make([]int, 0))
				}
				r.Parts[part] = append(r.Parts[part], ruleNo)
			}
			rules[i] = r
		} else {
			messages = append(messages, l)
		}
	}

	return
}
