package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"`
	gorm.Model
}

func OpenConnectionMySql() (*gorm.DB, error) {
	dsn := "root:root@tcp(localhost:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	messageError(err)
	return db, err
}

func AutoMigrateGORM(db *gorm.DB) {
	db.AutoMigrate(&Product{}, &Category{})
	categories := CreateCategoriesMySql(db)
	CreateProductCardiacMonitorMySql(db, categories)
}

func CreateCategoriesMySql(db *gorm.DB) []Category {
	categories := []Category{
		{
			Name: "Electronics",
		},
		{
			Name: "Heath",
		},
	}
	db.Create(categories)

	return categories
}

func CreateProductCardiacMonitorMySql(db *gorm.DB, categories []Category) Product {
	product := Product{
		Name:       "Cardiac monitor",
		Price:      340.00,
		Categories: categories,
	}

	db.Create(&product)
	return product
}

func selectAllProductsFromACategory(db *gorm.DB) {
	var categories []Category
	err := db.Model(&Category{}).Preload("Products").Find(&categories).Error
	messageError(err)

	for _, category := range categories {
		fmt.Println(category.Name, ":")

		for _, product := range category.Products {
			fmt.Println("-", product.Name)
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
