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

func TestLen(t *testing.T) {
	d := deck.New()
	left := d.Len()
	if left != 52 {
		t.Fatalf("expected 52 cards left, got %d", left)
	}
	d.Deal()
	left = d.Len()
	if left != 51 {
		t.Fatalf("expected 51 cards left, got %d", left)
	}
}

func TestLess(t *testing.T) {
	tt := []struct {
		name     string
		d        deck.Deck // A deck which should contain two cards to compare
		expected bool      // The expected result of Less(0, 1)
	}{
		{
			name: "A-Clubs, 2-Clubs",
			d: deck.Deck{
				deck.Card{Rank: deck.Ace, Suit: deck.Clubs},
				deck.Card{Rank: deck.Two, Suit: deck.Clubs},
			},
			expected: true,
		},
		{
			name: "A-Clubs, A-Diamonds",
			d: deck.Deck{
				deck.Card{Rank: deck.Ace, Suit: deck.Clubs},
				deck.Card{Rank: deck.Ace, Suit: deck.Diamonds},
			},
			expected: true,
		},
		{
			name: "2-Clubs, A-Diamonds",
			d: deck.Deck{
				deck.Card{Rank: deck.Two, Suit: deck.Clubs},
				deck.Card{Rank: deck.Ace, Suit: deck.Diamonds},
			},
			expected: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			a := tc.d.Less(0, 1)
			if a != tc.expected {
				t.Fatalf("expected %v, got %v", tc.expected, a)
			}
			// Test that the comparison also works in the opposite direction
			a = tc.d.Less(1, 0)
			if a == tc.expected {
				t.Fatalf("expected %v, got %v", !tc.expected, a)
			}
		})
	}
}

func TestValue(t *testing.T) {
	d := deck.New()
	expected := 1
	for _, c := range d {
		if c.Rank.Value != expected {
			t.Fatalf("expected value of %d, got %d", expected, c.Rank.Value)
		}
		expected++
		if expected > 13 {
			expected = 1
		}
	}
}
