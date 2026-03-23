package main

import (
	"fmt"
	"strconv"
)

func main() {

	// string → int
	str := "123"

	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Conversion error")
	} else {
		fmt.Println("String to Int:", num)
	}

	// int → string
	number := 456
	str2 := strconv.Itoa(number)

	fmt.Println("Int to String:", str2)

	// string -> float
	floatStr := "3.14"
	float,_ := strconv.ParseFloat(floatStr,64)

	fmt.Println("Int to String:", float)

}