package main

import "fmt"

func main() {
	x := 9
	if x == 10 {
		fmt.Println("ans is 10")
	} else if x < 10 {
		fmt.Println("smaller")
	} else if x > 10 {
		fmt.Println("Large")
	} else {
		fmt.Println("gg")
	}

	age := 25
	if age == 24 {
		fmt.Println("Age is 24")
	} else if age < 24 {
		fmt.Println("Age is less than 24")
	} else {
		fmt.Println("Focus")
	}
}
