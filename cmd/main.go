package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/products", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("products get ep"))
	}).Methods("GET")

	err := http.ListenAndServe(":4321", router)

	router.Handle(http.Handle())

	if err != nil {
		fmt.Println(err.Error())
	}
}
