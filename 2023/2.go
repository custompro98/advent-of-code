package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var MaxBlue = 14
var MaxRed = 12
var MaxGreen = 13

type Result struct {
	Blue  int
	Red   int
	Green int
}

type Game struct {
	Id      int
	Results []Result
}

func main() {
	lines, err := readLines(os.Args[1])
	games := make([]Game, 0)

	if err != nil {
		panic(fmt.Sprintf("failed to parse file: %v", err))
	}

	for _, line := range lines {
		game := Game{}

		gameResultsSplit := strings.Split(line, ":")
		gameStr := gameResultsSplit[0]
		resultStr := gameResultsSplit[1]

		gameNumStr := strings.Split(gameStr, " ")[1]
		resultStrs := strings.Split(resultStr, ";")

		gameNum, err := strconv.Atoi(gameNumStr)

		if err != nil {
			panic(fmt.Sprintf("could not parse game number :%v", err))
		}

		game.Id = gameNum

		for _, resultStr := range resultStrs {
			subResults := strings.Split(resultStr, ",")

			result := Result{}

			for _, subResult := range subResults {

				numColorPairs := strings.Split(subResult, " ")

				number, err := strconv.Atoi(numColorPairs[1])
				color := numColorPairs[2]

				if err != nil {
					panic(fmt.Sprintf("could not parse %v number :%v", color, err))
				}

				if color == "blue" {
					result.Blue = number
				} else if color == "red" {
					result.Red = number
				} else if color == "green" {
					result.Green = number
				}

			}
			game.Results = append(game.Results, result)
		}

		games = append(games, game)
	}

	sum := 0

	for _, game := range games {
		validResults := 0

		for _, result := range game.Results {
			if result.Blue <= MaxBlue && result.Red <= MaxRed && result.Green <= MaxGreen {
				validResults += 1
			}
		}

		if validResults == len(game.Results) {
			sum += game.Id
		}
	}

	fmt.Println("part 1: ", sum)

	gameMinimums := make([]Result, 0)

	for _, game := range games {
		gameMinimum := Result{}

		for _, result := range game.Results {
			if result.Blue > gameMinimum.Blue {
				gameMinimum.Blue = result.Blue
			}

			if result.Green > gameMinimum.Green {
				gameMinimum.Green = result.Green
			}

			if result.Red > gameMinimum.Red {
				gameMinimum.Red = result.Red
			}
		}

		gameMinimums = append(gameMinimums, gameMinimum)
	}

	powerSum := 0

	for _, minimum := range gameMinimums {
		powerSum += minimum.Blue * minimum.Green * minimum.Red
	}

	fmt.Println("part 2: ", powerSum)
}

func readLines(path string) ([]string, error) {
	out := make([]string, 0)

	file, err := os.Open(path)

	if err != nil {
		return out, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		out = append(out, scanner.Text())
	}

	return out, nil
}
