package main

import (
	"fmt"
	"os"
	"strings"
)

func readFile(path string) []string {
	c, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(c), "\n")
}

func printSection(section [][]string) {
	for _, line := range section {
		for _, c := range line {
			fmt.Print(c)
		}
		fmt.Println()
	}
}

// it's valid only with 1 or 0 difference
// so we can return with 3 if we know they are too different
func differenceCount(a, b []string) int {
	differenceCount := 0
	if len(a) != len(b) {
		return 3
	}
	for i, c := range a {
		if c != b[i] {
			if differenceCount > 1 {
				return 3
			}
			differenceCount++
		}
	}
	return differenceCount
}

func copyAndAppend(toAppend [][][]string, section [][]string) [][][]string {
	s := make([][]string, len(section))
	for i, line := range section {
		s[i] = make([]string, len(line))
		copy(s[i], line)
	}
	return append(toAppend, s)
}

func cleanSection(input [][]string) [][][]string {
	cleanedSections := [][][]string{}

	section := make([][]string, len(input))
	for i := range input {
		section[i] = make([]string, len(input[i]))
		copy(section[i], input[i])
	}

	//checking for horizontal mirrors
	for i := 0; i < len(section)-1; i++ {
		// fmt.Println("checking", section[i], "and", section[i+1])
		equals := differenceCount(section[i], section[i+1])
		if equals == 1 {
			//clean the smudge
			for k := 0; k < len(section[i]); k++ {
				if section[i][k] != section[i+1][k] {
					section[i][k] = section[i+1][k]
					break
				}
			}

			cleanedSections = copyAndAppend(cleanedSections, section)
			for i := range input {
				section[i] = make([]string, len(input[i]))
				copy(section[i], input[i])
			}
		} else if equals == 0 {
			smudgeFound := false
			// fmt.Println("possible horizontal mirror between", i, "and", i+1)
			isMirror := true
			//i-1 and i+2 then i-2 and i+3
			maxIndex := min(i, len(section)-2-i)
			// fmt.Println("i", i, "len", len(section), "len-i", len(section)-2-i, "maxIndex", maxIndex)
			for j := 1; j <= maxIndex; j++ {
				// fmt.Println("checking", i-j, "and", i+j+1)
				diffCount := differenceCount(section[i-j], section[i+j+1])
				if diffCount >= 2 {
					// fmt.Println("not equals, not a horizontal mirror")
					isMirror = false
					break
					// } else {
					// 	fmt.Println("equals")
				}

				if diffCount == 1 {
					if smudgeFound {
						//there is already a smudge, can have only one
						isMirror = false
						break
					} else {
						smudgeFound = true
						//clean the smudge
						for k := 0; k < len(section[i-j]); k++ {
							if section[i-j][k] != section[i+j+1][k] {
								section[i-j][k] = section[i+j+1][k]
								break
							}
						}
					}
				}
			}
			if isMirror && smudgeFound {
				// fmt.Println("mirror found between", i, "and", i+1)
				// mirror = i
				// return mirror + 1, true
				cleanedSections = copyAndAppend(cleanedSections, section)
				for i := range input {
					section[i] = make([]string, len(input[i]))
					copy(section[i], input[i])
				}
			}
		}
	}
	for i := range input {
		section[i] = make([]string, len(input[i]))
		copy(section[i], input[i])
	}

	//checking for vertical mirrors
	// fmt.Println("checking for vertical mirrors")
	for i := 0; i < len(section[0])-1; i++ {
		column1 := []string{}
		column2 := []string{}
		for _, line := range section {
			column1 = append(column1, line[i])
			column2 = append(column2, line[i+1])
		}
		fmt.Println("checking", column1, "and", column2)
		equals := differenceCount(column1, column2)
		fmt.Println("differences", equals)
		if equals == 1 {
			//clean the smudge
			for k := 0; k < len(column1); k++ {
				if column1[k] != column2[k] {
					column1[k] = column2[k]
					break
				}
			}
			for j := 0; j < len(section); j++ {
				section[j][i] = column1[j]
			}
			cleanedSections = copyAndAppend(cleanedSections, section)
			for i := range input {
				section[i] = make([]string, len(input[i]))
				copy(section[i], input[i])
			}
		} else if equals == 0 {
			isMirror := true
			smudgeFound := false
			//i-1 and i+2 then i-2 and i+3
			maxIndex := min(i, len(section[0])-2-i)
			// fmt.Println("i", i, "len", len(section[0]), "len-i", len(section[0])-2-i, "maxIndex", maxIndex)
			column1 := []string{}
			column2 := []string{}

			for j := 1; j <= maxIndex; j++ {
				// fmt.Println("checking", i-j, "and", i+j+1)
				for _, line := range section {
					column1 = append(column1, line[i-j])
					column2 = append(column2, line[i+j+1])
				}
				fmt.Println("checking", i-j, "and", i+j+1)
				diffCount := differenceCount(column1, column2)
				fmt.Println(diffCount)
				if diffCount >= 2 {
					// fmt.Println("not equals, not a horizontal mirror")
					isMirror = false
					break
					// } else {
					// 	fmt.Println("equals")
				}

				if diffCount == 1 {
					fmt.Println("found smudge")
					if smudgeFound {
						//there is already a smudge, can have only one
						isMirror = false
						break
					} else {
						smudgeFound = true
						//clean the smudge
						for k := 0; k < len(column1); k++ {
							fmt.Println(column1[k], "==", column2[k])
							if column1[k] != column2[k] {
								column1[k] = column2[k]
								break
							}
						}
						fmt.Println("cleaned smudge")

					}
				}
			}
			if isMirror && smudgeFound {
				// fmt.Println("mirror found between", i, "and", i+1)
				// mirror = i
				// return mirror + 1, true
				cleanedSections = copyAndAppend(cleanedSections, section)
				for i := range input {
					section[i] = make([]string, len(input[i]))
					copy(section[i], input[i])
				}
			}
		}
	}
	return cleanedSections
}

