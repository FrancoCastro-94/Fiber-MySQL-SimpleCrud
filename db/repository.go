package db

import (
	"fmt"

	"github.com/francocastro-94/Crud-Fiber/model"
)

func GetAll() []model.Product {
	openDB := dbConnection()
	findAllDB, err := openDB.Query("SELECT * FROM products.all_products ORDER BY id DESC")
	defer openDB.Close()
	if err != nil {
		panic(err.Error())
	}
	var product = model.Product{}
	var allProducts = []model.Product{}
	for findAllDB.Next() {
		var id int
		var name, description, image string
		err = findAllDB.Scan(&id, &name, &description, &image)
		if err != nil {
			panic(err.Error())
		}
		product.ID = id
		product.Name = name
		product.Description = description
		product.Image = image
		allProducts = append(allProducts, product)
	}
	return allProducts
}

func SaveOneProduct(product *model.Product) {
	openDB := dbConnection()
	defer openDB.Close()
	insertOne, err := openDB.Prepare("INSERT INTO products.all_products(name_product, product_description, image) VALUES(? , ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	insertOne.Exec(product.Name, product.Description, product.Image)
}

func DeleteById(id string) {
	openDB := dbConnection()
	defer openDB.Close()
	deleteOne, err := openDB.Prepare("DELETE FROM products.all_products WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	deleteOne.Exec(id)
}

func UpdateOneById(id string, product *model.Product) {
	update := findOneById(id)
	if product.Name != "" {
		update.Name = product.Name
	}
	if product.Image != "" {
		update.Image = product.Image
	}
	if product.Description != "" {
		update.Description = product.Description
	}
	openDB := dbConnection()
	defer openDB.Close()
	updateOne, err := openDB.Prepare("UPDATE products.all_products SET name_product = ?, product_description = ?, image = ?  WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	updateOne.Exec(update.Name, update.Description, update.Image, update.ID)
}

func GetOneById(id string) model.Product {
	product := findOneById(id)
	return product
}

func findOneById(id string) model.Product {
	openDB := dbConnection()
	findOne, err := openDB.Query("SELECT * FROM products.all_products WHERE id=?", id)
	defer openDB.Close()
	if err != nil {
		fmt.Println(err)
	}
	var product model.Product
	for findOne.Next() {
		var productId int
		var name, description, image string
		err = findOne.Scan(&productId, &name, &description, &image)
		if err != nil {
			panic(err.Error())
		}
		product = model.Product{
			ID:          productId,
			Name:        name,
			Description: description,
			Image:       image,
		}

	}
	return product
}
