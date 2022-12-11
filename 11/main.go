package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

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

type monkey struct {
	number        int
	holding       []float64
	operation     string
	value         float64
	division      float64
	monkey1       *monkey
	monkey1num    int
	monkey2       *monkey
	monkey2num    int
	totalItemHeld int
	module        int
}

func (m *monkey) fetch(number float64) {
	m.holding = append(m.holding, number)
}

func (m *monkey) check() {
	// fmt.Println("Monkey:", m.number)
	for _, worryLvl := range m.holding {
		// fmt.Println("  Monkey is holding", m.holding)
		m.holding = m.holding[1:]
		m.totalItemHeld++
		// fmt.Println("  Monkey inspects an item with a worry level of", worryLvl)
		// fmt.Print("    Worry level is ", worryLvl, m.operation, m.value, " = ")

		if m.operation == "+" {
			worryLvl += m.value
		} else if m.operation == "*" {
			worryLvl *= m.value
		} else if m.operation == "^" {
			worryLvl = worryLvl * worryLvl
		}
		// fmt.Println(worryLvl)
		// fmt.Println("    Monkey gets bored with item. Worry level is divided by 3 to", worryLvl)
		worryLvl = math.Mod(worryLvl, float64(m.module))

		if math.Mod(worryLvl, m.division) == 0 {
			// fmt.Println("    Current worry level is divisible by", m.division)
			// fmt.Println("    Item with worry level", worryLvl, "is thrown to monkey ", m.monkey1num)
			m.monkey1.fetch(worryLvl)
		} else {
			// fmt.Println("    Current worry level is not divisible by", m.division)
			// fmt.Println("    Item with worry level", worryLvl, "is thrown to monkey ", m.monkey2num)
			m.monkey2.fetch(worryLvl)
		}
		// fmt.Println()
	}
	// fmt.Println(m.holding)
}

func main() {
	rounds := 10000

	f, _ := ioutil.ReadFile("input")
	content := strings.Split(string(f), "\n")
	monkeys := make([]monkey, 0)
	m := monkey{}
	for _, c := range content {
		if strings.Contains(c, "Monkey") {
			m.number, _ = strconv.Atoi(strings.Split(c, " ")[1][:len(strings.Split(c, " ")[1])-1])
		} else if strings.Contains(c, "Starting") {
			c = strings.ReplaceAll(c, "  Starting items: ", "")
			l := strings.Split(c, ", ")
			for _, i := range l {
				val, _ := strconv.Atoi(i)
				v := float64(val)
				m.holding = append(m.holding, v)
			}
		} else if strings.Contains(c, "Operation:") {
			c = strings.ReplaceAll(c, "Operation: new = old", "")
			c = strings.TrimSpace(c)
			vals := strings.Split(c, " ")
			if vals[1] == "old" {
				m.operation = "^"
				m.value = 2
			} else {
				m.operation = vals[0]
				v, _ := strconv.Atoi(vals[1])
				m.value = float64(v)
			}
		} else if strings.Contains(c, "  Test: divisible by ") {
			c = strings.ReplaceAll(c, "  Test: divisible by ", "")
			d, _ := strconv.Atoi(c)
			m.division = float64(d)
		} else if strings.Contains(c, "    If true:") {
			c = strings.ReplaceAll(c, "    If true: throw to monkey ", "")
			m.monkey1num, _ = strconv.Atoi(c)
		} else if strings.Contains(c, "    If false:") {
			c = strings.ReplaceAll(c, "    If false: throw to monkey ", "")
			m.monkey2num, _ = strconv.Atoi(c)
		}
		if c == "" {
			monkeys = append(monkeys, m)
			m = monkey{}
		}
	}
	//linking monkeys
	for i := 0; i < len(monkeys); i++ {
		fmt.Printf("linking monkey %d to monkey %d and monkey %d\n", monkeys[i].number, monkeys[i].monkey1num, monkeys[i].monkey2num)
		monkeys[i].monkey1 = &monkeys[monkeys[i].monkey1num]
		monkeys[i].monkey2 = &monkeys[monkeys[i].monkey2num]
	}

	module := 1
	for _, m := range monkeys {
		module = module * int(m.division)
	}

	for i := 0; i < len(monkeys); i++ {
		monkeys[i].module = module
	}
	fmt.Println()
	fmt.Println()
	for r := 0; r < rounds; r++ {
		for i := 0; i < len(monkeys); i++ {
			monkeys[i].check()
		}
		// fmt.Println("end round", r)
		// for _, m := range monkeys {
		// 	fmt.Println("monkey", m.number, "is holding", m.holding)
		// }
	}
	high1 := 0
	high2 := 0
	for _, m := range monkeys {
		fmt.Println("monkey", m.number, "has held a total of", m.totalItemHeld, "items")
		if m.totalItemHeld >= high1 {
			high2 = high1
			high1 = m.totalItemHeld
		}
		if m.totalItemHeld >= high2 && m.totalItemHeld < high1 {
			high2 = m.totalItemHeld
		}
	}
	fmt.Println("result is:", high1*high2)
}
