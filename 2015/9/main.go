package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLines(filename string) []string {
	c, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(c), "\n")
}

type Edge struct {
	From     *Node
	To       *Node
	Distance int
}

type Node struct {
	Name string
	To   map[string]*Edge
}

var graph = make(map[string]*Node)

func solve1(lines []string) {
	// paths := make([]Node, 0)
	for _, line := range lines {
		words := strings.Split(line, " ")
		from := words[0]
		to := words[2]
		distance, _ := strconv.Atoi(words[4])
		if _, ok := graph[from]; !ok {
			n := new(Node)
			n.Name = from
			graph[from] = n
		}
		if _, ok := graph[to]; !ok {
			n := new(Node)
			n.Name = to
			graph[to] = n
		}
		edge := new(Edge)
		edge.From = graph[from]
		edge.To = graph[to]
		edge.Distance = distance
		if graph[from].To == nil {
			graph[from].To = make(map[string]*Edge)
		}
		graph[from].To[to] = edge
	}
	for _, node := range graph {
		for _, edge := range node.To {
			fmt.Println(node.Name, "to", edge.To.Name, "=", edge.Distance)
		}
	}
	fmt.Println(graph["London"].To["Belfast"].Distance)
	//generate all possible paths
	//must visit all nodes exactly once
	for _, node := range graph {
		explore(node)
	}

}

// paths := make([][]string, 0)
func explore(node *Node) {
	checked := make(map[string]bool)
	checked[node.Name] = true
	for _, edge := range node.To {
		if !checked[edge.To.Name] {
			checked[edge.To.Name] = true
			explore(edge.To)
		}
	}
	for _, edge := range node.From {
		if !checked[edge.From.Name] {
			checked[edge.From.Name] = true
			explore(edge.From)
		}
	}

}

func main() {
	lines := readLines("test.txt")
	solve1(lines)
}
