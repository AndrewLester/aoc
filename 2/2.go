package two

import (
	"aoc2022/input"
	"strings"
)

type Shape int

const (
	Rock Shape = iota
	Paper
	Scissors
)

var shapes = [...]Shape{Rock, Paper, Scissors}

var shapeStrings = map[string]Shape{
	"A": Rock,
	"X": Rock,
	"B": Paper,
	"Y": Paper,
	"C": Scissors,
	"Z": Scissors,
}

var scores = map[Shape]int{
	Rock:     1,
	Paper:    2,
	Scissors: 3,
}

func result(opponentShape Shape, playerShape Shape) int {
	if opponentShape == playerShape {
		return 0
	}
	if (playerShape == Rock && opponentShape == Scissors) || (playerShape == Paper && opponentShape == Rock) || (playerShape == Scissors && opponentShape == Paper) {
		return 1
	}
	return -1
}

func computePoints(opponent string, player string) int {
	var points int

	opponentShape := shapeStrings[opponent]
	playerShape := shapeStrings[player]

	if playerShape == opponentShape {
		points = 3
	} else if result(opponentShape, playerShape) == 1 {
		points = 6
	}

	points += scores[playerShape]
	return points
}

func Part1() int {
	text := input.ReadInput(2)

	lines := strings.Split(text, "\n")
	totalPoints := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		characters := strings.Split(line, " ")
		points := computePoints(characters[0], characters[1])
		totalPoints += points
	}
	return totalPoints
}

func Part2() int {
	text := input.ReadInput(2)

	lines := strings.Split(text, "\n")
	totalPoints := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		characters := strings.Split(line, " ")
		playerString := "X"
		playerMoves := [...]string{"X", "Y", "Z"}
		for i, shape := range shapes {
			matchResult := result(shapeStrings[characters[0]], shape)
			if (matchResult == -1 && characters[1] == "X") || (matchResult == 0 && characters[1] == "Y") || (matchResult == 1 && characters[1] == "Z") {
				playerString = playerMoves[i]
				break
			}
		}
		points := computePoints(characters[0], playerString)
		totalPoints += points
	}
	return totalPoints
}
