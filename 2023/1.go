package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	// "strings"
)

func main() {
	lines, err := readLines(os.Args[1])
    digits := make([]int, 0)

    if err != nil {
        panic(fmt.Sprintf("failed to parse file: %v", err))
    }

    r, err := regexp.Compile("([0-9])")

    if err != nil {
        panic(fmt.Sprintf("failed to make regex: %v", err))
    }

    for _, line := range lines {
        m := r.FindAllString(line, -1)

        // m[0] = strings.ReplaceAll(m[0], "one", "1")
        // m[0] = strings.ReplaceAll(m[0], "two", "2")
        // m[0] = strings.ReplaceAll(m[0], "three", "3")
        // m[0] = strings.ReplaceAll(m[0], "four", "4")
        // m[0] = strings.ReplaceAll(m[0], "five", "5")
        // m[0] = strings.ReplaceAll(m[0], "six", "6")
        // m[0] = strings.ReplaceAll(m[0], "seven", "7")
        // m[0] = strings.ReplaceAll(m[0], "eight", "8")
        // m[0] = strings.ReplaceAll(m[0], "nine", "9")
        //
        // m[len(m)-1] = strings.ReplaceAll(m[len(m)-1], "one", "1")
        // m[len(m)-1] = strings.ReplaceAll(m[len(m)-1], "two", "2")
        // m[len(m)-1] = strings.ReplaceAll(m[len(m)-1], "three", "3")
        // m[len(m)-1] = strings.ReplaceAll(m[len(m)-1], "four", "4")
        // m[len(m)-1] = strings.ReplaceAll(m[len(m)-1], "five", "5")
        // m[len(m)-1] = strings.ReplaceAll(m[len(m)-1], "six", "6")
        // m[len(m)-1] = strings.ReplaceAll(m[len(m)-1], "seven", "7")
        // m[len(m)-1] = strings.ReplaceAll(m[len(m)-1], "eight", "8")
        // m[len(m)-1] = strings.ReplaceAll(m[len(m)-1], "nine", "9")

        fmt.Println(line, m[0], m[len(m)-1])

        i, err := strconv.Atoi(fmt.Sprintf("%v%v", m[0], m[len(m)-1]))

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
