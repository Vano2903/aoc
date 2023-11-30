package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Node struct {
	name   string
	files  []int
	nodes  []*Node
	parent *Node
	sum    int
}

func main() {
	root := &Node{
		name:  "/",
		files: []int{},
		sum:   0,
		nodes: make([]*Node, 0),
		parent: &Node{
			name: "nothing",
		},
	}

	ptr := root

	f, _ := ioutil.ReadFile("input")
	for _, a := range strings.Split(string(f), "\n") {
		if strings.HasPrefix(a, "$ cd ") {
			a = strings.Replace(a, "$ cd ", "", 1)
			if a == "/" {
				fmt.Println("going to root")
				ptr = root
				continue
			} else if a == ".." {
				fmt.Println("going back by one (" + ptr.parent.name + ")")
				ptr = ptr.parent
				continue
			} else {
				fmt.Println("we are in " + ptr.name + " changing directory to " + a)
				for _, n := range ptr.nodes {
					fmt.Printf("checking %q with %q\n", a, n.name)
					if a == n.name {
						fmt.Println("found")
						ptr = n
						break
					}
				}
			}
		}
		if strings.HasPrefix(a, "$ ls") {
			fmt.Println("listing")
			continue
		}
		if strings.HasPrefix(a, "dir ") {
			a = strings.Replace(a, "dir ", "", 1)
			fmt.Println("found a directory of name " + a)
			n := &Node{
				name:   a,
				files:  []int{},
				nodes:  make([]*Node, 0),
				parent: ptr,
			}
			ptr.nodes = append(ptr.nodes, n)
		} else {
			fmt.Println("found a file " + a)
			val, _ := strconv.Atoi(strings.Split(a, " ")[0])
			ptr.files = append(ptr.files, val)
		}

	}

	fmt.Println()
	sumFiles(root, 0)
	printNodes(*root)
	fmt.Println("sum / = ", root.sum)
	fmt.Println("we need to free up at least", 30000000-(70000000-root.sum))
	// find the closest sum to 70000000
	sums := getSum(root)
	//sort sums
	for i := 0; i < len(sums); i++ {
		for j := i; j < len(sums); j++ {
			if sums[i] > sums[j] {
				sums[i], sums[j] = sums[j], sums[i]
			}
		}
	}

	fmt.Println(sums)
	for _, s := range sums {
		if s >= 30000000-(70000000-root.sum) {
			fmt.Println(s)
			break
		}
	}
}

// return a slice with all the sums of the given node recursively
func getSum(n *Node) []int {
	sums := make([]int, 0)
	sums = append(sums, n.sum)
	for _, c := range n.nodes {
		sums = append(sums, getSum(c)...)
	}
	return sums
}

func printNodes(n Node) {
	fmt.Printf("dir: %q, files sum: %d, parent: %q\n", n.name, n.sum, n.parent.name)
	for _, c := range n.nodes {
		printNodes(*c)
	}
}

func sumFiles(n *Node, sum int) int {
	for _, f := range n.files {
		sum += f
	}
	for _, c := range n.nodes {
		sum += sumFiles(c, 0)
	}
	n.sum = sum
	return sum
}
