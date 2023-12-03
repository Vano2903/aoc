package main

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

func solve1(input string) string {
	// fmt.Println("input is", input)
	// output := ""
	var buffer bytes.Buffer
	l := len(input)
	// c := input
	lastSeen := input[0]
	counter := 0
	for i := 0; i < l; i++ {
		// fmt.Println("i is", i, "l is", l)
		// r := rune(input[0])
		b := input[0]
		// fmt.Println("r is", string(r), "lastSeen is", string(lastSeen), "counter is", counter)
		// fmt.Print("checking if ", string(r), " is a number? ")
		if b == lastSeen {
			// fmt.Println("it is")
			counter++
		} else {
			// fmt.Println("is is not")
			// fmt.Printf("new char, adding %d%s to output\n", counter, string(lastSeen))
			buffer.WriteString(strconv.Itoa(counter))
			buffer.WriteByte(lastSeen)
			//  += fmt.Sprintf("%d%s", counter, string(lastSeen))
			lastSeen = b
			counter = 1
		}

		input = input[1:]
		// fmt.Println("line is now", line)
	}
	// output += fmt.Sprintf("%d%s", counter, string(lastSeen))
	buffer.WriteString(strconv.Itoa(counter))
	buffer.WriteByte(lastSeen)

	// fmt.Println("output of", c, "is", output)
	// fmt.Println()
	return buffer.String()
}

func main() {
	// input := "1"
	// inputs := []string{"1", "11", "21", "1211", "111221"}
	// for _, input := range inputs {
	// solve1(input)
	input := "1321131112"
	start := time.Now()
	a := solve1(input)
	fmt.Println("1) len of a is", len(a))
	for i := 0; i < 49; i++ {
		a = solve1(a)
		fmt.Printf("%d) len of a is %d\n", i+2, len(a))
	}
	fmt.Println("part 1:", len(a))
	fmt.Println("time taken:", time.Since(start))
	// }
}
