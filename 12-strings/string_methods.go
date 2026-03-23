package main

import (
	"fmt"
	"strings"
)

func main() {

	text := "  GoLang Learning  "

	fmt.Println("Original:", text)

	// 1. Length
	fmt.Println("Length:", len(text))

	// 2. To Upper
	fmt.Println("Upper:", strings.ToUpper(text))

	// 3. To Lower
	fmt.Println("Lower:", strings.ToLower(text))

	// 4. Trim spaces
	trimmed := strings.TrimSpace(text)
	fmt.Println("Trimmed:", trimmed)

	// 5. Contains
	fmt.Println("Contains 'Go':", strings.Contains(text, "Go"))

	// 6. Replace
	fmt.Println("Replace:", strings.Replace(text, "GoLang", "Golang", 1))

	// 7. Split
	words := strings.Split(trimmed, " ")
	fmt.Println("Split:", words)

	// 8. Join
	joined := strings.Join(words, "-")
	fmt.Println("Join:", joined)

	// 9. HasPrefix
	fmt.Println("HasPrefix 'Go':", strings.HasPrefix(trimmed, "Go"))

	// 10. HasSuffix
	fmt.Println("HasSuffix 'ing':", strings.HasSuffix(trimmed, "ing"))

}

