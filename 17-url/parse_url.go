package main

import (
	"fmt"
	"net/url"
)

func main() {

	rawURL := "https://example.com:8080/path?name=khushie&age=20"

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Path:", parsedURL.Path)

	// Query parameters
	queryParams := parsedURL.Query()
	fmt.Println("Name:", queryParams.Get("name"))
	fmt.Println("Age:", queryParams.Get("age"))

}