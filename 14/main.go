package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"strconv"
// 	"strings"
// )

// type coord struct {
// 	x, y int
// }

// func main() {
// 	f, _ := ioutil.ReadFile("inputtest")
// 	content := strings.Split(string(f), "\n")

// 	// cave := make([][]int, 0)

// 	smallestX := -1
// 	smallestY := -1
// 	biggestX := -1
// 	biggestY := -1

// 	for _, c := range content {
// 		c = strings.TrimSpace(c)
// 		//498,4 -> 498,6 -> 496,6
// 		coords := strings.Split(c, " -> ")
// 		for _, cd := range coords {
// 			xy := strings.Split(cd, ",")
// 			tmpx, _ := strconv.Atoi(xy[0])
// 			tmpy, _ := strconv.Atoi(xy[1])
// 			if tmpx < smallestX || smallestX == -1 {
// 				smallestX = tmpx
// 			}
// 			if tmpy < smallestY || smallestY == -1 {
// 				smallestY = tmpy
// 			}

// 			if tmpx > biggestX || biggestX == -1 {
// 				biggestX = tmpx
// 			}
// 			if tmpy > biggestY || biggestY == -1 {
// 				biggestY = tmpy
// 			}
// 		}
// 	}

// 	fmt.Println("biggesY:", biggestY)

// 	// smallestX -= 400
// 	// smallestX = 490
// 	// biggestX = 500
// 	// biggestY += 2

// 	smallestX--
// 	biggestX++
// 	biggestY++

// 	caveLength := (biggestX - smallestX) + 2 //+ 17 //- smallestX
// 	caveHeight := (biggestY) + 3             //+ 17
// 	fmt.Println("after analyzing cave system the further left rock is at x:", smallestX, "y:", smallestY)
// 	fmt.Println("we think the optimal cave size should be:", caveHeight, "long and", caveLength, "tall")

// 	mapping := make([][]int, 0)

// 	for x := 0; x < caveHeight; x++ {
// 		line := make([]int, caveLength)
// 		for i := 0; i < caveLength; i++ {
// 			line[i] = 0
// 		}
// 		mapping = append(mapping, line)
// 	}

// 	mapping[0][500-smallestX] = 1

// 	cave := cave{
// 		mapping: mapping,
// 		originX: smallestX,
// 		originY: smallestY,
// 	}

// 	for _, c := range content {
// 		c = strings.TrimSpace(c)
// 		numCoords := strings.Split(c, " -> ")
// 		p1 := strings.Split(numCoords[0], ",")
// 		tmpxp1, _ := strconv.Atoi(p1[0])
// 		tmpyp1, _ := strconv.Atoi(p1[1])
// 		lastCoord := coord{}
// 		lastCoord.x = tmpxp1
// 		lastCoord.y = tmpyp1

// 		for i := 0; i < len(numCoords)-1; i++ {
// 			p2 := strings.Split(numCoords[i+1], ",")
// 			tmpxp2, _ := strconv.Atoi(p2[0])
// 			tmpyp2, _ := strconv.Atoi(p2[1])
// 			cave.traceRock(lastCoord.x, lastCoord.y, tmpxp2, tmpyp2)
// 			lastCoord.x = tmpxp2
// 			lastCoord.y = tmpyp2
// 		}
// 	}

// 	//p2
// 	// for j := 0; j < 2; j++ {
// 	// 	line := make([]int, caveHeight)
// 	// 	for i := 0; i < caveHeight; i++ {
// 	// 		if j == 0 {
// 	// 			line[i] = 0
// 	// 		} else {
// 	// 			line[i] = 3
// 	// 		}
// 	// 	}
// 	// 	cave.mapping = append(cave.mapping, line)
// 	// }

// 	sand := sand{
// 		cave:     &cave,
// 		x:        500 - smallestX,
// 		y:        0,
// 		path:     make([]coord, 0),
// 		lastSand: nil,
// 	}

// 	sand.path = append(sand.path, coord{x: 500 - smallestX, y: 0})
// 	cave.printMap()

// 	// sand.DoMoveP2()
// 	sand.DoMove()

// 	// for _, c := range content {
// 	// 	c = strings.TrimSpace(c)

// 	// }
// }
