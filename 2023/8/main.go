package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func readLines(path string) []string {
	c, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(c), "\n")
}

type Node struct {
	name  string
	left  *Node
	right *Node
}

func (n *Node) String() string {
	return fmt.Sprintf("%s -> %s, %s", n.name, n.left.name, n.right.name)
}

func solve1(lines []string) {
	path := lines[0]
	graph := make(map[string]*Node)
	// root := new(Node)
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " = (")
		from := parts[0]
		// if
		to := strings.Split(parts[1], ", ")
		left := to[0]
		right := to[1][:len(to[1])-1]
		if _, ok := graph[from]; !ok {
			n := new(Node)
			n.name = from
			graph[from] = n
		}
		if _, ok := graph[left]; !ok {
			n := new(Node)
			n.name = left
			graph[left] = n
		}
		if _, ok := graph[right]; !ok {
			n := new(Node)
			n.name = right
			graph[right] = n
		}
		graph[from].left = graph[left]
		graph[from].right = graph[right]
	}
	// fmt.Println(path)

	pathIndex := 0
	steps := 0
	currentNode := graph["AAA"]
	for {
		pathIndex = pathIndex % len(path)

		if path[pathIndex] == 'R' {
			pathIndex += 1
			steps += 1
			currentNode = currentNode.right
		} else {
			pathIndex += 1
			steps += 1
			currentNode = currentNode.left
		}
		if currentNode.name == "ZZZ" {
			break
		}
	}

	// fmt.Println(graph["AAA"])
	fmt.Println(steps)
}

// it returns -1 if nothing was found (loop)
func stepsToFirstEndingNode(node *Node, path string) int {
	n := len(path)
	steps := 0
	checked := make(map[string]bool)
	for {
		key := fmt.Sprintf("%s%d", node.name, steps%n)
		if checked[key] {
			return -1
		}
		checked[key] = true
		if strings.HasSuffix(node.name, "Z") {
			return steps
		}

		if path[steps%n] == 'R' {
			node = node.right
		} else {
			node = node.left
		}
		steps += 1
	}
}

func solve2(lines []string) {
	path := lines[0]
	graph := make(map[string]*Node)
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " = (")
		from := parts[0]
		// if
		to := strings.Split(parts[1], ", ")
		left := to[0]
		right := to[1][:len(to[1])-1]
		if _, ok := graph[from]; !ok {
			n := new(Node)
			n.name = from
			graph[from] = n
		}
		if _, ok := graph[left]; !ok {
			n := new(Node)
			n.name = left
			graph[left] = n
		}
		if _, ok := graph[right]; !ok {
			n := new(Node)
			n.name = right
			graph[right] = n
		}
		graph[from].left = graph[left]
		graph[from].right = graph[right]
	}
	// fmt.Println(path)
	roots := make([]*Node, 0)
	for _, node := range graph {
		if strings.HasSuffix(node.name, "A") {
			roots = append(roots, node)
		}
	}

	paths := make([]int, 0)
	// steps := 0
	// pathIndex := make([]int, len(roots))
	// checked := make([]map[string]bool, len(roots))
	// for {

	// fmt.Println(len(roots))
	for _, root := range roots {
		pathCount := stepsToFirstEndingNode(root, path)
		if pathCount == -1 {
			// fmt.Println("hummm")
			// return
			continue
		}
		paths = append(paths, pathCount)
	}
	//
	// fmt.Println(paths)
	fmt.Println(LCM(paths[0], paths[1], paths[2:]...))

	// if checked[i] == nil {
	// 	checked[i] = make(map[string]bool)
	// }
	// // currentNode := root
	// // fmt.Println(roots[i])
	// pathIndex[i] = pathIndex[i] % len(path)

	// if path[pathIndex[i]] == 'R' {
	// 	if checked[i][roots[i].name+roots[i].right.name] {
	// 		fmt.Println("already checked, is a loop")
	// 		break
	// 	}
	// 	checked[i][roots[i].name+roots[i].right.name] = true
	// 	// fmt.Println("going right")
	// 	pathIndex[i] += 1
	// 	// steps += 1
	// 	roots[i] = roots[i].right
	// } else {
	// 	if checked[i][roots[i].name+roots[i].left.name] {
	// 		fmt.Println("already checked, is a loop")
	// 		break
	// 	}
	// 	checked[i][roots[i].name+roots[i].left.name] = true
	// 	// fmt.Println("going left")
	// 	pathIndex[i] += 1
	// 	roots[i] = roots[i].left
	// }
	// fmt.Println(checked[i])
	// if strings.HasSuffix(roots[i].name, "Z") {
	// 	break
	// }
	// fmt.Println(roots)
	// }
	// steps += 1
	// if steps == 100 {
	// 	return
	// }
	// done := true
	// for _, root := range roots {
	// 	// fmt.Println("checking if done", root.name)
	// 	if !strings.HasSuffix(root.name, "Z") {
	// 		// fmt.Println("not done", root.name)
	// 		done = false
	// 		break
	// 	}
	// }
	// if steps%100_000 == 0 {
	// 	fmt.Println(steps)
	// }
	// if done {
	// 	fmt.Println("took", steps, "steps")
	// 	return
	// }
}

//copied those 2 functions from
//https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func main() {
	l := readLines("input.txt")
	now := time.Now()
	solve1(l)
	fmt.Println("took", time.Since(now))
	now = time.Now()
	solve2(l)
	fmt.Println("took", time.Since(now))
}
