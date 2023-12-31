package main

import (
	"go-web-native/config"
	"go-web-native/controllers/ProductsController"
	"go-web-native/controllers/categoriesController"
	"go-web-native/controllers/homeController"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	//1. homepage
	http.HandleFunc("/", homeController.Welcome)

	//2. categories
	http.HandleFunc("/categories", categoriesController.Index)
	http.HandleFunc("/categories/add", categoriesController.Add)
	http.HandleFunc("/categories/edit", categoriesController.Edit)
	http.HandleFunc("/categories/delete", categoriesController.Delete)

	//2. products
	http.HandleFunc("/products", ProductsController.Index)
	http.HandleFunc("/products/add", ProductsController.Add)
	http.HandleFunc("/products/detail", ProductsController.Detail)
	http.HandleFunc("/products/edit", ProductsController.Edit)
	http.HandleFunc("/products/delete", ProductsController.Delete)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
