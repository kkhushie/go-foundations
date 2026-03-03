package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func main() {
    result := add(42, 13)
    fmt.Println("The sum is:", result)
}