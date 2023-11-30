package main

import (
	"fmt"
	"os"
	"strings"
)

var vowels = []string{"a", "e", "i", "o", "u"}

func readLines(filename string) []string {
	c, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(c), "\n")
}

func solve1(l []string) {
	//contains at least three vowels
	//contains at least one letter that appears twice in a row
	//does not contain the strings ab, cd, pq, or xy
	nice := 0
	for _, s := range l {
		if s == "" {
			continue
		}
		vowels := 0
		double := false
		var lastRead byte
		for i := 0; i < len(s); i++ {
			if strings.ContainsAny(string(s[i]), "aeiou") {
				vowels++
			}
			if s[i] == lastRead {
				double = true
			}
			if lastRead == 'a' && s[i] == 'b' {
				vowels = 0
				break
			}
			if lastRead == 'c' && s[i] == 'd' {
				vowels = 0
				break
			}
			if lastRead == 'p' && s[i] == 'q' {
				vowels = 0
				break
			}
			if lastRead == 'x' && s[i] == 'y' {
				vowels = 0
				break
			}
			lastRead = s[i]
		}
		if vowels >= 3 && double {
			nice++
		}
	}
	fmt.Printf("there are %d nice strings, result:\n%d\n", nice, nice)
}

func solve2(l []string) {
	nice := 0
	for _, s := range l {
		if s == "" {
			continue
		}
		criteria1 := false
		for i := 0; i < len(s); i++ {
			if i+1 >= len(s) {
				break
			}
			checking := string(s[i]) + string(s[i+1])
			// fmt.Println("checking", checking, "in", s)
			if strings.Count(s, checking) > 1 {
				criteria1 = true
				// fmt.Println(s, "has met the first criteria")
				break
			}
		}
		criteria2 := false
		for i := 0; i < len(s); i++ {
			if i+2 >= len(s) {
				break
			}
			if s[i] == s[i+2] {
				criteria2 = true
				// fmt.Println(s, "has met the second criteria")
				break
			}
		}
		if criteria2 && criteria1 {
			nice++
		}
	}
	fmt.Printf("there are %d nice strings, result:\n%d\n", nice, nice)

}

func main() {
	l := readLines("input.txt")
	// solve1(l)
	solve2(l)
}
