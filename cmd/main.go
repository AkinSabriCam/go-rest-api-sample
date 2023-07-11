package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-rest-api-sample/pkg/database"
	"go-rest-api-sample/pkg/handlers"
	"go-rest-api-sample/pkg/repositories"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	db := database.InitDb()
	repositories.NewProductRepository(db)
	handlers.InitRouter(router)

	err := http.ListenAndServe(":4321", router)

	if err != nil {
		fmt.Println(err.Error())
	}
}
