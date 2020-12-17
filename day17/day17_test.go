package day17_test

import (
	"testing"

	"github.com/arxeiss/advent-of-code-2020/day17"
)

var testInput = `.#.
..#
###`

func TestPart1(t *testing.T) {
	result, err := day17.Part1(testInput)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 112 {
		t.Errorf("Expected result to be 112, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result, err := day17.Part2(testInput)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 848 {
		t.Errorf("Expected result to be 848, got %d", result)
	}
}
