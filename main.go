package main

import (
	"fmt"

	"github.com/Conor-Fleming/deck/deck"
)

type Card = deck.Card

type Player struct {
	Dealer bool
	Hand   []deck.Card
}

// version 1
func main() {
	//Decks sets the amount of extra decks to add to the intial deck
	decks := 2
	deck := deck.New(deck.ExtraDecks(decks), deck.Shuffle)
	d, p1 := deal(deck)
	dealer := Player{
		Dealer: true,
		Hand:   d,
	}
	player1 := Player{
		Dealer: false,
		Hand:   p1,
	}

	fmt.Println("The dealers hand:\n", dealer.Hand[1:])
	fmt.Println("Player 1's hand:\n", player1.Hand)

	//player can hit or stand
	fmt.Println("Press 'h' to hit or 's' to stay")

}

func deal(deck []Card) ([]Card, []Card) {
	var p1 []Card
	var dealer []Card
	p1 = append(p1, deck[0])
	dealer = append(dealer, deck[1])
	p1 = append(p1, deck[2])
	dealer = append(dealer, deck[3])
	return dealer, p1
}

/*
fmt.Println("-----BlackJack-----")
	//Decks sets the amount of extra decks to add to the intial deck
	decks := 2

	quit := false
	//main game loop
	for !quit {
		//read hit or stay from user

		//

		fmt.Println("Press enter to play again or press q to quit")
		var input string
		fmt.Scan(&input)
		if input == "q" {
			quit = true
			fmt.Println("Thanks for playing!")
			break
		}
		if input == "\n" {
			fmt.Println("continue game")
		}
		fmt.Println("invalid input")
	}
}

func deal(deck []Card) ([]Card, []Card) {
	var p1 []Card
	var dealer []Card
	p1 = append(p1, deck[0])
	dealer = append(dealer, deck[1])
	p1 = append(p1, deck[2])
	dealer = append(dealer, deck[3])
	return dealer, p1
}

func gameLoop() {
	//deal cards
	deck := deck.New(deck.ExtraDecks(decks), deck.Shuffle)
	d, p1 := deal(deck)
	dealer := Player{
		Dealer: true,
		Hand:   d,
	}
	player1 := Player{
		Dealer: false,
		Hand:   p1,
	}

	fmt.Println("The dealers hand:\n", dealer.Hand[1:])
	fmt.Println("Player 1's hand:\n", player1.Hand)

	//player can hit or stand
	fmt.Println("Press 'h' to hit or 's' to stay")
}

func getHandTotal(hand []Card) int {
	total := 0
	for _, v := range hand {
		total += int(v.Value)
	}
	return total
}

*/
