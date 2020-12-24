package day24

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strings"
)

var cardinalDirections = regexp.MustCompile(`se|sw|ne|nw|e|w`)

type Coordinate struct {
	X, Y int
}

func Day24(part int) (err error) {
	result := 0

	content, err := ioutil.ReadFile(fmt.Sprintf("day24/input.txt"))
	if err != nil {
		return err
	}
	part1, part2, err := Solve(string(content))
	if err != nil {
		return err
	}

	if part == 1 {
		result = part1
	} else {
		part = 2
		result = part2
	}

	fmt.Printf("Done, result of part %d is %d \n", part, result)
	return nil
}

func Solve(input string) (part1, part2 int, err error) {
	floorTilesDirs, err := parseInput(input)
	if err != nil {
		return 0, 0, err
	}
	floorTiles := map[Coordinate]bool{}

	for _, tileDirs := range floorTilesDirs {
		coordinate := getCoordinate(tileDirs, Coordinate{0, 0})
		if _, exists := floorTiles[coordinate]; exists {
			delete(floorTiles, coordinate)
		} else {
			floorTiles[coordinate] = true
		}
	}

	part1 = len(floorTiles)

	for i := 0; i < 100; i++ {
		iterate(floorTiles)
	}

	part2 = len(floorTiles)

	return
}

func iterate(floorTiles map[Coordinate]bool) {
	minX, minY, maxX, maxY := math.MaxInt64, math.MaxInt64, math.MinInt64, math.MinInt64
	for tileCoordinate := range floorTiles {
		if minX > tileCoordinate.X {
			minX = tileCoordinate.X
		}
		if minY > tileCoordinate.Y {
			minY = tileCoordinate.Y
		}
		if maxX < tileCoordinate.X {
			maxX = tileCoordinate.X
		}
		if maxY < tileCoordinate.Y {
			maxY = tileCoordinate.Y
		}
	}
	minX--
	minY--
	maxX++
	maxY++

	toFlip := []Coordinate{}
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			currentCoordinate := Coordinate{x, y}
			neighbor := countNeighbors(floorTiles, currentCoordinate)
			_, isBlack := floorTiles[currentCoordinate]
			if isBlack && (neighbor == 0 || neighbor > 2) {
				toFlip = append(toFlip, currentCoordinate)
			} else if !isBlack && neighbor == 2 {
				toFlip = append(toFlip, currentCoordinate)
			}
		}
	}

	for _, flip := range toFlip {
		if _, isBlack := floorTiles[flip]; isBlack {
			delete(floorTiles, flip)
		} else {
			floorTiles[flip] = true
		}
	}

}

func countNeighbors(floorTiles map[Coordinate]bool, baseCoordinate Coordinate) (sum int) {
	directions := []string{"se", "sw", "ne", "nw", "e", "w"}
	for _, dir := range directions {
		if _, exists := floorTiles[getCoordinate([]string{dir}, baseCoordinate)]; exists {
			sum++
		}
	}
	return
}

func getCoordinate(tileDirs []string, coordinate Coordinate) Coordinate {
	for _, dir := range tileDirs {
		switch dir {
		case "ne":
			if coordinate.Y%2 == 0 {
				coordinate.X--
			}
			coordinate.Y++
		case "nw":
			if coordinate.Y%2 == 1 || coordinate.Y%2 == -1 {
				coordinate.X++
			}
			coordinate.Y++
		case "se":
			if coordinate.Y%2 == 0 {
				coordinate.X--
			}
			coordinate.Y--
		case "sw":
			if coordinate.Y%2 == 1 || coordinate.Y%2 == -1 {
				coordinate.X++
			}
			coordinate.Y--
		case "e":
			coordinate.X--
		case "w":
			coordinate.X++
		}
	}
	return coordinate
}

func parseInput(input string) (directions [][]string, err error) {
	lines := strings.Split(input, "\n")
	directions = make([][]string, 0)
	for _, l := range lines {
		if l == "" {
			continue
		}
		matches := cardinalDirections.FindAllStringSubmatch(l, -1)
		if matches == nil {
			return nil, fmt.Errorf("Invalid line '%s'", l)
		}
		lineDirs := []string{}
		for _, m := range matches {
			lineDirs = append(lineDirs, m[0])
		}
		directions = append(directions, lineDirs)
	}

	return
}
