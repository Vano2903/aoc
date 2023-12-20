package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func readFile(filename string) [][]string {
	c, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	m := make([][]string, 0)
	for _, line := range strings.Split(string(c), "\n") {
		m = append(m, strings.Split(line, ""))
	}
	return m
}

func printMap(m [][]string) {
	for _, line := range m {
		fmt.Println(strings.Join(line, ""))
	}
}

func printMapWithSeen(m [][]string, seen map[string]bool) {
	for y, line := range m {
		for x, t := range line {
			if seen[fmt.Sprintf("%d,%d", x, y)] {
				fmt.Print("#")
			} else {
				fmt.Print(t)
			}
		}
		fmt.Println()
	}
}

func traverse(m [][]string, x, y, direction int, seen map[string]bool, cycleDetector map[string]bool, print bool) {
	if x < 0 || y < 0 || x >= len(m[0]) || y >= len(m) {
		// fmt.Println("out of bounds")
		return
	}
	seenKey := fmt.Sprintf("%d,%d", x, y)
	cycleKey := fmt.Sprintf("%d,%d,%d", x, y, direction)
	// if _, ok := seen[seenKey]; ok {
	// 	//we've been here before
	// 	fmt.Println("we have been here before")
	// 	return
	// }
	seen[seenKey] = true
	if print {
		switch direction {
		case 0:
			fmt.Println(x, y, "up")
		case 1:
			fmt.Println(x, y, "right")
		case 2:
			fmt.Println(x, y, "down")
		case 3:
			fmt.Println(x, y, "left")
		}
		printMapWithSeen(m, seen)
		fmt.Println()
		// fmt.Println(cycleDetector)
		time.Sleep(50 * time.Millisecond)
	}
	if _, ok := cycleDetector[cycleKey]; ok {
		if print {
			fmt.Println(cycleKey)
			fmt.Println("this is a cycle")
		}
		return
	}
	cycleDetector[cycleKey] = true

	if direction == 0 { //up
		tile := m[y][x]
		if tile == "." || tile == "|" {
			//keep going up
			traverse(m, x, y-1, direction, seen, cycleDetector, print)
		} else if tile == "-" {
			//go right and left
			traverse(m, x+1, y, 1, seen, cycleDetector, print)
			traverse(m, x-1, y, 3, seen, cycleDetector, print)
		} else if tile == "/" {
			//go right
			traverse(m, x+1, y, 1, seen, cycleDetector, print)
		} else if tile == "\\" {
			//go left
			traverse(m, x-1, y, 3, seen, cycleDetector, print)
		}
	} else if direction == 1 { //right
		tile := m[y][x]
		if tile == "." || tile == "-" {
			//keep going right
			traverse(m, x+1, y, direction, seen, cycleDetector, print)
		} else if tile == "|" {
			//go up and down
			traverse(m, x, y-1, 0, seen, cycleDetector, print)
			traverse(m, x, y+1, 2, seen, cycleDetector, print)
		} else if tile == "/" {
			//go up
			traverse(m, x, y-1, 0, seen, cycleDetector, print)
		} else if tile == "\\" {
			//go down
			traverse(m, x, y+1, 2, seen, cycleDetector, print)
		}
	} else if direction == 2 { //down
		tile := m[y][x]
		if tile == "." || tile == "|" {
			//keep going down
			traverse(m, x, y+1, direction, seen, cycleDetector, print)
		} else if tile == "-" {
			//go right and left
			traverse(m, x+1, y, 1, seen, cycleDetector, print)
			traverse(m, x-1, y, 3, seen, cycleDetector, print)
		} else if tile == "/" {
			//go left
			traverse(m, x-1, y, 3, seen, cycleDetector, print)
		} else if tile == "\\" {
			//go right
			traverse(m, x+1, y, 1, seen, cycleDetector, print)
		}
	} else if direction == 3 { //left
		tile := m[y][x]
		if tile == "." || tile == "-" {
			//keep going left
			traverse(m, x-1, y, direction, seen, cycleDetector, print)
		} else if tile == "|" {
			//go up and down
			traverse(m, x, y-1, 0, seen, cycleDetector, print)
			traverse(m, x, y+1, 2, seen, cycleDetector, print)
		} else if tile == "/" {
			//go down
			traverse(m, x, y+1, 2, seen, cycleDetector, print)
		} else if tile == "\\" {
			//go up
			traverse(m, x, y-1, 0, seen, cycleDetector, print)
		}
	}
}

func solve1(m [][]string) {
	// printMap(m)
	cycleDetector := make(map[string]bool) //"x,y,direction"
	seen := make(map[string]bool)          //"x,y"
	traverse(m, 0, 0, 1, seen, cycleDetector, false)
	fmt.Println(len(seen))
	// printMapWithSeen(m, seen)
}

func solve2(m [][]string) {
	pos := make([]int, 0) //x,y,direction
	for y, line := range m {
		for x := range line {
			if y == 0 {
				pos = append(pos, []int{x, y, 2}...) //down
			} else if y == len(m)-1 {
				pos = append(pos, []int{x, y, 0}...) //up
			} else if x == 0 {
				pos = append(pos, []int{x, y, 1}...) //right
			} else if x == len(line)-1 {
				pos = append(pos, []int{x, y, 3}...) //left
			}
		}
	}

	highest := 0
	for i := 0; i < len(pos)-3; i += 3 {
		print := false
		// if pos[i] == 3 && pos[i+1] == 0 {
		// 	print = true
		// }
		cycleDetector := make(map[string]bool) //"x,y,direction"
		seen := make(map[string]bool)
		traverse(m, pos[i], pos[i+1], pos[i+2], seen, cycleDetector, print)
		// fmt.Println("from", pos[i], pos[i+1], pos[i+2], "seen", len(seen))
		if len(seen) > highest {
			highest = len(seen)
		}
	}

	// traverse(m, 0, 0, 1, seen, cycleDetector)
	fmt.Println(highest)

}

func main() {
	l := readFile("input")
	solve1(l)
	solve2(l)
}
