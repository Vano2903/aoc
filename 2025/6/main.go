package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func readInput(path string) []string {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\n")
	return lines
}

func sol1(input []string) int {
	solution := 0
	var cols [][]int
	for _, line := range input {
		elements := strings.Split(line, " ")
		counter := 0
		if strings.Contains(line, "*") {
			for _, el := range elements {
				if el == "" {
					continue
				}
				if el == "*" {
					// fmt.Println("multiplying cols:", cols[counter])
					result := 1
					for _, val := range cols[counter] {
						result *= val
					}
					// fmt.Println("intermediate solution:", result)
					solution += result
				} else {
					// fmt.Println("adding cols:", cols[counter])
					result := 0
					for _, val := range cols[counter] {
						result += val
					}
					// fmt.Println("intermediate solution:", result)
					solution += result
				}
				counter++

			}

		} else {
			for _, el := range elements {
				if el == "" {
					continue
				}

				val, err := strconv.Atoi(el)
				if err != nil {
					panic(err)
				}

				if counter >= len(cols) {
					cols = append(cols, []int{})
				}
				cols[counter] = append(cols[counter], val)
				counter++
			}
		}
	}

	return solution
}

func sol2(input []string) int {
	solution := 0

	var startOfEachOperation []int
	for _, line := range input[len(input)-1:] {
		for i, el := range line {
			if el != ' ' {
				startOfEachOperation = append(startOfEachOperation, i)
			}
		}
	}

	height := len(input) - 1
	numbers := input[:len(input)-1]
	operations := strings.Split(input[len(input)-1:][0], "")
	currentOperationCounter := 0
	for i := 0; i < len(startOfEachOperation); i++ {
		start := startOfEachOperation[i]
		var end int
		if i == len(startOfEachOperation)-1 {
			end = len(numbers[0])
		} else {
			end = startOfEachOperation[i+1] - 1
		}
		var currentNumbers []int
		for j := start; j < end; j++ {
			currentNumber := ""
			for h := 0; h < height; h++ {
				if numbers[h][j] != ' ' {
					currentNumber += string(numbers[h][j])
				}
			}
			val, _ := strconv.Atoi(currentNumber)
			currentNumbers = append(currentNumbers, val)
		}

		var result int
		if operations[startOfEachOperation[currentOperationCounter]] == "*" {
			result = 1
			for _, num := range currentNumbers {
				result *= num
			}
		} else {
			for _, num := range currentNumbers {
				result += num
			}
		}
		solution += result
		currentOperationCounter++
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
