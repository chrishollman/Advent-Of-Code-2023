package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/chrishollman/Advent-Of-Code-2023/utils"
)

const (
	limitBlue  = 14
	limitGreen = 13
	limitRed   = 12
)

//go:embed input.txt
var Input string

func PartOne(input string) int {
	lines := utils.EmbedStringToStringSlice(input)

	total := 0
	for gameNumber, rawRound := range lines {
		validRound := true

		outcomes := strings.Split(rawRound, " ")
		for i := 2; i < len(outcomes); i += 2 {
			num, _ := strconv.Atoi(outcomes[i])
			color := outcomes[i+1]
			switch {
			case strings.HasPrefix(color, "blue"):
				if num > limitBlue {
					validRound = false
				}
			case strings.HasPrefix(color, "green"):
				if num > limitGreen {
					validRound = false
				}
			case strings.HasPrefix(color, "red"):
				if num > limitRed {
					validRound = false
				}
			}

			if !validRound {
				break
			}
		}

		if validRound {
			total += gameNumber + 1
		}
	}

	return total
}

func PartTwo(input string) int {
	lines := utils.EmbedStringToStringSlice(input)

	total := 0
	for _, rawRound := range lines {
		minBlue, minGreen, minRed := 0, 0, 0

		outcomes := strings.Split(rawRound, " ")
		for i := 2; i < len(outcomes); i += 2 {
			num, _ := strconv.Atoi(outcomes[i])
			color := outcomes[i+1]
			switch {
			case strings.HasPrefix(color, "blue"):
				minBlue = max(minBlue, num)
			case strings.HasPrefix(color, "green"):
				minGreen = max(minGreen, num)
			case strings.HasPrefix(color, "red"):
				minRed = max(minRed, num)
			}
		}

		total += minBlue * minGreen * minRed
	}

	return total
}

func main() {
	fmt.Println("Day Two - Cube Conundrum")
	fmt.Println("########################")

	timer := time.Now()
	fmt.Printf("Part 1 Answer - %v\n", PartOne(Input))
	fmt.Printf("Part 1 Time   - %v\n", time.Since(timer))

	timer = time.Now()
	fmt.Printf("Part 2 Answer - %v\n", PartTwo(Input))
	fmt.Printf("Part 2 Time   - %v\n", time.Since(timer))
}
