package main

import (
	"fmt"
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

func OpenConnectionMySql() (*gorm.DB, error) {
	dsn := "root:root@tcp(localhost:3306)/godb"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	messageError(err)

	AutoMigrateGORM(db)
	return db, err
}

func AutoMigrateGORM(db *gorm.DB) {
	db.AutoMigrate(&Product{})
	//AutoMigrateCreateProduct(db)
}

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

func selectFirstProductById(db *gorm.DB, productId int) {
	var product Product
	db.First(&product, productId)
	fmt.Println(product)
}

func selectFirstProductByName(db *gorm.DB, productName string) {
	var product Product
	db.First(&product, "name = ?", productName)
	fmt.Println(product)
}

func selectAllProducts(db *gorm.DB) {
	var products []Product
	db.Find(&products)

	for _, product := range products {
		fmt.Println(product)
	}
}

func messageError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := OpenConnectionMySql()
	messageError(err)
	selectFirstProductById(db, 1)
	selectFirstProductByName(db, "Monitor")
	selectAllProducts(db)
}
