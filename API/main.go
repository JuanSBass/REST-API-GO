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

	http.ListenAndServe(":3000", r)
}