func solve2(input []string) {
	total := 0
	var section [][]string
	for _, line := range input {
		if line == "" {
			// sections := cleanSection(section)
			// fmt.Println("cleaned sections")
			// for _, section := range sections {
			// 	printSection(section)
			// 	fmt.Println()
			// }
			pos, isInLine := solveSectionWithSmudges(section)
			// if pos == -1 {
			// 	continue
			// }
			if isInLine {
				total += pos * 100
			} else {
				total += pos
			}

			// for _, section := range sections {
			// 	printSection(section)
			// 	fmt.Println()
			// }
			// fmt.Println()
			section = [][]string{}
			// break
		} else {
			section = append(section, strings.Split(line, ""))
		}
	}
	pos, isInLine := solveSectionWithSmudges(section)
	// if pos == -1 {
	// 	continue
	// }

	if isInLine {
		total += pos * 100
	} else {
		total += pos
	}
	// sections := cleanSection(section)

	// for _, section := range sections {
	// 	printSection(section)
	// 	fmt.Println()
	// 	pos, isInLine := solveSection(section)
	// 	if pos == -1 {
	// 		continue
	// 	}

	// 	if isInLine {
	// 		total += pos * 100
	// 	} else {
	// 		total += pos
	// 	}
	// }

	fmt.Println("total", total)
}

// true if mirror is in line
func solveSectionWithSmudges(section [][]string) (int, bool) {
	// fmt.Println("section")
	mirror := 0
	//checking for horizontal mirrors
	for i := 0; i < len(section)-1; i++ {
		// fmt.Println("checking", section[i], "and", section[i+1])
		diffCount := differenceCount(section[i], section[i+1])
		if diffCount < 2 {
			smudgeFound := diffCount == 1
			// fmt.Println("possible horizontal mirror between", i, "and", i+1)
			isMirror := true
			//i-1 and i+2 then i-2 and i+3
			maxIndex := min(i, len(section)-2-i)
			// fmt.Println("i", i, "len", len(section), "len-i", len(section)-2-i, "maxIndex", maxIndex)
			for j := 1; j <= maxIndex; j++ {
				// fmt.Println("checking", i-j, "and", i+j+1)
				diffCount := differenceCount(section[i-j], section[i+j+1])
				if diffCount >= 2 {
					//too different, not a mirror
					isMirror = false
					break
				}
				if diffCount == 1 {
					if smudgeFound {
						//there is already a smudge, can have only one
						isMirror = false
						break
					} else {
						smudgeFound = true
					}
				}
			}
			if isMirror && smudgeFound {
				// fmt.Println("horizontal mirror found between", i, "and", i+1)
				mirror = i
				return mirror + 1, true
			}
		}
	}

	//checking for vertical mirrors
	fmt.Println()
	fmt.Println()
	fmt.Println("checking for vertical mirrors")
	for i := 0; i < len(section[0])-1; i++ {
		column1 := []string{}
		column2 := []string{}
		for _, line := range section {
			column1 = append(column1, line[i])
			column2 = append(column2, line[i+1])
		}
		fmt.Println("checking difference between\n", column1, "\n", column2)
		diffCount := differenceCount(column1, column2)
		fmt.Println(diffCount)
		if diffCount < 2 {
			smudgeFound := diffCount == 1
			if smudgeFound {
				fmt.Println("found a possible mirror but it has a smudge at the start")
			} else {
				fmt.Println("found a possible mirror, exploring")
			}
			// fmt.Println("possible vertical mirror between", i, "and", i+1)
			isMirror := true
			//i-1 and i+2 then i-2 and i+3
			maxIndex := min(i, len(section[0])-2-i)
			fmt.Println("i", i, "len", len(section[0]), "len-i", len(section[0])-2-i, "maxIndex", maxIndex)

			for j := 1; j <= maxIndex; j++ {
				column1 := []string{}
				column2 := []string{}
				// fmt.Println("checking", i-j, "and", i+j+1)
				for _, line := range section {
					column1 = append(column1, line[i-j])
					column2 = append(column2, line[i+j+1])
				}
				fmt.Println("checking between\n", column1, "\n", column2)
				diffCount := differenceCount(column1, column2)
				fmt.Println("diffCount", diffCount)
				if diffCount >= 2 {
					fmt.Println("not equals, not a vertical mirror")
					isMirror = false
					break
				}
				if diffCount == 1 {
					if smudgeFound {
						fmt.Println("smudge already found2")
						//there is already a smudge, can have only one
						isMirror = false
						break
					} else {
						fmt.Println("found the first smudge")
						smudgeFound = true
					}
				}
			}
			if isMirror && smudgeFound {
				fmt.Println("vertical mirror found between", i, "and", i+1)
				mirror = i
				return mirror + 1, false
			}
		}
	}
	fmt.Println("no mirror found")
	printSection(section)
	fmt.Println()
	panic("no mirror found")
}

