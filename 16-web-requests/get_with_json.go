package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	var post Post

	err = json.NewDecoder(resp.Body).Decode(&post)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println("Post Title:", post.Title)
	fmt.Println("Post Body:", post.Body)

}