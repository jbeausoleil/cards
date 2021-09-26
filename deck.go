package main

import "fmt"

// Create a new type of 'deck' which is a slice of strings
type deck []string // 'deck' borrows or extends all properties of strings

// ----- Receiver Functions ----- //
// Create a new function that applies to 'deck' and prints out all the cards
func (d deck) print() { // d = actual copy of the deck being used // deck = Any variable of type 'deck' now gets access to the 'print' method
	for i, card := range d {
		fmt.Println(i, card)
	}
}