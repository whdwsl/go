package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "jack", "queen", "king", "ten":
		return 10
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		return 0
	}
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	var card1Value = ParseCard(card1)
	var card2Value = ParseCard(card2)
	switch card1Value + card2Value {
	case 21:
		return true
	default:
		return false
	}

}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	if isBlackjack && dealerScore < 10 {
		return "W"
	} else if !isBlackjack && dealerScore == 11 {
		return "P"
	} else {
		return "S"
	}
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	if handScore > 16 {
		return "S"
	} else if handScore < 12 {
		return "H"
	} else if 12 <= handScore && handScore <= 16 && dealerScore >= 7 {
		return "H"
	} else {
		return "S"
	}
}
