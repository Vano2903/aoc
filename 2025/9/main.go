package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type coord struct {
	row int
	col int
}

const (
	right = 0
	down  = 1
	left  = 2
	up    = 3
)

func readInput(path string) []coord {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\n")
	boxes := make([]coord, len(lines))
	for i, line := range lines {
		numbers := strings.Split(line, ",")
		x, _ := strconv.Atoi(numbers[0])
		y, _ := strconv.Atoi(numbers[1])
		boxes[i] = coord{
			col: x,
			row: y,
		}
	}
	return boxes
}

func calculateArea(coord1, coord2 coord) float64 {
	return (math.Abs(float64(coord1.col)-float64(coord2.col)) + 1) * (math.Abs(float64(coord1.row)-float64(coord2.row)) + 1)
}

func sol1(input []coord) int {
	solution := 0

	for i := range input {
		for j := i + 1; j < len(input); j++ {
			dist := calculateArea(input[i], input[j])
			solution = max(solution, int(dist))
		}
	}

	return solution
}

func sol2(input []coord) int {
	solution := 0

	for i := range input {
		for j := i + 1; j < len(input); j++ {
			A := input[i]
			C := input[j]

			if A.row > C.row {
				A, C = C, A
			}

			if A.col > C.col {
				A, C = coord{row: A.row, col: C.col}, coord{row: C.row, col: A.col}
			}

			B := coord{col: C.col, row: A.row}
			D := coord{col: A.col, row: C.row}

			if raycastIntersectionsCount(input, B, A, left) == 0 &&
				raycastIntersectionsCount(input, B, C, down) == 0 &&
				raycastIntersectionsCount(input, D, C, right) == 0 &&
				raycastIntersectionsCount(input, D, A, up) == 0 {
				area := calculateArea(A, C)
				solution = max(solution, int(area))
			}
		}
	}

	return solution
}

// direction: 0=right, 1=down, 2=left, 3=up
func raycastIntersectionsCount(input []coord, start, end coord, direction int) int {
	intersections := 0

	switch direction {
	case right:
		for i := range input {
			v1 := input[i]
			v2 := input[(i+1)%len(input)]

			if v2.row < v1.row {
				v1, v2 = v2, v1
			}
			if v1.col == v2.col && v1.col > start.col && v1.col < end.col {
				if start.row > v1.row && start.row <= v2.row {
					intersections++
				}
			}
		}

	case down:
		for i := range input {
			v1 := input[i]
			v2 := input[(i+1)%len(input)]

			if v2.col < v1.col {
				v1, v2 = v2, v1
			}
			if v1.row == v2.row && v1.row > start.row && v1.row < end.row {
				if start.col > v1.col && start.col <= v2.col {
					intersections++
				}
			}
		}

	case left:
		for i := range input {
			v1 := input[i]
			v2 := input[(i+1)%len(input)]

			if v2.row < v1.row {
				v1, v2 = v2, v1
			}
			if v1.col == v2.col && v1.col < start.col && v1.col > end.col {
				if start.row >= v1.row && start.row < v2.row {
					intersections++
				}
			}
		}
	case up:
		for i := range input {
			v1 := input[i]
			v2 := input[(i+1)%len(input)]

			if v2.col < v1.col {
				v1, v2 = v2, v1
			}
			if v1.row == v2.row && v1.row < start.row && v1.row > end.row {
				if start.col >= v1.col && start.col < v2.col {
					intersections++
				}
			}
		}
	}

	return intersections
}

func sol2V2(input []coord) int {
	solution := 0

	for i := range input {
		for j := i + 1; j < len(input); j++ {
			A := input[i]
			C := input[j]

			if A.row > C.row {
				A, C = C, A
			}

			if A.col > C.col {
				A, C = coord{row: A.row, col: C.col}, coord{row: C.row, col: A.col}
			}

			B := coord{col: C.col, row: A.row}
			D := coord{col: A.col, row: C.row}

			if !doesIntersect(input, B, A, left) &&
				!doesIntersect(input, B, C, down) &&
				!doesIntersect(input, D, C, right) &&
				!doesIntersect(input, D, A, up) {
				area := calculateArea(A, C)
				solution = max(solution, int(area))
			}
		}
	}

	return solution
}

// direction: 0=right, 1=down, 2=left, 3=up
func doesIntersect(input []coord, start, end coord, direction int) bool {
	switch direction {
	case right:
		for i := range input {
			v1 := input[i]
			v2 := input[(i+1)%len(input)]

			if v2.row < v1.row {
				v1, v2 = v2, v1
			}
			if v1.col == v2.col && v1.col > start.col && v1.col < end.col {
				if start.row > v1.row && start.row <= v2.row {
					return true
				}
			}
		}

	case down:
		for i := range input {
			v1 := input[i]
			v2 := input[(i+1)%len(input)]

			if v2.col < v1.col {
				v1, v2 = v2, v1
			}
			if v1.row == v2.row && v1.row > start.row && v1.row < end.row {
				if start.col > v1.col && start.col <= v2.col {
					return true
				}
			}
		}

	case left:
		for i := range input {
			v1 := input[i]
			v2 := input[(i+1)%len(input)]

			if v2.row < v1.row {
				v1, v2 = v2, v1
			}
			if v1.col == v2.col && v1.col < start.col && v1.col > end.col {
				if start.row >= v1.row && start.row < v2.row {
					return true
				}
			}
		}
	case up:
		for i := range input {
			v1 := input[i]
			v2 := input[(i+1)%len(input)]

			if v2.col < v1.col {
				v1, v2 = v2, v1
			}
			if v1.row == v2.row && v1.row < start.row && v1.row > end.row {
				if start.col >= v1.col && start.col < v2.col {
					return true
				}
			}
		}
	}

	return false
}

func main() {
	input := readInput("mora.txt")
	startSol1 := time.Now()
	solution1 := sol1(input)
	elapsed := time.Since(startSol1)
	fmt.Println("Part 1:")
	fmt.Println("solution:", solution1)
	fmt.Println("Time taken:", elapsed)
	startSol2 := time.Now()
	solution2 := sol2V2(input)
	elapsed2 := time.Since(startSol2)
	if solution2 != 1465767840 {
		panic("wrong solution")
	}
	fmt.Println("Part 2:")
	fmt.Println("Solution:", solution2)
	fmt.Println("Time taken:", elapsed2)
	fmt.Println("Total time:", elapsed+elapsed2)
}
