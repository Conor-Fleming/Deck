package main

import (
	"fmt"
	"strings"

	"github.com/Conor-Fleming/deck/deck"
)

type Card = deck.Card

type Hand struct {
	hand   []Card
	dealer bool
}

type Deck []Card

func (h Hand) String() string {
	stringsArr := make([]string, len(h.hand))
	if !h.dealer {
		for i := range h.hand {
			stringsArr[i] = h.hand[i].String()
		}
		return strings.Join(stringsArr, ", ")
	}
	return h.hand[1].String()
}

// version 1
func main() {
	//Decks sets the amount of extra decks to add to the intial deck
	deck := deck.New(deck.ExtraDecks(2), deck.Shuffle)

	// get input for how many players
	// ?

	// call deal function with that amount
	allPlayers, deck := deal(deck, 1)
	for i, v := range allPlayers {
		if i == len(allPlayers)-1 {
			continue
		}
		total, total2 := getTotal(v.hand)
		if total2 == 0 {
			fmt.Printf("Player %v: %v  Total-----> %v\n", i+1, v.String(), total)
			continue
		}
		fmt.Printf("Player %v: %v  Total-----> %v / %v\n", i+1, v.String(), total, total2)
	}
	fmt.Println("Dealer:", allPlayers[len(allPlayers)-1].String())
	fmt.Println()

	var move string
	var drawnCard Card

	//should look into making this into game loop function?
	for i := 0; i < len(allPlayers)-1; i++ {
		iter := 0
		fmt.Printf("Player %v's turn\n", i+1)
		for {
			if iter > 0 {
				total, total2 := getTotal(allPlayers[i].hand)
				if total2 == 0 {
					fmt.Printf("Player %v: %v  Total-----> %v\n", i+1, allPlayers[i].String(), total)
				} else {
					fmt.Printf("Player %v: %v  Total-----> %v / %v\n", i+1, allPlayers[i].String(), total, total2)
				}
			}
			fmt.Println("Hit or stand?")
			iter++
			fmt.Scanf("%s", &move)
			move = strings.ToLower(move)

			if move == "hit" {
				drawnCard, deck = drawCard(deck)
				fmt.Println(drawnCard.String())
				fmt.Println()
				allPlayers[i].hand = append(allPlayers[i].hand, drawnCard)
				continue
			}

			if move == "stand" {
				break
			}

			fmt.Println("Invalid entry")
		}
	}

	//need function to compare players hands with dealers hands and determine winner
}

func getTotal(hand []deck.Card) (int, int) {
	var total int
	var total2 int
	var aces int
	for _, v := range hand {
		if v.Value == 1 { //Ace card
			aces += 1
		}
		total += int(v.Value)
	}

	//this will need to be updated to account for multiple aces and the potential scores
	if aces == 0 {
		total2 = 0
	} else {
		total2 = total + (9 * aces)
	}

	return total, total2
}

func deal(deck Deck, players int) ([]*Hand, Deck) {
	// creating players and dealer hands
	HandArr := make([]*Hand, 0)

	for i := 0; i < players; i++ {
		player := Hand{
			hand:   nil,
			dealer: false,
		}
		HandArr = append(HandArr, &player)
	}
	dealer := Hand{
		hand:   nil,
		dealer: true,
	}
	HandArr = append(HandArr, &dealer)

	// iterate through players and dealer twice and deal a card on each iteration
	// similar to how a real game would be dealt
	for i := 0; i < 2; i++ {
		for _, v := range HandArr {
			card := deck[0]
			deck = deck[1:]
			v.hand = append(v.hand, card)
		}
	}

	return HandArr, deck
}

func drawCard(deck Deck) (Card, Deck) {
	card := deck[0]
	deck = deck[1:]
	return card, deck
}
