package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Create a new type of 'deck' which is a slice of strings
type deck []string // 'deck' borrows or extends all properties of strings

// ----- Normal Functions ----- //
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

func deal(d deck, handSize int) (deck, deck) { // call deal function with two arguments -- d of type deck, and handSize of type int
	return d[:handSize], d[handSize:] // return two values by separation of comma
}

// ----- Receiver Functions ----- //
// Create a new function that applies to deck and prints out all the cards
func (d deck) print() { // d = actual copy of the deck being used // deck = Any variable of type 'deck' now gets access to the 'print' method
	for i, card := range d {
		fmt.Println(i, card)
	}
	fmt.Println("There are", len(d), "cards in the deck.")
}

func (d deck) toString() string { // signature of string
	return strings.Join(d, ",") // convert back up the chain return a slice of strings
}

func (d deck) saveToFile (fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck { // read from file and create an object of type deck
	bs, err := ioutil.ReadFile(fileName) // bs = string of cards separated by comma
	if err != nil {
		// Option #1 - log the error and return a call to a newDeck()
		// Option #2 - log the error and quit the program
		fmt.Println("Error:", err)
		os.Exit(1) // if Exit() is anything other than 0 then exit
	}
	s := strings.Split(string(bs), ",") // split Ace of Spades, etc. to ["Ace of Spades", "etc."]
	return s // return slice of string as deck
}