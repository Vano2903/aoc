package main

import (
	"fmt"
	"math"
)

func splitInt(toSplit int, sequence []int) []int {
	if toSplit != 0 {
		i := toSplit % 10
		sequence = append([]int{i}, sequence...)
		return splitInt(toSplit/10, sequence)
	}
	return sequence
}

// 0 air
// 1 sand source
// 2 sand
// 3 rock
// 4 sand path
type cave struct {
	mapping [][]int
	originX int
	originY int
}

func (c cave) countSand() int {
	counter := 0
	for _, line := range c.mapping {
		for _, l := range line {
			if l == 2 {
				counter++
			}
		}
	}
	return counter
}

func (c *cave) traceRock(xStart, yStart, xEnd, yEnd int) {
	xMovementLength := int(math.Abs(float64(xEnd - xStart)))
	yMovementLength := int(math.Abs(float64(yEnd - yStart)))
	// fmt.Println("mov x:", xMovementLength)
	// fmt.Println("mov y:", yMovementLength)
	// type point struct {
	// 	x int
	// 	y int
	// }

	fmt.Println(xStart)
	fmt.Println(yStart)
	fmt.Println(xEnd)
	fmt.Println(yEnd)
	fmt.Println()

	// coords := make([]point, 0)

	//horizontal line
	if (xMovementLength > 0) && (yMovementLength == 0) {
		for x := 0; x <= xMovementLength; x++ {
			xCoord := 0
			if xEnd > xStart {
				xCoord = (xStart - c.originX) + x
			} else {
				xCoord = (xEnd - c.originX) + x
			}
			// fmt.Println("h x:", xCoord, "y:", xStart-c.originX)
			// fmt.Println("  x:", xStart-c.originX, "y:", yStart)

			// coords = append(coords, point{x: c.originX - xStart, y: c.originY - yCoord})
			c.mapping[yStart][xCoord] = 3
		}
	} else if (yMovementLength > 0) && (xMovementLength == 0) {
		for y := 0; y <= yMovementLength; y++ {
			yCoord := 0
			if yEnd > yStart {
				yCoord = (yStart) + y // - c.originY
			} else {
				yCoord = (yEnd) + y // - c.originY
			}
			// fmt.Println("h x:", yCoord, "y:", xStart-c.originX)
			// fmt.Println("  x:", xStart-c.originX, "y:", yStart)

			// coords = append(coords, point{x: c.originX - xStart, y: c.originY - yCoord})
			c.mapping[yCoord][xStart-c.originX] = 3
		}
	} else {
		//oblique line
		for x := 0; x < xMovementLength; x++ {
			for y := 0; x < yMovementLength; y++ {
				xCoord := 0
				if xEnd > xStart {
					xCoord = xStart + x
				} else {
					xCoord = xEnd + x
				}

				yCoord := 0
				if yEnd > yStart {
					yCoord = yStart + y
				} else {
					yCoord = yEnd + y
				}
				// fmt.Println("o x:", xCoord-c.originX, "y:", yCoord-c.originY)

				c.mapping[xCoord-c.originX][yCoord-c.originY] = 3
			}
		}
	}
}

func (c cave) printMap() {
	//print origin x in vetical
	verticalOriginX := splitInt(c.originX, []int{})
	for _, v := range verticalOriginX {
		fmt.Println("  ", v)
	}

	for i, line := range c.mapping {
		if i > 99 {
			fmt.Print(i)
		} else if i > 9 {
			fmt.Print(i, " ")
		} else {
			fmt.Print(i, "  ")
		}
		for _, l := range line {
			switch l {
			case 0:
				fmt.Print(".")
			case 1:
				fmt.Print("+")
			case 2:
				fmt.Print("o")
			case 3:
				fmt.Print("#")
			case 4:
				fmt.Print("~")
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
	fmt.Println()
}
