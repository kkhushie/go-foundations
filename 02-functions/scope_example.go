package main

import "fmt"

// Package level scope: accessible by any function in this file
var globalMessage = "I am global"

func displayScope() {
	// Local scope: only exists inside this function
	localSecret := "I am local"
	fmt.Println(globalMessage)
	fmt.Println(localSecret)
}

func main() {
	displayScope()
	fmt.Println("Main can see:", globalMessage)
}
