package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 12 {
		t.Errorf("Got %d cards.", len(d))
	}
}
