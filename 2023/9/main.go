package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func readLines(path string) []string {
	c, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(c), "\n")
}

// func solveForValues(indexA, indexB int, values []int, differences []int) {
// 	if indexA < 0 || indexB < 0 {
// 		fmt.Println("reached start, calcualting next")
// 		return
// 	}
// 	a := values[indexA]
// 	b := values[indexB]
// 	fmt.Println("solving for", a, b)
// 	differnce := a - b
// 	fmt.Println("difference", differnce)
// 	if differnce == 0 {
// 		fmt.Println("reached 0, calculating next")
// 	}
// 	differences = append(differences, differnce)
// 	solveForValues(indexA-1, indexB-1, values, differences)

// }

// func solveLine(line []int) {
// 	differences := []int{}
// 	solveForValues(len(line)-1, len(line)-2, line, differences)
// 	// for i := len(line) - 1; i >= 1; i-- {

// 	// fmt.Println(line[i], line[i-1])
// 	// difference := line[i] - line[i-1]
// 	// fmt.Println(difference)
// 	// if difference == 0 {
// 	// 	fmt.Println("reached 0, calculating next")
// 	// }
// 	// differences = append(differences, difference)
// 	// }

// }

func solve1(lines []string) {
	final := 0
	for _, line := range lines {
		differences := make([][]int, 0)
		l := strings.Split(line, " ")
		values := []int{}
		for _, c := range l {
			v, _ := strconv.Atoi(c)
			values = append(values, v)
		}
		differences = append(differences, []int{})

		differences[0] = append(differences[0], values...)
		differences = append(differences, []int{})
		for j := 0; j < len(values)-1; j++ {
			differences[1] = append(differences[1], values[j+1]-values[j])
		}
		for {
			index := len(differences) - 1
			differences = append(differences, []int{})
			last := true
			for j := 0; j < len(differences[index])-1; j++ {
				difference := differences[index][j+1] - differences[index][j]
				if difference != 0 {
					last = false
				}
				differences[index+1] = append(differences[index+1], difference)
			}
			if last {
				break
			}
		}

		//calcaulate the next
		toSum := 0
		for i := len(differences) - 2; i >= 0; i-- {
			toSum += differences[i][len(differences[i])-1]
		}
		final += toSum
	}

	fmt.Println("TOTAL", final)

}

func solve2(lines []string) {
	final := 0
	for _, line := range lines {
		differences := make([][]int, 0)
		l := strings.Split(line, " ")
		values := []int{}
		for _, c := range l {
			v, _ := strconv.Atoi(c)
			values = append(values, v)
		}
		differences = append(differences, []int{})

		differences[0] = append(differences[0], values...)
		differences = append(differences, []int{})
		for j := 0; j < len(values)-1; j++ {
			differences[1] = append(differences[1], values[j+1]-values[j])
		}
		for {
			index := len(differences) - 1
			differences = append(differences, []int{})
			last := true
			for j := 0; j < len(differences[index])-1; j++ {
				difference := differences[index][j+1] - differences[index][j]
				if difference != 0 {
					last = false
				}
				differences[index+1] = append(differences[index+1], difference)
			}
			if last {
				break
			}
		}

		//calcaulate the next
		toSum := 0
		for i := len(differences) - 2; i >= 0; i-- {
			toSum = differences[i][0] - toSum
		}
		final += toSum
	}
	fmt.Println("TOTAL", final)
}

func main() {
	l := readLines("input.txt")
	now := time.Now()
	solve1(l)
	fmt.Println("took", time.Since(now))
	now = time.Now()
	solve2(l)
	fmt.Println("took", time.Since(now))
}
