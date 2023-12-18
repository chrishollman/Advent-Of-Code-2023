package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/chrishollman/Advent-Of-Code-2023/utils"
)

//go:embed input.txt
var Input string

func PartOnePreOptimisation(input string) int {
	lines := utils.EmbedStringToStringSlice(input)

	var total = 0

	for _, line := range lines {
		split := strings.Split(line, "|")
		winningNumbersString := strings.Split(strings.TrimSpace(split[0][10:]), " ")
		cardNumbersString := strings.Split(strings.TrimSpace(split[1]), " ")

		var winningNumbers []int
		for _, n := range winningNumbersString {
			if num, err := strconv.Atoi(n); err == nil {
				winningNumbers = append(winningNumbers, num)
			}
		}

		cardValue := 0
		for _, n := range cardNumbersString {
			num, _ := strconv.Atoi(n)
			if slices.Contains(winningNumbers, num) {
				if cardValue == 0 {
					cardValue = 1
				} else {
					cardValue *= 2
				}
			}
		}

		total += cardValue
	}

	return total
}

func PartOne(input string) int {
	lines := utils.EmbedStringToStringSlice(input)
	total := 0

	for _, line := range lines {
		split := strings.Split(line, "|")
		winningNumbersString := strings.Fields(split[0])
		cardNumbersString := strings.Fields(split[1])

		winningNumbers := make(map[int]struct{})
		for _, n := range winningNumbersString {
			if num, err := strconv.Atoi(n); err == nil {
				winningNumbers[num] = struct{}{}
			}
		}

		cardValue := 1
		for _, n := range cardNumbersString {
			num, _ := strconv.Atoi(n)
			if _, exists := winningNumbers[num]; exists {
				cardValue *= 2
			}
		}

		if cardValue > 1 { // Only add if at least one number was found
			total += cardValue
		}
	}

	return total
}

func PartTwo(input string) int {
	lines := utils.EmbedStringToStringSlice(input)

	multipliers := map[string]int{}

	for i := 0; i < len(lines); i++ {
		multipliers[lines[i][:8]] = 1
	}

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		split := strings.Split(line, "|")
		winningNumbersString := strings.Split(strings.TrimSpace(split[0][10:]), " ")
		cardNumbersString := strings.Split(strings.TrimSpace(split[1]), " ")

		multiplier := multipliers[line[:8]]

		var winningNumbers []int
		for _, n := range winningNumbersString {
			if num, err := strconv.Atoi(n); err == nil {
				winningNumbers = append(winningNumbers, num)
			}
		}

		matchingNumbers := 0
		for _, n := range cardNumbersString {
			num, _ := strconv.Atoi(n)
			if slices.Contains(winningNumbers, num) {
				matchingNumbers++
			}
		}

		for j := 1; j <= matchingNumbers; j++ {
			multipliers[lines[i+j][:8]] = multipliers[lines[i+j][:8]] + multiplier
		}
	}

	sum := 0
	for _, v := range multipliers {
		sum += v
	}

	return sum
}

func main() {
	fmt.Println("Day Four - Scratchcards")
	fmt.Println("#######################")

	timer := time.Now()
	fmt.Printf("Part 1 Answer - %v\n", PartOne(Input))
	fmt.Printf("Part 1 Time   - %v\n", time.Since(timer))

	timer = time.Now()
	fmt.Printf("Part 2 Answer - %v\n", PartTwo(Input))
	fmt.Printf("Part 2 Time   - %v\n", time.Since(timer))
}
