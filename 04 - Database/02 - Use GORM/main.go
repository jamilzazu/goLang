package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

//func NewProduct(name string, price float64) *Product {
//	return &Product{
//		ID:    uuid.New().String(),
//		Name:  name,
//		Price: price,
//	}
//}

func AutoMigrateCreateProduct(db *gorm.DB) {
	products := []Product{
		{
			Name:  "Mouse",
			Price: 340.00,
		},
		{
			Name:  "Monitor",
			Price: 1000.00,
		},
		{
			Name:  "Keyboard",
			Price: 244.50,
		},
	}
	db.Create(products)
}

func AutoMigrateGORM(db *gorm.DB) {
	db.AutoMigrate(&Product{})
	AutoMigrateCreateProduct(db)
}

func OpenConnectionMySql() {
	dsn := "root:root@tcp(localhost:3306)/godb"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	messageError(err)

	AutoMigrateGORM(db)
}

func messageError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	OpenConnectionMySql()
}
