package main

import "fmt"

func main() {
	day := 0

	switch day {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("Monday")
	case 3:
		fmt.Println("Tuesday")
	default:
		fmt.Println("Not known")
	}

	temperature := 5
	switch {
	case temperature >= 0 && temperature <= 15:
		fmt.Println("Freezing")

	case temperature > 20:
		fmt.Println("Normal")

	case temperature > 25:
		fmt.Println("Hot")
	}

}
