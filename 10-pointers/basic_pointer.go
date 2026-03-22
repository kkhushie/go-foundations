package main

import "fmt"

func main(){
	x:=10

	//getting address
	fmt.Println("Value of x:",x)
	fmt.Println("Address of x:",&x)

	//pointer variable
	var ptr *int=&x

	fmt.Println("Value of ptr:",ptr)
	fmt.Println("Value at address ptr:",*ptr)

}