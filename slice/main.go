package main

import "fmt"

func main() {
	number := []int{1, 2, 3, 4, 5}
	fmt.Println(number)
	fmt.Printf("data type: %T\n", number)
	fmt.Println("length is: ", len(number))

	number = append(number, 6, 7, 8)
	fmt.Println(number)

	names := []string{}
	names = append(names, "Bibek")
	fmt.Println(names)

	phone := []int{}
	phone = append(phone, 9844878490)
	fmt.Println(phone)

	slice := []int{1, 2, 3, 4, 5}
	slice = append(slice, 6, 7, 8, 9, 10)
	fmt.Println("Slice: ", slice)
	fmt.Println("Length: ", len(slice))
	fmt.Println("Capacity : ", cap(slice))

	fmt.Println("Slice 2.......................................")
	slice2 := make([]int, 3, 5)
	slice2 = append(slice2, 11, 12, 13)
	fmt.Println("Length: ", len(slice2))
	fmt.Println("Capacity : ", cap(slice2))

	fmt.Println("Creating slice")

	fmt.Println("creating slice .......................................")
	arrayList := make([]int, 0)
	fmt.Println("Slice: ", arrayList)
	fmt.Println("Length: ", len(arrayList))
	fmt.Println("Capacity : ", cap(arrayList))

}
