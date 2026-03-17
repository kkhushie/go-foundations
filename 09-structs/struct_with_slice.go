package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {

	// slice of structs
	students := []Student{
		{Name: "Anoop", Age: 23},
		{Name: "Rahul", Age: 22},
		{Name: "Priya", Age: 24},
	}

	for _, student := range students {
		fmt.Println("Name:", student.Name, "Age:", student.Age)
	}

}