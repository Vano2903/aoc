package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(path string) string {
	c, err := os.ReadFile(path)
	if err != nil {

		panic(err)
	}
	return string(c)
}

func solve1(content string) {
	lines := strings.Split(content, "\n")
	t := strings.Split(lines[0], " ")
	d := strings.Split(lines[1], " ")
	times := []int{}
	for _, v := range t {
		time, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		times = append(times, time)
	}
	distances := []int{}
	for _, v := range d {
		distance, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		distances = append(distances, distance)
	}

	fmt.Println(times)
	fmt.Println(distances)

	total := 1

	// charged := 0
	for i, v := range times {
		winCounter := 0
		for j := 0; j <= v; j++ {
			travelled := j * (v - j)
			// fmt.Println("the boat charged for", j, "ms travelled", travelled, "mm")

			if travelled > distances[i] {
				winCounter++
			}
		}
		fmt.Println("winCounter", winCounter)
		total *= winCounter
	}
	fmt.Println("total", total)
}

func solve2(content string) {
	lines := strings.Split(content, "\n")
	t := strings.Split(lines[0], " ")
	d := strings.Split(lines[1], " ")
	times := []string{}
	for _, v := range t {
		_, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		times = append(times, v)
	}
	distances := []string{}
	for _, v := range d {
		_, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		distances = append(distances, v)
	}

	ts := strings.Join(times, "")
	ds := strings.Join(distances, "")
	// fmt.Println(ts)
	// fmt.Println(ds)
	time, _ := strconv.Atoi(ts)
	distance, _ := strconv.Atoi(ds)

	fmt.Println(time)
	fmt.Println(distance)

	total := 1

	// charged := 0
	// for i, v := range times {
	winCounter := 0
	for j := 0; j <= time; j++ {
		travelled := j * (time - j)
		// fmt.Println("the boat charged for", j, "ms travelled", travelled, "mm")

		if travelled > distance {
			winCounter++
		}
	}
	fmt.Println("winCounter", winCounter)
	total *= winCounter
	// }
	fmt.Println("total", total)
}

func main() {
	c := readFile("input.txt")
	solve2(c)
}
