package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file")
		return
	}

	// defer used to ensure file closes
	defer file.Close()

	fmt.Println("File created successfully")

}