package main

import (
	"fmt"
	"io"
	"net/http"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	req, _ := http.NewRequest("DELETE",
		"https://fakestoreapi.com/products/"+id,
		nil,
	)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		http.Error(w, "Error deleting product", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Fprintln(w, string(body))
}