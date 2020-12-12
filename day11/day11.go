package day11

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type (
	Seat          int
	NextSeatState func(inputSeats [][]Seat, row, col int) Seat
)

const (
	None Seat = iota
	Occupied
	Free
)

func (s Seat) String() string {
	switch s {
	case None:
		return "."
	case Occupied:
		return "#"
	case Free:
		return "L"
	}
	return "?"
}

func Day11(part int) (err error) {
	result := 0

	content, err := ioutil.ReadFile("day11/input.txt")
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
	seats, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	// for _, r := range seats {
	// 	fmt.Println(r)
	// }

	changed := true
	for changed {
		// fmt.Println("Iterating...")
		seats, changed = iterate(seats, getSimpleNextSeatState)
		// for _, r := range seats {
		// 	fmt.Println(r)
		// }
	}

	return countOccupied(seats), err
}

func Part2(input string) (result int, err error) {
	seats, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	// for _, r := range seats {
	// 	fmt.Println(r)
	// }

	changed := true
	for changed {
		// fmt.Println("Iterating...")
		seats, changed = iterate(seats, getExtendedNextSeatState)
		// for _, r := range seats {
		// 	fmt.Println(r)
		// }
	}

	return countOccupied(seats), err
}

func parseInput(input string) (seats [][]Seat, err error) {
	inputLines := strings.Split(input, "\n")
	seats = make([][]Seat, 0)

	for _, line := range inputLines {
		if line == "" {
			continue
		}
		seatLine := make([]Seat, 0)
		seatsInLine := strings.Split(line, "")

		for _, seat := range seatsInLine {
			if seat == "L" {
				seatLine = append(seatLine, Occupied)
			} else {
				seatLine = append(seatLine, None)
			}
		}
		seats = append(seats, seatLine)
	}

	return seats, nil
}

func iterate(inputSeats [][]Seat, getNextSeatState NextSeatState) (outputSeats [][]Seat, changed bool) {
	outputSeats = make([][]Seat, len(inputSeats))
	for r := 0; r < len(inputSeats); r++ {
		outputRow := make([]Seat, len(inputSeats[r]))
		for c := 0; c < len(inputSeats[r]); c++ {
			outputRow[c] = getNextSeatState(inputSeats, r, c)
			if outputRow[c] != inputSeats[r][c] {
				changed = true
			}
		}
		outputSeats[r] = outputRow
	}
	return outputSeats, changed
}

func getSimpleNextSeatState(inputSeats [][]Seat, row, col int) Seat {
	if inputSeats[row][col] == None {
		return None
	}
	adjacentOccupied := 0
	adjacentOccupied += isOccupied(inputSeats, row+1, col+0)
	adjacentOccupied += isOccupied(inputSeats, row+1, col+1)
	adjacentOccupied += isOccupied(inputSeats, row+0, col+1)
	adjacentOccupied += isOccupied(inputSeats, row-1, col+1)
	adjacentOccupied += isOccupied(inputSeats, row-1, col+0)
	adjacentOccupied += isOccupied(inputSeats, row-1, col-1)
	adjacentOccupied += isOccupied(inputSeats, row+0, col-1)
	adjacentOccupied += isOccupied(inputSeats, row+1, col-1)

	if inputSeats[row][col] == Free && adjacentOccupied == 0 {
		return Occupied
	}

	if inputSeats[row][col] == Occupied && adjacentOccupied >= 4 {
		return Free
	}

	return inputSeats[row][col]
}

func isOccupied(inputSeats [][]Seat, row, col int) int {
	if row < 0 || col < 0 || row >= len(inputSeats) || col >= len(inputSeats[row]) {
		return 0
	}
	if inputSeats[row][col] == None || inputSeats[row][col] == Free {
		return 0
	}

	return 1
}

func getExtendedNextSeatState(inputSeats [][]Seat, row, col int) Seat {
	if inputSeats[row][col] == None {
		return None
	}
	adjacentOccupied := 0
	adjacentOccupied += isOccupiedInDirection(inputSeats, row, col, +1, +0)
	adjacentOccupied += isOccupiedInDirection(inputSeats, row, col, +1, +1)
	adjacentOccupied += isOccupiedInDirection(inputSeats, row, col, +0, +1)
	adjacentOccupied += isOccupiedInDirection(inputSeats, row, col, -1, +1)
	adjacentOccupied += isOccupiedInDirection(inputSeats, row, col, -1, +0)
	adjacentOccupied += isOccupiedInDirection(inputSeats, row, col, -1, -1)
	adjacentOccupied += isOccupiedInDirection(inputSeats, row, col, +0, -1)
	adjacentOccupied += isOccupiedInDirection(inputSeats, row, col, +1, -1)

	if inputSeats[row][col] == Free && adjacentOccupied == 0 {
		return Occupied
	}

	if inputSeats[row][col] == Occupied && adjacentOccupied >= 5 {
		return Free
	}

	return inputSeats[row][col]
}

func isOccupiedInDirection(inputSeats [][]Seat, row, col, rowDir, colDir int) int {
	for {
		row, col = row+rowDir, col+colDir

		if row < 0 || col < 0 || row >= len(inputSeats) || col >= len(inputSeats[row]) {
			return 0
		}
		if inputSeats[row][col] == None {
			continue
		}

		if inputSeats[row][col] == Free {
			return 0
		}

		return 1
	}
}

func countOccupied(inputSeats [][]Seat) (result int) {
	for _, r := range inputSeats {
		for _, c := range r {
			if c == Occupied {
				result++
			}
		}
	}
	return result
}
