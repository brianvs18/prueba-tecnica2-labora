package main

import (
	"bufio"
	"ejercicio-poo/challenge/yu-gi-oh/common"
	"ejercicio-poo/challenge/yu-gi-oh/common/constants"
	"ejercicio-poo/challenge/yu-gi-oh/game"
	"fmt"
	"github.com/ttacon/chalk"
	"os"
	"strconv"
)

func Start() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		common.ClearScreen()
		fmt.Println(chalk.Cyan, "1. New Game")
		fmt.Println(chalk.Cyan, "2. Exit")
		scanner.Scan()
		choiceInput := scanner.Text()
		choice, err := strconv.Atoi(choiceInput)
		if err != nil {
			fmt.Println("Invalid option, please enter a number.")
			continue
		}

		switch choice {
		case 1:
			common.Sleep(constants.TwoSeconds)
			game.Start()
		case 2:
			fmt.Println(chalk.Yellow, "See you later...")
			os.Exit(0)
		}
	}
}
