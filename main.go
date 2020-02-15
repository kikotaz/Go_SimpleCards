package main

func main() {
	cards := newDeckFromFile("the_cards")
	cards.shuffle()
	cards.print()
	//cards := newDeck()
	//cards.saveToFile("the_cards")

	//hand, remaining := deal(cards, 3)
	//hand.print()
	//remaining.print()

}
