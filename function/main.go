package main

import "fmt"

func simpleFunction() {
	fmt.Println("Simple functions..")
}

func add(a int, b int) int {
	return a + b
}

func sub(a, b int) (result int) {
	result = a - b
	return
}

func multiply(a, b int) (ans int) {
	ans = a * b
	return ans
}

func main() {
	fmt.Println("Functions...")
	simpleFunction()

	ans := add(5, 5)
	fmt.Println(ans)

	sub := sub(10, 5)
	fmt.Println(sub)
}
