package get_projects

import (
	"ejercicio-poo/creative-projects-management-system/common/constants"
	"ejercicio-poo/creative-projects-management-system/entities"
	"ejercicio-poo/creative-projects-management-system/storage"
	"ejercicio-poo/creative-projects-management-system/usecases/queries/get-users"
	"github.com/olekukonko/tablewriter"
	"os"
	"strings"
)

func Run() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Name", "Description", "Members"})
	for _, project := range storage.Projects {
		if project.Status != constants.INACTIVE {
			for _, member := range project.Members {
				user, _ := get_users.FindById(member)
				projectData := []string{project.Id, project.Name, project.Description, user.Name}
				table.Append(projectData)
			}
		}
	}
	table.SetRowLine(true)
	table.SetAutoMergeCellsByColumnIndex([]int{0, 1, 2})
	table.Render()
}

func FindOne(value string) (*entities.Project, bool) {
	for _, project := range storage.Projects {
		if project.Id == value || strings.ToLower(project.Name) == strings.ToLower(value) {
			return project, true
		}
	}
	return &entities.Project{}, false
}
