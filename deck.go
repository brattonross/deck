package deck

import (
	"math/rand"
	"time"
)

// Suit represents the suit of a card.
type Suit rune

var suits = [...]Suit{'♣', '♦', '♥', '♠'}

// Rank represents the rank of a card.
type Rank string

var ranks = [...]Rank{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

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
		d[0], d[1] = d[1], d[0] // TODO: Randomize this
	default:
		for i := 0; i < len(d)-2; i++ {
			j := random.Intn(len(d)-1-i) + i
			d[i], d[j] = d[j], d[i]
		}
	}
}

// Deal deals a card from the deck, removing it from the deck.
func (d *Deck) Deal() Card {
	c := (*d)[0]
	*d = (*d)[1:]
	return c
}

// CardsLeft returns the number of cards left in the deck.
func (d Deck) CardsLeft() int {
	return len(d)
}
