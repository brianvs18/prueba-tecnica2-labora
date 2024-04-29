package main

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

type Product struct {
	Id    int
	Name  string
	Stock int
}

func addProduct(productStorage map[string]Product, name string, amount int) {
	for _, product := range productStorage {
		if product.Name == name {
			updatedProduct := Product{Id: product.Id, Name: product.Name, Stock: product.Stock + amount}
			productStorage[product.Name] = updatedProduct
			return
		}
	}

	newProduct := Product{Id: len(productStorage) + 1, Name: name, Stock: amount}
	productStorage[name] = newProduct
	fmt.Println("Product created successfully")
}

func getAllProducts(productStorage map[string]Product) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Name", "Stock"})

	for _, product := range productStorage {
		productData := []string{
			strconv.Itoa(product.Id),
			product.Name,
			strconv.Itoa(product.Stock),
		}
		table.Append(productData)
	}
	table.Render()
}

func updateProduct(productStorage map[string]Product, name string, amount int) {
	product, exists := productStorage[name]
	if exists {
		product.Name = name
		product.Stock = amount
		productStorage[name] = product
	} else {
		fmt.Println("Product does not exists")
	}
}

func deleteProduct(productStorage map[string]Product, name string) {
	_, exists := productStorage[name]
	if exists {
		delete(productStorage, name)
		fmt.Println("Product deleted successfully")
	}
}

//func main() {
//	productStorage := make(map[string]Product)
//
//	for {
//		fmt.Println("1. Crear")
//		fmt.Println("2. Mostrar")
//		fmt.Println("3. Actualizar")
//		fmt.Println("4. Eliminar")
//		fmt.Println("5. Salir")
//		var choice int
//		fmt.Scanln(&choice)
//		switch choice {
//		case 1:
//			var name string
//			var amount int
//			fmt.Println("Nombre")
//			fmt.Scanln(&name)
//			fmt.Println("Cantidad")
//			fmt.Scanln(&amount)
//			addProduct(productStorage, name, amount)
//		case 2:
//			getAllProducts(productStorage)
//		case 3:
//			var name string
//			var amount int
//			fmt.Println("Nombre")
//			fmt.Scanln(&name)
//			fmt.Println("Cantidad")
//			fmt.Scanln(&amount)
//			updateProduct(productStorage, name, amount)
//		case 4:
//			var name string
//			fmt.Println("Nombre")
//			fmt.Scanln(&name)
//			deleteProduct(productStorage, name)
//		case 5:
//			os.Exit(1)
//		default:
//			fmt.Println("Intentelo de nuevo")
//		}
//	}
//}
