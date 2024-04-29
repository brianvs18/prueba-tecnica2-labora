package main

import "fmt"

type Todo struct {
	Name        string
	Description string
	Owner       string
	Status      string
}

func CreateTodo(name string, description string, owner string) *Todo {
	return &Todo{
		Name:        name,
		Description: description,
		Owner:       owner,
		Status:      "Pending",
	}
}

func UpdateTodoStatus(todo *Todo, status string) {
	todo.Status = status
}

func FindAllPendingTodos(todos []*Todo) {
	for _, todo := range todos {
		if todo.Status == "Pending" {
			fmt.Printf("Name: %s \n, Description: %s \n, Owner: %s \n, Status: %s \n", todo.Name, todo.Description, todo.Owner, todo.Status)
		}
	}
}

func main() {
	//Crear algunas todos de ejemplo
	todoOne := CreateTodo("Implementar funcionalidad de inicio de sesión", "Implementar la funcionalidad de inicio de sesión en la aplicación web", "Juan")
	todoTwo := CreateTodo("Diseñar interfaz de usuario", "Diseñar la interfaz de usuario para la página de perfil del usuario", "María")
	todoThree := CreateTodo("Depurar errores en la base de datos", "Investigar y solucionar errores en la base de datos de producción", "Pedro")

	// Mostrar las todos pendientes
	// Mostrar las todos pendientes
	todos := []*Todo{todoOne, todoTwo, todoThree}
	FindAllPendingTodos(todos)

	// Actualizar el estado de una tarea
	UpdateTodoStatus(todoOne, "en progreso")
	fmt.Println("\nDespués de actualizar el estado de una tarea:")
	FindAllPendingTodos(todos)
}
