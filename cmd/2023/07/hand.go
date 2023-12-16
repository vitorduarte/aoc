package main

var cardValuesWithoutJoker = map[string]int{
	"A": 12,
	"K": 11,
	"Q": 10,
	"J": 9,
	"T": 8,
	"9": 7,
	"8": 6,
	"7": 5,
	"6": 4,
	"5": 3,
	"4": 2,
	"3": 1,
	"2": 0,
}

var cardValuesWithJoker = map[string]int{
	"A": 12,
	"K": 11,
	"Q": 10,
	"T": 8,
	"9": 7,
	"8": 6,
	"7": 5,
	"6": 4,
	"5": 3,
	"4": 2,
	"3": 1,
	"2": 0,
	"J": -1,
}

type Hand struct {
	cards string
	bid   int
	joker bool
}

type Hands []Hand

func (h *Hand) getHandValue() int {
	cardToCount := countCards(h.cards, h.joker)

	switch len(cardToCount) {
	case 1:
		// AAAAA - Five of a kind
		return 6
	case 2:
		for _, count := range cardToCount {
			// AA8AA - Four of a kind
			if count == 4 || count == 1 {
				return 5
			}

			// 23332 - Full house
			if count == 3 || count == 2 {
				return 4
			}
		}
	case 3:
		for _, count := range cardToCount {
			// TTT98 - Three of a kind
			if count == 3 {
				return 3
			}
			// 23432 - Two pair
			if count == 2 {
				return 2
			}
		}
	case 4:
		// 23456 - One pair
		return 1
	}

	// 23456 - High card
	return 0
}

func (a Hands) Len() int           { return len(a) }
func (a Hands) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Hands) Less(i, j int) bool { return isHandLess(a[i], a[j]) }

func isHandLess(first Hand, second Hand) bool {
	cardValues := cardValuesWithoutJoker
	if first.joker {
		cardValues = cardValuesWithJoker
	}

	fValue := first.getHandValue()
	sValue := second.getHandValue()

	if fValue < sValue {
		return true
	}

	if fValue == sValue {
		for i := 0; i < len(first.cards); i++ {
			fCardValue := cardValues[string(first.cards[i])]
			sCardValue := cardValues[string(second.cards[i])]

			if fCardValue < sCardValue {
				return true
			}

			if fCardValue > sCardValue {
				return false
			}
		}
	}

	return false
}

func countCards(cards string, joker bool) map[rune]int {
	cardToCount := map[rune]int{}

	for _, card := range cards {
		count, ok := cardToCount[card]
		if ok {
			cardToCount[card] = count + 1
			continue
		}

		cardToCount[card] = 1
	}

	if joker {
		return replaceJoker(cardToCount)
	}

	return cardToCount
}

func replaceJoker(cardToCount map[rune]int) map[rune]int {
	newCardToCount := map[rune]int{}

	jokerRune := 'J'
	jokerCount := 0

	var maxCard rune
	maxCardCount := 0

	for card, count := range cardToCount {
		if card != jokerRune {
			newCardToCount[card] = count

			if count > maxCardCount {
				maxCardCount = count
				maxCard = card
			}
			continue
		}

		jokerCount = count
	}

	newCardToCount[maxCard] = maxCardCount + jokerCount

	return newCardToCount
}
