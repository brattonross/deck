package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Suit represents the suit of a card.
type Suit struct {
	Name   string // The name of the Suit (e.g. Clubs)
	Symbol string // The symbol of the Suit (e.g. ♣)
}

// Suits.
var (
	Clubs    = Suit{Name: "Clubs", Symbol: "♣"}
	Diamonds = Suit{Name: "Diamonds", Symbol: "♦"}
	Hearts   = Suit{Name: "Hearts", Symbol: "♥"}
	Spades   = Suit{Name: "Spades", Symbol: "♠"}
)

// Enumeration of Suits.
var suits = [...]Suit{Clubs, Diamonds, Hearts, Spades}

// Rank represents the rank of a card.
type Rank struct {
	Name   string // The name of the Rank (e.g. Ace)
	Symbol string // The symbol of the Rank (e.g. A)
	Value  int    // The value of this rank
}

// Ranks.
var (
	Ace   = Rank{Name: "Ace", Symbol: "A", Value: 1} // TODO: Implement Aces High
	Two   = Rank{Name: "2", Symbol: "2", Value: 2}
	Three = Rank{Name: "3", Symbol: "3", Value: 3}
	Four  = Rank{Name: "4", Symbol: "4", Value: 4}
	Five  = Rank{Name: "5", Symbol: "5", Value: 5}
	Six   = Rank{Name: "6", Symbol: "6", Value: 6}
	Seven = Rank{Name: "7", Symbol: "7", Value: 7}
	Eight = Rank{Name: "8", Symbol: "8", Value: 8}
	Nine  = Rank{Name: "9", Symbol: "9", Value: 9}
	Ten   = Rank{Name: "10", Symbol: "10", Value: 10}
	Jack  = Rank{Name: "Jack", Symbol: "J", Value: 11}
	Queen = Rank{Name: "Queen", Symbol: "Q", Value: 12}
	King  = Rank{Name: "King", Symbol: "K", Value: 13}
)

// Enumeration of Ranks.
var ranks = [...]Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

// Card represents a single card from the deck.
type Card struct {
	Suit Suit
	Rank Rank
}

// Symbols returns the symbols of the Rank and Suit (e.g. A♠).
func (c Card) Symbols() string {
	return c.Rank.Symbol + c.Suit.Symbol
}

func (c Card) String() string {
	return fmt.Sprintf("%s of %s", c.Rank.Name, c.Suit.Name)
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
			d.Swap(0, 1)
		}
	default:
		for i := 0; i < len(d)-2; i++ {
			j := random.Intn(len(d)-1-i) + i
			d.Swap(i, j)
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

// Len returns the current number of cards in the deck.
// It is also one of the methods used to implement sort.Interface.
func (d Deck) Len() int {
	return len(d)
}

// Swap swaps the position of two cards in the deck.
// It is also one of the methods used to implement sort.Interface.
func (d Deck) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

// Less determines if the card at index i is lower in value than the card at index j.
// If two cards are the same rank but different suits, the choice is based on alphabetical order of the suits.
// Less is also one of the methods used to implement sort.Interface.
func (d Deck) Less(i, j int) bool {
	switch {
	case d[i].Suit.Symbol < d[j].Suit.Symbol:
		return true
	case d[i].Suit == d[j].Suit:
		return d[i].Rank.Value < d[j].Rank.Value
	}
	return false
}

// Sort sorts the remaining cards in the deck into order.
func (d Deck) Sort() {
	sort.Sort(d)
}
