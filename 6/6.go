package six

import (
	"aoc2022/input"
)

type Stack []string

// func firstStartOfX(text string, xLen int) int {
// 	buffer := map[string]int{}
// 	for i, c := range text {
// 		if _, ok := buffer[string(c)]; !ok && len(buffer) == xLen-1 {
// 			return i
// 		}

// 		if i > xLen-1 {
// 			oldChar := string(text[i-(xLen-1)])
// 			buffer[oldChar]--
// 			val := buffer[oldChar]
// 			if val == 0 {
// 				delete(buffer, oldChar)
// 			}
// 		}
// 		buffer[string(c)]++
// 	}

// 	return -1
// }

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
