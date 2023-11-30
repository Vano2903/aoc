package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var size = 1000

func readLines(filename string) []string {
	c, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(c), "\n")
}

func convertToInt(s string) (int, int) {
	split := strings.Split(s, ",")
	x, _ := strconv.Atoi(split[0])
	y, _ := strconv.Atoi(split[1])
	return x, y
}

func on(grid [][]int, start, end string) {
	sX, sY := convertToInt(start)
	eX, eY := convertToInt(end)
	// fmt.Printf("turning on lights from %dx,%dy to %dx,%dy\n", sX, sY, eX, eY)
	for i := sX; i <= eX; i++ {
		for j := sY; j <= eY; j++ {
			grid[j][i]++
		}
	}
	// fmt.Println(grid)
}
func off(grid [][]int, start, end string) {
	sX, sY := convertToInt(start)
	eX, eY := convertToInt(end)

	for i := sX; i <= eX; i++ {
		for j := sY; j <= eY; j++ {
			if grid[j][i] > 0 {

				grid[j][i]--
			}
		}
	}
}
func toggle(grid [][]int, start, end string) {
	sX, sY := convertToInt(start)
	eX, eY := convertToInt(end)

	for i := sX; i <= eX; i++ {
		for j := sY; j <= eY; j++ {
			grid[j][i] += 2
		}
	}
}

func initGrid() [][]int {
	grid := make([][]int, size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			grid[i] = make([]int, size)
			grid[i][j] = 0
		}
	}
	return grid
}

func printGrid(grid [][]int) {
	for i := 0; i < size; i++ {
		fmt.Println(grid[i])
	}
}

func solve1(l []string) {
	grid := initGrid()
	for _, line := range l {
		if line == "" {
			continue
		}
		values := strings.Split(line, " ")
		end := values[len(values)-1]
		start := values[len(values)-3]
		var command string
		if values[0] == "turn" {
			command = values[1]
		} else {
			command = values[0]
		}
		switch command {
		case "on":
			on(grid, start, end)
		case "off":
			off(grid, start, end)
		case "toggle":
			toggle(grid, start, end)
		}
		// printGrid(grid)
		// time.Sleep(100 * time.Millisecond)
	}
	count := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			count += grid[i][j]
		}
	}
	println(count)
}
func main() {
	l := readLines("input.txt")
	solve1(l)
}
