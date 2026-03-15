package main

import "fmt"

func main() {

	fmt.Println("---- Creating a Map ----")

	// map[keyType]valueType
	studentMarks := map[string]int{
		"Anoop":  90,
		"Rahul":  85,
		"Priya":  92,
	}

	fmt.Println("Student Marks:", studentMarks)

	fmt.Println("\n---- Accessing Values ----")

	fmt.Println("Marks of Anoop:", studentMarks["Anoop"])

	fmt.Println("\n---- Adding New Key ----")

	studentMarks["Riya"] = 88
	fmt.Println("Updated Map:", studentMarks)

	fmt.Println("\n---- Deleting a Key ----")

	delete(studentMarks, "Rahul")
	fmt.Println("After Deletion:", studentMarks)

	fmt.Println("\n---- Iterating Over Map ----")

	for key, value := range studentMarks {
		fmt.Println("Name:", key, "Marks:", value)
	}

}