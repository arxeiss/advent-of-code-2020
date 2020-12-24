package day24_test

import (
	"testing"

	"github.com/arxeiss/advent-of-code-2020/day24"
)

var testInput = `sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew`

func TestPart1(t *testing.T) {
	part1, part2, err := day24.Solve(testInput)
	t.Log("part1=", part1, "part2=", part2, "error:", err)
	if err != nil {
		t.Error(err)
	}
	if part1 != 10 {
		t.Errorf("Expected part1 to be 10, got %d", part1)
	}
	if part2 != 2208 {
		t.Errorf("Expected part2 to be 2208, got %d", part2)
	}
}
