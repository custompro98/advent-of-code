package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Scratchcard struct {
	id             string
	winningNumbers []string
	yourNumbers    []string
	isChecked      bool
	numWinners     int
}

func main() {
	lines, err := readLines(os.Args[1])

	if err != nil {
		panic(fmt.Sprintf("could not read file: %v", err))
	}

	var scratchcards []*Scratchcard

	for _, line := range lines {
		sc := Scratchcard{}

		sc.id = strings.Split(strings.Trim(strings.Split(line, ":")[0], " "), " ")[1]
		numberSet := strings.Split(line, ":")[1]
		winningNumbers := strings.Split(numberSet, "|")[0]
		yourNumbers := strings.Split(numberSet, "|")[1]

		for _, winningNumberStr := range strings.Split(winningNumbers, " ") {
			winningNumber := strings.Trim(winningNumberStr, " ")

			if winningNumber == "" {
				continue
			}

			sc.winningNumbers = append(sc.winningNumbers, winningNumber)
		}

		for _, yourNumberStr := range strings.Split(yourNumbers, " ") {
			yourNumber := strings.Trim(yourNumberStr, " ")

			if yourNumber == "" {
				continue
			}

			sc.yourNumbers = append(sc.yourNumbers, yourNumber)
		}

		scratchcards = append(scratchcards, &sc)
	}

	sum1 := 0

	for _, sc := range scratchcards {
		sum1 += sc.GetScore()
	}

	fmt.Println("part 1:", sum1)

	sum2 := countAllWinningCards(scratchcards)

	fmt.Println("part 2:", sum2)
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

func (sc *Scratchcard) GetScore() int {
	if !sc.isChecked {
		numWinners := 0

		for _, winningNumber := range sc.winningNumbers {
			for _, yourNumber := range sc.yourNumbers {
				if yourNumber == winningNumber {
					numWinners += 1
				}
			}
		}

		sc.numWinners = numWinners
		sc.isChecked = true
	}

	return int(math.Pow(2.0, float64(sc.numWinners-1)))
}

func countAllWinningCards(scratchcards []*Scratchcard) int {
	sum := 0

	for i := range scratchcards {
		sum += 1 + countNestedWinningCards(scratchcards, i)
	}

	return sum
}

func countNestedWinningCards(scratchcards []*Scratchcard, currentCardIdx int) int {
	currentCard := scratchcards[currentCardIdx]

	acc := 0

	for i := 0; i < currentCard.numWinners; i++ {
		acc += countNestedWinningCards(scratchcards, currentCardIdx+i+1)
	}

	return acc + currentCard.numWinners
}
