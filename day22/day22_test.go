package day22_test

import (
	"testing"

	"github.com/arxeiss/advent-of-code-2020/day22"
)

var testInput = `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`

func TestPart1(t *testing.T) {
	result, err := day22.Part1(testInput)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 306 {
		t.Errorf("Expected result to be 306, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result, err := day22.Part2(testInput)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 291 {
		t.Errorf("Expected result to be 291, got %d", result)
	}
}
