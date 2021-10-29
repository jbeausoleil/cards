package main // executable type package

////// notes //////

// Array = fixed length list of things
// Slice = an array that can grow or shrink
// Array and slice must contain data of the same type

//////////////////

func main() { // automatically called when running project
	// "string" is telling the go compiler that only a data type of string will be assigned to card
	// "string" is an example of a statically typed language
	// basic Go types: bool, string, int, float64
	// var card string = "Ace of Spades" // declares and assigns value to variable card // equivalent to variable := datatype
	//card := "Ace of Spades" // determine / infer what datatype to be assigned
	//card := newCard() // determine / infer what datatype to be assigned
	// do not need to use colon after initialization of variable
	// card = "Five of Diamonds"
	// cards := []string{"Ace of Diamonds", newCard()} // slice of type string
	// cards := deck{"Ace of Diamonds", newCard()} // replace []string with deck due to declaration in deck.go
	// cards = append(cards, "Six of Spades") // adds string cards // does not modify cards, but creates a new cards slice

	//cards := newDeck()

	// iterate over slice
	// for i, card := range cards { // i = index of this element in the array // card = current card of iteration // range cards = take the slice of "cards" and loop over iot
	// 	 fmt.Println(i, card)
	// }
	//cards.print() // replace loop with deck.go function print() because cards is of type deck (receiver)
	//fmt.Println(cards) // print out variable to terminal
	//hand, remainingCards := deal(cards, 5)
	//hand.print()
	//remainingCards.print()
	//err := cards.saveToFile("my_file")
	//if err != nil {
	//	log.Fatal(err)
	//}
	cards := newDeckFromFile("my_file")
	cards.shuffle()
	cards.print()
}

//func newCard() string { // inform the GO compiler that we will return data type "string"
//	return "Five of Diamonds"
//}
