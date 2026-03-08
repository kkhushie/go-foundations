package main

import "fmt"

func main() {

	var age int

	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	if age < 13 {
		fmt.Println("You are a child")
	} else if age >= 13 && age < 18 {
		fmt.Println("You are a teenager")
	} else if age >= 18 && age < 60 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a senior citizen")
	}

}