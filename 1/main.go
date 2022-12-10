package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input")

	lines := strings.Split(string(content), "\n")

	max := []int{0, 0, 0}
	currentCounter := 0
	for i, line := range lines {
		if len(line) == 0 {
			fmt.Printf("checking %d against %d\n", currentCounter, max)
			for j, v := range max {
				if currentCounter > v {
					max = append(max[:j+1], max[j:]...)
					max[j] = currentCounter
					fmt.Printf("found a new max in line %d\n", i)
					break
				}
			}
			currentCounter = 0
		} else {
			val, _ := strconv.Atoi(line)
			currentCounter += val
		}

	}

	fmt.Println(max)
	total := 0
	for _, v := range max[0:3] {
		fmt.Println("summing: ", v)
		total += v
	}
	fmt.Println(total)
}
