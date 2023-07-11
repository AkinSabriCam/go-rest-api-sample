package handlers

import (
	"github.com/gorilla/mux"
)

func InitRouter(router *mux.Router) {

	// Add Product Routings..
	router.HandleFunc("/api/products", GetProducts).Methods("GET")
	router.HandleFunc("/api/products", GetProductsById).Methods("GET")
	router.HandleFunc("/api/products", CreateProduct).Methods("POST")
	router.HandleFunc("/api/products", DeleteProduct).Methods("DELETE")
}
