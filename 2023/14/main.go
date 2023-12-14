package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func readFile(path string) [][]string {
	c, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	m := make([][]string, 0)
	for _, line := range strings.Split(string(c), "\n") {
		m = append(m, strings.Split(line, ""))
	}
	return m
}

func printMap(m [][]string) {
	for _, line := range m {
		for _, c := range line {
			fmt.Print(c)
		}
		fmt.Println()
	}
}

//=====================

func tiltEast(m [][]string) [][]string {
	l := len(m[0])
	// for y, line := range m {
	for y := l - 1; y >= 0; y-- {
		for x := len(m[y]) - 1; x >= 0; x-- {
			// for x, c := range m[y] {
			c := m[y][x]
			if c != "O" {
				continue
			}
			j := x

			for {
				if j+1 >= l || m[y][j+1] != "." {
					break
				}
				j++
			}
			if j == x {
				continue
			}
			m[y][j] = "O"
			m[y][x] = "."
		}
	}
	return m
}

func tiltWest(m [][]string) [][]string {
	for y, line := range m {
		for x, c := range line {
			if c != "O" {
				continue
			}
			j := x
			for {
				if j-1 < 0 || m[y][j-1] != "." {
					break
				}
				j--
			}
			if j == x {
				continue
			}
			m[y][j] = "O"
			m[y][x] = "."

		}
	}
	return m
}

func tiltSouth(m [][]string) [][]string {
	l := len(m)
	for y := l - 1; y >= 0; y-- {
		// for y, line := range m {
		for x, c := range m[y] {
			// for x, c := range line {
			if c != "O" {
				continue
			}
			j := y
			for {
				if j+1 >= l || m[j+1][x] != "." {
					break
				}
				j++
			}
			if j == y {
				continue
			}
			m[j][x] = "O"
			m[y][x] = "."
		}
	}
	return m
}

func tiltNorth(m [][]string) [][]string {
	for y, line := range m {
		for x, c := range line {
			if c != "O" {
				continue
			}
			// fmt.Println("checking north")
			j := y
			for {
				if j-1 < 0 || m[j-1][x] != "." {
					break
				}
				j--
			}
			// fmt.Println("the ball", y, x, "can go up to", y, j)
			if j == y {
				continue
			}
			m[j][x] = "O"
			m[y][x] = "."
			// printMap(m)
		}
	}
	return m
}

func calculateLoad(m [][]string) int {
	load := 0
	for i, line := range m {
		for _, c := range line {
			if c == "O" {
				load += len(m) - i
			}
		}
	}
	return load
}

func mapToString(m [][]string) string {
	s := ""
	for _, line := range m {
		s += strings.Join(line, "")
	}
	return s
}

func solve2(m [][]string) {
	type state struct {
		score      int
		cycleCount int
	}
	var key string
	cycleDetector := make(map[string]state)
	// cycleCounter := 1
	iterations := 1_000_000_000
	for i := 0; i < iterations; i++ {
		m = tiltNorth(m) //north
		m = tiltWest(m)  //west
		m = tiltSouth(m) //south
		m = tiltEast(m)  //east

		key = mapToString(m)
		if current, ok := cycleDetector[key]; ok {
			l := (i + 1) - current.cycleCount
			// fmt.Println("cycle detected")
			// fmt.Println("cycle count:", current.cycleCount)
			// fmt.Println("cycle score:", current.score)
			// fmt.Println("cycle length:", l)
			for _, s := range cycleDetector {
				if s.cycleCount >= current.cycleCount {
					if s.cycleCount%l == iterations%l {
						fmt.Println("solution 2:", s.score)
						return
					}
				}
			}
			fmt.Println("i dont think you are here either")
			break
		}
		cycleDetector[key] = state{score: calculateLoad(m), cycleCount: i + 1}
	}
	fmt.Println("i doubt you are here")
	fmt.Println("load:", calculateLoad(m))
}

func main() {
	m := readFile("input")
	now := time.Now()
	// solve1(m)
	fmt.Println("solution 1:", calculateLoad(tiltNorth(m)))
	fmt.Println("time:", time.Since(now))
	m = readFile("input")
	now = time.Now()
	solve2(m)
	fmt.Println("time:", time.Since(now))
}
