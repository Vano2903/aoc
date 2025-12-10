package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func readInput(path string) [][]int {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\n")
	lineNumbers := make([][]int, len(lines))
	for i, line := range lines {
		numbers := strings.Split(line, "")
		lineNumbers[i] = make([]int, len(numbers))
		for j, num := range numbers {
			lineNumbers[i][j], err = strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
		}
	}
	return lineNumbers
}

// func findBiggestConcatWithBiggestAndPointer(input []int, currentBiggest int, start int, end int, memory map[int]int) int {
// 	current := fmt.Sprintf("%d%d", input[start], input[end])
// }

// func findBiggestConcat(input []int) int {
// 	currentBiggest := -1
// 	start := 0
// 	memory := make(map[int]int)
// 	return findBiggestConcatWithBiggestAndPointer(input, currentBiggest, start, len(input)-1, memory)
// }

func findTwoBiggestIndexes(input []int) (int, int) {
	biggestIndex := -1
	secondBiggestIndex := -1
	biggestValue := -1
	secondBiggestValue := -1
	for i, val := range input {
		if val > biggestValue {
			secondBiggestValue = biggestValue
			secondBiggestIndex = biggestIndex
			biggestValue = val
			biggestIndex = i
		} else if val > secondBiggestValue {
			secondBiggestValue = val
			secondBiggestIndex = i
		}
	}
	return biggestIndex, secondBiggestIndex
}

func sol1(input [][]int) int {
	solution := 0
	for _, line := range input {
		biggestIndex, secondBiggestIndex := findTwoBiggestIndexes(line)

		currentIndex := biggestIndex
		if currentIndex == len(line)-1 {
			currentIndex = secondBiggestIndex
		}
		val := -1
		for i := currentIndex + 1; i < len(line); i++ {
			if line[i] > val {
				val = line[i]
			}
		}
		solution += line[currentIndex]*10 + val
	}
	return solution
}

func mapNumberWithIndexes(input []int) [][]int {
	mapped := make([][]int, 9)
	for i, val := range input {
		mapped[val-1] = append(mapped[val-1], i)
	}
	return mapped
}

func findBiggestNumber(input []int, mappedNumbers [][]int, indexes []int, currentHighest, size, start int) int {
	if len(indexes) == size {
		result := 0
		for i, index := range indexes {
			result += input[index] * int(math.Pow10(size-1-i))
		}

		return result
	}
	if currentHighest <= 0 {
		return 0
	}
	largestIndexes := mappedNumbers[currentHighest-1]
	if len(largestIndexes) == 0 {
		return findBiggestNumber(input, mappedNumbers, indexes, currentHighest-1, size, start)
	}

	for _, index := range largestIndexes {
		if index >= start && index <= len(input)-1 {
			newIndexes := append(indexes, index)
			newStart := index + 1
			result := findBiggestNumber(input, mappedNumbers, newIndexes, 9, size, newStart)
			if result != 0 {
				return result
			}
		}
	}
	return findBiggestNumber(input, mappedNumbers, indexes, currentHighest-1, size, start)
}

func sol2(input [][]int) int {
	solution := 0
	for _, line := range input {
		mapped := mapNumberWithIndexes(line)
		solution += findBiggestNumber(line, mapped, make([]int, 0), 9, 12, 0)
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
