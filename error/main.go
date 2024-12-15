package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("divide by zero error")
	}
	return a / b, nil
}

func main() {
	// ans, err := divide(10, 0)
	// if err != nil {
	// 	fmt.Println("Error handling")
	// }
	// fmt.Println("division is : ", ans)

	ans, _ := divide(10, 5)
	fmt.Println("divide is: ", ans)
}
