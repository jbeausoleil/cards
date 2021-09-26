package main // executable type package
import "fmt"

func main() { // automatically called when running project
	// "string" is telling the go compiler that only a data type of string will be assigned to card
	// "string" is an example of a statically typed language
	// basic Go types: bool, string, int, float64
	// var card string = "Ace of Spades" // declares and assigns value to variable card // equivalent to variable := datatype
	card := "Ace of Spades" // determine / infer what datatype to be assigned
	// do not need to use colon after initialization of variable
	card = "Five of Diamonds"
	fmt.Println(card) // print out variable to terminal
}
