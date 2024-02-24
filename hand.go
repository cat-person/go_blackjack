package main

import (
	"fmt"
	"strings"
)

var cardToValue = map[string]int{
	"A":  11,
	"K":  10,
	"Q":  10,
	"J":  10,
	"10": 10,
	"9":  9,
	"8":  8,
	"7":  7,
	"6":  6,
	"5":  5,
	"4":  4,
	"3":  3,
	"2":  2,
}

type Hand struct {
	isOpen bool
	cards  []string
}

func (hand *Hand) Add(card string) {
	hand.cards = append(hand.cards, card)
}

func (hand *Hand) Evaluate() int {
	result := 0
	acesCount := 0

	for cardIdx := 0; cardIdx < len(hand.cards); cardIdx++ {
		card := hand.cards[cardIdx]
		result += cardToValue[card]
		if card == "A" {
			acesCount++
		}
	}

	if len(hand.cards) == 2 && acesCount == 2 {
		return 21
	}

	if result > 21 {
		required := (result - 12) / 10
		result -= min(required, acesCount) * 10
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (hand Hand) String() string {
	if hand.isOpen {
		return fmt.Sprintf("[%s](%d)", strings.Join(hand.cards, ", "), hand.Evaluate())
	}
	return fmt.Sprintf("[%s](?)", strings.Join(hand.hideAllButFirst(), ", "))
}

func (hand Hand) hideAllButFirst() []string {
	result := make([]string, len(hand.cards))
	result[0] = hand.cards[0]
	for cardIdx := 1; cardIdx < len(result); cardIdx++ {
		result[cardIdx] = "?"
	}
	return result
}
