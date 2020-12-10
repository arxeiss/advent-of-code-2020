package day9_test

import (
	"testing"

	"github.com/arxeiss/advent-of-code-2020/day9"
)

var testInput = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`

func TestPart1(t *testing.T) {
	result, err := day9.Part1(testInput, 5)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 127 {
		t.Errorf("Expected result to be 127, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result, err := day9.Part2(testInput, 5)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 62 {
		t.Errorf("Expected result to be 62, got %d", result)
	}
}
