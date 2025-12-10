package main

import (
	"fmt"
	"math"
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

// Time taken: 26.26µs
func sol1(input []rotation) int {
	positions := 50
	zeroCounter := 0
	// fmt.Printf("The dial starts by pointing at %d\n", positions)
	for _, rot := range input {
		rot.steps %= 100
		if rot.direction == "L" {
			if (positions - rot.steps) < 0 {
				diff := math.Abs(float64(rot.steps - positions))
				positions = 100 - int(diff)
			} else {
				positions -= rot.steps
			}
		} else {
			if (positions + rot.steps) > 99 {
				diff := math.Abs(float64(positions + rot.steps - 99))
				positions = int(diff) - 1
			} else {
				positions += rot.steps
			}
		}
		if positions == 0 {
			zeroCounter++
		}

		// fmt.Printf("The dial is rotated %s%d to point at %d\n", rot.direction, rot.steps, positions)
	}
	return zeroCounter
}

// Time taken: 505.294µs
func sol2(input []rotation) int {
	positions := 50
	zeroCounter := 0
	for _, rot := range input {
		if rot.direction == "L" {
			for i := 0; i < rot.steps; i++ {
				positions--
				if positions < 0 {
					positions = 99
				}
				if positions == 0 {
					zeroCounter++
				}
			}
		} else {
			for i := 0; i < rot.steps; i++ {
				positions++
				if positions > 99 {
					positions = 0
				}
				if positions == 0 {
					zeroCounter++
				}
			}
		}
	}
	return zeroCounter
}

// negative zeros: 3540 positive zeros: 3318
// negative zeros: 2923 positive zeros: 3318

func sol2V2(input []rotation) int {
	positions := 50
	zeroCounter := 0

	for _, rot := range input {
		crosses := rot.steps / 100
		if crosses > 0 {
			rot.steps %= 100
			zeroCounter += crosses
		}
		if rot.direction == "L" {
			if (positions - rot.steps) <= 0 {
				if positions != 0 {
					zeroCounter += 1
				}
				positions = (positions - rot.steps) + 100
				if positions == 100 {
					positions = 0
				}
			} else {
				positions -= rot.steps
			}
		} else {
			if (positions + rot.steps) >= 100 {
				zeroCounter += 1
				positions = (positions + rot.steps) - 100
			} else {
				positions += rot.steps
			}
		}
	}

	return zeroCounter
}

// > 6835 && < 8026
// not 7443
func main() {
	input := readInput("test.txt")
	startSol1 := time.Now()
	zeroCount := sol1(input)
	elapsed := time.Since(startSol1)
	fmt.Println("Part 1:")
	fmt.Println("solution:", zeroCount)
	fmt.Println("Time taken:", elapsed)
	startSol2 := time.Now()
	finalPosition := sol2V2(input)
	elapsed2 := time.Since(startSol2)
	fmt.Println("Part 2:")
	fmt.Println("Solution:", finalPosition)
	fmt.Println("Time taken:", elapsed2)
	fmt.Println("Total time:", elapsed+elapsed2)
}
