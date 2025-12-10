package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func readInput(path string) [][]string {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\n")
	rows := make([][]string, len(lines))
	for i, line := range lines {
		rows[i] = make([]string, len(line))
		for j, char := range strings.Split(line, "") {
			rows[i][j] = char
		}
	}
	return rows
}

func printGrid(grid [][]string) {
	for _, row := range grid {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}
}

func sol1(input [][]string) int {
	solution := 0

	for i, row := range input {
		for j := range row {
			counter := 0
			if input[i][j] == "." {
				continue
			}

			if i-1 >= 0 && j-1 >= 0 && input[i-1][j-1] != "." {
				counter++
			}
			if i-1 >= 0 && input[i-1][j] != "." {
				counter++
			}
			if i-1 >= 0 && j+1 < len(row) && input[i-1][j+1] != "." {
				counter++
			}

			if j-1 >= 0 && input[i][j-1] != "." {
				counter++
			}
			if j+1 < len(row) && input[i][j+1] != "." {
				counter++
			}

			if i+1 < len(input) && j-1 >= 0 && input[i+1][j-1] != "." {
				counter++
			}
			if i+1 < len(input) && input[i+1][j] != "." {
				counter++
			}
			if i+1 < len(input) && j+1 < len(row) && input[i+1][j+1] != "." {
				counter++
			}

			if counter <= 3 {
				solution++
			}
		}
	}

	return solution
}

func sol2(input [][]string) int {
	solution := 0
	prevSolution := -1
	for {
		for i, row := range input {
			for j := range row {
				counter := 0
				if input[i][j] == "." {
					continue
				}
				if i-1 >= 0 && j-1 >= 0 && input[i-1][j-1] == "@" {
					counter++
				}
				if i-1 >= 0 && input[i-1][j] == "@" {
					counter++
				}
				if i-1 >= 0 && j+1 < len(row) && input[i-1][j+1] == "@" {
					counter++
				}

				if j-1 >= 0 && input[i][j-1] == "@" {
					counter++
				}
				if j+1 < len(row) && input[i][j+1] == "@" {
					counter++
				}

				if i+1 < len(input) && j-1 >= 0 && input[i+1][j-1] == "@" {
					counter++
				}
				if i+1 < len(input) && input[i+1][j] == "@" {
					counter++
				}
				if i+1 < len(input) && j+1 < len(row) && input[i+1][j+1] == "@" {
					counter++
				}

				if counter <= 3 {
					solution++
					input[i][j] = "."
				}
			}
		}
		if solution == prevSolution {
			break
		}
		prevSolution = solution
	}

	return solution
}

func main() {
	input := readInput("input.txt")
	startSol1 := time.Now()
	solution1 := sol1(input)
	elapsed := time.Since(startSol1)
	fmt.Println("Part 1:")
	fmt.Println("solution:", solution1)
	fmt.Println("Time taken:", elapsed)
	startSol2 := time.Now()
	solution2 := sol2(input)
	elapsed2 := time.Since(startSol2)
	fmt.Println("Part 2:")
	fmt.Println("Solution:", solution2)
	fmt.Println("Time taken:", elapsed2)
	fmt.Println("Total time:", elapsed+elapsed2)
}
