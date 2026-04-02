package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/products",GetProducts);
	http.HandleFunc("/create", CreateProduct)
	http.HandleFunc("/update", UpdateProduct)
	http.HandleFunc("/delete", DeleteProduct)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080",nil)
}