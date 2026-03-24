package main

import "fmt"

func main() {

	fmt.Println("Start")

	defer fmt.Println("Deferred Line 1")
	defer fmt.Println("Deferred Line 2")

	fmt.Println("End")

}

