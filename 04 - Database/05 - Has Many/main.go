package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func OpenConnectionMySql() (*gorm.DB, error) {
	dsn := "root:root@tcp(localhost:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	messageError(err)
	return db, err
}

func AutoMigrateGORM(db *gorm.DB) {
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})
	category := CreateCategoryElectronicsMySql(db)
	product := CreateProductMouseMySql(db, category.ID)
	CreateProductSerialNumberInElectronicsCategory(db, product.ID)
}

func CreateCategoryElectronicsMySql(db *gorm.DB) Category {
	category := Category{Name: "Electronics"}
	db.Create(&category)
	return category
}

func CreateProductMouseMySql(db *gorm.DB, categoryId int) Product {
	product := Product{
		Name:       "Mouse",
		Price:      340.00,
		CategoryID: categoryId,
	}

	db.Create(&product)
	return product
}

func CreateProductSerialNumberInElectronicsCategory(db *gorm.DB, productId int) {
	serialNumber := SerialNumber{Number: "123456", ProductID: productId}
	db.Create(&serialNumber)
}

func selectAllProductsFromACategory(db *gorm.DB) {
	var categories []Category
	err := db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	messageError(err)

	for _, category := range categories {
		fmt.Println(category.Name, ":")

		for _, product := range category.Products {
			fmt.Println("-", product.Name, "SerialNumber:", product.SerialNumber.Number)
		}
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
	selectAllProductsFromACategory(db)
}
