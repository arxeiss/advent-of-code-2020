package day25_test

import (
	"testing"

	"github.com/arxeiss/advent-of-code-2020/day25"
)

var testInput = `5764801
17807724`

func TestPart1(t *testing.T) {
	result, err := day25.Part1(testInput)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 14897079 {
		t.Errorf("Expected result to be 14897079, got %d", result)
	}
}
