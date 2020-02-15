package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Creating a new type of "deck"
type deck []string

//Creating a new deck to play
func newDeck() deck {
	//Creating new deck of cards containing all the 52 cards
	cards := deck{}

	//Defining two slices of Numbers and Sets to create a deck from them
	deckNumbers := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight",
		"Nine", "Ten", "Jack", "Queen", "King"}
	deckSets := []string{"Spades", "Hearts", "Diamonds", "Clubs"}

	//Looping through the slices to create the deck
	for _, set := range deckSets {
		for _, number := range deckNumbers {
			cards = append(cards, number+" of "+set)
		}
	}

	return cards
}

//This function will print the cards in the deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//Converting the deck back to slice of strings
func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

//Saving the deck to file as text by using byte slice
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

//Getting a new deck from a saved file
func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	cardsFromFile := strings.Split(string(bs), ", ")
	return deck(cardsFromFile)
}

func (d deck) shuffle() {
	//Creating new random source with new seed
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	//Performing the shuffling with the new random seed
	for i := range d {
		newIndex := r.Intn(len(d) - 1)

		d[i], d[newIndex] = d[newIndex], d[i]
	}
}
