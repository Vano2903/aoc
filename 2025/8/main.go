package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Box struct {
	x       int
	y       int
	z       int
	circuit int
}

func readInput(path string) []*Box {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\n")
	boxes := make([]*Box, len(lines))
	for i, line := range lines {
		numbers := strings.Split(line, ",")
		x, _ := strconv.Atoi(numbers[0])
		y, _ := strconv.Atoi(numbers[1])
		z, _ := strconv.Atoi(numbers[2])
		boxes[i] = &Box{
			x: x,
			y: y,
			z: z,
		}
	}
	return boxes
}

func distanceBetweenBoxes(box1, box2 *Box) float64 {
	return math.Sqrt(math.Pow(float64(box1.x)-float64(box2.x), 2) +
		math.Pow(float64(box1.y)-float64(box2.y), 2) +
		math.Pow(float64(box1.z)-float64(box2.z), 2))
}

type FullDistance struct {
	from     *Box
	to       *Box
	distance float64
}

func sol1(input []*Box) int {
	solution := 0

	totalDistances := make([]FullDistance, 0)
	for i := range input {
		for j := i + 1; j < len(input); j++ {
			dist := distanceBetweenBoxes(input[i], input[j])
			totalDistances = append(totalDistances, FullDistance{
				from:     input[i],
				to:       input[j],
				distance: dist,
			})
		}

	}
	slices.SortFunc(totalDistances, func(a, b FullDistance) int {
		return int(a.distance - b.distance)
	})

	circuits := make([][]*Box, len(input))
	for i, box := range input {
		box.circuit = i
		circuit := []*Box{box}
		circuits[i] = circuit
	}

	for i, distance := range totalDistances {
		box := distance.from
		closestBox := distance.to

		if box.circuit == closestBox.circuit {
			continue
		}

		circuit1 := circuits[box.circuit]
		circuit2 := circuits[closestBox.circuit]
		prevCircuit := closestBox.circuit
		for _, b := range circuit2 {
			b.circuit = box.circuit
			circuit1 = append(circuit1, b)
		}
		circuits[box.circuit] = circuit1
		circuits[prevCircuit] = []*Box{}

		if i == 1000-1 {
			break
		}
	}

	slices.SortFunc(circuits, func(a, b []*Box) int {
		return len(b) - len(a)
	})
	solution = len(circuits[0]) * len(circuits[1]) * len(circuits[2])

	return solution
}

func sol2(input []*Box) int {
	solution := 0

	totalDistances := make([]FullDistance, 0)
	for i := range input {
		for j := i + 1; j < len(input); j++ {
			dist := distanceBetweenBoxes(input[i], input[j])

			totalDistances = append(totalDistances, FullDistance{
				from:     input[i],
				to:       input[j],
				distance: dist,
			})
		}
	}
	slices.SortFunc(totalDistances, func(a, b FullDistance) int {
		return int(a.distance - b.distance)
	})

	circuits := make([][]*Box, len(input))
	for i, box := range input {
		box.circuit = i
		circuit := []*Box{box}
		circuits[i] = circuit
	}

	lastX := -1
	secondLastX := -1
	for _, distance := range totalDistances {
		box := distance.from
		closestBox := distance.to

		if box.circuit == closestBox.circuit {
			continue
		}
		lastX = box.x
		secondLastX = closestBox.x
		circuit1 := circuits[box.circuit]
		circuit2 := circuits[closestBox.circuit]
		prevCircuit := closestBox.circuit
		for _, b := range circuit2 {
			b.circuit = box.circuit
			circuit1 = append(circuit1, b)
		}
		circuits[box.circuit] = circuit1
		circuits[prevCircuit] = []*Box{}

		for circuit := range circuits {
			if len(circuits[circuit]) == len(input) {
				break
			}
		}
	}
	solution = lastX * secondLastX

	return solution
}

func main() {
	input := readInput("input.txt")
	startSol1 := time.Now()
	solution1 := sol1(input) // 96672
	elapsed := time.Since(startSol1)
	fmt.Println("Part 1:")
	fmt.Println("solution:", solution1)
	fmt.Println("Time taken:", elapsed)
	if solution1 != 96672 {
		panic("wrong solution")
	}
	startSol2 := time.Now()
	solution2 := sol2(input)
	elapsed2 := time.Since(startSol2)
	if solution2 != 22517595 {
		panic("wrong solution")
	}
	fmt.Println("Part 2:")
	fmt.Println("Solution:", solution2)
	fmt.Println("Time taken:", elapsed2)
	fmt.Println("Total time:", elapsed+elapsed2)
}
