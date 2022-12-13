package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

type point struct {
	x                 int
	height            int
	y                 int
	fromDir           string
	exludedDirections []string
	// excludedPoint     []point
	// visited           []point
}

func (p *point) excludeDirection(dir string) {
	p.exludedDirections = append(p.exludedDirections, dir)
}

func (p point) isDirectionExluded(dir string) bool {
	for _, ex := range p.exludedDirections {
		if ex == dir {
			return true
		}
	}
	return false
}

// func (p *point) excludePoint(toExclude point) {
// 	b.excludedPoint = append(b.excludedPoint, toExclude)
// }

// func (p point) isExcludedPoint(toCheck point) bool {
// 	for _, ex := range b.excludedPoint {
// 		if ex.x == toCheck.x && ex.y == toCheck.y {
// 			return true
// 		}
// 	}
// 	return false
// }

// func (p *point) addLastVisited(lastVisited point) {
// 	asaganaieie := point{
// 		x:      lastVisited.x,
// 		y:      lastVisited.y,
// 		height: lastVisited.height,
// 	}
// 	b.visited = append(b.visited, asaganaieie)
// }

// func (p point) isVisited(toCheck point) bool {
// 	for _, ex := range b.visited {
// 		if ex.x == toCheck.x && ex.y == toCheck.y {
// 			return true
// 		}
// 	}
// 	return false
// }

func (p *point) move(incx, incy, height int, dir string) {
	// b.addLastVisited(*p)
	p.x += incx
	p.y += incy
	p.height = height
	p.fromDir = dir
	p.exludedDirections = []string{}
}

// func (p *point) undo() {
// 	// b.excludePoint(*p)
// 	b.x = b.visited[len(b.visited)-1].x
// 	b.y = b.visited[len(b.visited)-1].y
// 	b.height = b.visited[len(b.visited)-1].height
// 	b.visited = b.visited[:len(b.visited)-1]
// }

type basculator struct {
	start      point
	end        point
	current    *point
	visited    []*point
	heatmap    [][]int
	pathLength int
}

func (b basculator) isVisited(toCheck point) bool {
	for _, ex := range b.visited {
		if ex.x == toCheck.x && ex.y == toCheck.y {
			return true
		}
	}
	return false
}

func (b *basculator) addVisited(toAdd point) {
	b.visited = append(b.visited, &toAdd)
}

func (b basculator) calculateHeightDiffFromCurrent(toCheck point) bool {
	// heightDiff := (b.heatmap[toCheck.y][toCheck.x] - b.current.height)
	// return heightDiff >= 0 && heightDiff <= 1

	heightDiff := (b.heatmap[toCheck.y][toCheck.x] - b.current.height)
	if b.heatmap[b.current.y][b.current.x] == 14 && b.heatmap[toCheck.y][toCheck.x] == 12 {
		return true
	} else {
		return heightDiff >= 0 && heightDiff <= 1
	}
}

func (b basculator) canGoRight() bool {
	if b.current.x+1 >= len(b.heatmap[b.current.y]) {
		//fmt.Println("  E: going right would exit the map")
		return false
	}

	toCheck := point{x: b.current.x + 1, y: b.current.y}

	// if b.current.isExcludedPoint(toCheck) {
	// 	//fmt.Println("  E: going right is an exluded movement")
	// 	return false
	// }

	if b.current.isDirectionExluded("right") {
		//fmt.Println("  E: going right is an exluded direction")
		return false
	}

	if b.isVisited(toCheck) {
		//fmt.Println("  E: going right would go back to the last visited point")
		return false
	}

	return b.calculateHeightDiffFromCurrent(toCheck)
}

func (b *basculator) moveRight() {
	b.addVisited(*b.current)
	b.pathLength++

	b.current.move(1, 0, b.heatmap[b.current.y][b.current.x+1], "right")
	//fmt.Println(">")
}

