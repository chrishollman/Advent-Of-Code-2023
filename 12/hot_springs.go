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

var memo = make(map[string]int)

func PartOne(input string) int {
	lines := utils.EmbedStringToStringSlice(input)

	total := 0

	for _, line := range lines {
		spl := strings.Split(line, " ")

		conditionRecord, pattern := spl[0], strToIntSlice(spl[1])
		total += count(conditionRecord, pattern)
	}

	return total
}

func PartTwo(input string) int {
	lines := utils.EmbedStringToStringSlice(input)

	total := 0

	for _, line := range lines {
		spl := strings.Split(line, " ")

		conditionRecord, pattern := spl[0], strToIntSlice(spl[1])

		bigConditionRecordSize := len(conditionRecord)*5 + 4
		bigPatternSize := len(pattern) * 5

		sb := strings.Builder{}
		sb.Grow(bigConditionRecordSize)
		for i := 0; i < 5; i++ {
			sb.WriteString(conditionRecord)
			if i < 4 {
				sb.WriteByte('?')
			}
		}
		bigConditionRecord := sb.String()

		bigRequirement := make([]int, 0, bigPatternSize)
		for i := 0; i < 5; i++ {
			bigRequirement = append(bigRequirement, pattern...)
		}

		total += count(bigConditionRecord, bigRequirement)
	}

	return total
}

func count(conditionRecord string, pattern []int) int {
	if len(conditionRecord) == 0 {
		if len(pattern) > 0 {
			return 0
		}
		return 1
	}

	hashKey := hash(conditionRecord, pattern)
	if value, ok := memo[hashKey]; ok {
		return value
	}
	if strings.ContainsRune(conditionRecord, '#') && len(pattern) == 0 {
		return setAndGet(hashKey, 0)
	}
	if !strings.ContainsRune(conditionRecord, '#') && len(pattern) == 0 {
		return setAndGet(hashKey, 1)
	}

	requiredPattern := getPatternRequirement(pattern)
	sb := strings.Builder{}
	sb.Grow(len(requiredPattern))
	patternIndex := 0

	for i := 0; i < len(conditionRecord); i++ {
		char := conditionRecord[i]
		switch char {
		case '.':
			if sb.Len() > 0 {
				sb.WriteByte('.')
			}
		case '#':
			sb.WriteByte('#')
		case '?':
			left := conditionRecord[:i] + "#" + conditionRecord[i+1:]
			right := conditionRecord[:i] + "." + conditionRecord[i+1:]
			return setAndGet(hashKey, count(left, pattern)+count(right, pattern))
		}

		if sb.Len() == len(requiredPattern) && sb.String() == requiredPattern {
			return setAndGet(hashKey, count(conditionRecord[i+1:], pattern[1:]))
		}

		if sb.Len() > len(requiredPattern) || strings.Count(sb.String(), "#") > pattern[patternIndex] {
			return setAndGet(hashKey, 0)
		}
	}

	return setAndGet(hashKey, 0)
}

func getPatternRequirement(requirement []int) string {
	if len(requirement) == 0 {
		return ""
	}
	num := requirement[0]
	sb := strings.Builder{}
	sb.Grow(num + len(requirement) - 1)
	for i := 0; i < num; i++ {
		sb.WriteByte('#')
	}
	if len(requirement) > 1 {
		sb.WriteByte('.')
	}
	return sb.String()
}

func hash(arrangement string, requirement []int) string {
	var sb strings.Builder
	sb.WriteString(arrangement)
	sb.WriteByte('-')
	for i, num := range requirement {
		if i > 0 {
			sb.WriteByte(',')
		}
		sb.WriteString(strconv.Itoa(num))
	}
	return sb.String()
}

func setAndGet(key string, value int) int {
	memo[key] = value
	return value
}

func strToIntSlice(in string) []int {
	tmp := strings.Split(in, ",")
	nums := make([]int, 0, len(tmp))
	for _, v := range tmp {
		num, _ := strconv.Atoi(v)
		nums = append(nums, num)
	}
	return nums
}

func main() {
	fmt.Println("Day Twelve - Hot Springs")
	fmt.Println("########################")

	timer := time.Now()
	fmt.Printf("Part 1 Answer - %v\n", PartOne(Input))
	fmt.Printf("Part 1 Time   - %v\n", time.Since(timer))

	timer = time.Now()
	fmt.Printf("Part 2 Answer - %v\n", PartTwo(Input))
	fmt.Printf("Part 2 Time   - %v\n", time.Since(timer))
}
