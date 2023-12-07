package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func readLines(path string) []string {
	c, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(c), "\n")
}

// all cards are the same
func fiveOfKind(sortedCards []byte) bool {
	for i := 0; i < 4; i++ {
		if sortedCards[i] != sortedCards[i+1] {
			return false
		}
	}
	return true
}

// there are 4 cards of the same kind
// aaaab
// baaaa
// aabaa
func fourOfKind(sortedCards []byte) bool {
	if sortedCards[0] == sortedCards[1] && sortedCards[1] == sortedCards[2] && sortedCards[2] == sortedCards[3] {
		return true
	}
	if sortedCards[1] == sortedCards[2] && sortedCards[2] == sortedCards[3] && sortedCards[3] == sortedCards[4] {
		return true
	}
	return false
}

// 3 of a kind and a pair
func fullHouse(sortedCards []byte) bool {
	//aaabb
	//bbaaa
	if sortedCards[0] == sortedCards[1] && sortedCards[1] == sortedCards[2] && sortedCards[3] == sortedCards[4] {
		return true
	}
	if sortedCards[0] == sortedCards[1] && sortedCards[2] == sortedCards[3] && sortedCards[3] == sortedCards[4] {
		return true
	}
	return false
}

func threeOfAKind(sortedCards []byte) bool {
	//aaabb
	//bbaaa
	//baaac
	if sortedCards[0] == sortedCards[1] && sortedCards[1] == sortedCards[2] {
		return true
	}
	if sortedCards[2] == sortedCards[3] && sortedCards[3] == sortedCards[4] {
		return true
	}
	if sortedCards[1] == sortedCards[2] && sortedCards[2] == sortedCards[3] {
		return true
	}
	return false
}

func twoPairs(sortedCards []byte) bool {
	if sortedCards[0] == sortedCards[1] && sortedCards[2] == sortedCards[3] {
		return true
	}
	if sortedCards[0] == sortedCards[1] && sortedCards[3] == sortedCards[4] {
		return true
	}
	if sortedCards[1] == sortedCards[2] && sortedCards[3] == sortedCards[4] {
		return true
	}
	return false
}

func onePair(sortedCards []byte) bool {
	for i := 0; i < 4; i++ {
		if sortedCards[i] == sortedCards[i+1] {
			return true
		}
	}
	return false
}

// all cards are different
func highCard(sortedCards []byte) bool {
	last := sortedCards[0]
	for i := 1; i < 5; i++ {
		if last == sortedCards[i] {
			return false
		}
	}
	return true
}

type Card struct {
	value string
	bet   int
}

var cardsValues = []byte{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}
var orderChecks = []func([]byte) bool{fiveOfKind, fourOfKind, fullHouse, threeOfAKind, twoPairs, onePair, highCard}

func solve1(l []string) {
	cards := make([]Card, len(l))
	for i, line := range l {
		game := strings.Split(line, " ")
		// card := []byte(game[0])
		card := game[0]
		bet, _ := strconv.Atoi(game[1])
		cards[i] = Card{card, bet}
	}
	sort.Slice(cards, func(i, j int) bool {
		iValue := []byte(cards[i].value)
		iValueSorted := []byte(cards[i].value)
		sort.Slice(iValueSorted, func(i, j int) bool { return iValueSorted[i] < iValueSorted[j] })

		jValue := []byte(cards[j].value)
		jValueSorted := []byte(cards[j].value)
		sort.Slice(jValueSorted, func(i, j int) bool { return jValueSorted[i] < jValueSorted[j] })

		for k := 0; k < len(orderChecks); k++ {
			iCheck := orderChecks[k](iValueSorted)
			jCheck := orderChecks[k](jValueSorted)
			if iCheck && jCheck {
				for k := 0; k < 5; k++ {
					if iValue[k] != jValue[k] {
						for l := 0; l < len(cardsValues); l++ {
							if cardsValues[l] == iValue[k] {
								//i wins
								return true
							} else if cardsValues[l] == jValue[k] {
								//j wins
								return false
							}
						}
					}
				}
			} else if iCheck && !jCheck {
				return true
			} else if !iCheck && jCheck {
				return false
			} else {
				continue
			}
		}
		fmt.Println("should not be here")
		return true
	})

	slices.Reverse[[]Card](cards)
	total := 0
	for i, card := range cards {
		// fmt.Println(card.bet, "*", (i + 1))
		total += card.bet * (i + 1)
	}
	fmt.Println(total)
}

