package main

import "fmt"
import "os"
import "strings"
import "strconv"
import "sort"
import "math"

func atoi(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return x
}

func main() {
	filename := "./input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	sum := part1(lines)
	fmt.Printf("sum: %v\n", sum)
	similarity := part2(lines)
	fmt.Printf("similarity: %v\n", similarity)
	if sum != 2580760 {
		panic("incorrect part1")
	}
	if similarity != 25358365 {
		panic("incorrect part2")
	}
}

func lists(lines []string) ([]int, []int) {
	as := []int{}
	bs := []int{}
	for i := range lines {
		vals := strings.Split(lines[i], "   ")
		if len(vals) != 2 {
			continue
		}
		a, b := atoi(vals[0]), atoi(vals[1])
		as = append(as, a)
		bs = append(bs, b)
	}
	if len(as) != len(bs) {
		panic("mismatched lists")
	}
	return as, bs
}

func part1(lines []string) int {
	as, bs := lists(lines)
	sort.Ints(as)
	sort.Ints(bs)
	sum := 0
	for i := 0; i < len(as); i++ {
		sum += int(math.Abs(float64(as[i] - bs[i])))
	}
	return sum
}

func part2(lines []string) int {
	as, bs := lists(lines)
	counts := map[int]int{}
	for _, b := range bs {
		val, exists := counts[b]
		if exists {
			counts[b] = val + 1
		} else {
			counts[b] = 1
		}
	}
	similarity := 0
	for _, a := range as {
		multiplier, exists := counts[a]
		if !exists {
			multiplier = 0
		}
		similarity += a * multiplier
	}
	return similarity
}
