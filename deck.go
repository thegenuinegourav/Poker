package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := [4]string{"heart", "spade", "club", "tile"}
	cardValues := [3]string{"Ace", "Two", "Three"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	fmt.Println("\n-----------------------")
	for i, card := range d {
		fmt.Println(i, card)
	}
	fmt.Println("-----------------------\n")
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveDeckToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func readDeckFromFile(filename string) deck {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error ocurred while reading file : ", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(b), ","))
}

func (d deck) shuffle() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for i := range d {
		newPosition := r.Intn(len(d))
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
