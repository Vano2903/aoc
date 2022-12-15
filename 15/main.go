package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

const (
	lineHeight = 2000000
)

var (
	line = []int{}
)

type point struct {
	x, y int
}

type sensor struct {
	p             point
	closestBeacon *beacon
	mapping       *mapping
	explored      []point
	latest        []point
}

func (s *sensor) checkSurroundingForBeacon(p point) bool {
	//up
	if p.y+1 == s.closestBeacon.p.y && p.x == s.closestBeacon.p.x {
		return true
	}
	//down
	if p.y-1 == s.closestBeacon.p.y && p.x == s.closestBeacon.p.x {
		return true
	}
	//left
	if p.y == s.closestBeacon.p.y && p.x-1 == s.closestBeacon.p.x {
		return true
	}
	//right
	if p.y == s.closestBeacon.p.y && p.x+1 == s.closestBeacon.p.x {
		return true
	}

	return false
}

func removeDuplicate(intSlice []point) []point {
	allKeys := make(map[point]bool)
	list := []point{}
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func (s *sensor) checkSurrounding(silent bool) bool {

	if len(s.explored) == 0 {
		if !silent {
			fmt.Println("first time exploring")
		}
		s.explored = append(s.explored, point{x: s.p.x, y: s.p.y})
		s.latest = append(s.latest, point{x: s.p.x, y: s.p.y})
	}

	found := false
	if !silent {
		fmt.Println("checking surrounding of", len(s.latest), "points")
	}
	for _, p := range s.latest {
		if s.checkSurroundingForBeacon(p) {
			if !silent {
				fmt.Println("beacon found at", p)
			}
			found = true
			break
		}
	}

	//save the latest points that are not duplicates
	tmp := []point{}
	if !silent {
		fmt.Println("saving latest points")
	}
	for _, p := range s.latest {
		tmp = append(tmp, point{x: p.x, y: p.y + 1})
		tmp = append(tmp, point{x: p.x, y: p.y - 1})
		tmp = append(tmp, point{x: p.x + 1, y: p.y})
		tmp = append(tmp, point{x: p.x - 1, y: p.y})
	}
	tmpLen := len(tmp)
	if !silent {
		fmt.Println("a total of ", tmpLen, "points were saved")
		fmt.Println("removing duplicates")
	}
	tmp = removeDuplicate(tmp)
	if !silent {
		fmt.Println("a total of ", tmpLen-len(tmp), "points were removed")
	}
	s.latest = tmp
	if !silent {
		fmt.Println("updating line")
	}
	s.updateLine()

	// fmt.Println(tmp)

	// time.Sleep(1 * time.Second)
	// fmt.Println()
	// for _, p := range s.explored {
	// 	s.explored = append(s.explored, point{x: p.x, y: p.y + 1})
	// 	s.explored = append(s.explored, point{x: p.x, y: p.y - 1})
	// 	s.explored = append(s.explored, point{x: p.x + 1, y: p.y})
	// 	s.explored = append(s.explored, point{x: p.x - 1, y: p.y})
	// }

	return found
}

func (s *sensor) exploreMap(silent bool) {
	for {
		print := len(s.latest)%100000 == 0
		if print {
			fmt.Println("explored:", len(s.latest))
		}
		if s.checkSurrounding(print) {
			fmt.Println("explored:", len(s.latest))
			fmt.Println("beacon found")
			fmt.Println("cleaning")
			fmt.Println()
			//update line
			s.explored = []point{}
			break
		}
		if !silent {
			s.mapping.updateMap()
			time.Sleep(100 * time.Millisecond)
			s.mapping.printMapping()
			fmt.Println()
		}
	}
}

func (s sensor) updateLine() {
	for _, e := range s.latest {
		//checking if the point exits the boundaries of the map
		if e.y != lineHeight || e.x < s.mapping.lowestX || e.x > s.mapping.biggestX || line[e.x-s.mapping.lowestX] != 0 {
			continue
		}
		line[e.x-s.mapping.lowestX] = 3
	}
}

type beacon struct {
	p point
}

// 0 air
// 1 sensor
// 2 beacon
// 3 explored
type mapping struct {
	mapp                        [][]int
	beacons                     []*beacon
	sensors                     []*sensor
	biggestX, biggestY, lowestX int
}

func (m *mapping) initMap() {
	//creating map of 0s
	for y := 0; y <= m.biggestY; y++ {
		m.mapp = append(m.mapp, []int{})
		for x := m.lowestX; x <= m.biggestX; x++ {
			m.mapp[y] = append(m.mapp[y], 0)
		}
	}

	//adding sensors
	for _, s := range m.sensors {
		m.mapp[s.p.y][s.p.x-m.lowestX] = 1
	}

	//adding beacons
	for _, b := range m.beacons {
		m.mapp[b.p.y][b.p.x-m.lowestX] = 2
	}
}

func (m mapping) initLine() {
	line = make([]int, 0)
	for x := m.lowestX; x <= m.biggestX; x++ {
		line = append(line, 0)
	}

	for _, s := range m.sensors {
		if s.p.y != lineHeight {
			continue
		}
		line[s.p.x-m.lowestX] = 1
	}

	for _, b := range m.beacons {
		if b.p.y != lineHeight {
			continue
		}
		line[b.p.x-m.lowestX] = 2
	}
}

func (m *mapping) updateLine() {
	for _, s := range m.sensors {
		for _, e := range s.explored {
			//checking if the point exits the boundaries of the map
			if e.y != lineHeight || e.x < m.lowestX || e.x > m.biggestX || line[e.x-m.lowestX] != 0 {
				continue
			}
			line[e.x-m.lowestX] = 3
		}
	}
}

func (m *mapping) updateMap() {
	for _, s := range m.sensors {
		for _, e := range s.explored {
			//checking if the point exits the boundaries of the map
			if e.y < 0 || e.y > m.biggestY || e.x < m.lowestX || e.x > m.biggestX || m.mapp[e.y][e.x-m.lowestX] != 0 {
				continue
			}
			m.mapp[e.y][e.x-m.lowestX] = 3
		}
	}
}

func (m mapping) printMapping() {
	for _, row := range m.mapp {
		for _, v := range row {
			switch v {
			case 0:
				fmt.Print(".")
			case 1:
				fmt.Print("S")
			case 2:
				fmt.Print("B")
			case 3:
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func (m mapping) printMap() {
	for y := 0; y <= m.biggestY; y++ {
		for x := m.lowestX; x <= m.biggestX; x++ {
			found := false
			for _, s := range m.sensors {
				if s.p.x == x && s.p.y == y {
					found = true
					break
				}
			}
			if found {
				fmt.Print("S")
				continue
			}
			for _, b := range m.beacons {
				if b.p.x == x && b.p.y == y {
					found = true
					break
				}
			}
			if found {
				fmt.Print("B")
				continue
			}
			fmt.Print(".")
		}
		fmt.Println()
	}
}

func (m mapping) printMapWithSurroundings(s []point) {
	for y := 0; y <= m.biggestY; y++ {
		for x := m.lowestX; x <= m.biggestX; x++ {
			found := false
			for _, s := range m.sensors {
				if s.p.x == x && s.p.y == y {
					found = true
					break
				}
			}
			if found {
				fmt.Print("S")
				continue
			}
			for _, b := range m.beacons {
				if b.p.x == x && b.p.y == y {
					found = true
					break
				}
			}
			if found {
				fmt.Print("B")
				continue
			}
			for _, p := range s {
				if p.x == x && p.y == y {
					found = true
					break
				}
			}
			if found {
				fmt.Print("#")
				continue
			}
			fmt.Print(".")
		}
		fmt.Println()
	}
}

func main() {
	f, _ := ioutil.ReadFile("input")
	content := strings.Split(string(f), "\n")

	biggestX := 0
	lowestX := 1_000_000_000_000
	biggestY := 0
	m := mapping{}
	sensors := []sensor{}
	for _, c := range content {
		c = strings.ReplaceAll(c, "\n", "")
		c = strings.ReplaceAll(c, "Sensor at ", "")
		c = strings.ReplaceAll(c, ": closest beacon is at", "")
		c = strings.ReplaceAll(c, ",", "")
		c = strings.ReplaceAll(c, "x=", "")
		c = strings.ReplaceAll(c, "y=", "")
		coords := strings.Split(c, " ")
		sx, _ := strconv.Atoi(coords[0])
		sy, _ := strconv.Atoi(coords[1])
		bx, _ := strconv.Atoi(coords[2])
		by, _ := strconv.Atoi(coords[3])
		if sx > biggestX {
			biggestX = sx
		}
		if sx < lowestX {
			lowestX = sx
		}
		if bx > biggestX {
			biggestX = bx
		}
		if bx < lowestX {
			lowestX = bx
		}
		if sy > biggestY {
			biggestY = sy
		}
		if by > biggestY {
			biggestY = by
		}
		m.beacons = append(m.beacons, &beacon{
			p: point{
				x: bx,
				y: by,
			},
		})
		sensors = append(sensors, sensor{
			p: point{
				x: sx,
				y: sy,
			},
			mapping: &m,
			closestBeacon: &beacon{
				p: point{
					x: bx,
					y: by,
				},
			},
		})
	}

	for i := range sensors {
		s := &sensors[i]
		m.sensors = append(m.sensors, s)
	}

	m.biggestX = biggestX
	m.biggestY = biggestY
	m.lowestX = lowestX

	// m.initMap()
	fmt.Println("creating line")
	m.initLine()
	fmt.Println("line created:", len(line))

	fmt.Println("exploring map")
	for i := range sensors {
		fmt.Println("sensor", i)
		sensors[i].exploreMap(true)
	}

	// m.updateMap()
	// m.printMapping()
	// m.printMapWithSurroundings(s.explored)
	// // fmt.Println(line)
	fmt.Println()
	// fmt.Println(m.mapp[lineHeight])
	fmt.Println(line)

	counter := 0
	fmt.Println(line)
	for _, l := range line {
		if l == 3 {
			counter++
		}
	}
	fmt.Println(counter)

	// start := point{
	// 	x: lowestX,
	// 	y: line,
	// }

	// end := point{
	// 	x: biggestX,
	// 	y: line,
	// }

	// toAnalize := []sensor{}
	// for _, s := range sensors {
	// 	if s.p.intersect(s.closestBeacon.p, start, end) {
	// 		toAnalize = append(toAnalize, s)
	// 	}
	// }
	// fmt.Println(toAnalize)

	// fmt.Println(m)
	// m.printMap()
	// m.initMap()

	// for i := range sensors {
	// 	sensors[i].exploreMap(true)
	// }
	// m.updateMap()
	// m.printMapping()
	// fmt.Println(len(m.mapp))

	// counter := 0
	// fmt.Println(m.mapp[line])
	// for _, l := range m.mapp[line] {
	// 	if l == 3 {
	// 		counter++
	// 	}
	// }
	// fmt.Println(counter)
}

/*


func (m mapping) printMap() {
	for y := 0; y <= m.biggestY; y++ {
		for x := m.lowestX; x <= m.biggestX; x++ {
			found := false
			for _, s := range m.sensors {
				if s.x == x && s.y == y {
					found = true
					break
				}
			}
			if found {
				fmt.Print("S")
				continue
			}
			for _, b := range m.beacons {
				if b.x == x && b.y == y {
					found = true
					break
				}
			}
			if found {
				fmt.Print("B")
				continue
			}
			fmt.Print(".")
		}
		fmt.Println()
	}
}

func (m mapping) printMapWithSurroundings() {
	for y := 0; y <= m.biggestY; y++ {
		for x := m.lowestX; x <= m.biggestX; x++ {
			found := false
			for _, s := range m.sensors {
				if s.x == x && s.y == y {
					found = true
					break
				}
			}
			if found {
				fmt.Print("S")
				continue
			}
			for _, b := range m.beacons {
				if b.x == x && b.y == y {
					found = true
					break
				}
			}
			if found {
				fmt.Print("B")
				continue
			}
			for _, p := range s {
				if p.x == x && p.y == y {
					found = true
					break
				}
			}
			if found {
				fmt.Print("#")
				continue
			}
			fmt.Print(".")
		}
		fmt.Println()
	}
}

*/
