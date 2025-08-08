package main

import "fmt"

func main() {

	// for loop
	for i := 1; i < 10; i++ {
		fmt.Println("Number is: ", i)
	}

	//range
	for i := range 10 {
		fmt.Println("Range number is: ", i)
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

	// loop in slice/array
	names := []string{}

	names = append(names, "bibek", "asmit", "jon")
	for index, value := range names {
		fmt.Println(index, " ", value)
	}

	data := "bibek"
	for index, char := range data {
		fmt.Printf("Index is %d , and value is %c\n", index, char)
	}

}