var cardsValues2 = []byte{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'}

// returns the best hand, the return value is the indes of the orderChecks that satisifies the hand
// returns -1 if there are no jokers
func getBestHandForJokers(sortedCards []byte) int {
	//J are wildcards so they can assume any value
	//QJJQ2 => QQQQ2 4 of a kind
	//QJJJ2 => QQQQ2 4 of a kind

	//no jokers
	if !slices.Contains[[]byte, byte](sortedCards, 'J') {
		return -1
	}

	sortedCardsString := string(sortedCards)

	jokers := strings.Count(sortedCardsString, "J")
	if jokers == 5 || jokers == 4 {
		return 0 //five of a kind
	}

	fullH := fullHouse(sortedCards)

	if jokers == 3 {
		//if full house -> 5 of a kind
		//if 3 of a kind -> 4 of a kind
		if fullH {
			return 0 //five of a kind
		} else {
			return 1 //four of a kind
		}
	}

	two := twoPairs(sortedCards)

	if jokers == 2 {
		//ABCJJ 1 pair -> 3 of a kind
		//AABJJ 2 pairs -> 4 of a kind
		//AAAJJ full house -> 5 of a kind
		if fullH {
			return 0 //five of a kind
		} else if two {
			return 1 //four of a kind
		} else {
			return 3 //three of a kind
		}
	}

	if jokers == 1 {
		//ABCDJ high card -> 1 pair
		//AABCJ 1 pair -> 3 of a kind
		//AABBJ 2 pairs -> full house
		//AAABJ 3 of a kind -> 4 of a kind
		//AAAAJ 4 of a kind -> 5 of a kind
		if fourOfKind(sortedCards) {
			return 0 //five of a kind
		} else if threeOfAKind(sortedCards) {
			return 1 //four of a kind
		} else if two {
			return 2 //full house
		} else if onePair(sortedCards) {
			return 3 //three of a kind
		} else {
			//high card
			return 5 //one pair
		}
	}

	fmt.Println("should not be here")
	return -1
}

func solve2(l []string) {
	cards := make([]Card, len(l))
	for i, line := range l {
		game := strings.Split(line, " ")
		// card := []byte(game[0])
		card := game[0]
		bet, _ := strconv.Atoi(game[1])
		cards[i] = Card{card, bet}
	}
	sort.Slice(cards, func(i, j int) bool {
		iValue := []byte(cards[i].value)
		iValueSorted := []byte(cards[i].value)
		sort.Slice(iValueSorted, func(i, j int) bool { return iValueSorted[i] < iValueSorted[j] })
		bestHandI := getBestHandForJokers(iValueSorted)

		jValue := []byte(cards[j].value)
		jValueSorted := []byte(cards[j].value)
		sort.Slice(jValueSorted, func(i, j int) bool { return jValueSorted[i] < jValueSorted[j] })
		bestHandJ := getBestHandForJokers(jValueSorted)

		if bestHandI == -1 {
			for k := 0; k < len(orderChecks); k++ {
				// fmt.Println(len(iValueSorted), len(jValueSorted))
				iCheck := orderChecks[k](iValueSorted)
				if iCheck {
					bestHandI = k
					break
				}
			}
		}

		if bestHandJ == -1 {
			for k := 0; k < len(orderChecks); k++ {
				jCheck := orderChecks[k](jValueSorted)
				if jCheck {
					bestHandJ = k
					break
				}
			}
		}

		// if bestHandI != -1 && bestHandJ != -1 {
		if bestHandI < bestHandJ {
			return true
		} else if bestHandI > bestHandJ {
			return false
		} else {
			for k := 0; k < 5; k++ {
				if iValue[k] != jValue[k] {
					for l := 0; l < len(cardsValues2); l++ {
						// fmt.Println(string(iValue[k]), string(jValue[k]), string(cardsValues2[l]))
						if cardsValues2[l] == iValue[k] {
							//i wins
							return true
						} else if cardsValues2[l] == jValue[k] {
							//j wins
							return false
						}
					}
				}
			}
		}
		fmt.Println("should not be here")
		return true
	})

	slices.Reverse[[]Card](cards)
	total := 0
	for i, card := range cards {
		total += card.bet * (i + 1)
	}
	fmt.Println(total)
}

func main() {
	l := readLines("input.txt")
	solve1(l)
	solve2(l)
}
