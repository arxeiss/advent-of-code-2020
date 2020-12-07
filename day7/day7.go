package day7

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var (
	bagRuleRegex      = regexp.MustCompile(`^([a-z ]+) bags? contains? (.*)\.$`)
	innerBagRuleRegex = regexp.MustCompile(`([0-9]+) ([a-z ]+) bags?$`)
)

type Bag struct {
	Name     string
	Contains map[string]int
}

func (bag *Bag) CanContains(bagName string) bool {
	_, contains := bag.Contains[bagName]
	return contains
}

func Day7(part int) (err error) {
	result := 0

	content, err := ioutil.ReadFile("day7/input.txt")
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
	bags, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	result = len(canHandleSelectedOnes(
		map[string]bool{"shiny gold": true},
		0,
		bags,
	))

	return result, err
}

func Part2(input string) (result int, err error) {
	bags, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	result = howManyBagsFitsInto("shiny gold", bags)

	return result, err
}

func parseInput(input string) (_ []*Bag, err error) {
	inputLines := strings.Split(input, "\n")
	bags := []*Bag{}

	for _, line := range inputLines {
		if line == "" {
			continue
		}
		outerBag := bagRuleRegex.FindStringSubmatch(line)
		if outerBag == nil {
			continue
		}
		bag := &Bag{
			Name:     outerBag[1],
			Contains: map[string]int{},
		}

		innerBags := strings.Split(outerBag[2], ",")
		for _, innerBag := range innerBags {
			if innerBag == "no other bags" {
				continue
			}
			innerBagDetails := innerBagRuleRegex.FindStringSubmatch(innerBag)
			if innerBagDetails == nil {
				continue
			}
			bag.Contains[innerBagDetails[2]], err = strconv.Atoi(innerBagDetails[1])
			if err != nil {
				return nil, err
			}
		}
		bags = append(bags, bag)
	}

	return bags, nil
}

func canHandleSelectedOnes(searchFor map[string]bool, level int, bags []*Bag) map[string]bool {
	possibilities := map[string]bool{}
	if searchFor == nil || len(searchFor) == 0 {
		return possibilities
	}
	for currentlySearched := range searchFor {
		for _, bag := range bags {
			if bag.CanContains(currentlySearched) {
				possibilities[bag.Name] = true
			}
		}
	}

	canHandleFromNextLevel := canHandleSelectedOnes(possibilities, level+1, bags)
	for nextLevel := range canHandleFromNextLevel {
		possibilities[nextLevel] = true
	}

	return possibilities
}

func howManyBagsFitsInto(searchFor string, bags []*Bag) (result int) {
	for _, bag := range bags {
		if bag.Name == searchFor {
			for bagName, amount := range bag.Contains {
				result += amount + amount*howManyBagsFitsInto(bagName, bags)
			}
		}
	}

	return result
}
