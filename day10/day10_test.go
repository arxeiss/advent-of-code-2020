package day10_test

import (
	"testing"

	"github.com/arxeiss/advent-of-code-2020/day10"
)

var testInput = `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`

func TestPart1(t *testing.T) {
	result, err := day10.Part1(testInput)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 220 {
		t.Errorf("Expected result to be 220, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result, err := day10.Part2(testInput)
	t.Log("result=", result, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if result != 19208 {
		t.Errorf("Expected result to be 19208, got %d", result)
	}
}
