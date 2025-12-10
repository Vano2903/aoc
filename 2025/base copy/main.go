package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type rotation struct {
	direction string // "L" or "R"
	steps     int
}

func readInput(path string) []rotation {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\n")
	rotations := make([]rotation, len(lines))
	for i, line := range lines {
		dir := string(line[0])
		var steps int
		fmt.Sscanf(line[1:], "%d", &steps)
		rotations[i] = rotation{direction: dir, steps: steps}
	}
	return rotations
}

func sol1(input []rotation) int {
	solution := 0

	return solution
}

func sol2(input []rotation) int {
	solution := 0

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
