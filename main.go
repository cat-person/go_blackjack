package main

import (
	"fmt"
)

func main() {
	deck := getDefaultDeck()
	deck.Shuffle()

	playersHands := map[string]Hand{}

addPlayersInput:
	for {
		var command string
		fmt.Printf("Add players by name or :end\n")
		fmt.Scan(&command)
		switch command {
		case ":end", "End", "end":
			break addPlayersInput
		default:
			playersHands[command] = Hand{true, []string{}}
		}
	}

	fmt.Println(playersHands)

	dealerHand := Hand{false, []string{}}

	for name, hand := range playersHands {
		hand.Add(deck.Draw())
		playersHands[name] = hand
	}

	dealerHand.Add(deck.Draw())
	for name, hand := range playersHands {
		hand.Add(deck.Draw())
		playersHands[name] = hand
	}
	dealerHand.Add(deck.Draw())

	fmt.Println(playersHands)

	fmt.Printf("Dealer's hand: %v\n", dealerHand)

	playersWhoWon := []string{}

	for name, hand := range playersHands {
		if hand.Evaluate() == 21 {
			// Autowin
			fmt.Printf("%s %v got 21 you have won 🎉\n", name, hand)
			playersWhoWon = append(playersWhoWon, name)
		}
	}

	for _, name := range playersWhoWon {
		delete(playersHands, name)
	}

	if len(playersHands) == 0 {
		fmt.Println("All won")
		return
	}

	playersWhoWonOrBusted := []string{}

	for name, hand := range playersHands {
		finishedHand := handleInputByPlayer(name, hand, &deck)

		playersHands[name] = finishedHand

		if 21 < finishedHand.Evaluate() {
			fmt.Printf("%s busted %v and lost\n", name, finishedHand)
			playersWhoWonOrBusted = append(playersWhoWonOrBusted, name)
		}

		if finishedHand.Evaluate() == 21 {
			fmt.Printf("%s %v got 21 you have won 🎉\n", name, finishedHand)
			playersWhoWonOrBusted = append(playersWhoWonOrBusted, name)
		}
	}

	for _, name := range playersWhoWonOrBusted {
		delete(playersHands, name)
	}

	for dealerHand.Evaluate() < 17 {
		dealerHand.Add(deck.Draw())
	}

	dealerHand.isOpen = true

	fmt.Println(dealerHand)
	fmt.Println(playersHands)

	for name, hand := range playersHands {
		if 21 < dealerHand.Evaluate() {
			fmt.Printf("Dealer busted %v, %s won!\n", name, dealerHand)
		} else if dealerHand.Evaluate() < hand.Evaluate() {
			fmt.Printf("%s won with hand: %v vs dealer's hand %v\n", name, hand, dealerHand)
		} else {
			fmt.Printf("%s Lost, your hand: %v vs dealer's hand %v\n", name, hand, dealerHand)
		}
	}
}

func handleInputByPlayer(name string, hand Hand, deck *Deck) Hand {
	for {
		var command string
		fmt.Printf("%s hand %v, Hit or stay ?\n", name, hand)
		fmt.Scan(&command)

		switch command {
		case "Hit", "hit":
			hand.Add(deck.Draw())
			if 21 <= hand.Evaluate() {
				return hand
			} else {
				fmt.Printf("Your hand: %v\n", hand)
			}
		case "Stay", "stay":
			return hand
		default:
			fmt.Print("Please enter one of the commands: ")
		}
	}
}
