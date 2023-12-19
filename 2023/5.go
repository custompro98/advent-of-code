package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var STANZA_SEEDS = "seeds"
var STANZA_SEED_TO_SOIL = "seed-to-soil map"
var STANZA_SOIL_TO_FERTILIZER = "soil-to-fertilizer map"
var STANZA_FERTILIZER_TO_WATER = "fertilizer-to-water map"
var STANZA_WATER_TO_LIGHT = "water-to-light map"
var STANZA_LIGHT_TO_TEMPERATURE = "light-to-temperature map"
var STANZA_TEMPERATURE_TO_HUMIDITY = "temperature-to-humidity map"
var STANZA_HUMIDITY_TO_LOCATION = "humidity-to-location map"

type SourceMap struct {
	title            string
	sourceStart      []int
	destinationStart []int
	length           []int
}

type Seed struct {
	seed        int
	soil        int
	fertilizer  int
	water       int
	light       int
	temperature int
	humidity    int
	location    int
}

func main() {
	lines, err := readLines(os.Args[1])

	if err != nil {
		panic(fmt.Sprintf("could not read file: %v", err))
	}

	stanzas := strings.Split(lines, "\n\n")

	seeds := make([]*Seed, 0)
	seedPairs := make([]*Seed, 0)
	seedToSoil := SourceMap{title: STANZA_SEED_TO_SOIL}
	soilToFertilizer := SourceMap{title: STANZA_SOIL_TO_FERTILIZER}
	fertilizerToWater := SourceMap{title: STANZA_FERTILIZER_TO_WATER}
	waterToLight := SourceMap{title: STANZA_WATER_TO_LIGHT}
	lightToTemperature := SourceMap{title: STANZA_LIGHT_TO_TEMPERATURE}
	temperatureToHumidity := SourceMap{title: STANZA_TEMPERATURE_TO_HUMIDITY}
	humidityToLocation := SourceMap{title: STANZA_HUMIDITY_TO_LOCATION}

	for _, stanza := range stanzas {
		splits := strings.Split(stanza, ":")

		title := splits[0]
		values := strings.Trim(strings.Trim(splits[1], " "), "\n")

		switch title {
		case STANZA_SEEDS:
			buildSeeds(&seeds, values)
			buildSeedPairs(&seedPairs, values)
		case STANZA_SEED_TO_SOIL:
			buildSourceMap(&seedToSoil, values)
		case STANZA_SOIL_TO_FERTILIZER:
			buildSourceMap(&soilToFertilizer, values)
		case STANZA_FERTILIZER_TO_WATER:
			buildSourceMap(&fertilizerToWater, values)
		case STANZA_WATER_TO_LIGHT:
			buildSourceMap(&waterToLight, values)
		case STANZA_LIGHT_TO_TEMPERATURE:
			buildSourceMap(&lightToTemperature, values)
		case STANZA_TEMPERATURE_TO_HUMIDITY:
			buildSourceMap(&temperatureToHumidity, values)
		case STANZA_HUMIDITY_TO_LOCATION:
			buildSourceMap(&humidityToLocation, values)
		default:
			fmt.Println("nothing to see here")
		}
	}

	for _, seed := range seeds {
		seed.soil = seedToSoil.GetValue(seed.seed)
		seed.fertilizer = soilToFertilizer.GetValue(seed.soil)
		seed.water = fertilizerToWater.GetValue(seed.fertilizer)
		seed.light = waterToLight.GetValue(seed.water)
		seed.temperature = lightToTemperature.GetValue(seed.light)
		seed.humidity = temperatureToHumidity.GetValue(seed.temperature)
		seed.location = humidityToLocation.GetValue(seed.humidity)
	}

	minLocation := math.MaxInt

	for _, seed := range seeds {
		if seed.location < minLocation {
			minLocation = seed.location
		}
	}

	fmt.Println("part 1", minLocation)

	for _, seed := range seedPairs {
		seed.soil = seedToSoil.GetValue(seed.seed)
		seed.fertilizer = soilToFertilizer.GetValue(seed.soil)
		seed.water = fertilizerToWater.GetValue(seed.fertilizer)
		seed.light = waterToLight.GetValue(seed.water)
		seed.temperature = lightToTemperature.GetValue(seed.light)
		seed.humidity = temperatureToHumidity.GetValue(seed.temperature)
		seed.location = humidityToLocation.GetValue(seed.humidity)
	}

	minLocation2 := math.MaxInt

	for _, seed := range seedPairs {
		if seed.location < minLocation2 {
			minLocation2 = seed.location
		}
	}

	fmt.Println("part 2", minLocation2)
}

func buildSeeds(seeds *[]*Seed, values string) {
	for _, seed := range strings.Split(values, " ") {
		seed, _ := strconv.Atoi(seed)
		*seeds = append(*seeds, &Seed{
			seed: seed,
		})
	}
}

func buildSeedPairs(seeds *[]*Seed, values string) {
	splits := strings.Split(values, " ")
	seedSourceMap := SourceMap{}

	for i := 0; i < len(splits); i++ {
		num, _ := strconv.Atoi(splits[i])

		if i%2 == 0 {
			// If seed range start
			seedSourceMap.sourceStart = append(seedSourceMap.sourceStart, num)
		} else {
			// If seed range length
			seedSourceMap.length = append(seedSourceMap.length, num)
		}
	}

	for idx, seed := range seedSourceMap.sourceStart {
		for i := 0; i < seedSourceMap.length[idx]; i++ {
			*seeds = append(*seeds, &Seed{
				seed: seed + i,
			})
		}
	}
}

func buildSourceMap(sm *SourceMap, lines string) {
	destinationStart := 0
	sourceStart := 0
	length := 0

	for _, line := range strings.Split(lines, "\n") {
		for idx, v := range strings.Split(line, " ") {
			num, _ := strconv.Atoi(v)

			if idx == 0 {
				destinationStart = num
			} else if idx == 1 {
				sourceStart = num
			} else if idx == 2 {
				length = num
			}
		}

		sm.destinationStart = append(sm.destinationStart, destinationStart)
		sm.sourceStart = append(sm.sourceStart, sourceStart)
		sm.length = append(sm.length, length)
	}
}

func (sm *SourceMap) GetValue(source int) int {
	for idx, sourceStart := range sm.sourceStart {
		if source >= sourceStart && source <= sourceStart+sm.length[idx] {
			diff := source - sourceStart
			return sm.destinationStart[idx] + diff
		}
	}
	return source
}

func readLines(path string) (string, error) {
	out := ""

	file, err := os.Open(path)

	if err != nil {
		return out, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		out = fmt.Sprintf("%v\n%v", out, scanner.Text())
	}

	return strings.Trim(out, "\n"), nil
}
