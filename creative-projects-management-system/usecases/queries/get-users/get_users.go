package get_users

import (
	"ejercicio-poo/creative-projects-management-system/entities"
	"ejercicio-poo/creative-projects-management-system/storage"
	"github.com/olekukonko/tablewriter"
	"os"
)

func Run() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Name", "Email"})

	for _, user := range storage.Users {
		userData := []string{user.Id, user.Name, user.Email}
		table.Append(userData)
	}
	table.Render()
}

func FindById(userId string) (*entities.User, bool) {
	for _, user := range storage.Users {
		if user.Id == userId {
			return user, true
		}
	}
	return &entities.User{}, false
}
