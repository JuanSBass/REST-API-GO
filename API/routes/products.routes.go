package routes

import (
	"encoding/json"
	"net/http"

	"github.com/JuanSBass/REST-API-GO/db"
	"github.com/JuanSBass/REST-API-GO/routes/models"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {

	var products []models.Product
	db.DB.Find(&products)
	json.NewEncoder(w).Encode((&products))

}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetProduct"))
}

func PostProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	json.NewDecoder(r.Body).Decode(&product)

	createdProduct := db.DB.Create(&product)
	err := createdProduct.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode((&product))
}
