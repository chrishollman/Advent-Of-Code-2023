package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"

	"github.com/chrishollman/Advent-Of-Code-2023/utils"
)

//go:embed input.txt
var Input string

func PartOne(input string) int {
	lines := utils.EmbedStringToStringSlice(input)

	sequence := strings.Split(lines[0], "")

	// Parse
	var associations = make(map[string][]string, len(lines))

	values := lines[2:]
	for _, value := range values {
		associations[value[0:3]] = []string{value[7:10], value[12:15]}
	}

	current := associations["AAA"]
	ctr, i := 0, 0
	for {
		var newKey string
		ctr++
		switch sequence[i] {
		case "L":
			newKey = current[0]
		case "R":
			newKey = current[1]
		}
		current = associations[newKey]
		if newKey == "ZZZ" {
			break
		}

		i++
		if i == len(sequence) {
			i = 0
		}
	}

	return ctr
}

func PartTwo(input string) int {
	lines := utils.EmbedStringToStringSlice(input)

	sequence := strings.Split(lines[0], "")

	// Parse
	var associations = make(map[string][]string, len(lines))
	var toFollow []string
	var steps []int

	values := lines[2:]
	for _, value := range values {
		associations[value[0:3]] = []string{value[7:10], value[12:15]}
		if value[2] == 'A' {
			toFollow = append(toFollow, value[0:3])
		}
	}

	for _, value := range toFollow {
		current := associations[value]
		ctr, i := 0, 0
		for {
			var newKey string
			ctr++
			switch sequence[i] {
			case "L":
				newKey = current[0]
			case "R":
				newKey = current[1]
			}
			current = associations[newKey]
			if newKey[2] == 'Z' {
				break
			}

			i++
			if i == len(sequence) {
				i = 0
			}
		}
		steps = append(steps, ctr)
	}

	// Brute force LCM
	return lcmOfList(steps)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcmOfList(numbers []int) int {
	l := numbers[0] * numbers[1] / gcd(numbers[0], numbers[1])

	for i := 2; i < len(numbers); i++ {
		l = l * numbers[i] / gcd(l, numbers[i])
	}

	return l
}

func main() {
	fmt.Println("Day Eight - Haunted Wasteland")
	fmt.Println("#############################")

	timer := time.Now()
	fmt.Printf("Part 1 Answer - %v\n", PartOne(Input))
	fmt.Printf("Part 1 Time   - %v\n", time.Since(timer))

	timer = time.Now()
	fmt.Printf("Part 2 Answer - %v\n", PartTwo(Input))
	fmt.Printf("Part 2 Time   - %v\n", time.Since(timer))
}
