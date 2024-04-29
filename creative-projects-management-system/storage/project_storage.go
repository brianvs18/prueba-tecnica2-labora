package storage

import (
	"ejercicio-poo/creative-projects-management-system/common/constants"
	"ejercicio-poo/creative-projects-management-system/entities"
)

var (
	Projects = []*entities.Project{
		{
			Id:          "ebDV8ycct7Yg0on975MTu",
			Name:        "Digital Frontier",
			Description: "A project to create digital services across the world",
			Members:     []string{"WbvcnLmqvUS3vaHVJe6Jo"},
			Status:      constants.ACTIVE,
		},
		{
			Id:          "69wrzyxs9usdfwDnU9yRf",
			Name:        "Sustain Sphere",
			Description: "A project to promote environmental projects across the world",
			Members:     []string{"D27NzTS8JQlzHwvLIZABv"},
			Status:      constants.ACTIVE,
		},
		{
			Id:          "ODayFm9R8Z9WERcgG4SkS",
			Name:        "Fit Forge",
			Description: "A project to impulse fitness across the world",
			Members:     []string{"W07NzREWJQlzHwvKeApKo", "D27NzTS8JQlzHwvLIZABv"},
			Status:      constants.ACTIVE,
		},
	}
)
