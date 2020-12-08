package day8

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var (
	instructionRegex = regexp.MustCompile(`^([a-z]+) ((?:\+|-)\d+)$`)
)

type Instruction struct {
	Name     string
	Argument int
}

func Day8(part int) (err error) {
	result := 0

	content, err := ioutil.ReadFile("day8/input.txt")
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
	instructions, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	var halted bool
	result, halted = getAcc(instructions)
	if !halted {
		return 0, fmt.Errorf("Expected to be halted")
	}

	return result, err
}

func Part2(input string) (_ int, err error) {
	instructions, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	for i := 0; i < len(instructions); i++ {
		if instructions[i].Name == "acc" {
			continue
		}

		instructionsCopy := make([]*Instruction, len(instructions))
		for j := 0; j < len(instructions); j++ {
			newInstructionName := instructions[j].Name
			if i == j {
				if newInstructionName == "jmp" {
					newInstructionName = "nop"
				} else if newInstructionName == "nop" {
					newInstructionName = "jmp"
				}
			}
			instructionsCopy[j] = &Instruction{
				Name:     newInstructionName,
				Argument: instructions[j].Argument,
			}
		}

		acc, halted := getAcc(instructionsCopy)
		if !halted {
			return acc, nil
		}
	}

	return 0, nil
}

func parseInput(input string) (instructions []*Instruction, err error) {
	inputLines := strings.Split(input, "\n")
	instructions = make([]*Instruction, 0)

	for _, line := range inputLines {
		args := instructionRegex.FindStringSubmatch(line)
		if args == nil {
			continue
		}
		ins := &Instruction{Name: args[1]}
		ins.Argument, err = strconv.Atoi(args[2])
		if err != nil {
			return nil, err
		}

		instructions = append(instructions, ins)
	}

	return instructions, nil
}

func getAcc(instructions []*Instruction) (acc int, halted bool) {
	i := 0
	for i < len(instructions) && instructions[i].Name != "halt" {
		addToI := 0
		switch instructions[i].Name {
		case "jmp":
			addToI = instructions[i].Argument
		case "acc":
			acc += instructions[i].Argument
			fallthrough
		case "nop":
			addToI = 1
		}
		instructions[i].Name = "halt"
		i += addToI
	}

	if i < len(instructions) {
		return acc, instructions[i].Name == "halt"
	}
	return acc, false
}
