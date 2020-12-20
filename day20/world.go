package day20

import (
	"fmt"
	"math"
	"strings"
)

type Side uint8

const (
	Top Side = iota
	Right
	Bottom
	Left
)

type Coordinate struct {
	X, Y int
}

type World map[Coordinate]*Tile

func (w World) MapRange() (minX, minY, maxX, maxY int) {
	minX, minY = math.MaxInt64, math.MaxInt64
	maxX, maxY = math.MinInt64, math.MinInt64
	for coordinate := range w {
		if minX > coordinate.X {
			minX = coordinate.X
		}
		if minY > coordinate.Y {
			minY = coordinate.Y
		}
		if maxX < coordinate.X {
			maxX = coordinate.X
		}
		if maxY < coordinate.Y {
			maxY = coordinate.Y
		}
	}
	return
}

type Tile struct {
	ID       int
	pixels   [][]byte
	Finished bool
	// Top, Right, Bottom, Left *Tile
}

var behindTheEdgeTile = &Tile{Finished: true}

func (t *Tile) Rotate() {
	newPixels := [][]byte{}

	for _, row := range t.pixels {
		newPixels = append(newPixels, make([]byte, len(row)))
	}
	for r, row := range t.pixels {
		for c, v := range row {
			newPixels[c][len(row)-1-r] = v
		}
	}

	t.pixels = newPixels
}

func (t *Tile) Flip() {
	newPixels := [][]byte{}

	for r, row := range t.pixels {
		newPixels = append(newPixels, make([]byte, len(row)))
		for c, v := range row {
			newPixels[r][len(row)-1-c] = v
		}
	}

	t.pixels = newPixels
}

func (t *Tile) GetSide(direction Side) string {
	c := 0
	switch direction {
	case Top:
		return fmt.Sprintf("%s", t.pixels[0])
	case Bottom:
		return fmt.Sprintf("%s", t.pixels[len(t.pixels)-1])
	case Right:
		c = len(t.pixels[0]) - 1
	}
	str := make([]byte, len(t.pixels))
	for r, row := range t.pixels {
		str[r] = row[c]
	}
	return string(str)
}

func (t *Tile) String() string {
	lines := []string{}
	for _, l := range t.pixels {
		lines = append(lines, fmt.Sprintf("%s", l))
	}
	return strings.Join(lines, "\n") + "\n"
}
