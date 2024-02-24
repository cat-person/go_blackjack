package main

import (
	"math/rand"
)

var defaultDeck = []string{
	"A", "A", "A", "A",
	"A", "A", "A", "A",
	"K", "K", "K", "K",
	"Q", "Q", "Q", "Q",
	"J", "J", "J", "J",
	"10", "10", "10", "10",
	"9", "9", "9", "9",
	"8", "8", "8", "8",
	"7", "7", "7", "7",
	"6", "6", "6", "6",
	"5", "5", "5", "5",
	"4", "4", "4", "4",
	"3", "3", "3", "3",
	"2", "2", "2", "2",
}

// Do not use it directly
type Deck struct {
	cards []string
}

func getDefaultDeck() Deck {
	return Deck{defaultDeck}
}

func (deck *Deck) Shuffle() {
	shuffledDeck := make([]string, len(deck.cards))

	copy(shuffledDeck, deck.cards)

	for i := range shuffledDeck {
		j := rand.Intn(i + 1)
		shuffledDeck[i], shuffledDeck[j] = shuffledDeck[j], shuffledDeck[i]
	}

	deck.cards = shuffledDeck
}

func (deck *Deck) Draw() string {
	cardIdx := len(deck.cards) - 1
	card := deck.cards[cardIdx]
	deck.cards = deck.cards[:cardIdx]
	return card
}
