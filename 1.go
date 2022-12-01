package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func part1(text string) int {
	elvesCalories := strings.Split(text, "\n\n")
	mostCalories := 0
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
		if calorieTotal > mostCalories {
			mostCalories = calorieTotal
		}
	}
	return mostCalories
}

func main() {
	data, err := ioutil.ReadFile("1.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(data)

	part1Result := part1(text)
	fmt.Println(part1Result)
}
