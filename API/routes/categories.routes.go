package routes

import (
	"encoding/json"
	"net/http"
	"github.com/JuanSBass/REST-API-GO/db"
	"github.com/JuanSBass/REST-API-GO/routes/models"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	var categories []models.Category
	db.DB.Find(&categories)
	json.NewEncoder(w).Encode(&categories)

}

func PostCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	var category models.Category

	json.NewDecoder(r.Body).Decode(&category)

	createdCategory := db.DB.Create(&category)
	err := createdCategory.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode((&category))

}
