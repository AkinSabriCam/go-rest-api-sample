package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"go-rest-api-sample/pkg/cache"
	"go-rest-api-sample/pkg/models"
	"go-rest-api-sample/pkg/repositories"
	"io"
	"net/http"
)

var redisCache = cache.NewRedis()

func GetProducts(writer http.ResponseWriter, request *http.Request) {
	data := redisCache.Get()
	writer.Write(data)
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

	repositories.ProductRepositoryInstant.Create(product)
	UpdateRedis(err)

	writer.WriteHeader(200)

}

func UpdateRedis(err error) {
	products := repositories.ProductRepositoryInstant.GetAll()
	productByArray, err := json.Marshal(products)
	if err != nil {
		panic("Could not serialize product array after creating" + err.Error())
	}
	redisCache.Put(productByArray)
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
