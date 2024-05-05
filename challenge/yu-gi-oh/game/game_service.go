package game

import (
	"bufio"
	"ejercicio-poo/challenge/yu-gi-oh/card"
	"ejercicio-poo/challenge/yu-gi-oh/common"
	"ejercicio-poo/challenge/yu-gi-oh/common/constants"
	"ejercicio-poo/challenge/yu-gi-oh/duel"
	"ejercicio-poo/challenge/yu-gi-oh/player"
	"fmt"
	"github.com/ttacon/chalk"
	"os"
	"strconv"
	"strings"
)

func Start() {
	scanner := bufio.NewScanner(os.Stdin)
	var username string

	for {
		common.ClearScreen()
		if username == "" {
			fmt.Println("Input your name!")
			scanner.Scan()
			username = scanner.Text()
		}
		common.ClearScreen()
		fmt.Printf("Your NAME is %s\n YES\n NO\n", username)
		scanner.Scan()
		nameValidation := scanner.Text()

		switch strings.ToLower(nameValidation) {
		case "yes":
			common.ClearScreen()
			common.Sleep(constants.TwoSeconds)
			fmt.Println(chalk.White, "Your duelist code has been recorded.")
			common.Sleep(constants.ThreeSeconds)
			player.CreatePlayer(username)
			prologueConversation()
			common.Sleep(constants.FourSeconds)
			options()
		case "no":
			username = ""
			fmt.Println("Input your name again!")
		default:
			fmt.Println("Invalid selection, please answer with YES or NO.")
		}
	}
}

func prologueConversation() {
	common.ClearScreen()
	common.Sleep(constants.ThreeSeconds)
	fmt.Println(chalk.White, "This is it...")
	common.Sleep(constants.ThreeSeconds)
	common.ClearScreen()
	fmt.Println(chalk.White, "I've found it at last...")
	common.Sleep(constants.ThreeSeconds)
	common.ClearScreen()
	fmt.Println(chalk.White, "The forbidden treasure of the ancient sorcerers...")
	common.Sleep(constants.ThreeSeconds)
	common.ClearScreen()
	fmt.Println(chalk.White, "Hahahahaha!")
	fmt.Println(chalk.White, "Hahahahaha...")
	fmt.Println(chalk.White, "Hahahahaha!!!")
	common.Sleep(constants.ThreeSeconds)
	common.ClearScreen()
	common.Sleep(constants.FourSeconds)
	fmt.Printf("My dear prince!\n Are you going to the city to play cards again!?\n")
	common.Sleep(constants.ThreeSeconds)
	common.ClearScreen()
	fmt.Printf("You are of royal blood!\n Walking the city streets dressed as a commoner..\n")
	common.Sleep(constants.ThreeSeconds)
	common.ClearScreen()
	fmt.Printf("...Have you no shame!?\n Quite frankly, I'm embarrassed!\n")
	common.Sleep(constants.ThreeSeconds)
	common.ClearScreen()
	fmt.Printf("Wait!!\n Stop, my prince!\n")
	common.Sleep(constants.ThreeSeconds)
	common.ClearScreen()
	fmt.Printf("Drat!\n He's gone...\n")
	common.Sleep(constants.ThreeSeconds)
	common.ClearScreen()
}

func options() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		//common.ClearScreen()
		fmt.Println(chalk.Yellow, "Select an option")
		fmt.Println(chalk.White, "1. Duel ground")
		fmt.Println(chalk.White, "2. Card shop")
		fmt.Println(chalk.White, "0. Exit")
		scanner.Scan()
		selection := scanner.Text()
		choice, err := strconv.Atoi(selection)
		if err != nil {
			fmt.Println("Invalid option, please enter a number.")
			continue
		}

		switch choice {
		case 1:
			humanPlayer := player.GetHumanPlayer()
			aiPlayer := player.CreatePlayer("COMP")
			aiPlayer.IsAI = true
			duel.Start(humanPlayer, aiPlayer)
		case 2:
			card.DisplayAll()
		case 0:
			common.ClearScreen()
			fmt.Println("See you soon")
			os.Exit(0)
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}
