package three

import (
	"aoc2022/input"
	"strings"
)

func priority(item string) int {
	code := int(item[0])

	// Uppercase
	if code <= 90 {
		return code - 38
	} else {
		return code - 96
	}
}

func Part1() int {
	text := input.ReadInput(3)

	lines := strings.Split(text, "\n")
	priorities := 0
	for _, line := range lines {
		compartmentOne := line[:len(line)/2]
		compartmentTwo := line[len(line)/2:]

		compartmentOneItems := map[string]struct{}{}

		for _, item := range strings.Split(compartmentOne, "") {
			compartmentOneItems[item] = struct{}{}
		}

		for _, item := range strings.Split(compartmentTwo, "") {
			if _, ok := compartmentOneItems[item]; ok {
				priorities += priority(item)
				break
			}
		}
	}
	return priorities
}

func Part2() int {
	text := input.ReadInput(3)

	lines := strings.Split(text, "\n")
	priorities := 0
	var elvesRucksacks []string
	for i, line := range lines {
		elvesRucksacks = append(elvesRucksacks, line)

		if i%3 != 2 {
			continue
		}

		elfOneItems := map[string]struct{}{}
		elfTwoItems := map[string]struct{}{}

		for _, item := range strings.Split(elvesRucksacks[0], "") {
			elfOneItems[item] = struct{}{}
		}

		for _, item := range strings.Split(elvesRucksacks[1], "") {
			elfTwoItems[item] = struct{}{}
		}

		for _, item := range strings.Split(elvesRucksacks[2], "") {
			_, elfOneHas := elfOneItems[item]
			_, elfTwohas := elfTwoItems[item]
			if elfOneHas && elfTwohas {
				priorities += priority(item)
				break
			}
		}

		elvesRucksacks = []string{}
	}
	return priorities
}
