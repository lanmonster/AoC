#! /bin/sh

NEW_DAY_NUM=$(find . -type d -name "day*" | wc -l | tr -d ' ' | awk '{print $1+1}')

mkdir "day$NEW_DAY_NUM"

curl "https://adventofcode.com/$(date +%Y)/day/$NEW_DAY_NUM/input" --cookie "session=$AOC_SESSION" -o "day$NEW_DAY_NUM/input.txt"
touch "day$NEW_DAY_NUM/test.txt" "day$NEW_DAY_NUM/test2.txt"

cat <<EOF >"day$NEW_DAY_NUM/day$NEW_DAY_NUM.jakt"
import relative parent::utils as utils

comptime part1_expected_value() => 0u64
comptime part2_expected_value() => 0u64

fn part1(lines: [String]) throws -> u64 {
    abort()
}
fn part2(lines: [String]) throws -> u64 {
    abort()
}
EOF

sed -i'.jakt' "s|^import utils.*|import day$NEW_DAY_NUM::day$NEW_DAY_NUM as d$NEW_DAY_NUM\n&|" aoc.jakt
sed -i'.jakt' "s|\( *\)else => {|\1\"$NEW_DAY_NUM\" => (d$NEW_DAY_NUM::part1_expected_value(), d$NEW_DAY_NUM::part1(lines), d$NEW_DAY_NUM::part2_expected_value(), d$NEW_DAY_NUM::part2(lines))\n&|" aoc.jakt
rm aoc.jakt.jakt