package main

import (
	"fmt"
	"log"
	"main/internal/db/postgres"
	"main/internal/entities"
	"main/internal/service/db"
	"main/internal/service/product"
	"os"
)

func main() {

	connectionString := os.Getenv("POSTGRES_CONNECTION_URL")
	dbConnection := postgres.NewPostgresConnection(connectionString)
	defer dbConnection.Close()

	dbService := db.NewPostgresDbService(dbConnection)
	productService := product.NewProductService(dbService)

	colors := []string{"white", "black", "red", "blue", "green", "pink", "cyan", "yellow", "grey", "gold", "biege"}
	products := []entities.Product{}

	// Save product within transaction
	for index, color := range colors {
		productName := fmt.Sprintf("product_%d", index)
		product, err := productService.CreateProduct(productName, map[string]string{
			"color":  color,
			"height": "10m",
		})

		if err != nil {
			log.Fatal(err)
		}

		products = append(products, *product)
	}

	// Get some product without transaction
	for _, pr := range products {
		product, err := productService.GetProduct(pr.Id)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("Product: %s with id - %s \n", product.Name, product.Id)
	}
}
