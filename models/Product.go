package models

import "web-application-golang/db"

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
