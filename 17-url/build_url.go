package main

import (
	"fmt"
	"net/url"
)

func main() {

	baseURL := "https://example.com/search"

	params := url.Values{}
	params.Add("query", "golang")
	params.Add("page", "1")

	finalURL := baseURL + "?" + params.Encode()

	fmt.Println("Generated URL:")
	fmt.Println(finalURL)

}