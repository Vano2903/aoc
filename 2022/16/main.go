package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type valve struct {
	name          string
	pressureValue int
	linkedValves  []*valve
}

func main() {
	f, _ := ioutil.ReadFile("inputtest")
	content := strings.Split(string(f), "\n")

	// valves := make(map[string]*valve)

	for _, line := range content {
		line = strings.ReplaceAll(line, "Valve ", "")
		line = strings.ReplaceAll(line, "has flow rate=", " ")
		line = strings.ReplaceAll(line, "; tunnels lead to valves", " ")
		line = strings.ReplaceAll(line, ",", "")
		vals := strings.Split(line, " ")
		fmt.Println(vals)
	}
}
