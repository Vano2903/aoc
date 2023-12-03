package main

import (
	"fmt"
	"math"
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

type Node struct {
	Name  string
	Edges map[string]*Edge
}

type Edge struct {
	Distance int
	Vertex   *Node
}

var nodes map[string]*Node

func parse(lines []string) {
	nodes = make(map[string]*Node)
	// paths := make([]Node, 0)
	for _, line := range lines {
		words := strings.Split(line, " ")
		from := words[0]
		to := words[2]
		distance, _ := strconv.Atoi(words[4])
		if _, ok := nodes[from]; !ok {
			v := new(Node)
			v.Name = from
			v.Edges = make(map[string]*Edge)
			nodes[from] = v
		}
		if _, ok := nodes[to]; !ok {
			v := new(Node)
			v.Name = to
			v.Edges = make(map[string]*Edge)
			nodes[to] = v
		}
		edgeForward := new(Edge)
		edgeForward.Distance = distance
		edgeForward.Vertex = nodes[to]
		edgeBackward := new(Edge)
		edgeBackward.Distance = distance
		edgeBackward.Vertex = nodes[from]
		nodes[from].Edges[to] = edgeForward
		nodes[to].Edges[from] = edgeBackward
	}
}

func solve1(lines []string) {
	parse(lines)
	// for _, v := range nodes {
	// 	//fmt.Println(v.Name)
	// 	for _, e := range v.Edges {
	// 		fmt.Println(" ", e.Vertex.Name, e.Distance)
	// 	}
	// }

	lenghts := make([]int, 0)
	for _, node := range nodes {
		l := exploreShortest(node, make(map[string]bool), "", 0)
		lenghts = append(lenghts, l)
	}
	l := math.MaxInt32
	for _, length := range lenghts {
		if length < l {
			l = length
		}
	}
	// l := explore(nodes["B"], make(map[string]bool), "", 0)

	fmt.Println("the best is", l)
}

func exploreShortest(node *Node, checked map[string]bool, spacing string, pathLength int) int {
	// fmt.Println(spacing+"exploring", node.Name, "with path length", pathLength)
	checked[node.Name] = true
	lengths := make([]int, 0)
	for _, edge := range node.Edges {
		// fmt.Println(spacing+"checked", checked)

		if _, ok := checked[edge.Vertex.Name]; ok {
			// fmt.Println(spacing+"already visited", edge.Vertex.Name)
			continue
		}
		// fmt.Println(spacing+"going from", node.Name, "to", edge.Vertex.Name)
		c := make(map[string]bool)
		for k, v := range checked {
			// if k != node.Name {
			// 	c[k] = v
			//
			c[k] = v
		}
		// c[node.Name] = true
		l := exploreShortest(edge.Vertex, c, spacing+"  ", edge.Distance)
		// fmt.Println(spacing+"got back from", edge.Vertex.Name, "with length", l)
		lengths = append(lengths, l)

		// checked[edge.Vertex.Name] =
	}
	// fmt.Println(spacing+"returning", lengths)
	//ret min of lengths
	if len(lengths) == 0 {
		if len(checked) == len(nodes) {
			// fmt.Println(spacing + "looks like we are done")
			return pathLength
		} else {
			// fmt.Println(spacing + "looks like we are missing some nodes")
			return math.MaxInt32
		}
	}
	l := math.MaxInt32
	for _, length := range lengths {
		if length < l {
			l = length
		}
	}
	return l + pathLength
}

func solve2(lines []string) {
	parse(lines)

	lenghts := make([]int, 0)
	for _, node := range nodes {
		l := exploreLongest(node, make(map[string]bool), "", 0)
		lenghts = append(lenghts, l)
	}
	l := 0
	for _, length := range lenghts {
		if length > l {
			l = length
		}
	}
	// l := exploreLongest(nodes["A"], make(map[string]bool), "", 0)

	fmt.Println("the worst is", l)
}

func exploreLongest(node *Node, checked map[string]bool, spacing string, pathLength int) int {
	//fmt.Println(spacing+"exploring", node.Name, "with path length", pathLength)
	checked[node.Name] = true
	lengths := make([]int, 0)
	for _, edge := range node.Edges {
		//fmt.Println(spacing+"checked", checked)

		if _, ok := checked[edge.Vertex.Name]; ok {
			//fmt.Println(spacing+"already visited", edge.Vertex.Name)
			continue
		}
		//fmt.Println(spacing+"going from", node.Name, "to", edge.Vertex.Name)
		c := make(map[string]bool)
		for k, v := range checked {
			// if k != node.Name {
			// 	c[k] = v
			//
			c[k] = v
		}
		// c[node.Name] = true
		l := exploreLongest(edge.Vertex, c, spacing+"  ", edge.Distance)
		//fmt.Println(spacing+"got back from", edge.Vertex.Name, "with length", l)
		lengths = append(lengths, l)

		// checked[edge.Vertex.Name] =
	}
	//fmt.Println(spacing+"returning", lengths)
	//ret min of lengths
	if len(lengths) == 0 {
		if len(checked) == len(nodes) {
			//fmt.Println(spacing + "looks like we are done")
			return pathLength
		} else {
			// //fmt.Println(spacing + "looks like we are missing some nodes")
			return 0
		}
	}
	l := 0
	for _, length := range lengths {
		if length > l {
			l = length
		}
	}
	return l + pathLength
}
func main() {
	lines := readLines("input.txt")
	solve1(lines)
	solve2(lines)
}
