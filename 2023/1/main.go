package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func readLine(filename string) []string {
	c, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(c), "\n")
}

func solve1(lines []string) {
	total := 0
	for _, line := range lines {
		firstNumber := 0
		for i := 0; i < len(line); i++ {
			r := rune(line[i])
			if unicode.IsNumber(r) {
				firstNumber, _ = strconv.Atoi(string(line[i]))
				break
			}
		}
		secondNumber := 0
		for i := len(line) - 1; i >= 0; i-- {
			r := rune(line[i])
			if unicode.IsNumber(r) {
				secondNumber, _ = strconv.Atoi(string(line[i]))
				break
			}
		}
		final, _ := strconv.Atoi(fmt.Sprintf("%d%d", firstNumber, secondNumber))
		total += final
	}
	fmt.Println(total)
}

func solve2(lines []string) {
	n := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	total := 0

	for _, line := range lines {
		copyL := line
		firstNumber := 0
		l := len(line)
		for i := 0; i < l; i++ {
			// fmt.Println("i is", i, "l is", l)
			r := rune(line[0])
			// fmt.Print("checking if ", string(r), " is a number? ")
			if unicode.IsNumber(r) {
				firstNumber, _ = strconv.Atoi(string(line[0]))
				// fmt.Println("it is")
				break
			} else {
				// fmt.Println("is is not")
			}
			for k, v := range n {
				// fmt.Println("checking if line", line, "has prefix", k)
				if strings.HasPrefix(line, k) {
					firstNumber = v
					// fmt.Println("has prefix")
					break
				}
			}
			if firstNumber != 0 {
				break
			}
			line = line[1:]
			// fmt.Println("line is now", line)
		}

		secondNumber := 0
		for i := len(line) - 1; i >= 0; i-- {
			r := rune(line[len(line)-1])
			// fmt.Println("checking if", string(r), "is a number")
			if unicode.IsNumber(r) {
				secondNumber, _ = strconv.Atoi(string(line[len(line)-1]))
				// fmt.Println("is a number")
				break
			}
			for k, v := range n {
				// fmt.Println("checking if line", line, "has suffix", k)
				if strings.HasSuffix(line, k) {
					// fmt.Println("has suffix")
					secondNumber = v
					break
				}
			}
			if secondNumber != 0 {
				break
			}
			line = line[:len(line)-1]
		}

		final, _ := strconv.Atoi(fmt.Sprintf("%d%d", firstNumber, secondNumber))
		fmt.Println("final for line", copyL, "is", final)
		total += final
	}
	fmt.Println(total)
}

func main() {
	l := readLine("input.txt")
	solve1(l)
	solve2(l)
}
