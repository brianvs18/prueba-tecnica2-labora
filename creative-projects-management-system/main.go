package main

import (
	"bufio"
	"ejercicio-poo/creative-projects-management-system/common"
	"ejercicio-poo/creative-projects-management-system/common/constants"
	"ejercicio-poo/creative-projects-management-system/usecases/commands/assign-member-to-project"
	"ejercicio-poo/creative-projects-management-system/usecases/commands/create-project"
	"ejercicio-poo/creative-projects-management-system/usecases/commands/delete-project"
	"ejercicio-poo/creative-projects-management-system/usecases/commands/update/project-description"
	"ejercicio-poo/creative-projects-management-system/usecases/commands/update/project-name"
	"ejercicio-poo/creative-projects-management-system/usecases/queries/get-projects"
	"ejercicio-poo/creative-projects-management-system/usecases/queries/get-users"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		InitMenu()
		scanner.Scan()
		choiceInput := scanner.Text()
		choice, err := strconv.Atoi(choiceInput)
		if err != nil {
			fmt.Println("Invalid option, please enter a number.")
			continue
		}

		switch choice {
		case 1:
			get_projects.Run()
			common.Sleep(constants.TwoSeconds)
		case 2:
			get_users.Run()
			common.Sleep(constants.TwoSeconds)
		case 3:
			create_project.Run(scanner)
			common.Sleep(constants.OneSecond)
		case 4:
			assign_member_to_project.Run(scanner)
			common.Sleep(constants.OneSecond)
		case 5:
			RunUpdate(scanner)
		case 6:
			delete_project.Run(scanner)
		case 7:
			fmt.Println("See you later")
			common.Sleep(constants.OneSecond)
			os.Exit(0)
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}

func RunUpdate(scanner *bufio.Scanner) {
	for {
		updateChoice, e := InitUpdateMenu(scanner)
		if e != nil {
			fmt.Println("Invalid option, please enter a number.")
			continue
		}
		switch updateChoice {
		case 1:
			project_description.Run(scanner)
		case 2:
			project_name.Run(scanner)
		case 3:
			fmt.Println("See you later")
			common.Sleep(constants.OneSecond)
			return
		}
	}
}

func InitUpdateMenu(scanner *bufio.Scanner) (int, error) {
	fmt.Println("What do you want to update?")
	fmt.Println("1. Project Description")
	fmt.Println("2. Project Name")
	fmt.Println("3. Return to main menu")
	scanner.Scan()
	updateChoiceInput := scanner.Text()
	updateChoice, e := strconv.Atoi(updateChoiceInput)
	return updateChoice, e
}

func InitMenu() {
	fmt.Println("Welcome to the Creative Projects Management System")
	fmt.Println("1. Display projects")
	fmt.Println("2. Display team members")
	fmt.Println("3. Create project")
	fmt.Println("4. Add member to project")
	fmt.Println("5. Update project")
	fmt.Println("6. Delete project")
	fmt.Println("7. Exit")
	fmt.Println("Please select one option: ")
}
