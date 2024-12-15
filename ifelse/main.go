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
}
