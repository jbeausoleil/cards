package main

import "fmt"

// Create a new type of 'deck' which is a slice of strings
type deck []string // 'deck' borrows or extends all properties of strings

// Create and return a new deck of playing cards
func newDeck() deck {
	cards := deck{} // create new variable of type deck

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// ----- Receiver Functions ----- //
// Create a new function that applies to deck and prints out all the cards
func (d deck) print() { // d = actual copy of the deck being used // deck = Any variable of type 'deck' now gets access to the 'print' method
	for i, card := range d {
		fmt.Println(i, card)
	}
	fmt.Println("There are", len(d), "cards in the deck.")
}

func deal(d deck, handSize int) (deck, deck) { // call deal function with two arguments -- d of type deck, and handSize of type int
	return d[:handSize], d[handSize:] // return two values by separation of comma
}
