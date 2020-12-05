package day5

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type BoardingPass struct {
	row, seat int
}

func (pass *BoardingPass) GetID() int {
	return pass.row*8 + pass.seat
}

func Day5(part int) (err error) {
	result := 0

	content, err := ioutil.ReadFile("day5/input.txt")
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
	boardingPasses, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	for _, pass := range boardingPasses {
		if result < pass.GetID() {
			result = pass.GetID()
		}
	}

	return result, err
}

func Part2(input string) (result int, err error) {
	boardingPasses, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	seatIDs := []int{}

	for _, pass := range boardingPasses {
		seatIDs = append(seatIDs, pass.GetID())
	}
	sort.Ints(seatIDs)
	maxID, err := Part1(input)
	if err != nil {
		return 0, err
	}

	for i, s := 0, seatIDs[0]; i <= maxID; i, s = i+1, s+1 {
		if seatIDs[i] != s {
			result = s
			break
		}
	}

	return result, err
}

func parseInput(input string) ([]*BoardingPass, error) {
	inputLines := strings.Split(input, "\n")
	boardingPasses := []*BoardingPass{}

	for _, line := range inputLines {
		if line == "" {
			continue
		}
		pass := &BoardingPass{row: 0, seat: 0}
		for i := 0; i < 7; i++ {
			if line[i] == 'B' {
				pass.row++
			}
			if i < 6 {
				pass.row <<= 1
			}
		}
		for i := 7; i < 10; i++ {
			if line[i] == 'R' {
				pass.seat++
			}
			if i < 9 {
				pass.seat <<= 1
			}
		}
		boardingPasses = append(boardingPasses, pass)
	}

	return boardingPasses, nil
}
