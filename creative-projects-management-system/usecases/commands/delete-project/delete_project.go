package delete_project

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
	fmt.Printf("Inactivating project with ID %s and name %s......\n", project.Id, project.Name)
	common.Sleep(constants.FourSeconds)
	project.Inactivate()
	fmt.Println("Project inactivated successfully")
}
