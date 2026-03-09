package main

import "fmt"

func main() {

	fmt.Println("---- Basic For Loop ----")

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("\n---- Infinite Loop (with break) ----")

	count := 1
	for {
		fmt.Println("Loop count:", count)

		if count == 3 {
			break
		}

		count++
	}

	fmt.Println("\n---- Range Loop with Slice ----")

	numbers := []int{10, 20, 30, 40}

	for index, value := range numbers {
		fmt.Println("Index:", index, "Value:", value)
	}

	fmt.Println("\n---- Range Loop with String ----")

	word := "GoLang"

	for index, char := range word {
		fmt.Printf("Index: %d Character: %c\n", index, char)
	}

}

// git add 07-loops/
// git commit -m "feat: add examples demonstrating for loops, infinite loops, and range iteration in Go"
// git push origin main