# Advent Of Code 2020 in Go

New project, new language, new Advent of Code in Go! To see the original tasks for all days and previous years, visit https://adventofcode.com/. All credits to them for this nice work!

- Year 2019 in Elixir: https://github.com/arxeiss/advent-of-code-2019
- Year 2021 in PHP: https://github.com/arxeiss/advent-of-code-2021
- Year 2022 in PHP: https://github.com/arxeiss/advent-of-code-2022

See each day for more information. I copied the instructions there as well.

## How to run

1. Install Go: https://golang.org/doc/install
1. Clone this repo
1. Run:
    - With makefile: `make run` and program will ask you which puzzle to run
    - Run all tests with `make tests`
    - Run program with specified day and puzzle part: `go run . [day] [part]`

## It is completed! But with some help

I was able to manage all day's puzzles to get all **`50 stars`**. But it wouldn't be possible without some help from friends or the community. Some days I had trouble understanding the task, to understand what I should do. Some days I managed to do the test input, but it was too slow for real input. Or someday I just had no idea how to do it.

### Big thanks to:
 - YouTube channel of [TurkeyDev](https://www.youtube.com/watch?v=9yHIE6nP50c&list=PLL7ab2XZ_GU3DuxWKiy-Rcrv0ACPFNx96), was a big help when I got really stuck
 - My friend [Kobzol](https://github.com/Kobzol), who helped me with performance during the days, when my code was slow.
- The people on [AoC Reddit](https://www.reddit.com/r/adventofcode/) for getting some hints.

### Days I got stuck and searched for the help

- **Day 13** was the first struggling day. I had no idea about the Chinese Remainder Theorem. Thanks TurkeyDev for the hint and see [Day 13](/day13) for more info and study material.
- On **Day 16**, I was implementing part 2 completely wrong, I had to use code from [CodingNagger](https://github.com/CodingNagger/advent-of-code-2020) as a debugger to find, where my mistake is. And also the inspiration.
- The task for **Day 17** wasn't clear for me. I used Reddit hints to understand what I should do.
- **Day 18** was really time-consuming. I wanted to do it the right way, not just the working way for AoC. Should be done properly with the [Shunting-Yard algorithm](https://en.wikipedia.org/wiki/Shunting-yard_algorithm).
- **Day 19** works just accidentally, see [Reddit Post](https://www.reddit.com/r/adventofcode/comments/kg9hal/2020_day_19_part_2_my_code_works_accidentally_i/)
- I failed totally with **Day 21**. My code was totally wrong and I had to get inspiration on how to start the elimination and then start from scratch.
- **Day 23** I used the help from **Kobzol**, as he suggested me to use the combination of `LinkedList` and `HashMap`. Otherwise, it would take many hours to solve.
- And maybe some more, I just don't remember now.

## Days

- [Day 1: Report Repair](/day1)
- [Day 2: Password Philosophy](/day2)
- [Day 3: Toboggan Trajectory](/day3)
- [Day 4: Passport Processing](/day4)
- [Day 5: Binary Boarding](/day5)
- [Day 6: Custom Customs](/day6)
- [Day 7: Handy Haversacks](/day7)
- [Day 8: Handheld Halting](/day8)
- [Day 9: Encoding Error](/day9)
- [Day 10: Adapter Array](/day10)
- [Day 11: Seating System](/day11)
- [Day 12: Rain Risk](/day12)
- [Day 13: Shuttle Search](/day13)
- [Day 14: Docking Data](/day14)
- [Day 15: Rambunctious Recitation](/day15)
- [Day 16: Ticket Translation](/day16)
- [Day 17: Conway Cubes](/day17)
- [Day 18: Operation Order](/day18)
- [Day 19: Monster Messages](/day19)
- [Day 20: Jurassic Jigsaw](/day20)
- [Day 21: Allergen Assessment](/day21)
- [Day 22: Crab Combat](/day22)
- [Day 23: Crab Cups](/day23)
- [Day 24: Lobby Layout](/day24)
- [Day 25: Combo Breaker](/day25)
