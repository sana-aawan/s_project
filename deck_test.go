package main

import "testing"

func TestnewDeskReadingFile(t *testing.T) {
	d := NewDeck()
	if len(d)!= 16{
		t.Errorf("Expected 16 cards, got" , len(d))
	}
}