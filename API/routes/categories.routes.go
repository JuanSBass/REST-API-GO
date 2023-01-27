package routes

import "net/http"

func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetCategories"))
}

func PostCategory(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PostCategory"))
}