package day11_test

import (
	"testing"

	"github.com/arxeiss/advent-of-code-2020/day11"
)

var testInput = `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`

func TestPart1(t *testing.T) {
	result, err := day11.Part1(testInput)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 37 {
		t.Errorf("Expected result to be 37, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result, err := day11.Part2(testInput)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 26 {
		t.Errorf("Expected result to be 26, got %d", result)
	}
}
