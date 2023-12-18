package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var Input string

type Interval struct {
	Min      int
	Max      int
	Modifier int
}

func (i *Interval) Contains(val int) bool {
	if val >= i.Min && val < i.Max {
		return true
	}
	return false
}

func transform(val int, intervals []*Interval) int {
	for _, interval := range intervals {
		if interval.Contains(val) {
			return val + interval.Modifier
		}
	}
	return val
}

func undoTransform(val int, intervals []*Interval) int {
	for _, interval := range intervals {
		target := val - interval.Modifier
		if interval.Contains(target) {
			return target
		}
	}

	return val
}

func generateSeedIntervals(input []string) []*Interval {
	intervals := make([]*Interval, len(input)/2)
	for i := 0; i < len(input); i += 2 {
		minimum, _ := strconv.Atoi(input[i])
		howMany, _ := strconv.Atoi(input[i+1])
		maximum := minimum + howMany
		intervals[i/2] = &Interval{Min: minimum, Max: maximum}
	}

	return intervals
}

func generateTransformIntervals(input string) []*Interval {
	numbers := strings.Fields(strings.Split(input, ":")[1])
	out := make([]*Interval, len(numbers)/3)

	for i := 0; i < len(numbers); i += 3 {
		minimum, _ := strconv.Atoi(numbers[i+1])
		maximum, _ := strconv.Atoi(numbers[i+2])
		modifier, _ := strconv.Atoi(numbers[i])
		out[i/3] = &Interval{Min: minimum, Max: minimum + maximum, Modifier: modifier - minimum}
	}

	return out
}

func PartOne(input string) int {
	splits := strings.Split(input, "\n\n")
	seedsStr := strings.Fields(splits[0][6:])

	seeds := make([]int, len(seedsStr))
	for i, v := range seedsStr {
		num, _ := strconv.Atoi(v)
		seeds[i] = num
	}

	a := generateTransformIntervals(splits[1])
	b := generateTransformIntervals(splits[2])
	c := generateTransformIntervals(splits[3])
	d := generateTransformIntervals(splits[4])
	e := generateTransformIntervals(splits[5])
	f := generateTransformIntervals(splits[6])
	g := generateTransformIntervals(splits[7])

	minimum := math.MaxInt
	for _, seed := range seeds {
		current := seed
		current = transform(current, a)
		current = transform(current, b)
		current = transform(current, c)
		current = transform(current, d)
		current = transform(current, e)
		current = transform(current, f)
		current = transform(current, g)

		if current < minimum {
			minimum = current
		}
	}

	return minimum
}

func PartTwo(input string) int {
	splits := strings.Split(input, "\n\n")
	seedsStr := strings.Fields(splits[0][6:])

	seeds := generateSeedIntervals(seedsStr)

	a := generateTransformIntervals(splits[1])
	b := generateTransformIntervals(splits[2])
	c := generateTransformIntervals(splits[3])
	d := generateTransformIntervals(splits[4])
	e := generateTransformIntervals(splits[5])
	f := generateTransformIntervals(splits[6])
	g := generateTransformIntervals(splits[7])

	i := 0
	for {
		current := i
		current = undoTransform(current, g)
		current = undoTransform(current, f)
		current = undoTransform(current, e)
		current = undoTransform(current, d)
		current = undoTransform(current, c)
		current = undoTransform(current, b)
		current = undoTransform(current, a)

		for _, interval := range seeds {
			if interval.Contains(current) {
				return i
			}
		}

		i++
	}
}

func main() {
	fmt.Println("Day Five - Fertilizer")
	fmt.Println("#####################")

	timer := time.Now()
	fmt.Printf("Part 1 Answer - %v\n", PartOne(Input))
	fmt.Printf("Part 1 Time   - %v\n", time.Since(timer))

	timer = time.Now()
	fmt.Printf("Part 2 Answer - %v\n", PartTwo(Input))
	fmt.Printf("Part 2 Time   - %v\n", time.Since(timer))
}
