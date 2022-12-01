package main

import (
	"fmt"
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

func part1(text string) int {
	elvesCaloriesNums := caloriesList(text)
	return elvesCaloriesNums[0]
}

func part2(text string) []int {
	elvesCaloriesNums := caloriesList(text)
	return elvesCaloriesNums[0:3]
}

func main() {
	text := ReadInput(1)

	part1Result := part1(text)
	fmt.Println("Part 1, ", part1Result)

	part2Result := part2(text)
	part2Total := part2Result[0] + part2Result[1] + part2Result[2]
	fmt.Println("Part 2, ", part2Total)
}
