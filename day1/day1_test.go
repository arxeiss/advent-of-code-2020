package day1_test

import (
	"testing"

	"github.com/arxeiss/advent-of-code-2020/day1"
)

var testInput = `1721
979
366
299
675
1456`

func TestPart1(t *testing.T) {
	a, b, result, err := day1.Part1(testInput)
	t.Log("a=", a, "b=", b, "result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 514579 {
		t.Errorf("Expected result 514579, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	a, b, c, result, err := day1.Part2(testInput)
	t.Log("a=", a, "b=", b, "c=", c, "result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 241861950 {
		t.Errorf("Expected result 241861950, got %d", result)
	}
}
