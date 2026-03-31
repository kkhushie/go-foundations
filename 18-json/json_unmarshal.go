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

	jsonString := `{"name":"Anoop","age":23,"email":"anoop@example.com"}`

	var user User

	// Convert JSON to struct
	err := json.Unmarshal([]byte(jsonString), &user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Parsed Struct:")
	fmt.Println("Name:", user.Name)
	fmt.Println("Age:", user.Age)
	fmt.Println("Email:", user.Email)

}