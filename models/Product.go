package models

import (
	"web-application-golang/db"
)

type Product struct {
	Id          int
	Description string
	Name        string
	Price       float64
	Quantity    int
}

func GetAllProducts() []Product {
	db := db.ConnectDb()

	selectAllProducts, err := db.Query(`
		select * 
		from tb_products
		order by id asc
	`)

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectAllProducts.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}

	defer db.Close()

	return products
}

func SaveProduct(name string, description string, price float64, quantity int) {
	db := db.ConnectDb()

	insertProduct, err := db.Prepare(`
		insert into tb_products(name, description, price, quantity) values (
			$1,
			$2,
			$3,
			$4
		)
	`)

	if err != nil {
		panic(err.Error())
	}

	insertProduct.Exec(name, description, price, quantity)

	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConnectDb()

	deleteProduct, err := db.Prepare(`
		delete  
		from tb_products
		where id=$1
	`)

	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)

	defer db.Close()
}

func GetProductById(id string) Product {
	db := db.ConnectDb()

	selectProduct, err := db.Query(`
		select * 
		from tb_products
		where id = $1
	`, id)

	if err != nil {
		panic(err.Error())
	}

	product := Product{}

	for selectProduct.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectProduct.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity
	}

	defer db.Close()

	return product
}

func UpdateProduct(id int, name string, description string, price float64, quantity int) {
	db := db.ConnectDb()

	updateProduct, err := db.Prepare(`
		update tb_products set name = $1, description = $2, price = $3,	quantity = $4
		where id = $5
	`)

	if err != nil {
		panic(err.Error())
	}

	updateProduct.Exec(name, description, price, quantity, id)

	defer db.Close()
}
