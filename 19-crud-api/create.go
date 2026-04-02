package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	jsonData := []byte(`{
		"title": "New Product",
		"price": 99.99
	}`)

	resp, err := http.Post(
		"https://fakestoreapi.com/products",
		"application/json",
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		http.Error(w, "Error creating product", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Fprintln(w, string(body))
}