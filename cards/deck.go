package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	d := deck{}
	cardSuits := []string{"Spade", "Heart", "Diamond", "Club"}
	cardsValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardsValues {
			d = append(d, value+" of "+suit)
		}
	}
	return d
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	err := os.WriteFile(filename, []byte(d.toString()), 0666)
	return err
}

func loadDeckFromFile(filename string) (deck, error) {
	d := deck{}
	content, err := os.ReadFile(filename)
	if err != nil {
		return d, err
	}
	d = strings.Split(string(content), ",")
	return d, err
}

func (d deck) shuffle() {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}
