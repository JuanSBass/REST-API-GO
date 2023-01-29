package main

import (
	"net/http"
	"github.com/JuanSBass/REST-API-GO/db"
	"github.com/JuanSBass/REST-API-GO/routes"
	"github.com/JuanSBass/REST-API-GO/routes/models"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConnection()

	db.DB.AutoMigrate(models.Product{})
	db.DB.AutoMigrate(models.Category{})

	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/categories", routes.GetCategories).Methods(("GET"))
	r.HandleFunc("/categories", routes.PostCategory).Methods(("POST"))

	//? Products Routes
	r.HandleFunc("/products/{categoryId}", routes.GetProductsPerCategory).Methods(("GET"))
	r.HandleFunc("/products", routes.GetProducts).Methods(("GET"))
	r.HandleFunc("/products", routes.PostProduct).Methods(("POST"))

	http.ListenAndServe(":3000", r)
}