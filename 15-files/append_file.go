package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	file.WriteString("Appending new line\n")

	fmt.Println("Data appended successfully")

}