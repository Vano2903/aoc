package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type IDRange struct {
	startI int
	endI   int
	startS string
	endS   string
}

func readInput(path string) []IDRange {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	line := strings.TrimSpace(string(input))
	ids := strings.Split(line, ",")
	idRanges := make([]IDRange, len(ids))
	for i, line := range ids {

		idSplits := strings.Split(line, "-")
		start := idSplits[0]
		end := idSplits[1]
		startI, _ := strconv.Atoi(start)
		endI, _ := strconv.Atoi(end)
		idRanges[i] = IDRange{startI: startI, endI: endI, startS: start, endS: end}
	}
	return idRanges
}

func sol1(input []IDRange) int {
	// invalidIDCounter := 0
	var invalidIDCounter int

	var currentID int
	currentIDStr := ""
	for _, r := range input {
		// fmt.Printf("Range: %d - %d, should check: %d numbers\n", r.startI, r.endI, r.endI-(r.startI+1))
		currentID = r.startI
		currentIDStr = r.startS
		for currentID < r.endI+1 {
			// fmt.Println("checking ID:", currentIDStr)
			// if len(currentIDStr)%2 != 0 {
			// 	currentID++
			// 	currentIDStr = strconv.Itoa(currentID)
			// 	continue
			// }
			mid := len(currentIDStr) / 2
			firstHalf := currentIDStr[:mid]
			secondHalf := currentIDStr[mid:]
			if firstHalf == secondHalf {
				invalidIDCounter += currentID
			}
			currentID++
			currentIDStr = strconv.Itoa(currentID)
		}
	}

	return invalidIDCounter
}

func checkRepetitionInNumberWithMiddlePoint(idStr string, mid int) bool {
	if mid == 0 {
		return false
	}

	if len(idStr)%mid != 0 {
		return checkRepetitionInNumberWithMiddlePoint(idStr, mid-1)
	}

	var sections []string
	for i := 0; i < len(idStr); i += mid {
		// end := min(i+mid, len(idStr))
		sections = append(sections, idStr[i:i+mid])
	}

	repetitions := true
	// fmt.Println("checking sections:", sections, "for idStr:", idStr, "with mid:", mid)
	for i := 0; i < len(sections)-1; i++ {
		// fmt.Printf("comparing section[%d]: %s and section[%d]: %s\n", i, sections[i], i+1, sections[i+1])
		if sections[i] != sections[i+1] {
			repetitions = false
			break
		}
	}

	// if repetitions {
	// 	fmt.Println("found repetition in idStr:", idStr, "with mid:", mid)
	// }

	if !repetitions {
		return checkRepetitionInNumberWithMiddlePoint(idStr, mid-1)
	}
	return true
}

func checkRepetitionInNumber(idStr string) bool {
	midPoint := len(idStr) / 2
	if midPoint == 0 {
		return false
	}
	if len(idStr)%midPoint != 0 {
		return checkRepetitionInNumberWithMiddlePoint(idStr, midPoint-1)
	}
	return checkRepetitionInNumberWithMiddlePoint(idStr, midPoint)
}

func sol2(input []IDRange) int {
	invalidIDCounter := 0

	currentID := 0
	currentIDStr := ""
	for _, r := range input {
		currentID = r.startI
		currentIDStr = strconv.Itoa(currentID)
		for currentID < r.endI+1 {
			if checkRepetitionInNumber(currentIDStr) {
				invalidIDCounter += currentID
			}
			currentID++
			currentIDStr = strconv.Itoa(currentID)
		}
	}

	return invalidIDCounter
}

// > 6835 && < 8026
// not 7443
func main() {
	input := readInput("mora.txt")
	startSol1 := time.Now()
	zeroCount := sol1(input)
	elapsed := time.Since(startSol1)
	// if zeroCount != 52316131093 {
	// 	panic("wrong solution for part 1: " + strconv.Itoa(zeroCount))
	// }
	fmt.Println("Part 1:")
	fmt.Println("solution:", zeroCount)
	fmt.Println("Time taken:", elapsed)
	startSol2 := time.Now()
	finalPosition := sol2(input)
	elapsed2 := time.Since(startSol2)
	// if finalPosition != 69564213293 {
	// 	panic("wrong solution for part 2: " + strconv.Itoa(finalPosition))
	// }
	fmt.Println("Part 2:")
	fmt.Println("Solution:", finalPosition)
	fmt.Println("Time taken:", elapsed2)
	fmt.Println("Total time:", elapsed+elapsed2)
}
