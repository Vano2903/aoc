package main

/*
package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	directionRight = iota
	directionUp
	directionLeft
	directionDown
)

func readMatrix(path string) [][]string {
	c, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(c), "\n")
	matrix := make([][]string, len(lines))
	for i, line := range lines {
		matrix[i] = strings.Split(line, "")
	}
	return matrix
}

func getKindOfStartingPoint(startX, startY int, pipes [][]string) string {
	top := "."
	bottom := "."
	right := "."
	left := "."
	if startX > 0 {
		top = pipes[startX-1][startY]
	}
	if startX < len(pipes)-1 {
		bottom = pipes[startX+1][startY]
	}
	if startY > 0 {
		left = pipes[startX][startY-1]
	}
	if startY < len(pipes[0])-1 {
		right = pipes[startX][startY+1]
	}
	// fmt.Println("top", top, "bottom", bottom, "left", left, "right", right)
	var validTop, validBottom, validLeft, validRight bool
	if top == "|" || top == "F" || top == "7" {
		validTop = true
	}
	if bottom == "|" || bottom == "L" || bottom == "J" {
		validBottom = true
	}
	if left == "-" || left == "F" || left == "L" {
		validLeft = true
	}
	if right == "-" || right == "7" || right == "J" {
		validRight = true
	}

	// fmt.Println("validTop", validTop, "validBottom", validBottom, "validLeft", validLeft, "validRight", validRight)
	if validTop {
		if validLeft {
			return "J"
		} else if validRight {
			return "L"
		} else if validBottom {
			return "|"
		}
	} else if validBottom {
		if validLeft {
			return "7"
		} else if validRight {
			return "F"
		}
	} else if validLeft {
		if validRight {
			return "-"
		}
	}
	fmt.Println("should not be here")
	return "ERRORE"
}

func solve1(input [][]string) {
	pipes := make([][]string, len(input))
	start := []int{0, 0}
	// fmt.Println(pipes)
	for i, line := range input {
		for j, char := range line {
			// if char == "." {
			// 	continue
			// }
			if char == "S" {
				start = []int{i, j}
				// break
			}
			if pipes[i] == nil {
				pipes[i] = make([]string, 0)
			}
			pipes[i] = append(pipes[i], char)
		}
		// fmt.Println(i, line)
	}
	kind := getKindOfStartingPoint(start[0], start[1], pipes)
	if kind == "ERRORE" {
		panic("ERRORE")
	}
	// fmt.Println(start)
	// fmt.Println(kind)
	// fmt.Println(pipes)
	pipes[start[0]][start[1]] = kind
	direction := 0
	//0 => right, 1 => up, 2 => left, 3 => down
	currentX, currentY := start[0], start[1]
	switch kind {
	case "J", "L", "|":
		direction = directionUp
		// fmt.Println("starting by going up")
	case "-":
		direction = directionRight
		// fmt.Println("starting by going right")
	case "F", "7":
		direction = directionDown
		// fmt.Println("starting by going down")
	}
	counter := 0
	for {
		counter++
		// fmt.Println("currentX", currentX, "currentY", currentY)
		switch direction {
		case directionRight:
			// fmt.Println("going right")
			currentY++
		case directionUp:
			// fmt.Println("going up")
			currentX--
		case directionLeft:
			// fmt.Println("going left")
			currentY--
		case directionDown:
			// fmt.Println("going down")
			currentX++
		}
		if currentX < 0 || currentX >= len(pipes) || currentY < 0 || currentY >= len(pipes[0]) {
			break
		}
		if currentX == start[0] && currentY == start[1] {
			// fmt.Println("back to start")
			break
		}
		// fmt.Println("currentX", currentX, "currentY", currentY)
		// fmt.Println("current char", pipes[currentX][currentY])
		if pipes[currentX][currentY] == "." {
			break
		}
		switch pipes[currentX][currentY] {
		case "J":
			if direction == directionRight {
				direction = directionUp
			} else if direction == directionDown {
				direction = directionLeft
			}
		case "L":
			if direction == directionDown {
				direction = directionRight
			} else if direction == directionLeft {
				direction = directionUp
			}
		case "F":
			if direction == directionUp {
				direction = directionRight
			} else if direction == directionLeft {
				direction = directionDown
			}
		case "7":
			if direction == directionRight {
				direction = directionDown
			} else if direction == directionUp {
				direction = directionLeft
			}
		case "|":
			if direction == directionUp {
				direction = directionUp
			} else if direction == directionDown {
				direction = directionDown
			}
		case "-":
			if direction == directionRight {
				direction = directionRight
			} else if direction == directionLeft {
				direction = directionLeft
			}
		}
	}
	fmt.Println("counter", counter/2)
}

func printGrid(grid [][]string, visisted map[int]map[int]bool) {
	for x, line := range grid {
		for y, char := range line {
			if visisted[x] != nil {
				if visisted[x][y] {
					fmt.Print("X")
					continue
				}
			}
			fmt.Print(char)
		}
		fmt.Println()
	}
}

func solve2(input [][]string) {
	pipes := make([][]string, len(input))
	start := []int{0, 0}
	// fmt.Println(pipes)
	for i, line := range input {
		for j, char := range line {
			// if char == "." {
			// 	continue
			// }
			if char == "S" {
				start = []int{i, j}
				// break
			}
			if pipes[i] == nil {
				pipes[i] = make([]string, 0)
			}
			pipes[i] = append(pipes[i], char)
		}
		// fmt.Println(i, line)
	}
	kind := getKindOfStartingPoint(start[0], start[1], pipes)
	if kind == "ERRORE" {
		panic("ERRORE")
	}
	// fmt.Println(start)
	// fmt.Println(kind)
	// fmt.Println(pipes)
	pipes[start[0]][start[1]] = kind

	// printGrid(pipes)
	// //allargo la matrice
	// for i := 0; i < len(pipes); i++ {

	// }

	direction := 0
	//0 => right, 1 => up, 2 => left, 3 => down
	currentX, currentY := start[0], start[1]
	coordinates := make(map[int]map[int]bool, 0)
	coordinates[currentX] = make(map[int]bool, 0)

	coordinates[currentX][currentY] = true
	switch kind {
	case "J", "L", "|":
		direction = directionUp
		// if kind == "J" || kind == "|" {
		// 	coordinates[currentX][currentY] = 1
		// }
		// coordinates = append(coordinates, []int{currentX, currentY})
		// fmt.Println("starting by going up")

	case "-":
		direction = directionRight
		// fmt.Println("starting by going right")
	case "F", "7":
		direction = directionDown

		// if kind == "7" {
		// 	coordinates[currentX][currentY] = 1
		// }

		// fmt.Println("starting by going down")
	}
	counter := 0
	// coordinatesX := make([]int, 0)
	// coordinatesY = append(coordinatesY, currentY)
	// coordinatesX = append(coordinatesX, currentX)
	for {
		counter++
		// fmt.Println("currentX", currentX, "currentY", currentY)
		switch direction {
		case directionRight:
			// fmt.Println("going right")
			currentY++
		case directionUp:
			// fmt.Println("going up")
			currentX--
		case directionLeft:
			// fmt.Println("going left")
			currentY--
		case directionDown:
			// fmt.Println("going down")
			currentX++
		}
		if currentX < 0 || currentX >= len(pipes) || currentY < 0 || currentY >= len(pipes[0]) {
			break
		}
		if currentX == start[0] && currentY == start[1] {
			// fmt.Println("back to start")
			break
		}
		// fmt.Println("currentX", currentX, "currentY", currentY)
		// fmt.Println("current char", pipes[currentX][currentY])
		if pipes[currentX][currentY] == "." {
			break
		}
		if coordinates[currentX] == nil {
			coordinates[currentX] = make(map[int]bool, 0)
		}
		coordinates[currentX][currentY] = true

		switch pipes[currentX][currentY] {
		case "J":
			if direction == directionRight {
				direction = directionUp
			} else if direction == directionDown {
				direction = directionLeft
			}
			// coordinatesX = append(coordinatesX, currentX)
			// coordinatesY = append(coordinatesY, currentY)
		case "L":
			if direction == directionDown {
				direction = directionRight
			} else if direction == directionLeft {
				direction = directionUp
			}
			// coordinates[currentX][currentY] = 1

			// coordinatesX = append(coordinatesX, currentX)
			// coordinatesY = append(coordinatesY, currentY)

		case "F":
			if direction == directionUp {
				direction = directionRight
			} else if direction == directionLeft {
				direction = directionDown
			}
			// coordinates[currentX][currentY] = 1

			// coordinatesX = append(coordinatesX, currentX)
			// coordinatesY = append(coordinatesY, currentY)
		case "7":
			if direction == directionRight {
				direction = directionDown
			} else if direction == directionUp {
				direction = directionLeft
			}
			// coordinates[currentX][currentY] = 1

			// coordinatesX = append(coordinatesX, currentX)
			// coordinatesY = append(coordinatesY, currentY)
			// coordinates[currentX][currentY] = 1
		case "|":
			if direction == directionUp {
				direction = directionUp
			} else if direction == directionDown {
				direction = directionDown
			}
			// coordinatesX = append(coordinatesX, currentX)
			// coordinatesY = append(coordinatesY, currentY)
			// coordinates[currentX][currentY] = 1
			// coordinates[currentX][currentY] = 1

		case "-":
			if direction == directionRight {
				direction = directionRight
			} else if direction == directionLeft {
				direction = directionLeft
			}
		}
	}

	//clean the matrix
	for i := 0; i < len(pipes); i++ {
		for j := 0; j < len(pipes[i]); j++ {
			if coordinates[i] != nil {
				if !coordinates[i][j] {
					pipes[i][j] = "."
				}
			} else {
				pipes[i][j] = "."
			}

		}
	}
	// printGrid(pipes, make(map[int]map[int]bool))
	// fmt.Println(coordinatesX)
	// fmt.Println(coordinatesY)
	// coord1, coord2 := 1,1
	// indexXBefore, indexXAfter := 0, 0
	// a := 0
	// coordSum := 0
	pointCounter := 0
	// for i := 0; i < len(input); i++ {
	// 	for j := 0; j < len(input[i]); j++ {
	// 		if input[i][j] == "." {
	// 			c := 0
	// 			if coordinates[i] != nil {
	// 				for k := range coordinates[i] {
	// 					// fmt.Println("k", k)
	// 					if k < j {
	// 						continue
	// 					}
	// 					c++
	// 				}
	// 				// if coordinates[i][j] == 1 {
	// 				// for k := 0; k < len(coordinates); k++ {
	// 				// 	for l := 0; l < len(coordinates[k]); l++ {
	// 				// 		if coordinates[k][l] < j {
	// 				// 			continue
	// 				// 		}
	// 				// 		c++
	// 				// 	}
	// 				// }
	// 			}
	// 			if c%2 == 1 {
	// 				// fmt.Println("i", i, "j", j)
	// 				pointCounter++
	// 			}
	// 		}
	// 	}
	// }
	// coordinatesX = []int{1, 5, 4, 3, 3, 1}
	// coordinatesY = []int{1, 1, 2, 3, 4, 4}
	// for i := 0; i < len(coordinatesY); i++ {
	// 	xi, yi := coordinatesX[i], coordinatesY[i]
	// 	xj, yj := 0, 0
	// 	if i == len(coordinatesY)-1 {
	// 		xj, yj = coordinatesX[0], coordinatesY[0]
	// 	} else {
	// 		xj, yj = coordinatesX[i+1], coordinatesY[i+1]
	// 	}
	// 	// fmt.Println("(", xi, "*", yj, ")-(", xj, "*", yi, ")")
	// 	coordSum += (xi * yj) - (xj * yi)
	// }
	// a := math.Abs(float64(coordSum)) / 2
	// fmt.Println("a", a)
	// fmt.Println("a*4", a*4)
	// fmt.Println("aaa", (a*4)-float64(counter))
	floodfill(0, 0, pipes, make(map[int]map[int]bool, 0))
	fmt.Println("counter", pointCounter)
}

func floodfill(currentX, currentY int, matrix [][]string, visited map[int]map[int]bool) {
	// fmt.Println("currentX", currentX, "currentY", currentY)

	if visited[currentX] == nil {
		visited[currentX] = make(map[int]bool, 0)
	}
	if visited[currentX][currentY] {
		// fmt.Println("already visited")
		return
	}
	visited[currentX][currentY] = true
	printGrid(matrix, visited)
	fmt.Println()

	if currentX > 0 {
		next := matrix[currentX-1][currentY]
		if next == "." { // || next == "|"|| next == "J" {
			floodfill(currentX-1, currentY, matrix, visited)
		}
	}
	if currentX < len(matrix)-1 {
		next := matrix[currentX+1][currentY]
		if next == "." {
			floodfill(currentX+1, currentY, matrix, visited)
		}
	}
	if currentY > 0 {
		next := matrix[currentX][currentY-1]
		if next == "." {
			// fmt.Println("going left")
			floodfill(currentX, currentY-1, matrix, visited)

		}
	}

	if currentY < len(matrix[0])-1 {
		next := matrix[currentX][currentY+1]
		if next == "." {
			// fmt.Println("going right")
			floodfill(currentX, currentY+1, matrix, visited)
		}
	}
	// fmt.Println()

}

func main() {
	matrix := readMatrix("test.txt")
	// solve1(matrix)
	solve2(matrix)
}


*/
