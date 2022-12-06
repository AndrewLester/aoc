package five

import (
	"aoc2022/input"
	"regexp"
	"strconv"
	"strings"
)

type Stack []string

func Part1() string {
	text := input.ReadInput(5)

	lines := strings.Split(text, "\n")

	readingMoves := false
	stacks := []Stack{}
	for _, line := range lines {
		if len(line) == 0 {
			if readingMoves {
				break
			}
			readingMoves = true
			continue
		}

		if !readingMoves {
			for i, c := range line {
				if i%4 == 0 {
					continue
				}
				if int(c) >= 49 && int(c) <= 57 {
					break
				}

				stackNum := i / 4
				crateChar := string(c)
				if crateChar != "[" && crateChar != "]" && crateChar != " " {
					if len(stacks) <= stackNum {
						for j := len(stacks); j < stackNum+1; j++ {
							stacks = append(stacks, []string{})
						}
					}
					stacks[stackNum] = append([]string{string(crateChar)}, stacks[stackNum]...)
				}
			}

			continue
		}

		re := regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)
		movesLine := re.ReplaceAllString(line, "$1 $2 $3")
		moves := strings.Split(movesLine, " ")

		count, _ := strconv.Atoi(moves[0])
		from, _ := strconv.Atoi(moves[1])
		to, _ := strconv.Atoi(moves[2])

		from--
		to--
		for i := 0; i < count; i++ {
			topOfFrom := stacks[from][len(stacks[from])-1]

			stacks[to] = append(stacks[to], topOfFrom)

			// Pop top off from
			stacks[from] = stacks[from][:len(stacks[from])-1]
		}
	}

	topOfStacks := ""
	for _, stack := range stacks {
		topOfStack := stack[len(stack)-1]
		topOfStacks += topOfStack
	}

	return topOfStacks
}

func Part2() string {
	text := input.ReadInput(5)

	lines := strings.Split(text, "\n")

	readingMoves := false
	stacks := []Stack{}
	for _, line := range lines {
		if len(line) == 0 {
			if readingMoves {
				break
			}
			readingMoves = true
			continue
		}

		if !readingMoves {
			for i, c := range line {
				if i%4 == 0 {
					continue
				}
				if int(c) >= 49 && int(c) <= 57 {
					break
				}

				stackNum := i / 4
				crateChar := string(c)
				if crateChar != "[" && crateChar != "]" && crateChar != " " {
					if len(stacks) <= stackNum {
						for j := len(stacks); j < stackNum+1; j++ {
							stacks = append(stacks, []string{})
						}
					}
					stacks[stackNum] = append([]string{string(crateChar)}, stacks[stackNum]...)
				}
			}

			continue
		}

		re := regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)
		movesLine := re.ReplaceAllString(line, "$1 $2 $3")
		moves := strings.Split(movesLine, " ")

		count, _ := strconv.Atoi(moves[0])
		from, _ := strconv.Atoi(moves[1])
		to, _ := strconv.Atoi(moves[2])

		from--
		to--

		topOfFrom := stacks[from][len(stacks[from])-count:]

		stacks[to] = append(stacks[to], topOfFrom...)

		// Pop top off from
		stacks[from] = stacks[from][:len(stacks[from])-count]
	}

	topOfStacks := ""
	for _, stack := range stacks {
		topOfStack := stack[len(stack)-1]
		topOfStacks += topOfStack
	}

	return topOfStacks
}
