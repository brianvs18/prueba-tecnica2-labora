package project_name

import (
	"bufio"
	"ejercicio-poo/creative-projects-management-system/common"
	"ejercicio-poo/creative-projects-management-system/common/constants"
	"ejercicio-poo/creative-projects-management-system/usecases/queries/get-projects"
	"fmt"
)

func Run(scanner *bufio.Scanner) {
	fmt.Println("Enter project ID or Name: ")
	scanner.Scan()
	searchValue := scanner.Text()
	project, exists := get_projects.FindOne(searchValue)

	if !exists {
		fmt.Println("Project not found")
		return
	}

	fmt.Println("Enter new project name")
	scanner.Scan()
	newName := scanner.Text()
	project.UpdateName(newName)
	fmt.Println("Updating project's name.......")
	common.Sleep(constants.FourSeconds)
	fmt.Println("Project name updated successfully")
}
