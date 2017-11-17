package deck_test

import (
	"reflect"
	"testing"

	"github.com/brattonross/deck"
)

func TestShuffle(t *testing.T) {
	d := deck.New()
	e := deck.New()
	d.Shuffle()
	if reflect.DeepEqual(d, e) {
		t.Fatalf("shuffled and non-shuffled deck match")
	}
}

func TestDeal(t *testing.T) {
	d := deck.New()
	c := d.Deal()
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
