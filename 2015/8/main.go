package main

import (
	"fmt"
	"os"
	"strings"
)

func readLines(filename string) []string {
	c, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(c), "\n")
}

func solve1(lines []string) int {
	//"" 2 chars 0 length
	//"abc" 5 chars 3 length
	//"aaa\"aaa" 10 chars 7 length (6 As and a ")
	//"\x27" 6 chars 1 length (a single character of hex code)
	//escape sequences are \\ (single backslash), \" (double quote), and \xAB (hexadecimal val)
	totalChars := 0
	inMemoryTotalChars := 0
	for _, l := range lines {
		if l == "" {
			continue
		}
		// fmt.Println(len(l))
		inMem := len(l) - 2
		for i := 0; i < len(l); i++ {
			if l[i] == '\\' {
				if l[i+1] == '\\' || l[i+1] == '"' {
					inMem--
					i++
				} else if l[i+1] == 'x' {
					inMem -= 3
					i += 3
				}
			}
		}
		inMemoryTotalChars += inMem
		totalChars += len(l)
	}
	// fmt.Println(totalChars)
	// fmt.Println(inMemoryTotalChars)
	fmt.Println("solution 1:", totalChars-inMemoryTotalChars)
	return totalChars
}

func solve2(lines []string) {
	totalChars := 0
	for _, l := range lines {
		if l == "" {
			continue
		}
		// fmt.Println(len(l))
		// inMem := len(l) - 2
		totalChars += len(l) + 4

		for i := 0; i < len(l); i++ {
			if l[i] == '\\' {
				if l[i+1] == '\\' || l[i+1] == '"' {
					totalChars += 2
					i++
				} else if l[i+1] == 'x' {
					totalChars++
					i += 3
				}
			}
		}
	}

	fmt.Println("solution 2:", totalChars-solve1(lines))
}

func main() {
	l := readLines("input.txt")
	solve1(l)
	solve2(l)
}