func (b basculator) canGoDown() bool {
	if b.current.y+1 >= len(b.heatmap) {
		//fmt.Println("  E: going down would exit the map")
		return false
	}

	toCheck := point{x: b.current.x, y: b.current.y + 1}

	// if b.current.isExcludedPoint(toCheck) {
	// 	//fmt.Println("  E: going down is an exluded movement")
	// 	return false
	// }

	if b.current.isDirectionExluded("down") {
		//fmt.Println("  E: going down is an exluded direction")
		return false
	}

	if b.isVisited(toCheck) {
		//fmt.Println("  E: going down would go back to the last visited point")
		return false
	}

	return b.calculateHeightDiffFromCurrent(toCheck)
}

func (b *basculator) moveDown() {
	b.addVisited(*b.current)
	b.pathLength++

	b.current.move(0, 1, b.heatmap[b.current.y+1][b.current.x], "down")
	//fmt.Println("v")
}

func (b basculator) canGoLeft() bool {
	if b.current.x-1 < 0 {
		//fmt.Println("  E: going left would exit the map")
		return false
	}

	toCheck := point{x: b.current.x - 1, y: b.current.y}

	// if b.current.isExcludedPoint(toCheck) {
	// 	//fmt.Println("  E: going left is an exluded movement")
	// 	return false
	// }

	if b.current.isDirectionExluded("left") {
		//fmt.Println("  E: going left is an exluded direction")
		return false
	}

	if b.isVisited(toCheck) {
		//fmt.Println("  E: going left would go back to the last visited point")
		return false
	}

	return b.calculateHeightDiffFromCurrent(toCheck)
}

func (b *basculator) moveLeft() {
	b.addVisited(*b.current)
	b.pathLength++

	b.current.move(-1, 0, b.heatmap[b.current.y][b.current.x-1], "left")
	//fmt.Println("<")
}

func (b basculator) canGoUp() bool {
	if b.current.y-1 < 0 {
		//fmt.Println("  E: going up would exit the map (", b.current.y-1, ")")
		return false
	}

	toCheck := point{x: b.current.x, y: b.current.y - 1}

	// if b.current.isExcludedPoint(toCheck) {
	// 	//fmt.Println("  E: going up is an exluded movement")
	// 	return false
	// }

	if b.current.isDirectionExluded("up") {
		//fmt.Println("  E: going up is an exluded direction")
		return false
	}

	if b.isVisited(toCheck) {
		//fmt.Println("  E: going up would go back to the last visited point")
		return false
	}

	return b.calculateHeightDiffFromCurrent(toCheck)
}

func (b *basculator) moveUp() {
	b.addVisited(*b.current)
	b.pathLength++

	b.current.move(0, -1, b.heatmap[b.current.y-1][b.current.x], "up")
	//fmt.Println("^")
}

func (b *basculator) undo() {
	b.current = b.visited[len(b.visited)-1]
	b.visited = b.visited[:len(b.visited)-1]
	b.pathLength--
}

func (b basculator) reachedDestination() bool {
	//fmt.Println("checking if reached destination")
	//fmt.Println("  current", b.current.x, b.current.y, b.current.height)
	//fmt.Println("  end", b.end.x, b.end.y)
	return b.current.x == b.end.x && b.current.y == b.end.y
}

func (b basculator) printSourroundings() {
	//fmt.Println("current", b.current.x, b.current.y)
	if b.canGoUp() {
		//fmt.Println("  up", b.heatmap[b.current.y-1][b.current.x])
	} else {
		//fmt.Println("  up", "no")
	}

	if b.canGoDown() {
		//fmt.Println("  down", b.heatmap[b.current.y+1][b.current.x])
	} else {
		//fmt.Println("  down", "no")
	}

	if b.canGoLeft() {
		//fmt.Println("  left", b.heatmap[b.current.y][b.current.x-1])
	} else {
		//fmt.Println("  left", "no")
	}

	if b.canGoRight() {
		//fmt.Println("  right", b.heatmap[b.current.y][b.current.x+1])
	} else {
		//fmt.Println("  right", "no")
	}
}

