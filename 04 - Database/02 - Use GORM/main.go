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
	gorm.Model
}

func OpenConnectionMySql() (*gorm.DB, error) {
	dsn := "root:root@tcp(localhost:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	messageError(err)
	return db, err
}

func AutoMigrateGORM(db *gorm.DB) {
	db.AutoMigrate(&Product{})
	AutoMigrateCreateProduct(db)
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

func updateProductMouse(db *gorm.DB, productId int) {
	var product Product
	db.First(&product, productId)
	product.Name = "New Mouse"
	db.Save(&product)

	selectFirstProductById(db, product.ID)
}

func deleteProductById(db *gorm.DB, productId int) {
	var product Product
	db.First(&product, productId)
	db.Delete(&product, productId)
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

func selectAllProductsWithPagination(db *gorm.DB) {
	var products []Product
	db.Limit(2).Offset(1).Find(&products)

	for _, product := range products {
		fmt.Println(product)
	}
}

func selectProductPriceIsGreaterThanTheValueEntered(db *gorm.DB, productPrice float64) {
	var products []Product
	db.Where("price > ?", productPrice).Find(&products)

	for _, product := range products {
		fmt.Println(product)
	}
}

func selectTheProductIfTheNameIsLike(db *gorm.DB, productName string) {
	var products []Product
	db.Where("name LIKE ?", "%"+productName+"%").Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}
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

	AutoMigrateGORM(db)

	selectFirstProductById(db, 1)
	selectFirstProductByName(db, "Monitor")
	selectAllProducts(db)
	selectAllProductsWithPagination(db)
	selectProductPriceIsGreaterThanTheValueEntered(db, 244.5)
	selectTheProductIfTheNameIsLike(db, "ouse")

	updateProductMouse(db, 1)
	deleteProductById(db, 1)
}
