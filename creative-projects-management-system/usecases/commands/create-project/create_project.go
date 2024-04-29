package create_project

import (
	"bufio"
	"ejercicio-poo/creative-projects-management-system/common"
	"ejercicio-poo/creative-projects-management-system/common/constants"
	"ejercicio-poo/creative-projects-management-system/entities"
	"ejercicio-poo/creative-projects-management-system/storage"
	"ejercicio-poo/creative-projects-management-system/usecases/queries/get-users"
	"fmt"
	"github.com/matoous/go-nanoid/v2"
	"strconv"
)

func Run(scanner *bufio.Scanner) {
	name, description, membersAmount, wrongInput := captureInputs(scanner)
	if wrongInput {
		return
	}
	var members []string
	for i := 0; i < membersAmount; i++ {
		for {
			fmt.Printf("Add member %d ID: ", i+1)
			scanner.Scan()
			memberId := scanner.Text()
			user, exists := get_users.FindById(memberId)

			if exists {
				members = append(members, user.Id)
				break
			} else {
				fmt.Println("Member not found, please try again.")
			}
		}
	}
	fmt.Println("Creating project........")
	common.Sleep(constants.FourSeconds)
	newProject := entities.ProjectBuilder(gonanoid.Must(), name, description, members, constants.ACTIVE)
	storage.Projects = append(storage.Projects, newProject)
	fmt.Println("Project created successfully")
}

func captureInputs(scanner *bufio.Scanner) (string, string, int, bool) {
	fmt.Println("Enter project name: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("Enter project description: ")
	scanner.Scan()
	description := scanner.Text()

	fmt.Println("How many team members do you want to add? ")
	scanner.Scan()
	membersAmountInput := scanner.Text()
	membersAmount, err := strconv.Atoi(membersAmountInput)
	if err != nil {
		fmt.Println("Invalid input, please enter a valid number.")
		return "", "", 0, true
	}
	return name, description, membersAmount, false
}
