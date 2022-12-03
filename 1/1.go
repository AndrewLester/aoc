package one

import (
	"aoc2022/input"
	"sort"
	"strconv"
	"strings"
)

func caloriesList(text string) []int {
	elvesCalories := strings.Split(text, "\n\n")
	elvesCaloriesNums := make([]int, len(elvesCalories))
	for _, elfCalories := range elvesCalories {
		calories := strings.Split(elfCalories, "\n")
		calorieTotal := 0
		for _, value := range calories {
			valueAsNumber, err := strconv.Atoi(value)
			if err != nil {
				continue
			}
			calorieTotal += valueAsNumber
		}
		elvesCaloriesNums = append(elvesCaloriesNums, calorieTotal)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elvesCaloriesNums)))

	return elvesCaloriesNums
}

func Part1() int {
	text := input.ReadInput(1)

	elvesCaloriesNums := caloriesList(text)
	return elvesCaloriesNums[0]
}

func Part2() int {
	text := input.ReadInput(1)

	elvesCaloriesNums := caloriesList(text)
	return elvesCaloriesNums[0] + elvesCaloriesNums[1] + elvesCaloriesNums[2]
}
