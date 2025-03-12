package main

import (
	"fmt"
	"time"
)

type order struct{
	id string
	amount int 
	status string
	createdAt time.Time
}

func newOrder (id string, amount int, status string, createdAt time.Time) *order {
	return & order{
		id: id,
		amount: amount,
		status: status,
		createdAt: createdAt,
	}
}

// func (o *order) getStatus() string {
// 	return o.status
// }

// func (o *order) setStatus(status string) {
// 	o.status = status
// }

func main() {
 myOrder := newOrder("123", 100, "pending", time.Now())
 fmt.Println(myOrder)
// fmt.Println(myOrder.amount)
// fmt.Println(myOrder.getStatus())
// myOrder.setStatus("shipped")
// fmt.Println(myOrder.getStatus())
}
