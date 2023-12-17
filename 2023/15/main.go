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

func hashPart(part string) int {
	current := 0
	for _, c := range part {
		current += int(c)
		current *= 17
		current = current % 256
	}
	return current
}

func solve1(line string) {
	total := 0
	parts := strings.Split(line, ",")
	for _, p := range parts {
		total += hashPart(p)
	}

	fmt.Println("total is:", total)
}

func printBoxes(boxes [][]lens) {
	for i, boxes := range boxes {
		if len(boxes) == 0 {
			continue
		}
		c := true
		for _, b := range boxes {
			if b.num != 0 {
				c = false
				break
			}
		}
		if c {
			continue
		}
		fmt.Printf("Box: %d:", i)
		for _, b := range boxes {
			if b.num == 0 {
				continue
			}
			fmt.Printf("[%s %d]", b.key, b.num)
		}
		fmt.Println()
	}
}

type lens struct {
	key string
	num int
}

func calculatePower(boxes [][]lens) int {
	total := 0
	for i, boxes := range boxes {
		position := 0
		for _, b := range boxes {
			if b.num == 0 {
				continue
			}
			total += (i + 1) * (position + 1) * b.num
			position++
		}
	}
	return total
}

func searchInBoxes(boxes []lens, lensKey string) int {
	for i, b := range boxes {
		if b.key == lensKey {
			return i
		}
	}
	return -1
}

func solve2v2(line string) {
	boxes := make([][]lens, 256)
	// currentIndex := 0

	for _, p := range strings.Split(line, ",") {
		elements := strings.Split(p, "")
		if elements[len(elements)-1] == "-" {
			key := strings.Join(elements[:len(elements)-1], "")
			hashKey := hashPart(key)
			if x := searchInBoxes(boxes[hashKey], key); x != -1 {
				boxes[hashKey][x].num = 0
			} else {
				boxes[hashKey] = append(boxes[hashKey], lens{key: key, num: 0})
			}
		} else {
			num, _ := strconv.Atoi(elements[len(elements)-1])
			key := strings.Join(elements[:len(elements)-2], "")
			hashKey := hashPart(key)
			if x := searchInBoxes(boxes[hashKey], key); x != -1 {
				if boxes[hashKey][x].num == 0 {
					//if it exists and
					//remove from boxes
					boxes[hashKey] = append(boxes[hashKey][:x], boxes[hashKey][x+1:]...)
					boxes[hashKey] = append(boxes[hashKey], lens{key: key, num: num})
				} else {
					//update
					boxes[hashKey][x].num = num
				}
			} else {
				//add
				boxes[hashKey] = append(boxes[hashKey], lens{key: key, num: num})
			}
		}
		// printBoxes(boxes)
		// fmt.Println()
	}
	fmt.Println("power", calculatePower(boxes))

}
func main() {
	line := readFile("input.txt")
	solve1(line)
	solve2v2(line)
}

//too high 
//4667352
//236057