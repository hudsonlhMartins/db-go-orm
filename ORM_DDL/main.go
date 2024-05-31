package main

import (
	"github.com/hudsonlhmartins/db-go-orm/ORM_DDL/db"
)

func main() {
	database := db.Db()
	database.AutoMigrate(&db.Categories{}, &db.Products{})
}
