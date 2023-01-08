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

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("insert into  products(id, name, price) values (?, ?, ?)")
	messageError(err)
	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	messageError(err)

	return nil
}

func messageError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/godb")
	messageError(err)
	defer db.Close()
	product := NewProduct("NoteBook", 1899.90)

	err = insertProduct(db, product)
	messageError(err)

}
