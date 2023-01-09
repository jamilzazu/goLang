package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func messageError(err error) {
	if err != nil {
		panic(err)
	}
}

func PessimisticLock(db *gorm.DB, err error) {
	newTransaction := db.Begin()
	var category Category
	err = newTransaction.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&category, 1).Error
	messageError(err)

	category.Name = "Update Category"
	newTransaction.Debug().Save(&category)
	newTransaction.Commit()
}

func main() {
	db, err := OpenConnectionMySql()
	messageError(err)
	AutoMigrateGORM(db)
	PessimisticLock(db, err)
}
