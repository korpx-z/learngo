package main

import "fmt"

//create a new type of deck
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{
		"Spades",
		"Hearts",
		"Diamonds",
		"Clubs",
	}
	cardValues := []string{
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
	}
	//loop through suits and append one of eavh cardValue to//  each suit
	// underscores tell golang that we're aware of unused va// riable
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

//reciever function vvv (method)
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
