package main

import "fmt"

func main() {
    var i int     
    var f float64 // 0
    var b bool    // false
    var s string  // "" (empty string)

    fmt.Println("Zero Values:")
    fmt.Printf("Int: %v, Float: %v, Bool: %v, String: [%v]\n", i, f, b, s)
}