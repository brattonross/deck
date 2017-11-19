package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Suit represents the suit of a card.
type Suit string

// Suits constants.
const (
	Clubs    Suit = "♣"
	Diamonds Suit = "♦"
	Hearts   Suit = "♥"
	Spades   Suit = "♠"
)

// Enumeration of Suits.
var suits = [...]Suit{Clubs, Diamonds, Hearts, Spades}

// Rank represents the rank of a card.
type Rank string

// Rank constants.
const (
	Ace   Rank = "A"
	Two   Rank = "2"
	Three Rank = "3"
	Four  Rank = "4"
	Five  Rank = "5"
	Six   Rank = "6"
	Seven Rank = "7"
	Eight Rank = "8"
	Nine  Rank = "9"
	Ten   Rank = "10"
	Jack  Rank = "J"
	Queen Rank = "Q"
	King  Rank = "K"
)

// Enumeration of Ranks.
var ranks = [...]Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

// Card represents a single card from the deck.
type Card struct {
	Suit Suit
	Rank Rank
}

func (c Card) String() string {
	return string(c.Suit) + string(c.Rank)
}

// Deck represents a deck of cards.
type Deck []Card

// New creates a new deck of cards.
func New() Deck {
	var d Deck
	for _, s := range suits {
		for _, r := range ranks {
			d = append(d, Card{s, r})
		}
	}
	return d
}

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

// Random returns a random card from the deck.
func (d Deck) Random() Card {
	return d[random.Intn(len(d))]
}

// Shuffle randomizes the order of the cards in the deck.
func (d Deck) Shuffle() {
	switch len(d) {
	case 0, 1:
		return
	case 2:
		// 50/50 chance to swap the two remaining cards
		if random.Intn(1) > 0 {
			d[0], d[1] = d[1], d[0]
		}
	default:
		for i := 0; i < len(d)-2; i++ {
			j := random.Intn(len(d)-1-i) + i
			d[i], d[j] = d[j], d[i]
		}
	}
}

// Deal deals a card from the deck, removing it from the deck.
func (d *Deck) Deal() (Card, error) {
	deck := *d
	if len(deck) == 0 {
		return Card{}, fmt.Errorf("unable to deal card: deck size is zero")
	}
	c := (deck)[0]
	*d = (deck)[1:]
	return c, nil
}

// CardsLeft returns the number of cards left in the deck.
func (d Deck) CardsLeft() int {
	return len(d)
}

// Sort sorts the remaining cards in the deck into order.
func (d Deck) Sort() {
	sort.SliceStable(d, func(i, j int) bool {
		return d[i].Rank < d[j].Rank && d[i].Suit < d[j].Suit
	})
}
