package six

import (
	"aoc2022/input"
)

func firstStartOfX(text string, xLen int) int {
	buffer := map[string]int{}

	for i, c := range text {
		buffer[string(c)]++

		if i >= xLen {
			oldChar := string(text[i-xLen])
			buffer[oldChar]--
			val := buffer[oldChar]
			if val == 0 {
				delete(buffer, oldChar)
			}
		}

		if len(buffer) == xLen {
			return i + 1
		}
	}

	return -1
}

func Part1() int {
	text := input.ReadInput(6)

	return firstStartOfX(text, 4)
}

func Part2() int {
	text := input.ReadInput(6)

	return firstStartOfX(text, 14)
}
