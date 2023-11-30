package main

import (
	"fmt"
	"os"
)

func readInput(filename string) []byte {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return content
}

func solve1(content []byte) {
	floor := 0
	for _, char := range content {
		if char == '(' {
			floor++
		} else {
			floor--
		}
	}
	fmt.Println("solution 1:")
	fmt.Println(floor)
}

func solve2(content []byte) {
	floor := 0
	for i, char := range content {
		// fmt.Println("char number:", i+1)
		// fmt.Println("floor:", floor)
		if char == '(' {
			floor++
		} else if char == ')' {
			floor--
		}
		if floor == -1 {
			fmt.Println("solution 2:")
			fmt.Println(i + 1)
			break
		}
	}
}

func main() {
	content := readInput("input.txt")
	solve1(content)
	solve2(content)
}
