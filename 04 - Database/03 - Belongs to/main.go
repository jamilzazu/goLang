package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
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
	category := CreateCategoryElectronicsMySql(db)
	CreateProductMouseMySql(db, category.ID)
}

func CreateCategoryElectronicsMySql(db *gorm.DB) Category {
	category := Category{Name: "Electronics"}
	db.Create(&category)
	return category
}

func CreateProductMouseMySql(db *gorm.DB, categoryId int) {
	product := Product{
		Name:       "Mouse",
		Price:      340.00,
		CategoryID: categoryId,
	}

	db.Create(&product)
}

func selectAllProductsThatHaveTheElectronicsCategory(db *gorm.DB) {
	var products []Product
	db.Preload("Category").Find(&products)

	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name)
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
	selectAllProductsThatHaveTheElectronicsCategory(db)
}
