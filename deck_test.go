package deck_test

import (
	"reflect"
	"testing"

	"github.com/brattonross/deck"
)

func TestShuffle(t *testing.T) {
	tt := []struct {
		name string
		a    deck.Deck
		e    deck.Deck
		fn   func(a, e deck.Deck) // A function which sets up and compares two decks, failing accordingly
	}{
		{
			name: "FullDeck",
			a:    deck.New(),
			e:    deck.New(),
			fn: func(a deck.Deck, e deck.Deck) {
				a.Shuffle()
				if len(a) != len(e) {
					t.Fatalf("expected length of %d, got %d", len(e), len(a))
				}
				if reflect.DeepEqual(a, e) {
					t.Fatalf("expected: %v, got %v", e, a)
				}
			},
		},
		{
			name: "OneLeft",
			a:    deck.Deck{deck.Card{Suit: deck.Clubs, Rank: deck.Ace}},
			e:    deck.Deck{deck.Card{Suit: deck.Clubs, Rank: deck.Ace}},
			fn: func(a deck.Deck, e deck.Deck) {
				a.Shuffle()
				if !reflect.DeepEqual(a, e) {
					t.Fatalf("expected: %v, got %v", e, a)
				}
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tc.fn(tc.a, tc.e)
		})
	}
}

func TestDeal(t *testing.T) {
	d := deck.New()
	c, _ := d.Deal()
	if len(d) == 52 {
		t.Fatal("size of deck was not altered after dealing")
	}
	if c == (deck.Card{}) {
		t.Fatal("dealt card was nil")
	}
}

func TestCardsLeft(t *testing.T) {
	d := deck.New()
	left := d.CardsLeft()
	if left != 52 {
		t.Fatalf("expected 52 cards left, got %d", left)
	}
	d.Deal()
	left = d.CardsLeft()
	if left != 51 {
		t.Fatalf("expected 51 cards left, got %d", left)
	}
}

func TestSort(t *testing.T) {

}
