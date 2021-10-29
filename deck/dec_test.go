package deck_test

import (
	"reflect"
	"testing"

	"github.com/aneshas/holdem/deck"
)

func TestDeckStructureShouldBeValid(t *testing.T) {
	cases := []struct {
		name string
		f    func(*testing.T, []deck.Card)
	}{
		{
			name: "all cards should be different",
			f:    checkDistinct,
		},
		{
			name: "deck contains correct ranks and suits",
			f:    checkRanksAndSuits,
		},
		{
			name: "new deck should be shuffled",
			f:    checkShuffling,
		},
	}

	cards := drawAllCards(t, buildNewDeck(t))

	for _, tc := range cases {
		t.Run(tc.name, func(tt *testing.T) {
			tc.f(tt, cards)
		})
	}
}

func checkDistinct(t *testing.T, cards []deck.Card) {
	m := make(map[deck.Card]bool)

	for _, c := range cards {
		m[c] = true
	}

	if len(m) != 52 {
		t.Error("deck contains duplicate cards")
	}
}

func checkRanksAndSuits(t *testing.T, cards []deck.Card) {
	ranks := []deck.Rank{
		deck.Ace,
		deck.Two,
		deck.Three,
		deck.Four,
		deck.Five,
		deck.Six,
		deck.Seven,
		deck.Eight,
		deck.Nine,
		deck.Ten,
		deck.Jack,
		deck.Queen,
		deck.King,
	}

	for _, r := range ranks {
		checkRank(t, r, cards)
	}

}

func checkRank(t *testing.T, rank deck.Rank, cards []deck.Card) {
	suits := []deck.Suit{
		deck.Spades, deck.Clubs, deck.Diamonds, deck.Hearts,
	}

	for _, s := range suits {
		if !hasRankAndSuit(rank, s, cards) {
			t.Errorf("no %d of %s present in deck", rank, s)
		}
	}
}

func hasRankAndSuit(rank deck.Rank, suit deck.Suit, cards []deck.Card) bool {
	for _, c := range cards {
		if c.Rank == rank && c.Suit == suit {
			return true
		}
	}

	return false
}

func checkShuffling(t *testing.T, cards []deck.Card) {
	newCards := drawAllCards(t, buildNewDeck(t))

	if reflect.DeepEqual(newCards, cards) {
		t.Error("new deck should be shuffled")
	}
}

func buildNewDeck(t *testing.T) *deck.Deck {
	return deck.New()
}

func drawAllCards(t *testing.T, d *deck.Deck) []deck.Card {
	var cards []deck.Card

	for i := 0; i < 52; i++ {
		card, err := d.DrawTop()
		if err != nil {
			t.Fatal("should be able to draw more cards")
		}

		cards = append(cards, card)
	}

	_, err := d.DrawTop()
	if err != deck.ErrNoMoreCards {
		t.Fatal("deck should contain only 52 cards")
	}

	return cards
}
