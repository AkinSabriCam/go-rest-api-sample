package handlers

import (
	"go-rest-api-sample/pkg/repositories"
	"net/http"
)

type ProductHandler struct {
	repository repositories.ProductRepository
}

func (handler ProductHandler) GetProducts(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("products get ep"))
}

func (handler ProductHandler) GetProductsById(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("products get ep"))
}

func (handler ProductHandler) Create(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("products get ep"))
}

func (handler ProductHandler) Delete(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("products get ep"))
}
