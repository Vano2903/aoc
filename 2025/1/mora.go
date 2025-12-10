package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := ReadFile(1)

	lines := SplitLines(input)

	// rotations := []int{}
	pointTo, totLoops := 50, 0
	for _, l := range lines {
		var rot int
		var err error
		if strings.HasPrefix(l, "L") {
			rot, err = strconv.Atoi(strings.ReplaceAll(l, "L", ""))
			if err != nil {
				panic(err)
			}
			rot *= -1
		} else {
			rot, err = strconv.Atoi(strings.ReplaceAll(l, "R", ""))
			if err != nil {
				panic(err)
			}
		}

		startPointTo := pointTo
		loops := 0
		// extra := false
		fmt.Printf("S=% +5d | R=% +5d | E=% +5d", pointTo, rot, pointTo+rot)
		pointTo += rot
		if pointTo < 0 {
			if pointTo > -100 {
				// extra = true
				loops = (pointTo/100)*-1 + 1
			} else {
				loops = (rot / 100) * -1
			}
			pointTo += (loops * 100)
			if pointTo < 0 {
				pointTo = 100 + pointTo

			}
		} else if pointTo > 100 {
			loops = pointTo / 100
			pointTo %= 100
		} else if pointTo == 100 {
			pointTo = 0
		}
		// if extra {
		//   loops++
		// }
		if pointTo == 0 {
			loops++
		}
		if startPointTo == 0 && (rot < 0 || pointTo == 0) {
			loops--
		}
		fmt.Printf(" => % +5d (%v)\n", pointTo, loops)
		totLoops += loops
		// fanculo colione, non ci riesco a farlo così

		//   deduction, induction := 0, 0
		//   if pointTo == 0 {
		//     deduction = 1
		//   }
		//   pointTo += rot
		//   if pointTo == 100 {
		//     induction = 1
		//   }

		//   if pointTo < 0 {
		//     loops += ((rot * -1) / 100) - deduction + induction
		//     pointTo = 100 - ((pointTo * -1) % 100)
		//     fmt.Printf("During % +5d looped %v times on 0 and landed on %v\n", rot, ((rot*-1)/100)+1-deduction+induction, pointTo)
		//   } else if pointTo >= 100 {
		//     loops += pointTo/100 - deduction + induction
		//     pointTo = pointTo % 100
		//     fmt.Printf("During % +5d looped %v times on 100 and landed on %v\n", rot, ((rot)/100)-deduction+induction, pointTo)
		//   } else if pointTo == 0 {
		//     loops++
		//     fmt.Printf("During % +5d landed without loops on 0\n", rot)
		//   } else {
		//     fmt.Printf("During % +5d landed without loops on %v\n", rot, pointTo)
		//   }
	}

	fmt.Printf("\n\nEnded on %v after %v loops\n", pointTo, totLoops)
	fmt.Println("Solution:", totLoops)
	if totLoops != 6858 {
		panic("wrong solution for part 1: " + strconv.Itoa(totLoops))
	}
}

func ReadFile(day int) string {
	file, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	return string(file)
}

func ReadTestFile(day int) string {
	file, err := os.ReadFile("test-day" + fmt.Sprint(day) + ".txt")
	if err != nil {
		panic(err)
	}
	return string(file)
}

func SplitLines(input string) []string {
	return strings.Split(input, "\n")
}
