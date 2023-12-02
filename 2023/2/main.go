package main

import (
	"fmt"
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

func solve2(lines []string) {
	total := 0
	for _, line := range lines {
		r := 0
		g := 0
		b := 0
		games := strings.Split(strings.Split(line, ": ")[1], "; ")
		for _, game := range games {
			rounds := strings.Split(game, ", ")
			for _, round := range rounds {
				a := strings.Split(round, " ")
				n, err := strconv.Atoi(a[0])
				if err != nil {
					panic(err)
				}
				color := a[1]
				switch color {
				case "red":
					if r < n {
						r = n
					}
				case "green":
					if g < n {
						g = n
					}
				case "blue":
					if b < n {
						b = n
					}
				}
			}
		}
		total += r * g * b
		// fmt.Println("line:", line)
		// fmt.Println("r:", r, "g:", g, "b:", b)
	}
	fmt.Println("totale:", total)
}

func solve1(lines []string) {
	total := 0
	c := 1
	r := 12
	g := 13
	b := 14
	for _, line := range lines {
		games := strings.Split(strings.Split(line, ": ")[1], "; ")
		possible := true
		for _, game := range games {
			rounds := strings.Split(game, ", ")
			for _, round := range rounds {
				a := strings.Split(round, " ")
				n, err := strconv.Atoi(a[0])
				if err != nil {
					panic(err)
				}
				color := a[1]
				switch color {
				case "red":
					if r < n {
						possible = false
						// break
					}
				case "green":
					if g < n {
						possible = false
						// break
					}
				case "blue":
					if b < n {
						possible = false
						// break
					}
				}
				if !possible {
					break
				}
			}
		}
		if possible {
			// fmt.Println("line:", line)
			total += c
		}
		c++
	}
	fmt.Println("possibili sono:", total)
}

func main() {
	l := readLines("input.txt")
	solve2(l)
}
