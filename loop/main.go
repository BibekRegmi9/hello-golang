package main

import "fmt"

func main() {

	// for loop
	for i := 1; i < 10; i++ {
		fmt.Println("Number is: ", i)
	}

	// Infinite loop
	counter := 0
	for {
		fmt.Println("Infinite counter: ", counter)
		counter++

		if counter == 100 {
			break
		}
	}

	names := []string{}

	names = append(names, "bibek", "asmit", "jon")
	for()

}
