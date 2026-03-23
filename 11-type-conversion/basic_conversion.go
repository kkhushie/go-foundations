package main

import "fmt"

func main(){
	var a int = 10
	var b float64 = 3.5

	// int → float64
	result1 := float64(a) + b
	fmt.Println("Int to Float:", result1)

	// float64 → int
	result2 := int(b)
	fmt.Println("Float to Int:", result2)

	// int → int64
	var x int = 100
	var y int64 = int64(x)

	fmt.Println("Int to Int64:", y)

}

