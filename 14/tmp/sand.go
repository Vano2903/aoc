package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var (
	isEnd = false
)

type sand struct {
	cave     *cave
	path     []coord
	lastSand *sand
	x        int
	y        int
}

func (s sand) canGoDown() bool {
	if len(s.cave.mapping) < s.y+1 {
		return true
	}
	return s.cave.mapping[s.y+1][s.x] == 0
}

func (s sand) canGoLeft() bool {
	if len(s.cave.mapping) < s.y+1 {
		return false
	}

	if s.x-1 < 0 {
		return true
	}

	return s.cave.mapping[s.y+1][s.x-1] == 0
}

func (s sand) canGoRight() bool {
	if len(s.cave.mapping) < s.y+1 {
		return false
	}

	if len(s.cave.mapping[s.y+1]) < s.x+1 {
		return true
	}

	return s.cave.mapping[s.y+1][s.x+1] == 0
}

func (s sand) isGoingDownTheEnd() bool {
	for _, line := range s.cave.mapping[s.y:] {
		for _, l := range line {
			if l != 0 {
				return false
			}
		}
	}
	return true
}

// move returns true if it was able to move, false if it can't move anymore
func (s *sand) move() bool {
	s.cave.mapping[s.y][s.x] = 0

	if s.canGoDown() {

		// if s.y+1 >= len(s.cave.mapping) {
		// 	s.cave.mapping[s.y][s.x] = 2
		// 	s.path = append(s.path, coord{x: s.x, y: s.y})
		// 	return false

		// }
		if s.isGoingDownTheEnd() {
			isEnd = true
			return false
		}

		s.y++
		s.cave.mapping[s.y][s.x] = 2
		s.path = append(s.path, coord{x: s.x, y: s.y})
		return true
	} else if s.canGoLeft() {
		if s.x-1 < 0 {
			fmt.Println("hey i can't go left but i could")
			fmt.Scanf("%s")
			s.path = append(s.path, coord{x: s.x, y: s.y})
			s.cave.mapping[s.y][s.x] = 2
			return false
		}
		s.y++
		s.x--
		s.cave.mapping[s.y][s.x] = 2
		s.path = append(s.path, coord{x: s.x, y: s.y})
		return true
	} else if s.canGoRight() {
		if s.x+1 > len(s.cave.mapping[s.y+1]) {
			s.path = append(s.path, coord{x: s.x, y: s.y})
			s.cave.mapping[s.y][s.x] = 2
			return false
		}

		s.y++
		s.x++
		s.cave.mapping[s.y][s.x] = 2
		s.path = append(s.path, coord{x: s.x, y: s.y})
		return true
	} else {
		// s.path = append(s.path, coord{x: s.x, y: s.y})
		s.cave.mapping[s.y][s.x] = 2
		return false
	}
}

func (s *sand) isSamePathAsLastOne() bool {
	if s.lastSand == nil {
		return false
	}
	if len(s.path) != len(s.lastSand.path) {
		return false
	}

	for i := range s.path {
		if s.path[i] != s.lastSand.path[i] {
			return false
		}
	}
	return true
}

var (
	counter = 0
	printed = false
)

// move recursevly until it finds a stop
func (s *sand) DoMove() {
	counter++
	if counter%10 == 0 {
		//clear terminal
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		s.cave.printMap()
		time.Sleep(25 * time.Millisecond)
	}
	if !s.move() {
		// fmt.Println("current y is at", s.y)
		// fmt.Println("map:", len(s.cave.mapping))
		// fmt.Println("i can't move anymore")
		// fmt.Scanf("%s")

		if isEnd {
			// if !printed {
			printed = true
			fmt.Println("the sand is:", s.cave.countSand())
			// os.Exit(0)
			// }
			return

		}

		if s.isSamePathAsLastOne() || s.y >= len(s.cave.mapping) {
			// os.Exit(0)
			return
		} else {
			// fmt.Println("current path is", s.path)
			s2 := sand{
				cave:     s.cave,
				x:        500 - s.cave.originX,
				y:        1,
				path:     make([]coord, 0),
				lastSand: s,
			}
			s2.DoMove()
		}
	}
	s.DoMove()
}

func (s *sand) DoMoveP2() {
	counter++
	if counter%1 == 0 {
		//clear terminal
		// cmd := exec.Command("clear")
		// cmd.Stdout = os.Stdout
		// cmd.Run()
		// s.cave.printMap()
		// time.Sleep(10 * time.Millisecond)
	}
	if !s.move() {
		// fmt.Println("current y is at", s.y)
		// fmt.Println("map:", len(s.cave.mapping))
		// fmt.Println("i can't move anymore")
		// fmt.Scanf("%s")

		if s.y == 0 && s.x == (500-s.cave.originX) {
			// if !printed {
			s.cave.printMap()

			printed = true
			fmt.Println("the sand is:", s.cave.countSand())
			os.Exit(0)
			// }
			return

		}

		if s.isSamePathAsLastOne() || s.y >= len(s.cave.mapping) {
			// os.Exit(0)
			return
		}
		// fmt.Println("current path is", s.path)
		s2 := sand{
			cave:     s.cave,
			x:        500 - s.cave.originX,
			y:        0,
			path:     make([]coord, 0),
			lastSand: s,
		}
		s2.DoMoveP2()
	}
	s.DoMoveP2()
}
