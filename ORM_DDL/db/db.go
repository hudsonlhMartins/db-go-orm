package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Categories struct {
	CategoryName        string
	CategoryDescription string
	Products            []Products `gorm:"foreignKey:CategoryID"`
	gorm.Model
}
type Products struct {
	ProductName     string
	ProductPrice    float64
	ProductQuantity int
	CategoryID      uint
	Category        Categories `gorm:"foreignKey:CategoryID;references:ID"`
	gorm.Model
}

func Db() *gorm.DB {
	dsn := "docker:docker@tcp(localhost:3306)/docker?parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db

}
