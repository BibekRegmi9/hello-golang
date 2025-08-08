package main

import (
	"fmt"
	"time"
)

type order struct {
	id        int
	amount    float32
	status    string
	createdAt time.Time // nanosec
}

func main() {
	newOrder := order{
		id:     1,
		amount: 300,
		status: "paid",
		// createdAt: ,
	}

	newOrder.createdAt = time.Now()

	fmt.Println(newOrder)
	fmt.Println(newOrder.amount)
}
