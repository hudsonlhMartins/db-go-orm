package main

import (
	"encoding/json"
	"os"

	"github.com/hudsonlhmartins/db-go-orm/ORM_DDL/db"
)

func main() {
	database := db.Db()

	// database.Create(&db.Products{
	// 	ProductName:     "Product 1",
	// 	ProductPrice:    10.5,
	// 	ProductQuantity: 100,
	// 	CategoryID:      1,
	// })

	var product db.Products

	res := database.Preload("Category").Find(&product, 1)
	if res.Error != nil {
		panic(res.Error)
	}

	err := json.NewEncoder(os.Stdout).Encode(product)
	if err != nil {
		panic(err)
	}
}
