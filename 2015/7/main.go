package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var wires map[string]*Wire

func readLine(filename string) []string {
	c, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(c), "\n")
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

type Wire struct {
	name     string
	value    uint16
	hasValue bool
	fromRule *Rule
	fromWire *Wire
}

func (w *Wire) String() string {
	if w.hasValue {
		return fmt.Sprintf("%s: %d", w.name, w.value)
	}
	if w.fromRule != nil {
		return w.fromRule.string()
	}
	if w.fromWire != nil {
		return fmt.Sprintf("%s -> %s", w.fromWire.name, w.name)
	}
	return fmt.Sprintf("unknown wire %s", w.name)
}

func (w *Wire) Relink() *Wire {
	// fmt.Println("possible unknown wire, relinking", w.name)
	if wires[w.name] == nil {
		fmt.Println("wire", w.name, "not found in wires")
		panic("wire not found")
	}
	return wires[w.name]
}

func (w *Wire) eval() (uint16, error) {
	// fmt.Println("eval", w)
	if w.hasValue {
		// fmt.Println("has value", w.value)
		return w.value, nil
	}
	var val, val1 uint16
	var err error
	if w.fromRule != nil {
		switch w.fromRule.op {
		case "NOT":
			toEval := w.fromRule.from[0]
			// fmt.Println("NOT", toEval)
			val, err = toEval.eval()
			if err != nil {
				toEval = toEval.Relink()
				val, err = toEval.eval()
				if err != nil {
					panic("should not happen")
				}
			}
			w.value = ^val
			w.hasValue = true
			return w.value, nil
		case "AND":
			if len(w.fromRule.from) != 2 {
				//do the end with the rule value
				toEval := w.fromRule.from[0]
				// fmt.Println(toEval, "AND", w.fromRule.value)
				val, err = toEval.eval()
				if err != nil {
					toEval = toEval.Relink()
					val, err = toEval.eval()
					if err != nil {
						panic("should not happen")
					}
				}

				w.value = val & w.fromRule.value
				w.hasValue = true
				return w.value, nil
			} else {
				toEval := w.fromRule.from[0]
				// fmt.Println(toEval, "AND", w.fromRule.from[1])
				val, err = toEval.eval()
				if err != nil {
					// fmt.Printf("toeval %q", toEval.name)
					toEval = toEval.Relink()
					val, err = toEval.eval()
					if err != nil {
						fmt.Println("toEval after err", toEval)
						// toEval.eval()
						fmt.Println("talking about", toEval.name)
						fmt.Printf("wire %p\n", wires["x"])
						fmt.Printf("toeval %p\n", toEval)
						fmt.Println(wires["x"])
						fmt.Println(toEval.hasValue)
						fmt.Println(val, err)
						panic("should not happen")
					}
				}

				toEval = w.fromRule.from[1]
				val1, err = toEval.eval()
				if err != nil {
					toEval = toEval.Relink()
					val1, err = toEval.eval()
					if err != nil {
						panic("should not happen")
					}
				}
				w.hasValue = true
				w.value = val & val1
			}
			return w.value, nil
		case "OR":
			if len(w.fromRule.from) != 2 {
				//do the end with the rule value
				//do the end with the rule value
				toEval := w.fromRule.from[0]
				// fmt.Println(toEval, "OR", w.fromRule.value)
				val, err = toEval.eval()
				if err != nil {
					toEval = toEval.Relink()
					val, err = toEval.eval()
					if err != nil {
						panic("should not happen")
					}
				}

				w.value = val | w.fromRule.value
				w.hasValue = true
				return w.value, nil
			} else {
				toEval := w.fromRule.from[0]
				// fmt.Println(toEval, "OR", w.fromRule.from[1])
				val, err = toEval.eval()
				if err != nil {
					toEval = toEval.Relink()
					val, err = toEval.eval()
					if err != nil {
						panic("should not happen")
					}
				}

				toEval = w.fromRule.from[1]
				val1, err = toEval.eval()
				if err != nil {
					// fmt.Println("possible unknown wire, relinking")
					toEval = toEval.Relink()
					val1, err = toEval.eval()
					if err != nil {
						panic("should not happen")
					}
				}
				w.value = val | val1
				w.hasValue = true
				return w.value, nil
			}
		case "LSHIFT":
			toEval := w.fromRule.from[0]
			// fmt.Println(toEval, "LSHIFT", w.fromRule.value)
			val, err = toEval.eval()
			if err != nil {
				toEval = toEval.Relink()
				val, err = toEval.eval()
				if err != nil {
					panic("should not happen")
				}
			}
			w.value = val << w.fromRule.value
			w.hasValue = true
			return w.value, nil
		case "RSHIFT":
			toEval := w.fromRule.from[0]
			// fmt.Println(toEval, "RSHIFT", w.fromRule.value)
			val, err = toEval.eval()
			if err != nil {
				toEval = toEval.Relink()
				val, err = toEval.eval()
				if err != nil {
					panic("should not happen")
				}
			}
			w.value = val >> w.fromRule.value
			w.hasValue = true
			return w.value, nil
		default:
			panic("unknown op")
		}
	}
	if w.fromWire != nil {
		// fmt.Println("from wire", w.fromWire)
		val, err = w.fromWire.eval()
		if err != nil {
			w.fromWire = w.fromWire.Relink()
			val, err = w.fromWire.eval()
			if err != nil {
				panic("should not happen")
			}
		}
		w.value = val
		w.hasValue = true
		return w.value, nil
	}
	// fmt.Println("unkonwn wire", w)
	return 0, fmt.Errorf("unknown wire %s", w.name)
}

