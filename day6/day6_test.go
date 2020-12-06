package day6_test

import (
	"testing"

	"github.com/arxeiss/advent-of-code-2020/day6"
)

var testInput = `abc

a
b
c

ab
ac

a
a
a
a

b`

func TestPart1(t *testing.T) {
	result, err := day6.Part1(testInput)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 11 {
		t.Errorf("Expected result to be 11, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result, err := day6.Part2(testInput)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 6 {
		t.Errorf("Expected result to be 6, got %d", result)
	}
}
