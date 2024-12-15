package main

import "fmt"

func main() {
	var name [5]string
	name[0] = "bibek"
	name[1] = "regmi"
	name[2] = "is"
	name[3] = "don"

	var nos [5]int
	nos[0] = 23
	nos[1] = 34
	nos[2] = 65
	nos[3] = 543

	var locations = [5]int{1, 2, 3, 4, 5}

	fmt.Println(name, nos)

	fmt.Println(locations)

	fmt.Println("Length is: ", len(locations))

	fmt.Println("my first name is: ", name[0])
}
