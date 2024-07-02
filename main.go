package main

func main() {
	//cards := deck{newCard(), newCard(), "Ace of Diamonds"}
	//// does not modify the slice but creates a new one
	//cards = append(cards, "Six of Spades")
	//cards := newDeck()
	//hand, remainingDeck := deal(cards, 3)
	//hand.print()
	//remainingDeck.print()
	//fmt.Println(cards.toString())
	//cards.saveToFile("my_cards")
	//cards := newDeckFromFile("my")
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
