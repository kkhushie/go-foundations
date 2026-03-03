package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter your full name: ")

    // Read until the newline character
    input, _ := reader.ReadString('\n')

    // Clean up the newline at the end
    input = strings.TrimSpace(input)

    fmt.Printf("Hello, %s! Welcome to Go programming.\n", input)
}