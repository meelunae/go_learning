package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Slices range syntax: slice[startIndexIncluding:upToNotIncluding]
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	deckStrings := []string(d)
	return strings.Join(deckStrings, ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func loadDeckFromFile(filename string) deck {
	bytes, err := os.ReadFile(filename)

	if err != nil { // Something went wrong
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// returns slice of strings that can be resolved to a deck type
	return strings.Split(string(bytes), ",")
}

// defines pseudo-randomness source from current timestamp using math/rand and time packages
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPos := r.Intn(len(d) - 1)
		// in-place swap syntax
		d[i], d[newPos] = d[newPos], d[i]
	}
}
