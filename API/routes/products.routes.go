package routes

import "net/http"

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetProducts"))
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetProduct"))
}