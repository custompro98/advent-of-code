package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var targets = []string{
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

type Match struct {
	startIdx int
	endIdx   int
	value    string
}

func main() {
	lines, err := readLines(os.Args[1])
	digits := make([]int, 0)

	if err != nil {
		panic(fmt.Sprintf("failed to parse file: %v", err))
	}

	for idx, line := range lines {
		matches := make([]Match, 0)

		for _, target := range targets {
			startIdx := strings.Index(line, target)

			if startIdx > -1 {
				matches = append(matches, Match{
					startIdx: startIdx,
					endIdx:   startIdx + len(target),
					value:    line[startIdx : startIdx+len(target)],
				})
			}

			startIdx = strings.LastIndex(line, target)

			if startIdx > -1 {
				matches = append(matches, Match{
					startIdx: startIdx,
					endIdx:   startIdx + len(target),
					value:    line[startIdx : startIdx+len(target)],
				})
			}
		}

		sort.Slice(matches, func(i, j int) bool {
			return matches[i].startIdx < matches[j].startIdx
		})

		first := line[matches[0].startIdx:matches[0].endIdx]
		last := line[matches[len(matches)-1].startIdx:matches[len(matches)-1].endIdx]

		if idx == 5 {
			fmt.Println("---")
			fmt.Println(line)
			fmt.Println(matches)
			fmt.Println(first, last)
			fmt.Println("---")
		}

		first = strings.ReplaceAll(first, "one", "1")
		first = strings.ReplaceAll(first, "two", "2")
		first = strings.ReplaceAll(first, "three", "3")
		first = strings.ReplaceAll(first, "four", "4")
		first = strings.ReplaceAll(first, "five", "5")
		first = strings.ReplaceAll(first, "six", "6")
		first = strings.ReplaceAll(first, "seven", "7")
		first = strings.ReplaceAll(first, "eight", "8")
		first = strings.ReplaceAll(first, "nine", "9")

		last = strings.ReplaceAll(last, "one", "1")
		last = strings.ReplaceAll(last, "two", "2")
		last = strings.ReplaceAll(last, "three", "3")
		last = strings.ReplaceAll(last, "four", "4")
		last = strings.ReplaceAll(last, "five", "5")
		last = strings.ReplaceAll(last, "six", "6")
		last = strings.ReplaceAll(last, "seven", "7")
		last = strings.ReplaceAll(last, "eight", "8")
		last = strings.ReplaceAll(last, "nine", "9")

		// fmt.Println(idx, first, last)

		i, err := strconv.Atoi(fmt.Sprintf("%v%v", first, last))

		if err != nil {
			panic(fmt.Sprintf("failed to convert to int: %v", err))
		}

		digits = append(digits, i)
	}

	sum := 0

	for _, i := range digits {
		sum += i
	}

	fmt.Println(sum)
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
