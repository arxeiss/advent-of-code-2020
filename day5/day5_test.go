package day5_test

import (
	"testing"

	"github.com/arxeiss/advent-of-code-2020/day5"
)

var testInput = `BFFFBBFRRR
FFFBBBFRRR
BBFFBBFRLL`

// ID 616 - 621, 620 missing
var test2Input = `BFFBBFBLLL
BFFBBFBLLR
BFFBBFBLRL
BFFBBFBLRR
BFFBBFBRLR
`

func TestPart1(t *testing.T) {
	result, err := day5.Part1(testInput)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 820 {
		t.Errorf("Expected result to be 820, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result, err := day5.Part2(test2Input)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 620 {
		t.Errorf("Expected result to be 620, got %d", result)
	}
}
