package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func readFile(path string) []string {
	c, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(c), "\n")
}

func countSols(springs []byte, values []int, groupSize int, cache map[string]int) int {
	if len(springs) == 0 {
		if len(values) == 1 && values[0] == groupSize {
			return 1
		}
		if len(values) == 0 && groupSize == 0 {
			return 1
		}
		return 0
	}

	v := ""
	for _, val := range values {
		v += strconv.Itoa(val)
	}
	key := string(springs) + string(v) + strconv.Itoa(groupSize)
	if sol, ok := cache[key]; ok {
		return sol
	}

	solutions := 0
	if springs[0] == '?' {
		//considering ? as #
		solutions += countSols(springs[1:], values, groupSize+1, cache)

		//considering ? as .
		if groupSize == 0 {
			//go to next position cause we are not in a group
			solutions += countSols(springs[1:], values, 0, cache)
		} else {
			if len(values) > 0 && values[0] == groupSize {
				solutions += countSols(springs[1:], values[1:], 0, cache)
			}
		}
	} else if springs[0] == '#' {
		solutions += countSols(springs[1:], values, groupSize+1, cache)
	} else {
		if groupSize == 0 {
			//go to next position cause we are not in a group
			solutions += countSols(springs[1:], values, 0, cache)
		} else {
			if len(values) > 0 && values[0] == groupSize {
				//group done, go to next value
				solutions += countSols(springs[1:], values[1:], 0, cache)
			}
		}
	}
	cache[key] = solutions

	return solutions
}

func getPossibilitiesForLine(line string) int {
	l := strings.Split(line, " ")
	v := strings.Split(l[1], ",")
	values := make([]int, len(v))
	for i, valString := range v {
		val, _ := strconv.Atoi(valString)
		values[i] = val
	}
	cache := make(map[string]int)
	springs := l[0]
	return countSols([]byte(springs), values, 0, cache)
}

func solve1(input []string) {
	total := 0
	for _, line := range input {
		total += getPossibilitiesForLine(line)
	}
	fmt.Println(total)
}

func solve2(input []string) {
	lines := make([]string, len(input))
	for i, line := range input {
		l := strings.Split(line, " ")
		for j := 0; j < 5; j++ {
			lines[i] += l[0] + "?"
		}
		lines[i] += " "

		for j := 0; j < 5; j++ {
			lines[i] += l[1] + ","
		}
	}

	// wg := new(sync.WaitGroup)
	total := 0
	for _, line := range lines {
		// wg.Add(1)
		// go func(wg *sync.WaitGroup, line string) {
		total += getPossibilitiesForLine(line)
		// wg.Done()
		// }(wg, line)
	}
	// wg.Wait()

	fmt.Println(total)
}

func main() {
	lines := readFile("input.txt")
	now := time.Now()
	solve1(lines)
	fmt.Println("Took:", time.Since(now))

	fmt.Println()
	now = time.Now()
	solve2(lines)
	fmt.Println("Took:", time.Since(now))
}
