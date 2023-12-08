package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction struct {
	xmod int
	ymod int
}

var (
	NW = Direction{-1, -1}
	N  = Direction{0, -1}
	NE = Direction{1, -1}
	E  = Direction{1, 0}
	SE = Direction{1, 1}
	S  = Direction{0, 1}
	SW = Direction{-1, 1}
	W  = Direction{-1, 0}
)

var directions = []Direction{NW, N, NE, E, SE, S, SW, W}

type Coordinate struct {
	rowno int
	colno int
}

type Number struct {
	coords []Coordinate
	value  int
}

type Gear struct {
	coord Coordinate
	partOne Number
	partTwo Number
}

type Schematic = [][]string

func main() {
	lines, err := readLines(os.Args[1])

	if err != nil {
		panic(fmt.Sprintf("could not read file: %v", err))
	}

	sch := make([][]string, len(lines))

	for idx := range sch {
		sch[idx] = make([]string, len(lines[idx]))
	}

	for rowno, line := range lines {
		for colno, char := range strings.Split(line, "") {
			sch[rowno][colno] = char
		}
	}

	numbers := make([]*Number, 0)
	partnos := make([]*Number, 0)

	rowno := 0
	colno := 0

	for rowno < len(sch) {
		for colno < len(sch[rowno]) {
			if isDigit(sch[rowno][colno]) {
				numStr := ""
				coords := make([]Coordinate, 0)

				for colno < len(sch[rowno]) && isDigit(sch[rowno][colno]) {
					numStr = fmt.Sprintf("%v%v", numStr, sch[rowno][colno])
					coords = append(coords, Coordinate{rowno, colno})
					colno += 1
				}

				value, _ := strconv.Atoi(numStr)

				numbers = append(numbers, &Number{coords, value})
			} else {
				colno += 1
			}
		}

		rowno += 1
		colno = 0
	}

	sum := 0

	for _, number := range numbers {
		if number.isPartNumber(sch) {
			partnos = append(partnos, number)
			sum += number.value
		}
	}

	fmt.Println("part 1:", sum)

	gears := make([]Gear, 0)

	rowno = 0
	colno = 0

	for rowno < len(sch) {
		for colno < len(sch[rowno]) {
			coord := Coordinate{rowno, colno}
			if coord.isGear(sch, partnos) {
				adjacencies := coord.partAdjacenices(sch, partnos)
				gears = append(gears, Gear{
					coord:   coord,
					partOne: *adjacencies[0],
					partTwo: *adjacencies[1],
				})
			}

			colno += 1
		}

		rowno += 1
		colno = 0
	}

	sum2 := 0

	for _, gear := range gears {
		sum2 += gear.partOne.value * gear.partTwo.value
	}

	fmt.Println("part 2:", sum2)
}

func isDigit(s string) bool {
	return (s == "0" ||
		s == "1" ||
		s == "2" ||
		s == "3" ||
		s == "4" ||
		s == "5" ||
		s == "6" ||
		s == "7" ||
		s == "8" ||
		s == "9")
}

func isDot(s string) bool {
	return s == "."
}

func isStar(s string) bool {
	return s == "*"
}

func isSymbol(s string) bool {
	return !isDot(s) && !isDigit(s)
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

func (n *Number) isPartNumber(sch Schematic) bool {
	for _, coord := range n.coords {
		if coord.isAdjacentToSymbol(sch) {
			return true
		}
	}

	return false
}

func (c *Coordinate) isAdjacentToSymbol(sch Schematic) bool {
	for _, d := range directions {
		newCoord := Coordinate{c.rowno + d.ymod, c.colno + d.xmod}
		if newCoord.isValid(sch) && isSymbol(sch[newCoord.rowno][newCoord.colno]) {
			return true
		}
	}

	return false
}

func (c *Coordinate) isGear(sch Schematic, partnos []*Number) bool {
	if (!isStar(sch[c.rowno][c.colno])) {
		return false
	}

	parts := c.partAdjacenices(sch, partnos)

	return len(parts) == 2
}

func (c *Coordinate) partAdjacenices(sch Schematic, partnos []*Number) []*Number {
	adjacencies := make([]*Number, 0)

	for _, partno := range partnos {
		for _, coord := range partno.coords {
			if c.isAdjacent(&coord, sch) && !isPartAlreadyCounted(adjacencies, partno) {
				adjacencies = append(adjacencies, partno)
			}
		}
	}

	return adjacencies
}

func isPartAlreadyCounted(partnos []*Number, part *Number) bool {
	for _, partno := range partnos {
		if partno == part {
			return true
		}
	}

	return false
}

func (c *Coordinate) isAdjacent(oc *Coordinate, sch Schematic) bool {
	for _, d := range directions {
		nc := Coordinate{oc.rowno + d.ymod, oc.colno + d.xmod}
		if nc.isValid(sch) && c.rowno == nc.rowno && c.colno == nc.colno {
			return true
		}
	}

	return false
}

func (c *Coordinate) isValid(sch Schematic) bool {
	return c.rowno > 0 && c.colno > 0 && c.rowno < len(sch) && c.colno < len(sch[c.rowno])
}
