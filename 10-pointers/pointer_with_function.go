package main

import "fmt"

// function that modifies value using pointer
func updateValue(x *int) {
	*x = *x + 10
}

func main(){

	x := 20

	fmt.Println("Before:", x)

	updateValue(&x)

	fmt.Println("After:", x)


}