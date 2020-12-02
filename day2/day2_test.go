package day2_test

import (
	"testing"

	"github.com/arxeiss/advent-of-code-2020/day2"
)

var testInput = `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
1-5 c: eeee`

func TestPart1(t *testing.T) {
	result, err := day2.Part1(testInput)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 2 {
		t.Errorf("Expected result to be 2, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result, err := day2.Part2(testInput)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 1 {
		t.Errorf("Expected result 1, got %d", result)
	}
}
