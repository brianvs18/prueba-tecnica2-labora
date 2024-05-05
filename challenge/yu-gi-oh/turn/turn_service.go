package turn

import (
	"bufio"
	"ejercicio-poo/challenge/yu-gi-oh/common"
	"ejercicio-poo/challenge/yu-gi-oh/common/constants"
	"ejercicio-poo/challenge/yu-gi-oh/player"
	"fmt"
	"os"
	"strings"
)

func PlayerTurn(player *player.Player) bool {
	fmt.Println("It's your turn. Type 'draw' to draw a card or 'quit' to end the duel:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())

	switch input {
	case "draw":
		common.ClearScreen()
		player.DrawCard()
		common.Sleep(constants.TwoSeconds)
		fmt.Printf("You drew %s with %d attack \n", player.Card.Name, player.Card.Atk)
		return true
	case "quit":
		common.ClearScreen()
		fmt.Println("You chose to quit the duel.")
		common.Sleep(constants.TwoSeconds)
		return false
	default:
		fmt.Println("Invalid input. Please type 'draw' or 'quit'.")
		common.Sleep(constants.OneSecond)
		common.ClearScreen()
		return PlayerTurn(player)
	}
}

func AITurn(ai *player.Player) {
	ai.DrawCard()
	fmt.Printf("Opponent draws %s with %d attack\n", ai.Card.Name, ai.Card.Atk)
}
