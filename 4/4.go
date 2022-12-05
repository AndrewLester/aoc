package two

import (
	"aoc2022/input"
	"strconv"
	"strings"
)

func Part1() int {
	text := input.ReadInput(4)

	lines := strings.Split(text, "\n")

	fullyContained := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		elfRanges := strings.Split(line, ",")

		elf1Range := strings.Split(elfRanges[0], "-")
		elf1Start, _ := strconv.Atoi(elf1Range[0])
		elf1End, _ := strconv.Atoi(elf1Range[1])

		elf2Range := strings.Split(elfRanges[1], "-")
		elf2Start, _ := strconv.Atoi(elf2Range[0])
		elf2End, _ := strconv.Atoi(elf2Range[1])

		// (same start time AND elf 2 is after elf 1) OR elf 2 starts before elf1
		if (elf1Start == elf2Start && elf2End > elf1End) || elf2Start < elf1Start {
			elf1Start, elf2Start, elf1End, elf2End = elf2Start, elf1Start, elf2End, elf1End
		}

		if elf1Start <= elf2Start && elf1End >= elf2End {
			fullyContained++
		}
	}
	return fullyContained
}

func Part2() int {
	text := input.ReadInput(4)

	lines := strings.Split(text, "\n")

	overlapping := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		elfRanges := strings.Split(line, ",")

		elf1Range := strings.Split(elfRanges[0], "-")
		elf1Start, _ := strconv.Atoi(elf1Range[0])
		elf1End, _ := strconv.Atoi(elf1Range[1])

		elf2Range := strings.Split(elfRanges[1], "-")
		elf2Start, _ := strconv.Atoi(elf2Range[0])
		elf2End, _ := strconv.Atoi(elf2Range[1])

		if elf1Start > elf2Start {
			elf1Start, elf2Start, elf1End, elf2End = elf2Start, elf1Start, elf2End, elf1End
		}

		if elf2Start <= elf1End {
			overlapping++
		}
	}
	return overlapping
}
