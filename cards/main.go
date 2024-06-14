package main

import "fmt"

func main() {
	// deck := newDeck()
	// deck.print()
	// deck.saveToFile("_MyCardDeck")

	filename := "_MyCardDeck"
	deck, err := loadDeckFromFile(filename)
	if err != nil {
		fmt.Println("Unbale to load the deck form the file: ", filename)
	}
	deck.print()
	deck.shuffle()
	deck.print()
}
