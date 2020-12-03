package day3_test

import (
	"testing"

	"github.com/arxeiss/advent-of-code-2020/day3"
)

var testInput = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

func TestPart1(t *testing.T) {
	result, err := day3.Part1(testInput)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 7 {
		t.Errorf("Expected result to be 7, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result, err := day3.Part2(testInput)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 336 {
		t.Errorf("Expected result 336, got %d", result)
	}
}
