package day14_test

import (
	"testing"

	"github.com/arxeiss/advent-of-code-2020/day14"
)

var testInput = `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`

func TestPart1(t *testing.T) {
	result, err := day14.Part1(testInput)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 165 {
		t.Errorf("Expected result to be 165, got %d", result)
	}
}

var test2Input = `mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`

func TestPart2(t *testing.T) {
	result, err := day14.Part2(test2Input)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 208 {
		t.Errorf("Expected result to be 208, got %d", result)
	}
}
