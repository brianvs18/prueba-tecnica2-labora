package get_projects

import (
	"ejercicio-poo/creative-projects-management-system/common/constants"
	"ejercicio-poo/creative-projects-management-system/entities"
	"reflect"
	"testing"
)

func TestFindOneProjectExists(t *testing.T) {
	projectExistsId := "ebDV8ycct7Yg0on975MTu"
	expectedProject := entities.ProjectBuilder(projectExistsId,
		"Digital Frontier",
		"A project to create digital services across the world",
		[]string{"WbvcnLmqvUS3vaHVJe6Jo"},
		constants.ACTIVE)
	project, _ := FindOne(projectExistsId)

	if !reflect.DeepEqual(*project, *expectedProject) {
		t.Errorf("FindOne(%s) \n expected %s \n got %s", projectExistsId, expectedProject, project)
	}
}

func TestFindOneProjectEmpty(t *testing.T) {
	projectExistsId := "ebDV8ycct7Yg0on975RT34"
	emptyProject := entities.Project{}
	project, _ := FindOne(projectExistsId)

	if !reflect.DeepEqual(*project, emptyProject) {
		t.Errorf("FindOne(%s) \n expected %s \n got %s", projectExistsId, emptyProject, project)
	}
}
