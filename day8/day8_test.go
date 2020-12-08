package day8_test

import (
	"testing"

	"github.com/arxeiss/advent-of-code-2020/day8"
)

var testInput = `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`

func TestPart1(t *testing.T) {
	result, err := day8.Part1(testInput)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 5 {
		t.Errorf("Expected result to be 5, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result, err := day8.Part2(testInput)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 8 {
		t.Errorf("Expected result to be 8, got %d", result)
	}
}
