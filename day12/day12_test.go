package day12_test

import (
	"testing"

	"github.com/arxeiss/advent-of-code-2020/day12"
)

var testInput = `F10
N3
F7
R90
F11`

func TestPart1(t *testing.T) {
	result, err := day12.Part1(testInput)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 25 {
		t.Errorf("Expected result to be 25, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result, err := day12.Part2(testInput)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 286 {
		t.Errorf("Expected result to be 286, got %d", result)
	}
}
