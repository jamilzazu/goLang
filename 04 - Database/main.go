package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func insertNewProduct(db *sql.DB, err error) *Product {
	product := NewProduct("NoteBook", 1899.90)
	err = insertProduct(db, product)
	messageError(err)
	return product
}

func updateNewProduct(db *sql.DB, product *Product, err error) {
	product.Price = 201.0
	err = updateProduct(db, product)
	messageError(err)
}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("insert into  products(id, name, price) values (?, ?, ?)")
	messageError(err)
	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	messageError(err)

	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	messageError(err)
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	messageError(err)

	return nil
}

func OpenConnectionMySql() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/godb")
	messageError(err)
	return db, err
}

func closeConnectionMySql(db *sql.DB) {
	defer db.Close()
}

func messageError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := OpenConnectionMySql()
	product := insertNewProduct(db, err)
	updateNewProduct(db, product, err)
	closeConnectionMySql(db)
}
