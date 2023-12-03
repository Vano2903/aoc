package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func readFile(path string) []string {
	c, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(c), "\n")
}

func solve1(lines []string) {
	sum := 0
	for posY, v := range lines {
		isInNumber := false
		numberChars := make([]string, 0)
		valid := false
		for posX, char := range v {
			// fmt.Println("current char:", string(char), "is number:", unicode.IsNumber(char), "chars:", strings.Join(numberChars, ""))
			if unicode.IsNumber(char) {
				isInNumber = true
				numberChars = append(numberChars, string(char))
			} else {
				isInNumber = false
			}
			if isInNumber {
				if posY != 0 && posX != 0 {
					topLeft := rune(lines[posY-1][posX-1])
					if !unicode.IsNumber(topLeft) && topLeft != '.' {
						//fmt.Println(strings.Join(numberChars, ""), "is valid cause of top left")
						valid = true
					}
				}
				if posY != 0 {
					top := rune(lines[posY-1][posX])
					if !unicode.IsNumber(top) && top != '.' {
						//fmt.Println(strings.Join(numberChars, ""), "is valid cause of top")
						valid = true
					}
				}

				if posY != 0 && posX != len(v)-1 {
					topRight := rune(lines[posY-1][posX+1])
					if !unicode.IsNumber(topRight) && topRight != '.' {
						//fmt.Println(strings.Join(numberChars, ""), "is valid cause of top right")
						valid = true
					}
				}

				if posX != 0 {
					left := rune(v[posX-1])
					if !unicode.IsNumber(left) && left != '.' {
						//fmt.Println(strings.Join(numberChars, ""), "is valid cause of left")
						valid = true
					}
				}
				if posX != len(v)-1 {
					right := rune(v[posX+1])
					if !unicode.IsNumber(right) && right != '.' {
						//fmt.Println(strings.Join(numberChars, ""), "is valid cause of right")
						valid = true
					}
				}

				if posY != len(lines)-1 && posX != 0 {
					bottomLeft := rune(lines[posY+1][posX-1])
					if !unicode.IsNumber(bottomLeft) && bottomLeft != '.' {
						//fmt.Println(strings.Join(numberChars, ""), "is valid cause of bottom left")
						valid = true
					}
				}
				if posY != len(lines)-1 {
					bottom := rune(lines[posY+1][posX])
					if !unicode.IsNumber(bottom) && bottom != '.' {
						//fmt.Println(strings.Join(numberChars, ""), "is valid cause of bottom")
						valid = true
					}
				}

				if posY != len(lines)-1 && posX != len(v)-1 {
					bottomRight := rune(lines[posY+1][posX+1])
					if !unicode.IsNumber(bottomRight) && bottomRight != '.' {
						//fmt.Println(strings.Join(numberChars, ""), "is valid cause of bottom right")
						valid = true
					}
				}
			}
			if !isInNumber && len(numberChars) > 1 {
				if valid {
					n := strings.Join(numberChars, "")
					number, _ := strconv.Atoi(n)
					sum += number
					fmt.Println(number, "is valid")
				} else {
					n := strings.Join(numberChars, "")
					fmt.Println(n, "is not valid")
				}
			}
			if !isInNumber {
				numberChars = make([]string, 0)
				valid = false
			}
		}
		if len(numberChars) > 1 {
			if valid {
				n := strings.Join(numberChars, "")
				number, _ := strconv.Atoi(n)
				sum += number
				fmt.Println(number, "is valid")
			} else {
				n := strings.Join(numberChars, "")
				fmt.Println(n, "is not valid")
			}
		}
	}
	fmt.Println(sum)
}

func main() {
	l := readFile("a.txt")
	solve1(l)
}