func (b basculator) printHeatmap() {
	for x := 0; x < len(b.heatmap); x++ {
		for y := 0; y < len(b.heatmap[x]); y++ {
			if b.current.x == y && b.current.y == x {
				//fmt.Print(" . ")
			} else {
				if b.heatmap[x][y] >= 10 {
					//fmt.Print(b.heatmap[x][y])
					//fmt.Print(" ")
				} else {
					//fmt.Print(" ", b.heatmap[x][y])
					//fmt.Print(" ")
				}
			}
		}
		//fmt.Println()
	}
}

func (b *basculator) p1() int {
	bestPath := -1
	var counter uint64
	for {
		// ////fmt.Println("visited:", b.current.visited[0])
		////fmt.Println("user is at", b.current.x, b.current.y, "with height", b.current.height)
		// b.printHeatmap()

		// //fmt.Println()
		// b.printSourroundings()
		// //fmt.Println()
		// b.printSourroundings()
		//fmt.Println(b.current.exludedDirections)
		if b.current.x == b.start.x && b.current.y == b.start.y {
			if !b.canGoUp() && !b.canGoDown() && !b.canGoLeft() && !b.canGoRight() {
				//fmt.Println("  W: FORSE HO FINITO")
				return bestPath
			}
		}
		if b.reachedDestination() {
			//save path
			if bestPath == -1 || b.pathLength < bestPath {
				bestPath = b.pathLength
			}
			// //fmt.Println("okay i am removing going", lastDir, "from the last visited point (", b.current.visited[len(b.current.visited)-2], ")")
			if counter%100000 == 0 {
				fmt.Println("hey i found a path :) (", counter/100000, ")")

			}
			counter++
			//fmt.Println("  path length:", b.pathLength)
			//erease current path
			b.visited[len(b.visited)-1].excludeDirection(b.current.fromDir)

			b.undo()
			//fmt.Println(b.current.exludedDirections)

			//fmt.Println("==========================================================================================")
			// time.Sleep(1 * time.Second)
		} else if b.canGoRight() {
			//go right
			b.moveRight()
		} else if b.canGoDown() {
			//go down
			b.moveDown()

		} else if b.canGoUp() {
			//go up
			b.moveUp()

		} else if b.canGoLeft() {
			//go left
			b.moveLeft()

		} else {
			//fmt.Println("non posso ne scendere ne salire... NE SCENDERE... NE SALIRE!")
			// no more moves
			// going back to last visited and excluding this move from the next moves

			//erease current path
			b.visited[len(b.visited)-1].excludeDirection(b.current.fromDir)
			b.undo()
			//fmt.Println(b.current.exludedDirections)
		}
		// time.Sleep(40 * time.Millisecond)
		// cmd := exec.Command("clear") //Linux example, its tested
		// cmd.Stdout = os.Stdout
		// cmd.Run()
	}
}

func main() {
	f, _ := ioutil.ReadFile("input")
	content := strings.Split(string(f), "\n")
	heatmap := make([][]int, 0)
	start, end := point{}, point{}
	for x, c := range content {
		c = strings.TrimSpace(c)
		line := strings.Split(c, "")
		heatline := make([]int, 0)
		for y, l := range line {
			val := (l[0] - 'a')

			if unicode.IsUpper(rune(l[0])) {
				if l == "S" {
					start.x = y
					start.y = x
					start.height = 0
					val = 0
				}
				if l == "E" {
					end.x = y
					end.y = x
					end.height = 25
					val = 25
				}
			}
			heatline = append(heatline, int(val))

		}
		heatmap = append(heatmap, heatline)
	}
	////fmt.Println(heatmap)
	////fmt.Println(start)
	b := basculator{start: start, end: end, current: &start, heatmap: heatmap}
	fmt.Println("OI MEGA MAGICO (FORSE):", b.p1())

}
