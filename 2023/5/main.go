package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func readLines(path string) []string {
	c, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(c), "\n")
}

type GardenThing struct {
	from         string
	to           string
	sources      []int //from, range, from, range
	destinations []int //to, range, ...
}

func (g *GardenThing) fillGaps() {
	l := len(g.sources)
	lastFound := 0
	for i := 0; i < l; i += 2 {
		from := g.sources[i]
		// r := g.sources[i+1]
		if lastFound < from {
			g.sources = append(g.sources, lastFound, from-lastFound)
			g.destinations = append(g.destinations, lastFound, from-lastFound)
		}
		lastFound = from + g.sources[i+1]
	}
}

func (g *GardenThing) String() string {
	return fmt.Sprintf("from: %s, to: %s, sources: %v, destinations: %v", g.from, g.to, g.sources, g.destinations)
}

func solve1(lines []string) {
	s := strings.Split(lines[0], " ")[1:]
	seeds := make([]int, len(s))
	for i, v := range s {
		seeds[i], _ = strconv.Atoi(v)
	}
	fmt.Println("seeds", seeds)
	garden := make([]GardenThing, 0)
	g := GardenThing{}
	for _, l := range lines[2:] {
		if l == "" {
			fmt.Println()
			g.fillGaps()
			fmt.Println(g)
			garden = append(garden, g)
			g = GardenThing{}
			continue
		}
		data := strings.Split(l, " ")
		if _, err := strconv.Atoi(data[0]); err != nil {
			kind := strings.Split(data[0], "-to-")
			from := kind[0]
			to := kind[1]
			g = GardenThing{from: from, to: to}
			// fmt.Println("from", from, "to", to)
		} else {
			values := make([]int, len(data))
			for i, v := range data {
				values[i], _ = strconv.Atoi(v)
			}
			if g.sources == nil {
				g.sources = make([]int, 0)
			}
			if g.destinations == nil {
				g.destinations = make([]int, 0)
			}

			g.sources = append(g.sources, values[1], values[2])
			g.destinations = append(g.destinations, values[0], values[2])

			// for i := 0; i < values[2]; i++ {
			// g.values = append(g.values, values[0]+i, values[1]+i)

			// }
			// fmt.Println("values", values)
		}
		// fmt.Println(l)
	}
	g.fillGaps()
	fmt.Println(g)
	garden = append(garden, g)
	g = GardenThing{}

	// gerarchy := []string{"seed", "soil", "fertilizer", "water", "light", "temperature", "humidity", "location"}
	// fmt.Println(garden["seed"])
	location := math.MaxInt64
	for _, s := range seeds {
		nextValue := s
		fmt.Printf("seed %d,", s)
		for _, g := range garden {
			for i := 0; i < len(g.sources); i += 2 {
				// fmt.Println("checking if", nextValue, "is between", g.sources[i], "and", g.sources[i]+g.sources[i+1]-1)
				if nextValue >= g.sources[i] && nextValue <= g.sources[i]+g.sources[i+1]-1 {
					index := nextValue - g.sources[i]
					// fmt.Println("index", index)
					nextValue = g.destinations[i] + index
					// fmt.Println("it is, the next value is", nextValue)
					break
				}
			}
			fmt.Printf(" %s %d,", g.to, nextValue)

			// break
		}
		fmt.Println()
		fmt.Println("final location", nextValue)
		if nextValue < location {
			location = nextValue
		}
		// break
	}
	fmt.Println("final location", location)

}

func solve2(lines []string) {
	s := strings.Split(lines[0], " ")[1:]
	seeds := make([]int, len(s))
	for i, v := range s {
		seeds[i], _ = strconv.Atoi(v)
	}
	fmt.Println("seeds", seeds)

	// return
	garden := make([]GardenThing, 0)
	g := GardenThing{}
	for _, l := range lines[2:] {
		if l == "" {
			fmt.Println()
			g.fillGaps()
			fmt.Println(g)
			garden = append(garden, g)
			g = GardenThing{}
			continue
		}
		data := strings.Split(l, " ")
		if _, err := strconv.Atoi(data[0]); err != nil {
			kind := strings.Split(data[0], "-to-")
			from := kind[0]
			to := kind[1]
			g = GardenThing{from: from, to: to}
			// fmt.Println("from", from, "to", to)
		} else {
			values := make([]int, len(data))
			for i, v := range data {
				values[i], _ = strconv.Atoi(v)
			}
			if g.sources == nil {
				g.sources = make([]int, 0)
			}
			if g.destinations == nil {
				g.destinations = make([]int, 0)
			}

			g.sources = append(g.sources, values[1], values[2])
			g.destinations = append(g.destinations, values[0], values[2])

			// for i := 0; i < values[2]; i++ {
			// g.values = append(g.values, values[0]+i, values[1]+i)

			// }
			// fmt.Println("values", values)
		}
		// fmt.Println(l)
	}
	g.fillGaps()
	fmt.Println(g)
	garden = append(garden, g)
	g = GardenThing{}

	// gerarchy := []string{"seed", "soil", "fertilizer", "water", "light", "temperature", "humidity", "location"}
	// fmt.Println(garden["seed"])
	location := math.MaxInt64
	for seed := 0; seed < len(seeds); seed += 2 {
		for s := seeds[seed]; s <= seeds[seed]+seeds[seed+1]; s++ {
			nextValue := s
			// fmt.Printf("seed %d,", s)
			for _, g := range garden {
				for i := 0; i < len(g.sources); i += 2 {
					// fmt.Println("checking if", nextValue, "is between", g.sources[i], "and", g.sources[i]+g.sources[i+1]-1)
					if nextValue >= g.sources[i] && nextValue <= g.sources[i]+g.sources[i+1]-1 {
						index := nextValue - g.sources[i]
						// fmt.Println("index", index)
						nextValue = g.destinations[i] + index
						// fmt.Println("it is, the next value is", nextValue)
						break
					}
				}
				// fmt.Printf(" %s %d,", g.to, nextValue)

				// break
			}
			// fmt.Println()
			// fmt.Println("final location", nextValue)
			if nextValue < location {
				location = nextValue
			}
			// break
		}
	}
	fmt.Println("final location", location)

}

/*
{seed soil [50 2 52 48 0 50] [98 2 50 48 0 50]}
{soil fertilizer [0 37 37 2 39 15] [15 37 52 2 0 15]}
{fertilizer water [49 8 0 42 42 7 57 4 0 49 49 8] [53 8 11 42 0 7 7 4 0 49 49 8]}
{water light [88 7 18 70 0 88] [18 7 25 70 0 88]}
{light temperature [45 23 81 19 68 13 0 45 68 13] [77 23 45 19 64 13 0 45 68 13]}
{temperature humidity [0 1 1 69] [69 1 0 69]}
{humidity location [60 37 56 4 0 60] [56 37 93 4 0 60]}
*/

func main() {
	lines := readLines("input.txt")
	solve2(lines)
}
