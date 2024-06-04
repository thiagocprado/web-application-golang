package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"web-application-golang/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.GetAllProducts()

	temp.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("nome")
		description := r.FormValue("descricao")
		price := r.FormValue("preco")
		quantity := r.FormValue("quantidade")

		convertedPrice, err := strconv.ParseFloat(price, 64)

		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		convertedQuantity, err := strconv.Atoi(quantity)

		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		models.SaveProduct(name, description, convertedPrice, convertedQuantity)
	}

	http.Redirect(w, r, " /", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.DeleteProduct(id)
	http.Redirect(w, r, "/", 301)
}
