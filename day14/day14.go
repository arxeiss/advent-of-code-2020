package day14

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var (
	maskRegex   = regexp.MustCompile("mask = ([X01]+)")
	assignRegex = regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)
)

type Instruction struct {
	Mask    string
	Address int
	Value   int
}

func Day14(part int) (err error) {
	result := 0

	content, err := ioutil.ReadFile("day14/input.txt")
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

	memory := map[int]int{}
	currentMask := ""
	for _, ins := range instructions {
		if len(ins.Mask) > 0 {
			currentMask = strings.TrimLeft(ins.Mask, "X")
		} else {
			var maskAnd, maskOr int
			if maskAnd, err = binToInt(strings.ReplaceAll(currentMask, "X", "1")); err != nil {
				return 0, err
			}
			if maskOr, err = binToInt(strings.ReplaceAll(currentMask, "X", "0")); err != nil {
				return 0, err
			}
			memory[ins.Address] = (ins.Value & maskAnd) | maskOr
		}
	}

	for _, v := range memory {
		result += v
	}

	return result, nil
}

func Part2(input string) (result int, err error) {
	instructions, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	memory := map[int]int{}
	currentMask := ""
	for _, ins := range instructions {
		if len(ins.Mask) > 0 {
			currentMask = ins.Mask
		} else {
			// Replace all 0->1 and then X->0 and do &
			numToAnd, err := binToInt(strings.ReplaceAll(strings.ReplaceAll(currentMask, "0", "1"), "X", "0"))
			if err != nil {
				return 0, err
			}
			// Replace X->0 and do |
			numToOr, err := binToInt(strings.ReplaceAll(currentMask, "X", "0"))
			if err != nil {
				return 0, err
			}
			// The replaces above and & and | operations below change address in a way, that it contains all 1 from mask
			// and 0 where mask is X
			// Example:
			// 101010 <- original address
			// X1001X <- mask
			//
			// 011110 <- mask with replaces 0->1 and X->0
			// 001010 <- address & mask with replaces ^^
			//
			// 010010 <- mask with X->0
			// 011010 <- modified address | mask with replaces ^^
			newAddr := (ins.Address & numToAnd) | numToOr

			// Replace all 1->0 in mask, so later the | can be performed with only X changed to 0 and 1 and only
			// those bits are going to be changed, no others
			err = handleAddressAndAssign(&memory, strings.ReplaceAll(currentMask, "1", "0"), newAddr, ins.Value)
			if err != nil {
				return 0, err
			}
		}
	}

	for _, v := range memory {
		result += v
	}

	return result, nil
}

func handleAddressAndAssign(memory *map[int]int, mask string, address, value int) error {
	if firstX := strings.IndexByte(mask, 'X'); firstX >= 0 {
		if err := handleAddressAndAssign(memory, strings.Replace(mask, "X", "0", 1), address, value); err != nil {
			return err
		}
		if err := handleAddressAndAssign(memory, strings.Replace(mask, "X", "1", 1), address, value); err != nil {
			return err
		}
	} else {
		maskInt, err := binToInt(mask)
		if err != nil {
			return err
		}

		(*memory)[address|maskInt] = value
	}

	return nil
}

func binToInt(bin string) (val int, err error) {
	for i := 0; i < len(bin); i++ {
		if bin[i] == '1' {
			val += 1
		} else if bin[i] != '0' {
			return 0, fmt.Errorf("bin '%s' has invalid characters", bin)
		}
		if i < len(bin)-1 {
			val <<= 1
		}
	}
	return val, nil
}

func parseInput(input string) (instructions []*Instruction, err error) {
	inputLines := strings.Split(input, "\n")

	for _, line := range inputLines {
		if line == "" {
			continue
		}
		instruction := &Instruction{}
		if matches := maskRegex.FindStringSubmatch(line); matches != nil {
			instruction.Mask = matches[1]
		} else if matches := assignRegex.FindStringSubmatch(line); matches != nil {
			if instruction.Address, err = strconv.Atoi(matches[1]); err != nil {
				return nil, err
			}
			if instruction.Value, err = strconv.Atoi(matches[2]); err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("Unexpected line '%s'", line)
		}
		instructions = append(instructions, instruction)
	}

	return
}
