package main

import (
	"fmt"
	"strings"
)

func validateFirstRule(input string) bool {
	for i := 0; i < len(input)-2; i++ {
		if input[i]+1 == input[i+1] && input[i]+2 == input[i+2] {
			return true
		}
	}
	return false
}

func validateSecondRule(input string) bool {
	invalidLetters := []string{"i", "o", "l"}
	for _, letter := range invalidLetters {
		if strings.Contains(input, letter) {
			return false
		}
	}
	return true
}

func validateThirdRule(input string) bool {
	pairsCount := 0
	pairs := []string{"aa", "bb", "cc", "dd", "ee", "ff", "gg", "hh", "ii", "jj", "kk", "ll", "mm", "nn", "oo", "pp", "qq", "rr", "ss", "tt", "uu", "vv", "ww", "xx", "yy", "zz"}
	for _, pair := range pairs {
		if strings.Count(input, pair) > 0 {
			pairsCount++
		}
	}
	return pairsCount >= 2
}

func validatePassword(input string) bool {
	return validateFirstRule(input) && validateSecondRule(input) && validateThirdRule(input)
}

func generateNextPasswordNoChecks(input string) string {
	inputBytes := []byte(input)
	for i := len(inputBytes) - 1; i >= 0; i-- {
		if inputBytes[i] == 'z' {
			inputBytes[i] = 'a'
		} else {
			inputBytes[i]++
			break
		}
	}
	return string(inputBytes)
}

func solve1(input string) string {
	// if !validatePassword(input) {
	// 	fmt.Println("input is invalid")
	// 	return
	// }
	for {
		input = generateNextPasswordNoChecks(input)

		if validatePassword(input) {
			// fmt.Println("next password:", input)
			return input
		}
	}
}

func solve2(input string) string {
	input = solve1(input)
	fmt.Println("first password:", input)
	return solve1(input)
}

func main() {
	input := "vzbxkghb"
	// input := "zzz"
	// fmt.Println(generateNextPasswordNoChecks(input))
	fmt.Println("solve 1:", solve1(input))
	fmt.Println("solve 2:", solve2(input))
}
