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

const Gear = '*'

//go:embed input.txt
var Input string

var Numbers = []uint8{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
var Symbols = []uint8{'#', '*', '$', '-', '%', '+', '@', '&', '=', '/'}

func GetKey(row, col int) string {
	return strconv.Itoa(row) + "," + strconv.Itoa(col)
}

func GetAdjacentPoints(row, col, h, w int) []string {
	var results []string

	possibles := [][]int{
		{row - 1, col - 1}, {row - 1, col + 0}, {row - 1, col + 1},
		{row - 0, col - 1} /*{row, col}*/, {row - 0, col + 1},
		{row + 1, col - 1}, {row + 1, col + 0}, {row + 1, col + 1},
	}

	for _, value := range possibles {
		newRow, newCol := value[0], value[1]
		if newRow < 0 || newRow > h || newCol < 0 || newRow > w {
			continue
		}
		results = append(results, GetKey(newRow, newCol))
	}

	return results
}

func IsNumber(in uint8) bool {
	return slices.Contains(Numbers, in)
}

func IsSymbol(in uint8) bool {
	return slices.Contains(Symbols, in)
}

func PartOnePreOptimisation(input string) int {
	lines := utils.EmbedStringToStringSlice(input)

	h, w := len(lines), len(lines[0])

	total := 0

	var symbolLocations []string
	for row := 0; row < h; row++ {
		for col := 0; col < w; col++ {
			if char := lines[row][col]; IsSymbol(char) {
				symbolLocations = append(symbolLocations, GetKey(row, col))
			}
		}
	}

	var (
		nextToSymbolCache = make(map[string]bool, len(symbolLocations))
		shouldBeIncluded  bool
		strNumBuffer      strings.Builder
	)

	for row := 0; row < h; row++ {
		for col := 0; col < w; col++ {
			char := lines[row][col]

			if IsNumber(char) {
				strNumBuffer.WriteByte(char)
				adjacentPoints := GetAdjacentPoints(row, col, h, w)
				for _, point := range adjacentPoints {
					if val, ok := nextToSymbolCache[point]; ok {
						if val {
							shouldBeIncluded = true
							break
						}
					}
					val := slices.Contains(symbolLocations, point)
					nextToSymbolCache[point] = val
					if val {
						shouldBeIncluded = true
						break
					}
				}
			}

			if !IsNumber(char) || col == w {
				// We were in a number
				if strNumBuffer.Len() > 0 {
					if shouldBeIncluded {
						num, _ := strconv.Atoi(strNumBuffer.String())
						total += num
					}
					strNumBuffer.Reset()
					shouldBeIncluded = false
				}
			}
		}
	}

	return total
}

func PartOne(input string) int {
	lines := utils.EmbedStringToStringSlice(input)

	h, w := len(lines), len(lines[0])

	total := 0

	symbolLocations := make(map[string]struct{})
	for row := 0; row < h; row++ {
		for col := 0; col < w; col++ {
			if IsSymbol(lines[row][col]) {
				symbolLocations[GetKey(row, col)] = struct{}{}
			}
		}
	}

	for row := 0; row < h; row++ {
		currentNumber := 0
		isAdjacentToSymbol := false

		for col := 0; col < w; col++ {
			char := lines[row][col]

			if IsNumber(char) {
				currentNumber = currentNumber*10 + int(char-'0')
				if !isAdjacentToSymbol {
					adjacentPoints := GetAdjacentPoints(row, col, h, w)
					for _, point := range adjacentPoints {
						if _, found := symbolLocations[point]; found {
							isAdjacentToSymbol = true
							break
						}
					}
				}
			} else {
				if isAdjacentToSymbol {
					total += currentNumber
				}
				currentNumber = 0
				isAdjacentToSymbol = false
			}
		}

		// Check if the last number in the row was adjacent to a symbol
		if isAdjacentToSymbol {
			total += currentNumber
		}
	}

	return total
}

func PartTwoPreOptimisation(input string) int {
	lines := utils.EmbedStringToStringSlice(input)

	h, w := len(lines), len(lines[0])

	total := 0

	var gearLocations []string
	for row := 0; row < h; row++ {
		for col := 0; col < w; col++ {
			if char := lines[row][col]; char == Gear {
				gearLocations = append(gearLocations, GetKey(row, col))
			}
		}
	}

	var (
		gearAdjacent         = make(map[string][]int, len(gearLocations))
		pointsToAddNumbersTo []string
		shouldBeIncluded     bool
		strNumBuffer         strings.Builder
	)

	for row := 0; row < h; row++ {
		for col := 0; col < w; col++ {
			char := lines[row][col]

			if IsNumber(char) {
				strNumBuffer.WriteByte(char)
				adjacentPoints := GetAdjacentPoints(row, col, h, w)
				for _, point := range adjacentPoints {
					if slices.Contains(gearLocations, point) {
						shouldBeIncluded = true
						if !slices.Contains(pointsToAddNumbersTo, point) {
							pointsToAddNumbersTo = append(pointsToAddNumbersTo, point)
						}
					}
				}
			}

			if !IsNumber(char) || col == w {
				if strNumBuffer.Len() > 0 {
					if shouldBeIncluded {
						num, _ := strconv.Atoi(strNumBuffer.String())
						for _, point := range pointsToAddNumbersTo {
							if val, ok := gearAdjacent[point]; ok {
								gearAdjacent[point] = append(val, num)
							} else {
								gearAdjacent[point] = []int{num}
							}
						}
					}
					pointsToAddNumbersTo = []string{}
					strNumBuffer.Reset()
					shouldBeIncluded = false
				}
			}
		}
	}

	for _, v := range gearAdjacent {
		if len(v) == 2 {
			total += v[0] * v[1]
		}
	}

	return total
}

func PartTwo(input string) int {
	lines := utils.EmbedStringToStringSlice(input)

	h, w := len(lines), len(lines[0])

	total := 0

	gearLocations := make(map[string]struct{})
	for row := 0; row < h; row++ {
		for col := 0; col < w; col++ {
			if lines[row][col] == Gear {
				gearLocations[GetKey(row, col)] = struct{}{}
			}
		}
	}

	gearAdjacent := make(map[string][]int)

	for row := 0; row < h; row++ {
		currentNumber := 0
		pointsToAddNumbersTo := make(map[string]struct{})

		for col := 0; col < w; col++ {
			char := lines[row][col]

			if IsNumber(char) {
				currentNumber = currentNumber*10 + int(char-'0')
				adjacentPoints := GetAdjacentPoints(row, col, h, w)
				for _, point := range adjacentPoints {
					if _, found := gearLocations[point]; found {
						pointsToAddNumbersTo[point] = struct{}{}
					}
				}
			} else {
				if len(pointsToAddNumbersTo) > 0 {
					for point := range pointsToAddNumbersTo {
						gearAdjacent[point] = append(gearAdjacent[point], currentNumber)
					}
				}
				currentNumber = 0
				pointsToAddNumbersTo = make(map[string]struct{})
			}
		}

		// Check if the last number in the row should be added
		if len(pointsToAddNumbersTo) > 0 {
			for point := range pointsToAddNumbersTo {
				gearAdjacent[point] = append(gearAdjacent[point], currentNumber)
			}
		}
	}

	for _, nums := range gearAdjacent {
		if len(nums) == 2 {
			total += nums[0] * nums[1]
		}
	}

	return total
}

func main() {
	fmt.Println("Day Three - Gear Ratios")
	fmt.Println("#######################")

	timer := time.Now()
	fmt.Printf("Part 1 Answer - %v\n", PartOne(Input))
	fmt.Printf("Part 1 Time   - %v\n", time.Since(timer))

	timer = time.Now()
	fmt.Printf("Part 2 Answer - %v\n", PartTwo(Input))
	fmt.Printf("Part 2 Time   - %v\n", time.Since(timer))
}
