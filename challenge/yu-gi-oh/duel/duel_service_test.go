package duel

import (
	"ejercicio-poo/challenge/yu-gi-oh/card"
	"testing"
)

func TestCompareCards(t *testing.T) {
	playerCard := &card.Card{Atk: 1000}
	opponentCard := &card.Card{Atk: 500}
	expectedResult := "player"
	result := CompareCards(playerCard, opponentCard)
	if result != expectedResult {
		t.Errorf("CompareCards \n expected %s \n got %s", expectedResult, result)

	}
}
