package main

import "fmt"

// defining a struct
type Student struct {
	Name string
	Age  int
	City string
}

func main() {

	// creating struct instance
	s1 := Student{
		Name: "Anoop",
		Age:  23,
		City: "Surat",
	}

	fmt.Println("Student:", s1)

	// accessing fields
	fmt.Println("Name:", s1.Name)
	fmt.Println("Age:", s1.Age)

	// modifying field
	s1.Age = 24
	fmt.Println("Updated Age:", s1.Age)

}