package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/chrishollman/Advent-Of-Code-2023/utils"
)

//go:embed input.txt
var Input string

var strToIntMapper = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func PartOne(input string) int {
	lines := utils.EmbedStringToStringSlice(input)

	total := 0
	for _, line := range lines {
		var ans int

		// Front
		for i := 0; i < len(line); i++ {
			r := line[i]
			if r > 48 && r < 58 {
				ans += int(r-48) * 10
				break
			}
		}

		// Back
		for i := len(line) - 1; i >= 0; i-- {
			r := line[i]
			if r > 48 && r < 58 {
				ans += int(r - 48)
				break
			}
		}

		total += ans
	}
	return total
}

func PartTwo(input string) int {
	lines := utils.EmbedStringToStringSlice(input)

	total := 0
	for _, line := range lines {
		idxFront, idxBack := math.MaxInt, math.MinInt
		front, back := 0, 0

		// Front
		for i := 0; i < len(line); i++ {
			r := line[i]
			if r > 48 && r < 58 {
				idxFront = i
				front = int(r-48) * 10
				break
			}
		}

		// Back
		for i := len(line) - 1; i >= 0; i-- {
			r := line[i]
			if r > 48 && r < 58 {
				idxBack = i
				back = int(r - 48)
				break
			}
		}

		// Words
		for k, v := range strToIntMapper {
			// Front
			idx := strings.Index(line, k)
			if idx != -1 && idx < idxFront {
				idxFront = idx
				front = v * 10
			}
			// Back
			idx = strings.LastIndex(line, k)
			if idx != -1 && idx > idxBack {
				idxBack = idx
				back = v
			}
		}

		total += front + back
	}
	return total
}

func main() {
	fmt.Println("Day One - Trebuchet")
	fmt.Println("###################")

	timer := time.Now()
	fmt.Printf("Part 1 Answer - %v\n", PartOne(Input))
	fmt.Printf("Part 1 Time   - %v\n", time.Since(timer))

	timer = time.Now()
	fmt.Printf("Part 2 Answer - %v\n", PartTwo(Input))
	fmt.Printf("Part 2 Time   - %v\n", time.Since(timer))
}
