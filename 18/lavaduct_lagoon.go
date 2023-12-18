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

type Point struct {
	X int
	Y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Shoelace formula => https://w.wiki/8YJc
func shoelaceFormula(p []*Point) int {
	sum := 0
	n := len(p)
	for i := range p {
		j := (i + 1) % n
		sum += (p[i].X * p[j].Y) - (p[j].X * p[i].Y)
	}

	return abs(sum) / 2
}

func PartOne(input string) int {
	lines := utils.EmbedStringToStringSlice(input)

	totalDistance := 0
	history := make([]*Point, len(lines))
	current := &Point{X: 0, Y: 0}
	for idx, line := range lines {
		splitInput := strings.Fields(line)

		direction := splitInput[0]
		distance, _ := strconv.Atoi(splitInput[1])

		switch direction {
		case "U":
			current.Y += distance
		case "D":
			current.Y -= distance
		case "L":
			current.X -= distance
		case "R":
			current.X += distance
		}

		totalDistance += distance
		history[idx] = &Point{X: current.X, Y: current.Y}
	}

	return totalDistance/2 + shoelaceFormula(history) + 1
}

func PartTwo(input string) int {
	lines := utils.EmbedStringToStringSlice(input)

	totalDistance := 0
	history := make([]*Point, len(lines))
	current := &Point{X: 0, Y: 0}
	for idx, line := range lines {
		splitInput := strings.Fields(line)

		direction := splitInput[2][7]
		distanceTmp, _ := strconv.ParseInt(splitInput[2][2:7], 16, 64)
		distance := int(distanceTmp)

		switch direction {
		case '3': // 'U'
			current.Y += distance
		case '1': // 'D'
			current.Y -= distance
		case '2': // 'L'
			current.X -= distance
		case '0': // 'R'
			current.X += distance
		}

		totalDistance += distance
		history[idx] = &Point{X: current.X, Y: current.Y}
	}

	// Picks Theorem => https://w.wiki/8YJS
	return totalDistance/2 + shoelaceFormula(history) + 1
}

func main() {
	fmt.Println("Day Eighteen - Lavaduct Lagoon")
	fmt.Println("##############################")

	timer := time.Now()
	fmt.Printf("Part 1 Answer - %v\n", PartOne(Input))
	fmt.Printf("Part 1 Time   - %v\n", time.Since(timer))

	timer = time.Now()
	fmt.Printf("Part 2 Answer - %v\n", PartTwo(Input))
	fmt.Printf("Part 2 Time   - %v\n", time.Since(timer))
}
