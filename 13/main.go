package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	f, _ := ioutil.ReadFile("inputtest")
	content := strings.Split(string(f), "\n")
	for _, c := range content {
		c = strings.TrimSpace(c)
		fmt.Println(c)
	}
}
