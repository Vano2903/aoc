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
			if !isInNumber && len(numberChars) > 0 {
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
		if len(numberChars) > 0 {
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

type val struct {
	val int
	key string
}

func solve2(lines []string) {
	gears := make(map[string][]val)
	keys := make([]string, 0)
	for posY, v := range lines {
		isInNumber := false
		numberChars := make([]string, 0)
		inGear := false
		posKey := ""
		for posX, char := range v {
			// fmt.Println("current char:", string(char), "is number:", unicode.IsNumber(char), "chars:", strings.Join(numberChars, ""))
			if unicode.IsNumber(char) {
				isInNumber = true
				numberChars = append(numberChars, string(char))
				posKey = fmt.Sprintf("%d-%d", posY+1, posX+1)
			} else {
				isInNumber = false
			}
			if isInNumber {
				if posY != 0 && posX != 0 {
					// fmt.Println()
					y := posY - 1
					x := posX - 1
					topLeft := rune(lines[y][x])
					if topLeft == '*' {
						inGear = true
						key := fmt.Sprintf("%d-%d", y+1, x+1)
						keys = append(keys, key)
						if len(gears[key]) == 0 {
							gears[key] = make([]val, 0)
						}
					}
				}
				if posY != 0 {
					y := posY - 1
					x := posX

					top := rune(lines[y][x])
					if top == '*' {
						inGear = true
						key := fmt.Sprintf("%d-%d", y+1, x+1)
						keys = append(keys, key)
						if len(gears[key]) == 0 {
							gears[key] = make([]val, 0)
						}
					}
				}

				if posY != 0 && posX != len(v)-1 {
					y := posY - 1
					x := posX + 1
					topRight := rune(lines[y][x])
					if topRight == '*' {
						inGear = true
						key := fmt.Sprintf("%d-%d", y+1, x+1)
						keys = append(keys, key)
						if len(gears[key]) == 0 {
							gears[key] = make([]val, 0)
						}
					}
				}

				if posX != 0 {
					y := posY
					x := posX - 1
					left := rune(lines[y][x])
					if left == '*' {
						inGear = true
						key := fmt.Sprintf("%d-%d", y+1, x+1)
						keys = append(keys, key)
						if len(gears[key]) == 0 {
							gears[key] = make([]val, 0)
						}
					}
				}
				if posX != len(v)-1 {
					y := posY
					x := posX + 1
					right := rune(v[posX+1])
					if right == '*' {
						inGear = true
						key := fmt.Sprintf("%d-%d", y+1, x+1)
						keys = append(keys, key)
						if len(gears[key]) == 0 {
							gears[key] = make([]val, 0)
						}
					}
				}

				if posY != len(lines)-1 && posX != 0 {
					y := posY + 1
					x := posX - 1
					bottomLeft := rune(lines[posY+1][posX-1])
					if bottomLeft == '*' {
						inGear = true
						key := fmt.Sprintf("%d-%d", y+1, x+1)
						keys = append(keys, key)
						if len(gears[key]) == 0 {
							gears[key] = make([]val, 0)
						}
					}
				}
				if posY != len(lines)-1 {
					y := posY + 1
					x := posX
					bottom := rune(lines[y][x])
					if bottom == '*' {
						inGear = true
						key := fmt.Sprintf("%d-%d", y+1, x+1)
						keys = append(keys, key)
						if len(gears[key]) == 0 {
							gears[key] = make([]val, 0)
						}
					}

				}

				if posY != len(lines)-1 && posX != len(v)-1 {
					y := posY + 1
					x := posX + 1
					bottomRight := rune(lines[y][x])
					if bottomRight == '*' {
						inGear = true
						key := fmt.Sprintf("%d-%d", y+1, x+1)
						keys = append(keys, key)
						if len(gears[key]) == 0 {
							gears[key] = make([]val, 0)
						}
					}
				}
			}
			if !isInNumber && len(numberChars) > 0 {
				if inGear {
					n := strings.Join(numberChars, "")
					number, _ := strconv.Atoi(n)
					for _, key := range keys {
						gears[key] = append(gears[key], val{number, posKey})
						// fmt.Println(number, "is valid")
					}
					// } else {
					// 	n := strings.Join(numberChars, "")
					// fmt.Println(n, "is not valid")
				}
			}
			if !isInNumber {
				keys = make([]string, 0)
				numberChars = make([]string, 0)
				inGear = false
			}
		}
		if inGear {
			n := strings.Join(numberChars, "")
			number, _ := strconv.Atoi(n)
			for _, key := range keys {
				gears[key] = append(gears[key], val{number, posKey})
			}
		}
		keys = make([]string, 0)
	}

	rateo := 0
	for k, v := range gears {
		// if len()
		vals := removeDuplicates(v)
		if len(vals) == 2 {
			fmt.Println("valid gear in postion", k, "with values", vals)
			rateo += vals[0].val * vals[1].val
		}
		// fmt.Println("k", k, "v", v)
		// fmt.Println(sum)
	}
	fmt.Println("rateo", rateo)
}

func removeDuplicates(strList []val) []val {
	list := []val{}
	for _, item := range strList {
		if !contains(list, item) {
			list = append(list, item)
		}
	}
	return list
}

func contains(s []val, e val) bool {
	for _, a := range s {
		if a.key == e.key && a.val == e.val {
			return true
		}
	}
	return false
}

func main() {
	l := readFile("input.txt")
	// solve1(l)
	solve2(l)
	//539713
	//539651

	//83290270
	//84084546
	//84084546
	//84084546
}
