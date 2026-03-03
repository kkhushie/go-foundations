package main

import "fmt"

func swap(x, y string) (string, string) {
    return y, x
}

func main() {
    first := "Hello"
    second := "World"

    a, b := swap(first, second)
    fmt.Println("Swapped:", a, b)

    onlySecond, _ := swap("Left", "Right")
    fmt.Println("Only interested in:", onlySecond)
}