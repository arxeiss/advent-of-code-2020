package day17

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

type Coordinate struct {
	X, Y, Z, W int
}

func NewCoordinate(x, y, z, w int) Coordinate {
	return Coordinate{
		X: x,
		Y: y,
		Z: z,
		W: w,
	}
}

type World map[Coordinate]bool

func (world World) Size(handleW bool) (xB, xE, yB, yE, zB, zE, wB, wE int) {
	xB, yB, zB, wB = math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64
	xE, yE, zE, wE = math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64

	for coordinate := range world {
		// Handle beginning
		if coordinate.X < xB {
			xB = coordinate.X
		}
		if coordinate.Y < yB {
			yB = coordinate.Y
		}
		if coordinate.Z < zB {
			zB = coordinate.Z
		}
		if coordinate.W < wB {
			wB = coordinate.W
		}
		// Handle end
		if coordinate.X > xE {
			xE = coordinate.X
		}
		if coordinate.Y > yE {
			yE = coordinate.Y
		}
		if coordinate.Z > zE {
			zE = coordinate.Z
		}
		if coordinate.W > wE {
			wE = coordinate.W
		}
	}
	return
}

func (world World) CountActiveNeighbors(xS, yS, zS, wS int, handleW bool) (neighbors int) {
	for x := xS - 1; x <= xS+1; x++ {
		for y := yS - 1; y <= yS+1; y++ {
			for z := zS - 1; z <= zS+1; z++ {
				if handleW {
					for w := wS - 1; w <= wS+1; w++ {
						if x == xS && y == yS && z == zS && w == wS {
							continue
						}
						if _, active := world[NewCoordinate(x, y, z, w)]; active {
							neighbors++
						}
					}
				} else {
					if x == xS && y == yS && z == zS {
						continue
					}
					if _, active := world[NewCoordinate(x, y, z, 0)]; active {
						neighbors++
					}
				}
			}
		}
	}

	return
}

func Day17(part int) (err error) {
	result := 0

	content, err := ioutil.ReadFile("day17/input.txt")
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
	world, err := parseInput(input)

	for i := 0; i < 6; i++ {
		world = iterateWorld(world, false)
	}

	return len(world), nil
}

func Part2(input string) (result int, err error) {
	world, err := parseInput(input)

	for i := 0; i < 6; i++ {
		world = iterateWorld(world, true)
	}

	return len(world), nil
}

func iterateWorld(world World, handleW bool) World {
	newWorld := make(World)
	xB, xE, yB, yE, zB, zE, wB, wE := world.Size(handleW)

	addToWorld := func(x, y, z, w int) {
		activeNeighbor := world.CountActiveNeighbors(x, y, z, w, handleW)
		if _, active := world[NewCoordinate(x, y, z, w)]; active {
			if activeNeighbor == 2 || activeNeighbor == 3 {
				newWorld[NewCoordinate(x, y, z, w)] = true
			}
		} else {
			if activeNeighbor == 3 {
				newWorld[NewCoordinate(x, y, z, w)] = true
			}
		}
	}

	for x := xB - 1; x <= xE+1; x++ {
		for y := yB - 1; y <= yE+1; y++ {
			for z := zB - 1; z <= zE+1; z++ {
				if handleW {
					for w := wB - 1; w <= wE+1; w++ {
						addToWorld(x, y, z, w)
					}
				} else {
					addToWorld(x, y, z, 0)
				}
			}
		}
	}

	return newWorld
}

func parseInput(input string) (world World, err error) {
	inputLines := strings.Split(input, "\n")
	world = make(World)

	for y, line := range inputLines {
		for x, c := range line {
			if c == '#' {
				world[NewCoordinate(x, y, 0, 0)] = true
			}
		}
	}

	return
}
