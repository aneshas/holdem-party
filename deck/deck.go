package deck

import (
	"errors"
	"math/rand"
	"time"
)

var (
	ErrNoMoreCards = errors.New("no more cards left in the deck")
)

type Rank int

const (
	Ace Rank = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	None
	Jack
	Queen
	King
)

type Suit string

const (
	Spades   Suit = "SPADES"
	Hearts   Suit = "HEARTS"
	Diamonds Suit = "DIAMONDS"
	Clubs    Suit = "CLUBS"
)

type Card struct {
	Rank Rank
	Suit Suit
}

func New() *Deck {
	cards := getCards()

	shuffle(cards)

	return &Deck{
		cards: cards,
	}
}

func getCards() []Card {
	var cards []Card

	ranks := []Rank{
		Ace,
		Two,
		Three,
		Four,
		Five,
		Six,
		Seven,
		Eight,
		Nine,
		Ten,
		Jack,
		Queen,
		King,
	}

	for _, rank := range ranks {
		cards = append(cards, get(rank)...)
	}

	return cards
}

func get(rank Rank) []Card {
	return []Card{
		{
			Suit: Spades,
			Rank: rank,
		},
		{
			Suit: Hearts,
			Rank: rank,
		},
		{
			Suit: Diamonds,
			Rank: rank,
		},
		{
			Suit: Clubs,
			Rank: rank,
		},
	}
}

func shuffle(cards []Card) {
	for i := 0; i < 7; i++ {
		for i := 0; i < 52; i++ {
			s := rand.NewSource(time.Now().UnixNano())
			r := rand.New(s)

			other := r.Int31n(52)

			temp := cards[other]
			cards[other] = cards[i]
			cards[i] = temp
		}
	}
}

type Deck struct {
	cards []Card
}

func (d *Deck) DrawTop() (Card, error) {
	if len(d.cards) <= 0 {
		return Card{}, ErrNoMoreCards
	}

	nc := d.cards[0]
	d.cards = d.cards[1:]

	return nc, nil
}
