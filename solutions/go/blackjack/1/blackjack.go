package blackjack

import "slices"

var cardValues = map[string]int{
	"ace":   11,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {

	return cardValues[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	if card1 == "ace" && card2 == "ace" {
		return "P"
	}
	specialCards := []string{"ace", "ten", "jack", "queen", "king"}
	myCardSum := cardValues[card1] + cardValues[card2]

	switch {
	case myCardSum == 21:
		if !slices.Contains(specialCards, dealerCard) {
			return "W"
		} else {
			return "S"
		}

	case myCardSum >= 17 && myCardSum <= 20:
		return "S"
	case myCardSum >= 12 && myCardSum <= 16:
		if cardValues[dealerCard] >= 7 {
			return "H"
		} else {
			return "S"
		}
	default:
		return "H"

	}

}
