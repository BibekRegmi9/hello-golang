package main

import "fmt"

func counter() func() int {
	count := 1

	return func() int {
		count++
		return count
	}

}

func main() {
	counter := counter()
	fmt.Println(counter())

}
