package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "King", "Queen"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (cards deck) print() {
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func deal(cards deck, handSize int) (deck, deck) {
	return cards[:handSize], cards[handSize:]
}

func (cards deck) toString() string {
	return strings.Join([]string(cards), ",")
}

func (cards deck) toFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(cards.toString()), 0666)
}

func fromFile(fileName string) deck {
	deckString, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	s := strings.Split(string(deckString), ",")
	return deck(s)
}

func (cards deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range cards {
		newPosition := r.Intn(len(cards) - 1)
		cards[i], cards[newPosition] = cards[newPosition], cards[i]
	}
}
