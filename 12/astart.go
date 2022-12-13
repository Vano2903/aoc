/*
BIG DISCLAIMER: this code sucks, it's late and idc about factoring it so that's what you get
also some functions repeats cause at first i was trying to implement a* and then i said "fuck it"
time to do djiakstra
*/

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
	"time"
	"unicode"
)

type point struct {
	x      int
	y      int
	height int
}

func distance(heatmapHeight, toCheckHeight point) int {
	heightDifference := toCheckHeight.height - heatmapHeight.height
	// fmt.Println("  checking point x:", toCheckHeight.x, "y:", toCheckHeight.y, "height:", toCheckHeight.height)
	// fmt.Println("  against point x:", heatmapHeight.x, "y:", heatmapHeight.y, "height:", heatmapHeight.height)
	// fmt.Println("  height difference:", heightDifference)

	if heightDifference > 1 {
		// fmt.Println("  W: osbstacle")
		return 100000000
	}

	return 1
}

func findLowest(list []point, measure map[point]float64) point {
	var lowest point
	lowest.height = 10000000000
	for _, l := range list {
		if measure[l] < float64(lowest.height) {
			lowest = l
		}
	}
	return lowest
}

func findLowestInt(list []point, measure map[point]int) point {
	var lowest point
	lowest.height = 10000000000
	for _, l := range list {
		if measure[l] < lowest.height {
			lowest = l
		}
	}
	return lowest
}

func remove(set []point, toRemove point) []point {
	//remove to remove from set
	for i, v := range set {
		if v.x == toRemove.x && v.y == toRemove.y {
			return append(set[:i], set[i+1:]...)
		}
	}
	return set
}

func getNeighbours(graph [][]int, current point) []point {
	neighbours := []point{}
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if current.x+i < 0 || current.x+i >= len(graph[0]) {
				continue
			}

			if current.y+j < 0 || current.y+j >= len(graph) {
				continue
			}

			if i == 0 && j == 0 {
				continue
			}

			neighbours = append(neighbours,
				point{
					y:      current.x + i,
					x:      current.y + j,
					height: graph[current.y+j][current.x+i],
				},
			)
		}
	}

	return neighbours
}

func getNeighboursPoint(graph [][]point, current point) []point {
	neighbours := []point{}
	//check top
	if current.y-1 >= 0 {
		neighbours = append(neighbours, graph[current.y-1][current.x])
	}
	//check bottom
	if current.y+1 < len(graph) {
		neighbours = append(neighbours, graph[current.y+1][current.x])
	}

	//check left
	if current.x-1 >= 0 {
		neighbours = append(neighbours, graph[current.y][current.x-1])
	}

	//check right
	if current.x+1 < len(graph[0]) {
		neighbours = append(neighbours, graph[current.y][current.x+1])
	}

	return neighbours
}

// Returns the estimated cost of the cheapest path from the current node to the goal node
func heuristic(current, goal point) float64 {
	// Here, we use the Manhattan distance as the heuristic function
	// return math.Abs(float64(current.x-goal.x)) + math.Abs(float64(current.y-goal.y))
	return 1
}

func printHeatmap(heatmap [][]point, current, check point, neighbours []point) {
	for x := 0; x < len(heatmap); x++ {
		for y := 0; y < len(heatmap[x]); y++ {
			if current.x == y && current.y == x {
				fmt.Print(" . ")
				continue
			}
			if check.x == y && check.y == x {
				fmt.Print(" C ")
				continue
			}
			print := true
			for _, n := range neighbours {
				if n.x == y && n.y == x {
					fmt.Print(" N ")
					print = false
				}
			}

			if print {
				if heatmap[x][y].height >= 10 {
					fmt.Print(heatmap[x][y].height)
					fmt.Print(" ")
				} else {
					fmt.Print(" ", heatmap[x][y].height)
					fmt.Print(" ")
				}
			}
		}
		fmt.Println()
	}
}

