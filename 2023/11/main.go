package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

func readFile(path string) [][]string {
	c, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(c), "\n")
	matrix := make([][]string, len(lines))
	for i, line := range lines {
		matrix[i] = strings.Split(line, "")
	}
	return matrix
}

type Point struct {
	row, col int
}


func insertSlice(a [][]string, index int, value []string) [][]string {
	if len(a) == index {
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...)
	a[index] = value
	return a
}

func insert(a []string, index int, value string) []string {
	if len(a) == index {
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...)
	a[index] = value
	return a
}

func expandUniverse(universe [][]string) [][]string {
	for y := 0; y < len(universe[0]); y++ {
		hasStar := false
		for x := 0; x < len(universe); x++ {
			if universe[x][y] != "." {
				hasStar = true
				break
			}
		}
		if !hasStar {
			for x := 0; x < len(universe); x++ {
				universe[x] = insert(universe[x], y, ".")
			}
			y++
		}
	}

	for x := 0; x < len(universe); x++ {
		hasStar := false
		for _, c := range universe[x] {
			if c != "." {
				hasStar = true
				break
			}
		}
		if !hasStar {
			fmt.Println("add line in line: ", x)
			line := make([]string, len(universe[0]))
			for i := range line {
				line[i] = "."
			}
			universe = insertSlice(universe, x, line)
			x++
		}
	}
	return universe
}

func solve1(matrix [][]string) {
	matrix = expandUniverse(matrix)
	stars := []Point{}
	for row, line := range matrix {
		for col, c := range line {
			if c != "." {
				stars = append(stars, Point{row: row, col: col})
			}
		}
	}

	pairs := []Point{}
	for i, star := range stars {
		for j := i + 1; j < len(stars); j++ {
			pairs = append(pairs, Point{row: star.row, col: star.col}, Point{row: stars[j].row, col: stars[j].col})
		}
	}
	fmt.Println(len(pairs))

	total := 0
	for i := 0; i < len(pairs)-1; i += 2 {
		start := pairs[i]
		end := pairs[i+1]
		distance := math.Abs(float64(end.col-start.col)) + math.Abs(float64(end.row-start.row))
		total += int(distance)
	}
	fmt.Println("Total:", total)
}

func solve2(matrix [][]string) {
	toIncrease := 1000000
	toIncrease--

	starsBefore := []Point{}
	for row, line := range matrix {
		for col, c := range line {
			if c != "." {
				starsBefore = append(starsBefore, Point{row: row, col: col})
			}
		}
	}
	stars := make([]Point, len(starsBefore))
	copy(stars, starsBefore)

	for y := 0; y < len(matrix[0]); y++ {
		hasStar := false
		for x := 0; x < len(matrix); x++ {
			if matrix[x][y] != "." {
				hasStar = true
				break
			}
		}
		if !hasStar {
			shifted := 0
			for i := range starsBefore {
				if starsBefore[i].col > y {
					shifted++
					stars[i].col += toIncrease
				}
			}
		}
	}
	fmt.Println()

	for x := 0; x < len(matrix); x++ {
		hasStar := false
		for _, c := range matrix[x] {
			if c != "." {
				hasStar = true
				break
			}
		}
		if !hasStar {
			shifted := 0
			for i := range starsBefore {
				if starsBefore[i].row > x {
					shifted++
					stars[i].row += toIncrease
				}
			}
		}
	}

	pairs := []Point{}
	for i, star := range stars {
		for j := i + 1; j < len(stars); j++ {
			pairs = append(pairs,
				Point{row: star.row, col: star.col},
				Point{row: stars[j].row, col: stars[j].col})
		}
	}

	total := 0
	for i := 0; i < len(pairs)-1; i += 2 {
		start := pairs[i]
		end := pairs[i+1]
		distance := math.Abs(float64(end.col-start.col)) + math.Abs(float64(end.row-start.row))
		total += int(distance)
	}
	fmt.Println("Total:", total)
}

func main() {
	file := "input.txt"
	matrix := readFile(file)
	m2 := readFile(file)
	now := time.Now()
	solve1(matrix)
	fmt.Println("Time:", time.Since(now))
	now = time.Now()
	solve2(m2)
	fmt.Println("Time:", time.Since(now))
}
