package main

import (
	"fmt"
	"myprofile/profile"
)

func main() {
	fmt.Println("User Name:", profile.Name)

	profile.Projects = append(profile.Projects, "Impact Leaders")

	profile.GetSummary()
}