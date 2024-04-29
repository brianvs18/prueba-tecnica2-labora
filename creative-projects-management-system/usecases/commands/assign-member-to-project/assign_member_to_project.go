package assign_member_to_project

import (
	"bufio"
	"ejercicio-poo/creative-projects-management-system/usecases/queries/get-projects"
	"ejercicio-poo/creative-projects-management-system/usecases/queries/get-users"
	"fmt"
)

func Run(scanner *bufio.Scanner) {
	for {
		fmt.Println("Enter project ID: ")
		scanner.Scan()
		projectId := scanner.Text()
		fmt.Println("Enter member ID: ")
		scanner.Scan()
		memberId := scanner.Text()

		project, projectExists := get_projects.FindOne(projectId)
		user, userExists := get_users.FindById(memberId)

		if projectExists && userExists {
			project.AddMember(user.Id)
			fmt.Printf("Member with ID %s added to project successfully\n", memberId)
			break
		} else {
			fmt.Println("Project or Member not found, try again")
		}
	}
}
