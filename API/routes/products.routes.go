package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JuanSBass/REST-API-GO/db"
	"github.com/JuanSBass/REST-API-GO/routes/models"
	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	var products []models.Product
	db.DB.Find(&products)
	json.NewEncoder(w).Encode((&products))

}

func GetProductsPerCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080") //? Da acceso al cliente
	params := mux.Vars(r)
	var product []models.Product

	fmt.Print(params)

	db.DB.Where("categoryId = ?", params["categoryId"]).Find(&product)

	// if product[] == 0 {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	w.Write([]byte("el producto no existe"))
	// 	return
	// }

	json.NewEncoder(w).Encode((&product))
}

func PostProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	var product models.Product

	json.NewDecoder(r.Body).Decode(&product)

	createdProduct := db.DB.Create(&product)
	err := createdProduct.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode((&product))
}


// func GetProductPerCategory(w http.ResponseWriter, r *http.Request){
// 	var product models.Product

// 	req := mux.Vars(r)
// 	json.NewEncoder(w).Encode(&req)
// }