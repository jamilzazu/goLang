package main

import (
	"database/sql"
	"fmt"
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

func selectAllProducts(db *sql.DB, err error) {
	products, err := selectAllProductsMySql(db)
	messageError(err)
	for _, product := range products {
		fmt.Printf("Id:%v, Product: %v, has the price of %.2f \n", product.ID, product.Name, product.Price)
	}
}

func selectProductById(db *sql.DB, productId string, err error) {
	productSelected, err := selectProductByIdMySql(db, productId)
	messageError(err)
	fmt.Printf("Product: %v, has the price of %.2f", productSelected.Name, productSelected.Price)
}

func insertNewProduct(db *sql.DB, err error) *Product {
	product := NewProduct("NoteBook", 1899.90)
	err = insertProductMySql(db, product)
	messageError(err)
	return product
}

func updateNewProduct(db *sql.DB, product *Product, err error) {
	product.Price = 201.0
	err = updateProductMySql(db, product)
	messageError(err)
}

func selectAllProductsMySql(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("select id, name, price from products")
	messageError(err)
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var product Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		messageError(err)

		products = append(products, product)
	}

	return products, nil
}

func selectProductByIdMySql(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ?")
	messageError(err)
	defer stmt.Close()

	var product Product
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)
	messageError(err)

	return &product, nil
}

func insertProductMySql(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("insert into  products(id, name, price) values (?, ?, ?)")
	messageError(err)
	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	messageError(err)
	return nil
}

func updateProductMySql(db *sql.DB, product *Product) error {
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
	selectProductById(db, product.ID, err)
	selectAllProducts(db, err)
	closeConnectionMySql(db)
}
