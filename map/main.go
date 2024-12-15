package main

import "fmt"

func main() {
	// initialize map
	studentGrades := make(map[string]int)

	// adding data
	studentGrades["Bibek"] = 69
	studentGrades["Hari"] = 75
	studentGrades["Binod"] = 13
	studentGrades["Astha"] = 65
	studentGrades["John"] = 67

	fmt.Println("Marks of Bibek is: ", studentGrades["Bibek"])

	// updating
	studentGrades["Bibek"] = 79
	fmt.Println("Marks of Bibek is: ", studentGrades["Bibek"])

	// deleting
	studentGrades["toDelete"] = 100
	fmt.Println(studentGrades)
	delete(studentGrades, "toDelete")
	fmt.Println(studentGrades)

	// looping in map
	for index, value := range studentGrades {
		fmt.Printf("Key is %s and marks is %d\n", index, value)
	}

}
