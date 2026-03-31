package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	user := User{
		Name:  "Alice",
		Age:   30,
		Email: "khushipal1470@gmail.com",
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println("JSON Output:", string(jsonData))
}