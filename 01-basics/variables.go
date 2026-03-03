package main

import "fmt"

func main() {
    var username string = "khush"
    
    var age = 25 
    
    isLearning := true 

    const Version = "1.0.0"

    fmt.Printf("User: %s, Age: %d, Learning: %v\n", username, age, isLearning)
    fmt.Println("App Version:", Version)
}