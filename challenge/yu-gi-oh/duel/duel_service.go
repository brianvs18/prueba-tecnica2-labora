package duel

import (
	"ejercicio-poo/challenge/yu-gi-oh/card"
	"ejercicio-poo/challenge/yu-gi-oh/common"
	"ejercicio-poo/challenge/yu-gi-oh/common/constants"
	"ejercicio-poo/challenge/yu-gi-oh/player"
	"ejercicio-poo/challenge/yu-gi-oh/turn"
	"ejercicio-poo/challenge/yu-gi-oh/world"
	"fmt"
)

func CompareCards(playerCard, opponentCard *card.Card) string {
	if playerCard.Atk > opponentCard.Atk {
		return "player"
	} else if opponentCard.Atk > playerCard.Atk {
		return "opponent"
	}
	return "draw"
}

func announceResult(winner string, playerCard, opponentCard *card.Card) {
	fmt.Println("Battle in course....")
	common.Sleep(constants.ThreeSeconds)
	switch winner {
	case "player":
		fmt.Printf("You win this round with %s (%d Attack) against %s (%d Attack)!\n", playerCard.Name, playerCard.Atk, opponentCard.Name, opponentCard.Atk)
	case "opponent":
		fmt.Printf("Opponent wins this round with %s (%d Attack) against %s (%d Attack)!\n", opponentCard.Name, opponentCard.Atk, playerCard.Name, playerCard.Atk)
	case "draw":
		fmt.Println("The round is a draw, both cards are equal!")
	}
}

func Start(player *player.Player, opponent *player.Player) {
	currentWorld := world.GetRandomWorld(world.Storage)
	fmt.Printf("Dueling in the %s world!\n", currentWorld.Name)
	common.Sleep(constants.TwoSeconds)
	for {
		if !turn.PlayerTurn(player) {
			fmt.Println("The duel has ended early!")
			common.Sleep(constants.TwoSeconds)
			return
		}
		turn.AITurn(opponent)

		playerBonus := currentWorld.BonusIncrement(player.Card.Attribute)
		opponentBonus := currentWorld.BonusIncrement(opponent.Card.Attribute)
		player.Card.IncrementAtk(playerBonus)
		opponent.Card.IncrementAtk(opponentBonus)

		fmt.Printf("Your card: %s with %d attack (after bonus)\n", player.Card.Name, player.Card.Atk)
		common.Sleep(constants.ThreeSeconds)
		fmt.Printf("Opponent's card: %s with %d attack (after bonus)\n", opponent.Card.Name, opponent.Card.Atk)

		winner := CompareCards(player.Card, opponent.Card)
		if winner != "draw" {
			announceResult(winner, player.Card, opponent.Card)
			common.Sleep(constants.FiveSeconds)
			break
		}

		fmt.Println("The duel is a draw, re-drawing new cards...")
		common.Sleep(constants.FourSeconds)
		//common.ClearScreen()
		player.ClearCard()
		opponent.ClearCard()
	}
}
