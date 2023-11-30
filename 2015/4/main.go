package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// func testHash(input string, numThreads, start int, wg *sync.WaitGroup, result chan int) {
// 	var b bytes.Buffer
// 	counter := start
// 	var hexsum string
// 	fmt.Println("starting routine", start)
// 	for {
// 		b.WriteString(input + string(counter))
// 		md5sum := md5.Sum(b.Bytes())
// 		hexsum = hex.EncodeToString(md5sum[:])
// 		found := true
// 		for i := 0; i < 5; i++ {
// 			if hexsum[i] != '0' {
// 				found = false
// 				break
// 			}
// 		}
// 		if found {
// 			break
// 		}
// 		counter += numThreads + start
// 		if counter%100000 == 0 {
// 			fmt.Printf("routine %d is at %d, is value of %s\n", start, counter, hexsum)
// 		}
// 	}
// 	fmt.Printf("found hash %s with counter of %d", hexsum, counter)
// 	result <- counter
// 	wg.Done()
// }

// func solve1(input string) {
// 	cores := runtime.NumCPU()
// 	// counter := 0
// 	wg := sync.WaitGroup{}
// 	wg.Add(cores)
// 	result := make(chan int)
// 	for i := 0; i < cores; i++ {
// 		go testHash(input, cores, i, &wg, result)
// 	}
// 	go func(result chan int) {
// 		for i := range result {
// 			fmt.Println("Result")
// 			fmt.Println(i)
// 		}
// 	}(result)
// 	wg.Wait()
// }

var foundOne bool
var currentRoutines int

func checkValue(input string, counter int, wg *sync.WaitGroup) {
	toCheck := fmt.Sprintf("%s%d", input, counter)
	md5sum := md5.Sum([]byte(toCheck))
	hexsum := hex.EncodeToString(md5sum[:])
	found := true
	for i := 0; i < 6; i++ {
		if hexsum[i] != '0' {
			found = false
			break
		}
	}
	if found {
		fmt.Printf("found hash %s with counter of %d, string is %s\nresult: %d\n", hexsum, counter, toCheck, counter)
		foundOne = true
	}
	wg.Done()
	currentRoutines--
}

func solve1(input string) {
	cores := runtime.NumCPU()
	maxRoutines := cores * 2000
	wg := sync.WaitGroup{}
	// wg.Add(cores)
	counter := 0
	for {
		wg.Add(1)
		currentRoutines++
		go checkValue(input, counter, &wg)
		counter++
		if currentRoutines > maxRoutines {
			wg.Wait()
		}
		if foundOne {
			wg.Wait() //wait for all routines cause multiple might find the result
			//so they all print and we can see which one is the lowest on the console
			break
		}
		// if counter%100000 == 0 {
		// 	fmt.Printf("counter is at %d\n", counter)
		// }
	}

	// counter := 0
	// wg := sync.WaitGroup{}
	// wg.Add(cores)
	// result := make(chan int)
	// for i := 0; i < cores; i++ {
	// 	go testHash(input, cores, i, &wg, result)
	// }
	// go func(result chan int) {
	// 	for i := range result {
	// 		fmt.Println("Result")
	// 		fmt.Println(i)
	// 	}
	// }(result)
	// wg.Wait()

}
func main() {
	// input := "bgvyzdsv"
	start := time.Now()
	solve1("bgvyzdsv")
	fmt.Println("took", time.Since(start))
}