func dijkstra(graph [][]point, start, end point) int {
	dist := make(map[point]int, 0)
	previous := make(map[point]*point, 0)
	points := []point{}
	for y, line := range graph {
		for x, l := range line {
			p := point{
				x:      x,
				y:      y,
				height: l.height,
			}
			dist[p] = 1000000
			previous[p] = nil
			points = append(points, p)
		}
	}

	dist[start] = 0

	for {
		if len(points) == 0 {
			// fmt.Println("hey maybe i found something")
			break
		}

		current := findLowestInt(points, dist)
		// fmt.Println("  current x:", current.x, "y:", current.y, "height:", current.height, "dist:", dist[current], "previous:", previous[current])
		points = remove(points, current)
		// if dist[current] >= 1000 {
		// 	fmt.Println("found a dead end")
		// 	break
		// }

		neighbours := getNeighboursPoint(graph, current)
		for _, neighbour := range neighbours {
			// printHeatmap(graph, current, neighbour, neighbours)
			alt := dist[current] + distance(current, neighbour)
			if alt < dist[neighbour] {
				// fmt.Println("updating", neighbour, "from", dist[neighbour], "to", alt)
				dist[neighbour] = alt
				previous[neighbour] = &current
				points = append([]point{neighbour}, points...)
			}
			// fmt.Println()
		}
		// fmt.Println()
	}

	// fmt.Println(dist)
	// for k, v := range dist {
	// 	fmt.Println("dist from x:", start.x, "y:", start.y, "to x:", k.x, "y:", k.y, "is", v)
	// }
	return dist[end]
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

	for _, v := range heatmap {
		for _, p := range v {
			if p < 10 {
				fmt.Print(" ", p, " ")
			} else {
				fmt.Print(p, " ")
			}
		}
		fmt.Println()
	}

	// fmt.Println(astar(heatmap, start, end))
	heatmapPoint := make([][]point, 0)
	for x, line := range heatmap {
		heatline := make([]point, 0)
		for y, l := range line {
			heatline = append(heatline, point{
				x:      y,
				y:      x,
				height: l,
			})
		}
		heatmapPoint = append(heatmapPoint, heatline)
	}

	startingPoints := []point{}
	for _, line := range heatmapPoint {
		for _, p := range line {
			if p.height == 0 {
				startingPoints = append(startingPoints, p)
			}
		}
	}

	fmt.Println("points to check:", len(startingPoints))

	//waigroup of max 20 goroutines
	var wg sync.WaitGroup
	maxgoroutines := 30
	goroutines := 0
	var paths []int

	for _, s := range startingPoints {
		wg.Add(1)
		time.Sleep(100 * time.Millisecond)
		goroutines++
		go func(s point, counter *int, paths *[]int, wg *sync.WaitGroup) {
			// fmt.Print("calculatin dijkstra for point x:", s.x, " y:", s.y, "...")
			// fmt.Println("starting routine", *counter)
			p := dijkstra(heatmapPoint, s, end)
			if p == 1000000 {
				fmt.Println("no path found")
			} else {
				fmt.Println(p)
			}
			//append to paths
			*paths = append(*paths, p)
			// fmt.Println("ending routine", *counter)
			*counter--
			wg.Done()
		}(s, &goroutines, &paths, &wg)

		fmt.Println("goroutines:", goroutines, "max:", maxgoroutines)
		if goroutines >= maxgoroutines {
			fmt.Println("waiting...")
			wg.Wait()
			fmt.Println("done")
			goroutines = 0
		}
	}
	wg.Wait()
	fmt.Println()
	fmt.Println()
	fmt.Println(paths)
	fmt.Println()
	fmt.Println("lowest ever:", findLowestPaths(paths))
}

func findLowestPaths(p []int) int {
	lowest := 1000000
	for _, v := range p {
		if v < lowest {
			lowest = v
		}
	}
	return lowest
}
