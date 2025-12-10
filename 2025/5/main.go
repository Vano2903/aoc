package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func readInput(path string) ([][]int, []int) {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\n")
	var sets [][]int
	var ids []int
	for _, line := range lines {
		if line == "" {
			break
		}
		values := strings.Split(line, "-")
		start, _ := strconv.Atoi(values[0])
		end, _ := strconv.Atoi(values[1])
		sets = append(sets, []int{start, end})
	}
	for _, line := range lines[len(sets)+1:] {
		if line == "" {
			break
		}
		var id int
		fmt.Sscanf(line, "%d", &id)
		ids = append(ids, id)
	}
	return sets, ids
}

func mergeTwoSets(set1, set2 []int) ([]int, bool) {
	if set1[1] < set2[0]-1 || set2[1] < set1[0]-1 {
		return nil, false
	}
	newStart := min(set2[0], set1[0])
	newEnd := max(set2[1], set1[1])
	return []int{newStart, newEnd}, true
}

func mergeSets(sets [][]int) [][]int {
	if len(sets) == 0 {
		return sets
	}
	merged := true
	for merged {
		merged = false
		newSets := [][]int{}
		used := make([]bool, len(sets))
		for i := 0; i < len(sets); i++ {
			if used[i] {
				continue
			}
			currentSet := sets[i]
			for j := i + 1; j < len(sets); j++ {
				if used[j] {
					continue
				}
				if mergedSet, ok := mergeTwoSets(currentSet, sets[j]); ok {
					currentSet = mergedSet
					used[j] = true
					merged = true
				}
			}
			newSets = append(newSets, currentSet)
			used[i] = true
		}
		sets = newSets
	}
	return sets

}

func sol1(sets [][]int, ids []int) int {
	solution := 0
	for _, id := range ids {
		for _, set := range sets {
			start, end := set[0], set[1]
			if id >= start && id <= end {
				solution++
				break
			}
		}
	}
	return solution
}

func sol2(sets [][]int) int {
	solution := 0
	mergedSets := mergeSets(sets)
	for _, set := range mergedSets {
		start, end := set[0], set[1]
		solution += end - start + 1
	}

	return solution
}

func main() {
	sets, ids := readInput("input.txt")
	startSol1 := time.Now()
	solution1 := sol1(sets, ids)
	elapsed := time.Since(startSol1)
	fmt.Println("Part 1:")
	fmt.Println("solution:", solution1)
	fmt.Println("Time taken:", elapsed)
	startSol2 := time.Now()
	solution2 := sol2(sets)
	elapsed2 := time.Since(startSol2)
	fmt.Println("Part 2:")
	fmt.Println("Solution:", solution2)
	fmt.Println("Time taken:", elapsed2)
	fmt.Println("Total time:", elapsed+elapsed2)
}
