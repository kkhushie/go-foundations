package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	// Formatting time
	formatted := now.Format("02-01-2006 15:04:05")
	fmt.Println("Formatted Time:", formatted)

	// Sleep for 2 seconds
	fmt.Println("Waiting for 2 seconds...")
	time.Sleep(2 * time.Second)

	fmt.Println("Done!")

}