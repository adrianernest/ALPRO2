package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Domino struct {
	sideOne int
	sideTwo int
	value   int
	isBalak bool
}

type DominoSet struct {
	cards [28]Domino
	count int
}

func initializeDominoSet(set *DominoSet) {
	idx := 0
	for a := 0; a <= 6; a++ {
		for b := a; b <= 6; b++ {
			card := Domino{
				sideOne: a,
				sideTwo: b,
				value:   a + b,
				isBalak: a == b,
			}
			set.cards[idx] = card
			idx++
		}
	}
	set.count = 28
}

func shuffleCards(set *DominoSet) {
	rand.Seed(time.Now().UnixNano())
	for i := set.count - 1; i > 0; i-- {
		randomIndex := rand.Intn(i + 1)
		set.cards[i], set.cards[randomIndex] = set.cards[randomIndex], set.cards[i]
	}
}

func drawCard(set *DominoSet) Domino {
	if set.count == 0 {
		return Domino{-1, -1, -1, false}
	}
	set.count--
	return set.cards[set.count]
}

func getCardSide(card Domino, side int) int {
	if side == 1 {
		return card.sideOne
	} else if side == 2 {
		return card.sideTwo
	}
	return -1
}

func getCardValue(card Domino) int {
	return card.value
}

func main() {
	var dominoSet DominoSet

	initializeDominoSet(&dominoSet)
	shuffleCards(&dominoSet)

	fmt.Println("Displaying the first 5 drawn cards:")
	for i := 0; i < 5; i++ {
		card := drawCard(&dominoSet)
		fmt.Printf("Card %d: (%d,%d), Value: %d, IsBalak: %v\n",
			i+1, card.sideOne, card.sideTwo, getCardValue(card), card.isBalak)
	}
}
