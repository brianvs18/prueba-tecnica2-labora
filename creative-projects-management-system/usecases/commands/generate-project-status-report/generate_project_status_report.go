package generate_project_status_report

import (
	"bufio"
	get_projects "ejercicio-poo/creative-projects-management-system/usecases/queries/get-projects"
	"fmt"
	"os"
	"time"
)

func Run(scanner *bufio.Scanner) {
	fileName := fmt.Sprintf("creative-projects-management-system/reports/project_registry_%s.txt", time.Now().Format("20060102_150405"))
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error al crear el archivo de registro: %v\n", err)
		return
	}
	defer file.Close()

	fmt.Print("Escriba el ID o nombre del proyecto: ")
	scanner.Scan()
	value := scanner.Text()

	project, exists := get_projects.FindOne(value)
	if !exists {
		fmt.Println("Proyecto no encontrado")
		return
	}
	_, err = file.WriteString(fmt.Sprintf("Report: %v", project))
	if err != nil {
		fmt.Printf("Error al escribir en el archivo de registro: %v\n", err)
		return
	}
	fmt.Printf("Reporte %s creado exitosamente.\n", fileName)
}
