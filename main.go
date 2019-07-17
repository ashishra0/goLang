package main

import "fmt"

func main() {
	cards := []string{newCard()}
	cards = append(cards, "ace of spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of diamonds"
}
