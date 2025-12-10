package main

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strconv"
	"strings"
	"time"
)

func readInput(path string) [][]string {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\n")
	var grid [][]string
	for _, line := range lines {
		row := strings.Split(line, "")
		grid = append(grid, row)
	}
	return grid
}

func printGrid(grid [][]string) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func sol1(input [][]string) int {
	solution := 0
	start := slices.Index(input[0], "S")
	if start == -1 {
		panic("no start found")
	}

	input[0][start] = "|"
	for i := 1; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i-1][j] == "." {
				continue
			}
			if input[i-1][j] == "|" {
				if input[i][j] == "^" {
					solution++
					if j > 0 {
						input[i][j-1] = "|"
					}
					if j < len(input[i])-1 {
						input[i][j+1] = "|"
					}
				} else {
					input[i][j] = "|"
				}
			}
		}
		// printGrid(input)
		// time.Sleep(100 * time.Millisecond)
		// cmd := exec.Command("clear") //Linux example, its tested
		// cmd.Stdout = os.Stdout
		// cmd.Run()
	}

	return solution
}

func sol2(input [][]string) int {
	solution := 0
	start := slices.Index(input[0], "S")
	if start == -1 {
		panic("no start found")
	}

	input[0][start] = "1"
	for i := 1; i < len(input); i++ {
		time.Sleep(100 * time.Millisecond)
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
		printGrid(input)
		for j := 0; j < len(input[i]); j++ {
			if input[i-1][j] == "." {
				continue
			}
			if input[i-1][j] != "^" && input[i-1][j] != "." {
				before, _ := strconv.Atoi(input[i-1][j])
				if input[i][j] == "^" {
					// solution++
					if j > 0 {
						left, err := strconv.Atoi(input[i][j-1])
						if err != nil {
							left = 0
						}
						input[i][j-1] = strconv.Itoa(before + left)
					}
					if j < len(input[i])-1 {
						right, err := strconv.Atoi(input[i][j+1])
						if err != nil {
							right = 0
						}
						input[i][j+1] = strconv.Itoa(before + right)
					}
				} else {
					current, err := strconv.Atoi(input[i][j])
					if err != nil {
						current = 0
					}
					input[i][j] = strconv.Itoa(before + current)
				}
			}
		}
	}

	for j := 0; j < len(input[len(input)-1]); j++ {
		if input[len(input)-1][j] != "." && input[len(input)-1][j] != "^" {
			val, err := strconv.Atoi(input[len(input)-1][j])
			if err != nil {
				panic(err)
			}
			solution += val
		}
	}

	return solution
}

func main() {
	input := readInput("input-test.txt")
	duplicate := make([][]string, len(input))
	for i := range input {
		duplicate[i] = make([]string, len(input[i]))
		copy(duplicate[i], input[i])
	}
	// printGrid(input)
	startSol1 := time.Now()
	solution1 := sol1(input)
	elapsed := time.Since(startSol1)
	fmt.Println("Part 1:")
	fmt.Println("solution:", solution1)
	fmt.Println("Time taken:", elapsed)
	startSol2 := time.Now()
	solution2 := sol2(duplicate)
	elapsed2 := time.Since(startSol2)
	fmt.Println("Part 2:")
	fmt.Println("Solution:", solution2)
	fmt.Println("Time taken:", elapsed2)
	fmt.Println("Total time:", elapsed+elapsed2)
}
