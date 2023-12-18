package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/chrishollman/Advent-Of-Code-2023/utils"
)

//go:embed input.txt
var Input string

func PartOne(input string) int {
	lines := utils.EmbedStringToStringSlice(input)

	timesStr := strings.Fields(lines[0][5:])
	distancesStr := strings.Fields(lines[1][9:])

	times := make([]int, len(timesStr))
	for i, v := range timesStr {
		num, _ := strconv.Atoi(v)
		times[i] = num
	}

	distances := make([]int, len(distancesStr))
	for i, v := range distancesStr {
		num, _ := strconv.Atoi(v)
		distances[i] = num
	}

	raceTotals := 1
	for race := 0; race < len(times); race++ {

		variations := 0
		for attempt := 0; attempt < times[race]; attempt++ {
			result := attempt * (times[race] - attempt)
			if result > distances[race] {
				variations++
			}
		}

		raceTotals *= variations
	}

	return raceTotals
}

func PartTwo(input string) int {
	lines := utils.EmbedStringToStringSlice(input)

	timeStr := strings.Replace(lines[0][5:], " ", "", -1)
	distanceStr := strings.Replace(lines[1][9:], " ", "", -1)

	raceTime, _ := strconv.Atoi(timeStr)
	raceDistance, _ := strconv.Atoi(distanceStr)

	variations := 0
	for attempt := 0; attempt < raceTime; attempt++ {
		result := attempt * (raceTime - attempt)
		if result > raceDistance {
			variations++
		}
		if variations > 0 && result < raceDistance {
			break
		}
	}

	return variations
}

func main() {
	fmt.Println("Day Six - Wait for it")
	fmt.Println("#####################")

	timer := time.Now()
	fmt.Printf("Part 1 Answer - %v\n", PartOne(Input))
	fmt.Printf("Part 1 Time   - %v\n", time.Since(timer))

	timer = time.Now()
	fmt.Printf("Part 2 Answer - %v\n", PartTwo(Input))
	fmt.Printf("Part 2 Time   - %v\n", time.Since(timer))
}
