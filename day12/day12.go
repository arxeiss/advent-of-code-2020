package day12

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
	Left
	Right
	Forward
)

type Movement struct {
	MoveDir Direction
	Steps   int
}

type Ship struct {
	X, Y        int
	WaypointX   int
	WaypointY   int
	PointingDir Direction
}

func (ship *Ship) doMove(dir Direction, steps int) error {
	switch dir {
	case North:
		ship.Y += steps
	case East:
		ship.X += steps
	case South:
		ship.Y -= steps
	case West:
		ship.X -= steps
	default:
		return fmt.Errorf("doMove accept only North, East, South, West directions")
	}
	return nil
}

func (ship *Ship) Move(movement *Movement) error {
	switch movement.MoveDir {
	case North, East, South, West:
		return ship.doMove(movement.MoveDir, movement.Steps)
	case Left, Right:
		steps := int(movement.Steps / 90)
		if movement.MoveDir == Left {
			steps = 4 - steps
		}
		ship.PointingDir = Direction((int(ship.PointingDir) + steps) % 4)
	default:
		return ship.doMove(ship.PointingDir, movement.Steps)
	}
	return nil
}

func (ship *Ship) WaypointMove(movement *Movement) error {
	switch movement.MoveDir {
	case North:
		ship.WaypointY += movement.Steps
	case East:
		ship.WaypointX += movement.Steps
	case South:
		ship.WaypointY -= movement.Steps
	case West:
		ship.WaypointX -= movement.Steps
	case Left, Right:
		steps := int(movement.Steps / 90)
		if movement.MoveDir == Left {
			steps = 4 - steps
		}
		for i := 0; i < steps; i++ {
			ship.WaypointX, ship.WaypointY = ship.WaypointY, -ship.WaypointX
		}
	default:
		ship.X += ship.WaypointX * movement.Steps
		ship.Y += ship.WaypointY * movement.Steps
	}
	return nil
}

func (ship *Ship) Position() int {
	x, y := ship.X, ship.Y
	if x < 0 {
		x *= -1
	}
	if y < 0 {
		y *= -1
	}
	return x + y
}

func Day12(part int) (err error) {
	result := 0

	content, err := ioutil.ReadFile("day12/input.txt")
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
	movements, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	ship := &Ship{X: 0, Y: 0, PointingDir: East}
	for _, m := range movements {
		if err := ship.Move(m); err != nil {
			return 0, err
		}
	}

	return ship.Position(), err
}

func Part2(input string) (result int, err error) {
	movements, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	ship := &Ship{X: 0, Y: 0, PointingDir: East, WaypointX: 10, WaypointY: 1}
	for _, m := range movements {
		if err := ship.WaypointMove(m); err != nil {
			return 0, err
		}
	}

	return ship.Position(), err
}

func parseInput(input string) (movements []*Movement, err error) {
	inputLines := strings.Split(input, "\n")
	movements = make([]*Movement, 0)

	for _, line := range inputLines {
		if line == "" {
			continue
		}
		nextMove := &Movement{}

		switch line[0] {
		case 'N':
			nextMove.MoveDir = North
		case 'E':
			nextMove.MoveDir = East
		case 'S':
			nextMove.MoveDir = South
		case 'W':
			nextMove.MoveDir = West
		case 'L':
			nextMove.MoveDir = Left
		case 'R':
			nextMove.MoveDir = Right
		case 'F':
			nextMove.MoveDir = Forward
		}

		nextMove.Steps, err = strconv.Atoi(line[1:])
		if err != nil {
			return nil, err
		}

		movements = append(movements, nextMove)
	}

	return movements, nil
}
