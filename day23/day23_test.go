package day23_test

import (
	"os"
	"testing"

	"github.com/arxeiss/advent-of-code-2020/day23"
)

var testInput = `389125467`

func TestPart1Rounds10(t *testing.T) {
	result, err := day23.Part1(testInput, 10)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 92658374 {
		t.Errorf("Expected result to be 92658374, got %d", result)
	}
}

func TestPart1Rounds100(t *testing.T) {
	result, err := day23.Part1(testInput, 100)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 67384529 {
		t.Errorf("Expected result to be 67384529, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	// Skip in CI
	if os.Getenv("GITHUB_SHA") != "" {
		t.Skip()
	}
	result, err := day23.Part2(testInput)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 149245887792 {
		t.Errorf("Expected result to be 149245887792, got %d", result)
	}
}
