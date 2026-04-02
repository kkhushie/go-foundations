package main

import (
	"fmt"
	"io"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	var url string

	if id != "" {
		url = "https://fakestoreapi.com/products/" + id
	} else {
		url = "https://fakestoreapi.com/products"
	}

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Fprintln(w, string(body))
}