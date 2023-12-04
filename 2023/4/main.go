package main

import (
	"fmt"
	"os"
	"strings"
)

func readLines(path string) []string {
	c, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(c), "\n")
}

func parseLine(line string) ([]string, []string) {
	parts := strings.Split(line, "|")
	return strings.Split(parts[0], " ")[2:], strings.Split(parts[1], " ")
}

func count(list []string, item string) int {
	counter := 0
	for _, s := range list {
		if s == item {
			counter++
		}
	}
	return counter
}

func solve1(lines []string) {
	totalPoints := 0
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		winningsToClean, foundsToClean := parseLine(line)
		winnings := []string{}
		founds := []string{}
		for _, f := range foundsToClean {
			if f == "" {
				continue
			}
			founds = append(founds, f)
		}
		for _, w := range winningsToClean {
			if w == "" {
				continue
			}
			winnings = append(winnings, w)
		}
		// fmt.Println("Winnings:", winnings)
		// fmt.Println("Founds:", founds)
		points := 0
		first := true
		for _, f := range founds {
			counts := count(winnings, f)
			// fmt.Println("Found", f, "times", counts)
			if first && counts > 0 {
				points = 1
				first = false
			} else {

				if counts > 0 {
					// points += 1
					// counts--
					for i := 0; i < counts; i++ {
						points *= 2
					}
					// fmt.Println("Points:", points)
				}
			}
		}
		fmt.Println("Card ", i+1, ":", points)
		totalPoints += points
	}
	fmt.Println("Total points:", totalPoints)
}

type c struct {
	copies int
	value  int
}

func isIn(list []string, item string) bool {
	for _, s := range list {
		if s == item {
			return true
		}
	}
	return false
}

func solve2(lines []string) {
	cards := make([]c, len(lines))
	for i, line := range lines {
		cards[i].copies += 1
		if len(line) == 0 {
			continue
		}
		winningsToClean, foundsToClean := parseLine(line)
		winnings := []string{}
		founds := []string{}
		for _, f := range foundsToClean {
			if f == "" {
				continue
			}
			founds = append(founds, f)
		}
		for _, w := range winningsToClean {
			if w == "" {
				continue
			}
			winnings = append(winnings, w)
		}
		// fmt.Println("Winnings:", winnings)
		// fmt.Println("Founds:", founds)
		// poin
		counts := 0
		for _, f := range founds {
			// counts := count(winnings, f)
			if isIn(winnings, f) {
				counts++
			}
			// fmt.Println("Found", f, "times", counts)
			// cards[i].value = counts
		}
		// fmt.Println("Card", i+1, ":", counts, ", there are", cards[i].copies, "copies")
		for j := 1; j < counts+1; j++ {
			for k := 0; k < cards[i].copies; k++ {
				cards[i+j].copies += 1
			}
			// fmt.Println("incrementing", j)
			// cards[i+j].copies += cards[i+j].copies + cards[i].copies
		}
		// fmt.Println(cards)
	}
	// for _, card := range cards {
	// }
	points := 0
	for _, card := range cards {
		points += card.copies
	}
	fmt.Println("Total points:", points)

}

func main() {
	lines := readLines("input.txt")
	// solve1(lines)
	solve2(lines)
}
