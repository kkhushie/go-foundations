package profile

import "fmt"

var Name = "Anoop"
var Projects = []string{"Not Your College", "Trader Raj", "Invex"}
var Interests = []string{"YouTube Design", "Web Development", "Ancient Mythology"}

func GetSummary() {
	fmt.Printf("--- Profile for %s ---\n", Name)
	fmt.Printf("Projects: %v\n", Projects)
	fmt.Printf("Interests: %v\n", Interests)
}