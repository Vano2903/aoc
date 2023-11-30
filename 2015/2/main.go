package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) []byte {
	c, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return c
}

func readLine(filename string) []string {
	c := readFile(filename)
	lines := strings.Split(string(c), "\n")
	return lines
}

func resultLine(w, l, h int) int {
	c1 := l * w
	c2 := w * h
	c3 := h * l
	total := (c1 * 2) + (c2 * 2) + (c3 * 2) + min(c1, c2, c3)
	return total
}

func solve1(l []string) {
	//2*l*w + 2*w*h + 2*h*l
	//lxwxh
	final := 0
	for _, line := range l {
		if line == "" {
			continue
		}
		dimensions := strings.Split(line, "x")
		l, _ := strconv.Atoi(dimensions[0])
		w, _ := strconv.Atoi(dimensions[1])
		h, _ := strconv.Atoi(dimensions[2])
		total := resultLine(l, w, h)
		final += total
	}
	fmt.Println("solution 1:")
	fmt.Println(final)
}

func solve2(l []string) {
	final := 0
	for _, line := range l {
		if line == "" {
			continue
		}
		dimensions := strings.Split(line, "x")
		l, _ := strconv.Atoi(dimensions[0])
		w, _ := strconv.Atoi(dimensions[1])
		h, _ := strconv.Atoi(dimensions[2])
		// total := resultLine(l, w, h)
		shortest1 := min(l, w, h)
		shortest2 := 0
		switch shortest1 {
		case l:
			shortest2 = min(w, h)
		case w:
			shortest2 = min(l, h)
		case h:
			shortest2 = min(l, w)
		}
		ribbon := (shortest1*2 + shortest2*2) + (l * w * h)
		final += ribbon
	}
	fmt.Println("solution 2:")
	fmt.Println(final)
}

func main() {
	c := readLine("input.txt")
	solve1(c)
	solve2(c)
}
