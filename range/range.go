package main

import "fmt"

func main() {
	// range
	nums := []int{2, 4, 8, 16, 32}

	for _, num := range nums {
		fmt.Println(num)
	}

	sum := 0
	for i, no := range nums {
		fmt.Println("Index:", i, "Value:", no)
		sum += no
	}
	fmt.Println("Sum:", sum)

	m := map[string]string{"fname": "Brian", "lname": "lara"}
	for k, v := range m {
		fmt.Println("Key:", k, "Value:", v)
	}

}
