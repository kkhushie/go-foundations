package main

import (
	"fmt"
	"os"
)

func main() {

	info, err := os.Stat("test.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File Name:", info.Name())
	fmt.Println("File Size:", info.Size(), "bytes")
	fmt.Println("Is Directory:", info.IsDir())

}