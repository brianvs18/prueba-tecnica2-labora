package get_users

import (
	"ejercicio-poo/creative-projects-management-system/entities"
	"testing"
)

func TestFindByIdWhenUserExistsReturnUser(t *testing.T) {
	existingUserId := "WbvcnLmqvUS3vaHVJe6Jo"
	expectedUser := entities.UserBuilder(existingUserId, "John Doe", "john@email.com")
	user, _ := FindById(existingUserId)

	if *user != *expectedUser {
		t.Errorf("FindById(%s) \n expected %s \n got %s", existingUserId, expectedUser, user)
	}
}

func TestFindByIdWhenUserNotExistsReturnEmptySlice(t *testing.T) {
	existingUserId := "InvalidUserId"
	expectedUser := &entities.User{}
	user, _ := FindById(existingUserId)

	if *user != *expectedUser {
		t.Errorf("FindById(%s) \n expected %s \n got %s", existingUserId, expectedUser, user)
	}
}
