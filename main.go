package main

func main() {
	cards := deck{newCard()}
	cards = append(cards, "ace of spades")
	cards.print()
}

func newCard() string {
	return "Five of diamonds"
}
