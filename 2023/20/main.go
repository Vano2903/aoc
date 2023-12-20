package main

import (
	"fmt"
	"os"
	"strings"
)

func readFile(path string) []string {
	c, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	return strings.Split(string(c), "\n")
}

const print bool = false

var highCounter int
var lowCounter int

var pulsesQueue []pulseQueueItem

type Pulse string
type kind string

const (
	High    Pulse = "high"
	Low     Pulse = "low"
	Nothing Pulse = "nothing"

	FlipFlop    kind = "flip"
	Conjunction kind = "conj"
	Broadcaster kind = "broad"
	Untyped     kind = "untyped"
)

type pulseQueueItem struct {
	p    Pulse
	from string
	to   *module
}

type module struct {
	kind       kind
	isOn       bool
	name       string
	tmpOutputs []string
	storage    map[string]Pulse
	// inputs []Pulse
	Outputs []*module
}

func (m module) String() string {
	outs := make([]string, len(m.Outputs))
	for i, o := range m.Outputs {
		outs[i] = o.name
	}
	return fmt.Sprintf("%s %s -> %s", m.kind, m.name, outs)
}

func (m *module) ConnectOutput(mod *module) {
	m.Outputs = append(m.Outputs, mod)
}

func (m *module) ConnectInput(mod *module) {
	if m.kind == Conjunction {
		if m.storage == nil {
			m.storage = make(map[string]Pulse)
		}
		m.storage[mod.name] = Low
	}
}

func (m *module) Process(from string, p Pulse) {
	switch p {
	case High:
		highCounter++
	case Low:
		lowCounter++
	}
	switch m.kind {
	case Broadcaster:
		// fmt.Println("broadcasting", p, "from", name)
		m.broadcast(p)
	case FlipFlop:
		// fmt.Println("flipflop", name, "got", p)
		if p == Low {
			toSend := Nothing
			if m.isOn {
				toSend = Low
				m.isOn = false
			} else {
				toSend = High
				m.isOn = true
			}
			m.broadcast(toSend)

			// } else {
			// 	fmt.Println("flipflop not doing anything on high")
		}
	case Conjunction:
		if m.storage == nil {
			m.storage = make(map[string]Pulse)
		}
		// fmt.Println("conjunction", name, "got", p)
		// fmt.Println("storage", m.storage)
		m.storage[from] = p
		toSend := Low
		for _, v := range m.storage {
			if v == Low {
				// fmt.Println("there is a low in the storage, sending high")
				toSend = High
				break
			}
		}
		m.broadcast(toSend)
	case Untyped:
		// if p == Low {
		// 	singleLowPulse++
		// }
	}
}

func (m *module) broadcast(p Pulse) {
	for _, o := range m.Outputs {
		if print {
			switch p {
			case High:
				fmt.Println(m.name, "-high ->", o.name)
			case Low:
				fmt.Println(m.name, "-low ->", o.name)
			}
		}
		pulsesQueue = append(pulsesQueue, pulseQueueItem{p, m.name, o})
	}
}

func solve1(lines []string) {
	all := make(map[string]*module)
	br := new(module)
	br.kind = Broadcaster
	br.name = "broadcaster"
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " -> ")
		if parts[0] == "broadcaster" {
			to := strings.Split(parts[1], ", ")
			br.tmpOutputs = to
			all["broadcast"] = br
		} else {
			t := parts[0][:1]
			name := parts[0][1:]
			to := strings.Split(parts[1], ", ")
			m := new(module)
			m.name = name
			m.tmpOutputs = to
			switch t {
			case "%":
				m.kind = FlipFlop
			case "&":
				m.kind = Conjunction
			default:
				panic("unknown type")
			}
			// fmt.Println("connecting", name, "to", to)
			all[name] = m
		}
	}

	for _, m := range all {
		for _, o := range m.tmpOutputs {
			if o == "" {
				continue
			}
			if all[o] == nil {
				m := new(module)
				m.name = o
				m.kind = Untyped
				all[o] = m
			}
			m.ConnectOutput(all[o])
			all[o].ConnectInput(m)
		}
	}
	// fmt.Println("output", all["output"])
	// fmt.Println(br.Outputs[0].Outputs[0].Outputs[0].Outputs[0].Outputs[0].name)

	startPulse := Low
	for i := 0; i < 1000; i++ {
		if print {
			fmt.Println("\nstarting pulse", i+1)
		}
		pulsesQueue = append(pulsesQueue, pulseQueueItem{startPulse, br.name, br})
		for len(pulsesQueue) > 0 {
			item := pulsesQueue[0]
			pulsesQueue = pulsesQueue[1:]

			pulse := item.p
			m := item.to
			from := item.from
			m.Process(from, pulse)
		}
	}
	// lowCounter++
	fmt.Println("solve1", highCounter*lowCounter)
}

func solve2(lines []string) {
	all := make(map[string]*module)
	br := new(module)
	br.kind = Broadcaster
	br.name = "broadcaster"
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " -> ")
		if parts[0] == "broadcaster" {
			to := strings.Split(parts[1], ", ")
			br.tmpOutputs = to
			all["broadcast"] = br
		} else {
			t := parts[0][:1]
			name := parts[0][1:]
			to := strings.Split(parts[1], ", ")
			m := new(module)
			m.name = name
			m.tmpOutputs = to
			switch t {
			case "%":
				m.kind = FlipFlop
			case "&":
				m.kind = Conjunction
			default:
				panic("unknown type")
			}
			all[name] = m
		}
	}

	un := new(module)
	un.kind = Untyped
	for _, m := range all {
		for _, o := range m.tmpOutputs {
			if o == "" {
				continue
			}
			if all[o] == nil {
				un.name = o
				all[o] = un
			}
			m.ConnectOutput(all[o])
			all[o].ConnectInput(m)
		}
	}

	cycleCounter := make(map[string]int)
	var before string
	for _, m := range all {
		for _, o := range m.Outputs {
			if o.kind == Untyped {
				before = m.name
			}
		}
		if before != "" {
			break
		}
	}

	for _, m := range all {
		for _, o := range m.Outputs {
			if o.name == before {
				cycleCounter[m.name] = 0
			}
		}
	}

	startPulse := Low
	i := 1
	for {
		if print {
			fmt.Println("\nstarting pulse", i)
		}
		pulsesQueue = append(pulsesQueue, pulseQueueItem{startPulse, br.name, br})
		for len(pulsesQueue) > 0 {
			item := pulsesQueue[0]
			pulsesQueue = pulsesQueue[1:]

			pulse := item.p
			m := item.to
			from := item.from
			m.Process(from, pulse)

			for _, p := range pulsesQueue {
				if j, ok := cycleCounter[p.from]; ok && p.p == High && j == 0 {
					cycleCounter[p.from] = i
				}
			}
		}
		done := true
		for _, v := range cycleCounter {
			if v == 0 {
				done = false
				break
			}
		}
		if done {
			break
		}
		i++
	}
	var vals []int
	for _, v := range cycleCounter {
		vals = append(vals, v)
	}
	fmt.Println("solve2", LCM(vals[0], vals[1], vals[2], vals[3]))
}

func main() {
	l := readFile("input")
	solve1(l)
	solve2(l)
}

//100.000.000 low
//10.000.000.000 low
//884.814.263.280 low
//884814263280
//240914003753369 RIGHT (fucking missed the +1)

//find the conjunctions that connect to rx
//calculate how many cycles for each to get return an high pulse
//find the lcm of all of them

// lcm from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
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
