package main

import "fmt"

func changeNum(num int) {
	num = 5
	fmt.Println("in change num", num)
}

func main() {
	num := 1
	fmt.Println("Memory ", &num)
	changeNum(num)
	fmt.Println("After change num in main", num)
}
