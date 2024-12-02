package main

import (
    "strings"
    "strconv"
    "fmt"
    "os"
    "bufio"
    "math"
)

type Report struct {
    Levels []int
}

func NewReport(line string) Report {
    parts := strings.Split(line, " ")
    levels := make([]int, len(parts))
    for i, part := range parts {
        levels[i] = atoi(part)
    }
    return Report{levels}
}

func (r Report) WithoutOneLevel() []Report {
    reports := make([]Report, len(r.Levels) - 1)
    fmt.Printf("starting with %v\n", r.Levels)
    for i := 0; i < len(r.Levels) - 1; i++ {
        newLevels := append(r.Levels[:i], r.Levels[i+1:]...)
        fmt.Printf("i: %v, new levels: %v\n", i, newLevels)
        reports[i] = Report{newLevels}
    }
    return reports
}

func (r Report) IsSafe() (bool, int) {
    for i := 1; i < len(r.Levels); i++ {
        diff := abs(r.Levels[i-1] - r.Levels[i])
        if diff < 1 || diff > 3 {
            return false, i
        }
    }
    return r.IsMonotonic()
}

func (r Report) IsDecreasing() (bool, int) {
    if len(r.Levels) < 2 {
        return true, -1
    }
    
    for i := 1; i < len(r.Levels); i++ {
        if r.Levels[i] > r.Levels[i-1] {
            return false, i
        }
    }

    return true, -1
 
}

func (r Report) IsIncreasing() (bool, int) {
    if len(r.Levels) < 2 {
        return true, -1
    }
    
    for i := 1; i < len(r.Levels); i++ {
        if r.Levels[i] < r.Levels[i-1] {
            return false, i
        }
    }

    return true, -1
 
}


func (r Report) IsMonotonic() (bool, int) {
    if len(r.Levels) < 2 {
        return true, -1
    }

    isDecreasing, x := r.IsDecreasing()
    if isDecreasing {
        return true, -1
    }
    isIncreasing, _ := r.IsIncreasing()
    if isIncreasing {
        return true, -1
    }
    return false, x
}

func (r *Report) removeAt(i int) {
    r.Levels = append(r.Levels[:i], r.Levels[i+1:]...)
}


func atoi(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return x
}

func abs(x int) int {
    return int(math.Abs(float64(x)))
}

func main() {
	filename := "./input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
    defer file.Close()
    //part1(file)
    file.Seek(0,0)
    part2(file)
}

func part1(file *os.File) {
    scanner := bufio.NewScanner(file) 
    count := 0
    for scanner.Scan() {
        report := NewReport(scanner.Text())
        if isSafe, _ := report.IsSafe(); isSafe {
            fmt.Println(report)
            count++
        }
    }
    fmt.Printf("safe reports: %v\n", count)
    
}

func part2(file *os.File) {
    scanner := bufio.NewScanner(file) 
    count := 0
    for scanner.Scan() {
        report := NewReport(scanner.Text())
        isSafe, i := report.IsSafe()
        if isSafe {
            count++
            continue
        }
        if !isSafe && i >= 0 {
            lessOne := report.WithoutOneLevel()
            fmt.Println(lessOne)
        }
    }
    fmt.Printf("safe reports: %v\n", count)
}
