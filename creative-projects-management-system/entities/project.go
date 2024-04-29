package entities

import (
	"ejercicio-poo/creative-projects-management-system/common/constants"
	"fmt"
)

type Project struct {
	Id          string
	Name        string
	Description string
	Members     []string
	Status      constants.StatusType
}

func ProjectBuilder(id string, name string, description string, members []string, status constants.StatusType) *Project {
	return &Project{
		Id:          id,
		Name:        name,
		Description: description,
		Members:     members,
		Status:      status,
	}
}

func (p *Project) AddMember(memberId string) {
	p.Members = append(p.Members, memberId)
}

func (p *Project) UpdateDescription(description string) {
	p.Description = description
}

func (p *Project) UpdateName(name string) {
	p.Name = name
}

func (p *Project) Inactivate() {
	p.Status = constants.INACTIVE
}

func (p *Project) String() string {
	return fmt.Sprintf("Project{Id: %s, Name: %s, Description: %s, Members: %s, Status: %s}", p.Id, p.Name, p.Description, p.Members, p.Status)
}
