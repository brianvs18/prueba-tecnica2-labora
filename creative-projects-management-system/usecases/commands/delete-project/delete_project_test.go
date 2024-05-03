package delete_project

import (
	"bufio"
	"bytes"
	"ejercicio-poo/creative-projects-management-system/common/constants"
	"ejercicio-poo/creative-projects-management-system/entities"
	get_projects "ejercicio-poo/creative-projects-management-system/usecases/queries/get-projects"
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {

	input := "ebDV8ycct7Yg0on975MTu"
	expectedProject := entities.ProjectBuilder(input,
		"Digital Frontier",
		"A project to create digital services across the world",
		[]string{"WbvcnLmqvUS3vaHVJe6Jo"},
		constants.ACTIVE)
	scan := bufio.NewScanner(bytes.NewBufferString(input))

	//var output bytes.Buffer
	Run(scan)
	project, _ := get_projects.FindOne(input)

	if reflect.DeepEqual(*project, *expectedProject) {
		t.Errorf("FindOne(%s) \n expected %s \n got %s", input, expectedProject, project)
	}

}
