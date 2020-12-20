package day20

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var tileRegex = regexp.MustCompile(`Tile (\d+):`)

func Day20(part int) (err error) {
	result := 0

	content, err := ioutil.ReadFile(fmt.Sprintf("day20/input.txt"))
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
	tiles, err := parseInput(input)

	world := World{}
	baseTile := tiles[0]
	baseTile.Finished = true
	world[Coordinate{0, 0}] = baseTile

	findNeighborTiles(world, tiles, 0, 0)

	minX, minY, maxX, maxY := world.MapRange()

	result = world[Coordinate{minX, minY}].ID *
		world[Coordinate{minX, maxY}].ID *
		world[Coordinate{maxX, minY}].ID *
		world[Coordinate{maxX, maxY}].ID

	return result, err
}

func Part2(input string) (result int, err error) {
	tiles, err := parseInput(input)

	world := World{}
	baseTile := tiles[0]
	baseTile.Finished = true
	world[Coordinate{0, 0}] = baseTile

	findNeighborTiles(world, tiles, 0, 0)

	minX, minY, maxX, maxY := world.MapRange()

	// Just the content, without borders
	singleTileWidth, singleTileHeight := len(baseTile.pixels[0])-2, len(baseTile.pixels)-2
	finalTileWidth, finalTileHeight := (maxX-minX+1)*singleTileWidth, (maxY-minY+1)*singleTileHeight
	finalTile := &Tile{Finished: true, pixels: make([][]byte, finalTileHeight)}
	for i := 0; i < finalTileWidth; i++ {
		finalTile.pixels[i] = make([]byte, finalTileWidth)
	}

	for tileR, tileMapR := 0, maxY; tileMapR >= minY; tileMapR, tileR = tileMapR-1, tileR+1 {
		for tileC, tileMapC := 0, minX; tileMapC <= maxX; tileMapC, tileC = tileMapC+1, tileC+1 {
			tile := world[Coordinate{tileMapC, tileMapR}]
			for r, row := range tile.pixels {
				if r == 0 || r == len(tile.pixels)-1 {
					continue
				}
				for c, col := range row {
					if c == 0 || c == len(row)-1 {
						continue
					}
					finalTile.pixels[tileR*singleTileHeight+r-1][tileC*singleTileWidth+c-1] = col
				}
			}
		}
	}

	monsterRegex := regexp.MustCompile(fmt.Sprintf(
		`(.{18})#(.{%d})#(.{4})##(.{4})##(.{4})###(.{%d})#(.{2})#(.{2})#(.{2})#(.{2})#(.{2})#(.{3})`,
		1+finalTileWidth-20,
		1+finalTileWidth-20,
	))
	removedMonster := findAndRemoveMonster(finalTile, monsterRegex)
	result = strings.Count(removedMonster, "#")

	return result, nil
}

func findAndRemoveMonster(finalTile *Tile, monsterRegex *regexp.Regexp) string {
	for f := 0; f < 2; f++ {
		for i := 0; i < 4; i++ {
			tileAsLine := strings.ReplaceAll(finalTile.String(), "\n", "")
			matches := monsterRegex.FindAllStringSubmatch(tileAsLine, -1)
			if matches != nil {
				// Regex does not match overlapping monsters, run again until none are found
				for matches != nil {
					tileAsLine = monsterRegex.ReplaceAllString(
						tileAsLine,
						"${1}O${2}O${3}OO${4}OO${5}OOO${6}O${7}O${8}O${9}O${10}O${11}O${12}",
					)
					matches = monsterRegex.FindAllStringSubmatch(tileAsLine, -1)
				}
				return tileAsLine
			}

			finalTile.Rotate()
		}
		finalTile.Flip()
	}

	return ""
}

func findNeighborTiles(world World, tiles []*Tile, x, y int) {
	// Find on the right side
	if _, exists := world[Coordinate{x + 1, y}]; !exists {
		if match := findMatch(world, tiles, x, y, x+1, y, Right, Left); match != nil {
			match.Finished = true
			world[Coordinate{x + 1, y}] = match

			findNeighborTiles(world, tiles, x+1, y)
		}
	}

	// Find on the top side
	if _, exists := world[Coordinate{x, y + 1}]; !exists {
		if match := findMatch(world, tiles, x, y, x, y+1, Top, Bottom); match != nil {
			match.Finished = true
			world[Coordinate{x, y + 1}] = match

			findNeighborTiles(world, tiles, x, y+1)
		}
	}

	// Find on the left side
	if _, exists := world[Coordinate{x - 1, y}]; !exists {
		if match := findMatch(world, tiles, x, y, x-1, y, Left, Right); match != nil {
			match.Finished = true
			world[Coordinate{x - 1, y}] = match

			findNeighborTiles(world, tiles, x-1, y)
		}
	}

	// Find on the bottom side
	if _, exists := world[Coordinate{x, y - 1}]; !exists {
		if match := findMatch(world, tiles, x, y, x, y-1, Bottom, Top); match != nil {
			match.Finished = true
			world[Coordinate{x, y - 1}] = match

			findNeighborTiles(world, tiles, x, y-1)
		}
	}
}

func findMatch(world World, tiles []*Tile, baseX, baseY, findingX, findingY int, fromSide, toSide Side) *Tile {
	toMatch := world[Coordinate{baseX, baseY}].GetSide(fromSide)
	for _, tile := range tiles {
		if tile.Finished {
			continue
		}

		for f := 0; f < 2; f++ {
			for i := 0; i < 4; i++ {
				if tile.GetSide(toSide) == toMatch {
					return tile
				}
				tile.Rotate()
			}
			tile.Flip()

		}
	}

	return nil
}

func parseInput(input string) (tiles []*Tile, err error) {
	tileInputs := strings.Split(input, "\n\n")

	for _, tileInput := range tileInputs {
		tileLines := strings.Split(tileInput, "\n")
		if len(tileLines) < 2 {
			continue
		}
		tile := &Tile{}
		matches := tileRegex.FindStringSubmatch(tileLines[0])
		if matches == nil {
			return nil, fmt.Errorf("Invalid Tile ID line '%s'", tileLines[0])
		}
		if tile.ID, err = strconv.Atoi(matches[1]); err != nil {
			return nil, err
		}

		for i := 1; i < len(tileLines); i++ {
			if tileLines[i] == "" {
				continue
			}
			pixelLine := make([]byte, 0)
			for c := 0; c < len(tileLines[i]); c++ {
				pixelLine = append(pixelLine, tileLines[i][c])
			}
			tile.pixels = append(tile.pixels, pixelLine)
		}

		tiles = append(tiles, tile)
	}

	return
}
