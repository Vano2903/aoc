package main

import (
	"os"
)

func readFile(filename string) []byte {
	c, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return c
}

func solve1(content []byte) {
	houses := make(map[int]map[int]bool)
	x, y := 0, 0
	houses[x] = make(map[int]bool)
	houses[x][y] = true
	for _, c := range content {
		switch c {
		case '<':
			y--
			if houses[x] == nil {
				houses[x] = make(map[int]bool)
			}
			houses[x][y] = true
		case '>':
			y++
			if houses[x] == nil {
				houses[x] = make(map[int]bool)
			}
			houses[x][y] = true
		case '^':
			x--
			if houses[x] == nil {
				houses[x] = make(map[int]bool)
			}
			houses[x][y] = true
		case 'v':
			x++
			if houses[x] == nil {
				houses[x] = make(map[int]bool)
			}
			houses[x][y] = true
		}
	}
	housesCount := 0
	for _, v := range houses {
		for _, housesY := range v {
			if housesY {
				housesCount++
			}
		}
	}
	println(housesCount)
}

func solve2(content []byte) {
	houses := make(map[int]map[int]bool)
	x, y := 0, 0
	santaX, santaY := 0, 0
	robotX, robotY := 0, 0
	houses[x] = make(map[int]bool)
	houses[x][y] = true
	for i, c := range content {
		if i%2 == 0 {
			x, y = santaX, santaY
		} else {
			x, y = robotX, robotY
		}
		switch c {
		case '<':
			y--
			if houses[x] == nil {
				houses[x] = make(map[int]bool)
			}
			houses[x][y] = true
		case '>':
			y++
			if houses[x] == nil {
				houses[x] = make(map[int]bool)
			}
			houses[x][y] = true
		case '^':
			x--
			if houses[x] == nil {
				houses[x] = make(map[int]bool)
			}
			houses[x][y] = true
		case 'v':
			x++
			if houses[x] == nil {
				houses[x] = make(map[int]bool)
			}
			houses[x][y] = true
		}
		if i%2 == 0 {
			santaX, santaY = x, y

		} else {
			robotX, robotY = x, y
		}
	}
	housesCount := 0
	for _, v := range houses {
		for _, housesY := range v {
			if housesY {
				housesCount++
			}
		}
	}
	println(housesCount)
}

func main() {
	c := readFile("input.txt")
	// solve1(c)
	solve2(c)
}
