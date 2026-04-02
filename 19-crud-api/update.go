package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	jsonData := []byte(`{
		"title": "Updated Product",
		"price": 120.00
	}`)

	req, _ := http.NewRequest("PUT",
		"https://fakestoreapi.com/products/"+id,
		bytes.NewBuffer(jsonData),
	)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		http.Error(w, "Error updating product", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Fprintln(w, string(body))
}