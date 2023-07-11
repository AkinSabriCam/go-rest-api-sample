package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"go-rest-api-sample/pkg/models"
	"go-rest-api-sample/pkg/repositories"
	"io"
	"net/http"
)

func GetProducts(writer http.ResponseWriter, request *http.Request) {
	data := repositories.ProductRepositoryInstant.GetAll()

	content, err := json.Marshal(data)
	if err != nil {
		panic("Could not serialize data to byte array " + err.Error())
	}
	writer.Write(content)
	writer.WriteHeader(200)
}

func GetProductsById(writer http.ResponseWriter, request *http.Request) {

	var idString = request.FormValue("id")
	id, err := uuid.Parse(idString)

	if err != nil {
		panic("Could not serialize data to byte array " + err.Error())
	}

	data := repositories.ProductRepositoryInstant.GetById(id)
	content, err := json.Marshal(data)

	if err != nil {
		panic("Could not serialize data to byte array " + err.Error())
	}
	writer.Write(content)
	writer.WriteHeader(200)
}

func CreateProduct(writer http.ResponseWriter, request *http.Request) {

	requestContent, err := io.ReadAll(request.Body)

	if err != nil {
		panic("Could not read request.Body " + err.Error())
	}

	var product models.Product
	err = json.Unmarshal(requestContent, &product)
	if err != nil {
		panic("Could not serialize data to byte array " + err.Error())
	}

	fmt.Println(product)
	repositories.ProductRepositoryInstant.Create(product)
	writer.WriteHeader(200)

}

func DeleteProduct(writer http.ResponseWriter, request *http.Request) {
	var idString = request.FormValue("id")
	id, err := uuid.Parse(idString)

	if err != nil {
		panic("Could not serialize data to byte array " + err.Error())
	}

	repositories.ProductRepositoryInstant.Delete(id)
	writer.WriteHeader(200)
}
