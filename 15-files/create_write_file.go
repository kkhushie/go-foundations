package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	file.WriteString("Hello, Go File Handling!\n")

	fmt.Println("File created and written successfully")

}