func equalLines(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, c := range a {
		if c != b[i] {
			return false
		}
	}
	return true
}

// true if mirror is in line
func solveSection(section [][]string) (int, bool) {
	// fmt.Println("section")
	mirror := 0
	//checking for horizontal mirrors
	for i := 0; i < len(section)-1; i++ {
		// fmt.Println("checking", section[i], "and", section[i+1])
		equals := equalLines(section[i], section[i+1])
		if equals {
			// fmt.Println("possible horizontal mirror between", i, "and", i+1)
			isMirror := true
			//i-1 and i+2 then i-2 and i+3
			maxIndex := min(i, len(section)-2-i)
			// fmt.Println("i", i, "len", len(section), "len-i", len(section)-2-i, "maxIndex", maxIndex)
			for j := 1; j <= maxIndex; j++ {
				// fmt.Println("checking", i-j, "and", i+j+1)
				if !equalLines(section[i-j], section[i+j+1]) {
					// fmt.Println("not equals, not a horizontal mirror")
					isMirror = false
					break
					// } else {
					// 	fmt.Println("equals")
				}
			}
			if isMirror {
				// fmt.Println("mirror found between", i, "and", i+1)
				mirror = i
				return mirror + 1, true
			}
		}
	}

	//checking for vertical mirrors
	// fmt.Println("checking for vertical mirrors")
	for i := 0; i < len(section[0])-1; i++ {
		column1 := []string{}
		column2 := []string{}
		for _, line := range section {
			column1 = append(column1, line[i])
			column2 = append(column2, line[i+1])
		}
		// fmt.Println("checking", column1, "and", column2)
		equals := equalLines(column1, column2)
		if equals {
			// fmt.Println("possible vertical mirror between", i, "and", i+1)
			isMirror := true
			//i-1 and i+2 then i-2 and i+3
			maxIndex := min(i, len(section[0])-2-i)
			// fmt.Println("i", i, "len", len(section[0]), "len-i", len(section[0])-2-i, "maxIndex", maxIndex)
			column1 := []string{}
			column2 := []string{}

			for j := 1; j <= maxIndex; j++ {
				// fmt.Println("checking", i-j, "and", i+j+1)
				for _, line := range section {
					column1 = append(column1, line[i-j])
					column2 = append(column2, line[i+j+1])
				}

				if !equalLines(column1, column2) {
					// fmt.Println("not equals, not a vertical mirror")
					isMirror = false
					break
					// } else {
					// 	fmt.Println("equals")
				}
			}
			if isMirror {
				// fmt.Println("mirror found between", i, "and", i+1)
				mirror = i
				return mirror + 1, false
			}
		}
	}

	// fmt.Println("mirror at column", mirror+1)
	fmt.Println("should not be here")
	return -1, false
}

func solve1(input []string) {
	total := 0
	var section [][]string
	for _, line := range input {
		if line == "" {
			pos, isInLine := solveSection(section)
			if isInLine {
				total += pos * 100
			} else {
				total += pos
			}
			// fmt.Println()
			section = [][]string{}
		} else {
			section = append(section, strings.Split(line, ""))
		}
	}
	pos, isInLine := solveSection(section)
	if isInLine {
		total += pos * 100
	} else {
		total += pos
	}

	fmt.Println("total", total)
}

func main() {
	l := readFile("input.txt")
	// solve1(l)
	solve2(l)
	// fmt.Println("test")
	// l = readFile("input.txt")
	// solve2(l)
	// fmt.Println("real (should be less than 43666 and grater than 37962)")
}

//45673 too high
//43666 too high
//43579 wrong
//57176
//37962 too low
