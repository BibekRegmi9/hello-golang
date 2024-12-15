package main

import "fmt"

func sum(nums ...int) int {
	sum := 0

	for _, num := range nums {
		sum = sum + num
		fmt.Println(sum)
	}

	return sum
}

func main() {
	ans := sum(1, 2, 3, 4, 5, 6)
	fmt.Println(ans)
}