type Rule struct {
	to    *Wire
	from  []*Wire
	op    string
	value uint16
}

func (r *Rule) string() string {
	switch r.op {
	case "NOT":
		return fmt.Sprintf("NOT %s -> %s", r.from[0].name, r.to.name)
	case "AND":
		if len(r.from) == 1 {
			return fmt.Sprintf("%s AND %d -> %s", r.from[0].name, r.value, r.to.name)
		}
		return fmt.Sprintf("%s AND %s -> %s", r.from[0].name, r.from[1].name, r.to.name)
	case "OR":
		if len(r.from) == 1 {
			return fmt.Sprintf("%s OR %d -> %s", r.from[0].name, r.value, r.to.name)
		}
		return fmt.Sprintf("%s OR %s -> %s", r.from[0].name, r.from[1].name, r.to.name)
	case "LSHIFT":
		return fmt.Sprintf("%s LSHIFT %d -> %s", r.from[0].name, r.value, r.to.name)
	case "RSHIFT":
		return fmt.Sprintf("%s RSHIFT %d -> %s", r.from[0].name, r.value, r.to.name)
	default:
		panic("unknown op")
	}
}

func isNumber(s string) (uint16, bool) {
	n, err := strconv.Atoi(s)
	return uint16(n), err == nil
}

func addWireIfNotExists(name string) *Wire {
	w2, ok := wires[name]
	if !ok {
		w2 = new(Wire)
		w2.name = name
		wires[name] = w2
	}
	return w2
}

func generateWires(lines []string) {
	wires = make(map[string]*Wire)
	ordered := make([]*Wire, 0)
	for _, line := range lines {
		i := strings.Split(line, " -> ")
		w := new(Wire)
		to := i[1]
		ordered = append(ordered, w)
		w.name = to
		wires[to] = w
		from := strings.Split(i[0], " ")

		if len(from) == 1 {
			//its a number
			value, err := strconv.Atoi(from[0])
			if err != nil {
				//its a wire
				w2 := addWireIfNotExists(from[0])
				w.fromWire = w2
			} else {
				w.value = uint16(value)
				w.hasValue = true
				w.fromRule = nil
			}
		} else {
			r := new(Rule)
			r.to = w
			if from[0] == "NOT" {
				number, isNum := isNumber(from[1])
				if isNum {
					r.value = number
				} else {
					w2 := addWireIfNotExists(from[1])
					r.from = append(r.from, w2)
				}
				r.op = "NOT"
				w.fromRule = r
			} else {
				switch from[1] {
				case "AND":
					number, isNum := isNumber(from[0])
					if isNum {
						r.value = number
					} else {
						w2 := addWireIfNotExists(from[0])
						r.from = append(r.from, w2)
					}
					number, isNum = isNumber(from[2])
					if isNum {
						r.value = number
					} else {
						w2 := addWireIfNotExists(from[2])
						r.from = append(r.from, w2)
					}
					r.op = "AND"
					w.fromRule = r
				case "OR":
					number, isNum := isNumber(from[0])
					if isNum {
						r.value = number
					} else {
						w2 := addWireIfNotExists(from[0])
						r.from = append(r.from, w2)
					}
					number, isNum = isNumber(from[2])
					if isNum {
						r.value = number
					} else {
						w2 := addWireIfNotExists(from[2])
						r.from = append(r.from, w2)
					}
					r.op = "OR"
					w.fromRule = r

				case "LSHIFT":
					w2 := addWireIfNotExists(from[0])
					value, err := strconv.Atoi(from[2])
					if err != nil {
						panic(err)
					}
					r.from = append(r.from, w2)
					r.op = "LSHIFT"
					r.value = uint16(value)
					w.fromRule = r

				case "RSHIFT":
					w2 := addWireIfNotExists(from[0])
					value, err := strconv.Atoi(from[2])
					if err != nil {
						panic(err)
					}
					r.from = append(r.from, w2)
					r.op = "RSHIFT"
					r.value = uint16(value)
					w.fromRule = r

				default: //its a wire NOOP
					panic("should not happen here, it's in the len(1) stage")
				}
			}
		}
	}
}

func solve1(lines []string) {
	generateWires(lines)
	// fmt.Println(wires["c"])
	// fmt.Println(wires["c"].value)
	// fmt.Println(wires["c"].eval())

	// fmt.Println("copy a:", copyWires["a"])
	val, _ := wires["a"].eval()
	fmt.Println("solution 1:", val)
	// fmt.Println()
	// fmt.Println("b:", wires["b"])
	// fmt.Println("copy b:", copyWires["b"])
	generateWires(lines)
	wires["b"].value = val
	wires["b"].hasValue = true
	wires["b"].fromRule = nil
	wires["b"].fromWire = nil
	val, _ = wires["a"].eval()
	fmt.Println("solution 2:", val)

	//check links

	// for _, w := range ordered {
	// 	// fmt.Println(w)
	// 	val, err := w.eval()
	// 	if err != nil {
	// 		fmt.Println("possible unknown wire, relinking")
	// 		w.fromWire = wires[w.fromWire.name]
	// 		val, err = w.eval()
	// 		if err != nil {
	// 			panic("should not happen")
	// 		}
	// 	}
	// 	w.value = val
	// 	fmt.Println(w.name+":", val)
	// }
}

func main() {
	lines := readLine("input.txt")
	solve1(lines)
}
