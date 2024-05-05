package main

import (
	"ejercicio-poo/challenge/yu-gi-oh/common"
	"ejercicio-poo/challenge/yu-gi-oh/common/constants"
	"fmt"
	"github.com/ttacon/chalk"
)

func main() {
	for {
		var choice string
		printMenu()
		fmt.Scanln(&choice)
		switch choice {
		case "s":
			common.Sleep(constants.TwoSeconds)
			Start()
		default:
			fmt.Println("Wrong choice, try again!")
		}
	}
}

func printMenu() {
	fmt.Println(chalk.Yellow, "Yu-Gi-Oh!")
	fmt.Println(chalk.Yellow, chalk.Bold.TextStyle("Forbidden Memories"))
	fmt.Println()
	fmt.Println(chalk.Green, chalk.Bold.TextStyle("Press start button (s)"))
	fmt.Println()
	fmt.Println()
	fmt.Println(chalk.White, chalk.Bold.TextStyle("1996 KAZUKI TAKAHASHI"))
}
