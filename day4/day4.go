package day4

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

type Passport map[string]string

var (
	requiredFields  = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	heightRegex     = regexp.MustCompile(`^([0-9]+)(cm|in)$`)
	hairRegex       = regexp.MustCompile(`^#[0-9a-f]{6}$`)
	eyeRegex        = regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	passportIDRegex = regexp.MustCompile(`^[0-9]{9}$`)
)

func Day4(part int) (err error) {

	result := 0

	content, err := ioutil.ReadFile("day4/input.txt")
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
	passports, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	for _, pass := range passports {
		if hasPassportAllRequiredFields(pass) {
			result++
		}
	}

	return result, err
}

func Part2(input string) (result int, err error) {
	passports, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	for _, pass := range passports {
		if hasPassportValidValues(pass) {
			result++
		}
	}

	return result, err
}

func parseInput(input string) ([]Passport, error) {
	inputLines := strings.Split(input, "\n")
	passports := []Passport{}
	currentPassport := Passport{}

	for _, line := range inputLines {
		if line == "" {
			if len(currentPassport) > 0 {
				passports = append(passports, currentPassport)
			}
			currentPassport = Passport{}
			continue
		}
		attributes := strings.Split(line, " ")
		for _, attr := range attributes {
			if attr == "" {
				continue
			}
			entryAttr := strings.SplitN(attr, ":", 2)
			currentPassport[entryAttr[0]] = entryAttr[1]
		}
	}
	// If there is no new line at the end
	if len(currentPassport) > 0 {
		passports = append(passports, currentPassport)
	}

	return passports, nil
}

func hasPassportAllRequiredFields(passport Passport) bool {
	for _, field := range requiredFields {
		if _, ok := passport[field]; !ok {
			return false
		}
	}
	return true
}

func hasPassportValidValues(passport Passport) bool {
	if !hasPassportAllRequiredFields(passport) {
		return false
	}

	// As the numbers here have always same length, they can be kept in string
	if passport["byr"] < "1920" || passport["byr"] > "2002" {
		return false
	}
	if passport["iyr"] < "2010" || passport["iyr"] > "2020" {
		return false
	}
	if passport["eyr"] < "2020" || passport["eyr"] > "2030" {
		return false
	}

	parsedHeight := heightRegex.FindStringSubmatch(passport["hgt"])
	if parsedHeight == nil {
		return false
	}
	// If in, the number must be at least 59 and at most 76.
	if parsedHeight[2] == "in" && (parsedHeight[1] < "59" || parsedHeight[1] > "76") {
		return false
	}
	// If cm, the number must be at least 150 and at most 193.
	if parsedHeight[2] == "cm" && (parsedHeight[1] < "150" || parsedHeight[1] > "193") {
		return false
	}

	if !hairRegex.MatchString(passport["hcl"]) {
		return false
	}
	if !eyeRegex.MatchString(passport["ecl"]) {
		return false
	}
	if !passportIDRegex.MatchString(passport["pid"]) {
		return false
	}

	return true
}
