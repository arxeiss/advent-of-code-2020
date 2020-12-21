# --- Day 21: Allergen Assessment ---

See [How to run](https://github.com/arxeiss/advent-of-code-2020/#how-to-run) chapter to run this puzzle

> :warning: **SPOILER ALERT** :warning: - The code contains solution for the whole task. Try first to solve it **yourself**. :link: https://adventofcode.com/2020/day/21

## Preface

:information_source: NOTE: :information_source: I totally failed with this day. My first code was non-deterministic and was producing random results. Here are the reasons:

1. My algorithm was bad from the beginning. I was searching for a single ingredient which appeared most times at 1 allergen. But this leads to the situation, that there were 2 or more ingredients with the same amount of the same allergen. Normally this program would produce always the same bad result, but:
1. Go `key, val := range map` is not-deterministic either and looping randomly through a map. So this causes that my nested loops break always at a different number.

I watched the video from [TurkeyDev on YouTube](https://www.youtube.com/watch?v=tXwh0y0PyPw) then, which used a very clever algorithm and I inspired myself by it.

## --- Part 1 ---

You reach the train's last stop and the closest you can get to your vacation island without getting wet. There
aren't even any boats here, but nothing can stop you now: you build a raft. You just need a few days' worth of
food for your journey.

You don't speak the local language, so you can't read any ingredients lists. However, sometimes, allergens are
listed in a language you **do** understand. You should be able to use this information to determine which
ingredient contains which allergen and work out which foods are safe to
take with you on your trip.

You start by compiling a list of foods (your puzzle input), one food per line. Each line includes that food's
**ingredients list** followed by some or all of the allergens the food contains.

Each allergen is found in exactly one ingredient. Each ingredient contains zero or one allergen. **Allergens
aren't always marked**; when they're listed (as in `(contains nuts, shellfish)` after an
ingredients list), the ingredient that contains each listed allergen will be **somewhere in the corresponding
ingredients list**. However, even if an allergen isn't listed, the ingredient that contains that allergen
could still be present: maybe they forgot to label it, or maybe it was labeled in a language you don't know.

For example, consider the following list of foods:

```
mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)
```

The first food in the list has four ingredients (written in a language you don't understand):
`mxmxvkd`, `kfcds`, `sqjhc`, and `nhms`. While the food might
contain other allergens, a few allergens the food definitely contains are listed afterward: `dairy`
and `fish`.

The first step is to determine which ingredients **can't possibly** contain any of the allergens in any
food in your list. In the above example, none of the ingredients `kfcds`, `nhms`,
`sbzzf`, or `trh` can contain an allergen. Counting the number of times any of these
ingredients appear in any ingredients list produces **`5`**: they all appear once each except
`sbzzf`, which appears twice.

Determine which ingredients cannot possibly contain any of the allergens in your list. **How many times do
any of those ingredients appear?**

## --- Part 2 ---

Now that you've isolated the inert ingredients, you should have enough information to figure out which
ingredient contains which allergen.

In the above example:

- `mxmxvkd` contains `dairy`.
- `sqjhc` contains `fish`.
- `fvjkl` contains `soy`.

Arrange the ingredients **alphabetically by their allergen** and separate them by commas to produce your
**canonical dangerous ingredient list**. (There should **not be any spaces** in your canonical
dangerous ingredient list.) In the above example, this would be **`mxmxvkd,sqjhc,fvjkl`**.

Time to stock your raft with supplies. **What is your canonical dangerous ingredient list?**
