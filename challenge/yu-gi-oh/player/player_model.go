package player

import "ejercicio-poo/challenge/yu-gi-oh/card"

type Player struct {
	Name string
	Card *card.Card
	IsAI bool
}

func (p *Player) DrawCard() {
	p.Card = card.GetRandomCard()
}

func (p *Player) ClearCard() {
	p.Card = &card.Card{}
}
