package main

import (
	"fmt"
	"log"
	"main/internal/db/postgres"
	"main/internal/service/product"
	"main/internal/service/transaction"
	"os"
)

func main() {

	connectionString := os.Getenv("POSTGRES_CONNECTION_URL")
	dbConnection := postgres.NewPostgresConnection(connectionString)
	defer dbConnection.Close()

	txService := transaction.NewPostgresTransactionService(dbConnection)

	productService := product.NewProductService(txService)

	colors := []string{"white", "black", "red", "blue", "green", "pink", "cyan", "yellow", "grey", "gold", "biege"}

	for index, color := range colors {
		productName := fmt.Sprintf("product_%d", index)
		product, err := productService.CreateProduct(productName, map[string]string{
			"color":  color,
			"height": "10m",
		})

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(product)
	}

}
