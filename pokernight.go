package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type CardRank struct {
	symbol string
	rank   int
}

const (
	FourOfAKind = "Four of a Kind"
	FullHouse   = "Full House"
	Triple      = "Triple"
	TwoPairs    = "Two Pairs"
	APair       = "A Pair"
	HighCard    = "High Card"
)

var combinationRank = map[int]string{1: FourOfAKind, 2: FullHouse, 3: Triple, 4: TwoPairs, 5: APair, 6: HighCard}

func main() {
	fmt.Println("Welcome to Poker Night with LARVIS")
	fmt.Println("Drumrolls")
	time.Sleep(2 * time.Second)
	deck := newDeck()
	//shuffles the deck of cards
	shuffledDeck := shuffle(deck)
	//assigns 5 cards to both players from the top of shuffled deck
	hand1, hand2 := distributeCards(shuffledDeck)
	//checks which hand has higher combination
	compareHands(hand1, hand2)
	// //Print cards and declare winner
	fmt.Println("Hand of Player 1 ", hand1)
	fmt.Println("Hand of Player 2 = ", hand2)
	//fmt.Println("Final Result = ", finalResult)
}
func newDeck() []CardRank {
	cards := "23456789TJQKA"
	deck := make([]CardRank, 0, len(cards))
	r := 2
	for _, symbol := range cards {
		for i := 0; i < 4; i++ {
			deck = append(deck, CardRank{symbol: string(symbol), rank: r})
		}
		r = r + 1
	}
	return deck
}
func shuffle(deck []CardRank) []CardRank {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })
	return deck
}
func distributeCards(shuffledDeck []CardRank) ([5]CardRank, [5]CardRank) {
	var Hand1, Hand2 [5]CardRank
	for i := 0; i < 5; i++ {
		Hand1[i] = shuffledDeck[rand.Intn(len(shuffledDeck))]
		Hand2[i] = shuffledDeck[rand.Intn(len(shuffledDeck))]
	}
	return Hand1, Hand2
}

// Assign rank on the basis of combination values with FourOfAKind=1,FullHouse=1,Triple=3,TwoPairs=4,APair=5 and HighCard=6
func checkCombinationValue(hand [5]CardRank) int {
	frequency := checkFrequencyOfSymbol(hand)
	//store only values in an int array for easy comparison
	f := make([]int, 0, len(frequency))
	for _, value := range frequency {
		f = append(f, value)
	}
	var rank = 0
	switch len(f) {
	//Four of a kind or Full house
	case 2:
		if f[0] == 4 || f[1] == 4 {
			rank = 1
		} else if (f[0] == 3 && f[1] == 2) || (f[0] == 2 && f[1] == 3) {
			rank = 2
		}
	//Triple or Two Pairs
	case 3:
		for i := 0; i < len(f); i++ {
			if f[i] == 3 {
				rank = 3
			} else if f[i] == 2 {
				rank = 4
			}
		}
	//A Pair
	case 4:
		rank = 5
	//High Card
	case 5:
		rank = 6
	}
	return rank
}

// checks occurence of a symbol in hand[]
func checkFrequencyOfSymbol(hand [5]CardRank) map[string]int {
	frequency := make(map[string]int)
	for _, count := range hand {
		frequency[count.symbol]++
	}
	return frequency
}
func compareHands(hand1 [5]CardRank, hand2 [5]CardRank) {
	r1 := checkCombinationValue(hand1)
	fmt.Println("Rank of Player 1 = ", r1)
	r2 := checkCombinationValue(hand2)
	fmt.Println("Rank of Player 2 = ", r2)
	time.Sleep(2 * time.Second)
	if r1 == r2 {
		fmt.Println("It's a TIE. Let's have a TIE Breaker.")
		time.Sleep(2 * time.Second)
		resolveTie(hand1, hand2, r1)
	} else if r1 < r2 {
		//return ("Hand1 wins")
		fmt.Println("Hand1 Wins with", combinationRank[r1])
	} else if r1 > r2 {
		//return ("Hand2 wins")
		fmt.Println("Hand2 wins with", combinationRank[r2])
	}
}

// In case of tie, compares each element of slice for the higher value
func resolveTie(hand1 [5]CardRank, hand2 [5]CardRank, ranktie int) {
	fMap1 := checkFrequencyOfSymbol(hand1)
	fMap2 := checkFrequencyOfSymbol(hand2)
	var highestRank1, highestRank2 int
	switch ranktie {
	case 6:
		sortAndCompare(hand1, hand2)
	case 5:
		for symbol, frequency := range fMap1 {
			if frequency == 2 {
				for _, card := range hand1 {
					if card.symbol == symbol {
						highestRank1 = card.rank
						break
					}
				}
			}
		}
		for symbol, frequency := range fMap2 {
			if frequency == 2 {
				for _, card := range hand2 {
					if card.symbol == symbol {
						highestRank2 = card.rank
						break
					}
				}
			}
		}
		if highestRank1 == highestRank2 {
			sortAndCompare(hand1, hand2)
		} else if highestRank1 > highestRank2 {
			fmt.Println("Player 1 wins with a high card of", highestRank1)
		} else if highestRank2 > highestRank1 {
			fmt.Println("Player 2 wins with a high card of", highestRank2)

		}
	case 4:
	case 3:
	case 2:
	case 1:
		findHighestCard(ranktie, fMap1, fMap2, hand1, hand2)
	}
}

// Sorts the slice and compare each element according to their rank value
func sortAndCompare(hand1 [5]CardRank, hand2 [5]CardRank) {
	sort.Slice(hand1[:], func(i, j int) bool {
		return hand1[i].rank > hand1[j].rank
	})
	sort.Slice(hand2[:], func(i, j int) bool {
		return hand2[i].rank > hand2[j].rank
	})
	for i := 0; i < 5; i++ {
		if hand1[i].rank > hand2[i].rank {
			fmt.Println("Player 1 wins with a high card of", hand1[i].symbol)
			break
		} else if hand1[i].rank < hand2[i].rank {
			fmt.Println("Player 2 wins with a high card of", hand2[i].symbol)
			break
		}
	}
}

// Find highest card by comparing value of cards having same frequency
func findHighestCard(r int, fMap1 map[string]int, fMap2 map[string]int, hand1 [5]CardRank, hand2 [5]CardRank) {
	var highestRank1, highestRank2 int
	for symbol, frequency := range fMap1 {
		if frequency == r {
			for _, card := range hand1 {
				if card.symbol == symbol {
					highestRank1 = card.rank
					break
				}
			}
		}
	}
	for symbol, frequency := range fMap2 {
		if frequency == r {
			for _, card := range hand2 {
				if card.symbol == symbol {
					highestRank2 = card.rank
					break
				}
			}
		}
	}
	if highestRank1 > highestRank2 {
		fmt.Println("Player 1 wins with a high card of", highestRank1)
	} else if highestRank2 > highestRank1 {
		fmt.Println("Player 2 wins with a high card of", highestRank2)
	}
}
