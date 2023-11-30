package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	nodes = 9
)

var (
	visited map[int]map[int]bool
)

type node struct {
	head *node
	tail *node
	x    int
	y    int
}

func addVisited(x, y int) {
	if visited[x] == nil {
		visited[x] = make(map[int]bool)
	}
	visited[x][y] = true
}

func (n *node) moveUp() {
	n.move(0, 1)
	if n.tail == nil {
		addVisited(n.x, n.y)
	}
}

func (n *node) moveDown() {
	n.move(0, -1)
	if n.tail == nil {
		addVisited(n.x, n.y)
	}
}

func (n *node) moveLeft() {
	n.move(-1, 0)
	if n.tail == nil {
		addVisited(n.x, n.y)
	}
}

func (n *node) moveRight() {
	n.move(1, 0)
	if n.tail == nil {
		addVisited(n.x, n.y)
	}
}

func (n *node) moveFromHead() {
	if n.head.x == n.x {
		if n.head.y > n.y {
			n.moveUp()
		} else if n.head.y < n.y {
			n.moveDown()
		}
	} else if n.head.y == n.y {
		if n.head.x > n.x {
			n.moveRight()
		} else if n.head.x < n.x {
			n.moveLeft()
		}
	}

	//diagonal movements
	if n.head.x > n.x && n.head.y > n.y {
		//top right
		n.move(1, 1)
		if n.tail == nil {
			addVisited(n.x, n.y)
		}
	} else if n.head.x > n.x && n.head.y < n.y {
		//bottom right
		n.move(1, -1)
		if n.tail == nil {
			addVisited(n.x, n.y)
		}
	} else if n.head.x < n.x && n.head.y > n.y {
		//top left
		n.move(-1, 1)
		if n.tail == nil {
			addVisited(n.x, n.y)
		}
	} else if n.head.x < n.x && n.head.y < n.y {
		//bottom left
		n.move(-1, -1)
		if n.tail == nil {
			addVisited(n.x, n.y)
		}
	}
}

func (n *node) move(increaseX, increaseY int) {
	n.x += increaseX
	n.y += increaseY
	if n.tail == nil {
		return
	}

	//check if tail is around n
	for i := n.x - 1; i <= n.x+1; i++ {
		for j := n.y - 1; j <= n.y+1; j++ {
			if i == n.tail.x && j == n.tail.y {
				return
			}
		}
	}

	//if not move tail
	n.tail.moveFromHead()
}

func main() {
	visited = make(map[int]map[int]bool)
	head := &node{nil, nil, 0, 0}
	last := head
	for i := 0; i < nodes; i++ {
		last.tail = &node{last, nil, 0, 0}
		last = last.tail
	}

	addVisited(0, 0)

	file, _ := ioutil.ReadFile("input")
	content := strings.Split(string(file), "\n")
	for _, c := range content {
		coor := strings.Split(c, " ")
		switch coor[0] {
		case "U":
			mov, _ := strconv.Atoi(coor[1])
			for i := 0; i < mov; i++ {
				head.moveUp()
			}
		case "D":
			mov, _ := strconv.Atoi(coor[1])
			for i := 0; i < mov; i++ {
				head.moveDown()
			}
		case "L":
			mov, _ := strconv.Atoi(coor[1])
			for i := 0; i < mov; i++ {
				head.moveLeft()
			}
		case "R":
			mov, _ := strconv.Atoi(coor[1])
			for i := 0; i < mov; i++ {
				head.moveRight()
			}
		}
	}

	//count all the visited nodes
	count := 0
	for _, v := range visited {
		for _, vv := range v {
			if vv {
				count++
			}
		}
	}
	fmt.Println(count)
}
