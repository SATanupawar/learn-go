package main

import "fmt"

type orderStatus string

const (
	received orderStatus = "received"
	processing  = "processing"
	shipped = "shipped"
	delivered = "delivered"
)

func changeStatus(status orderStatus) {
	fmt.Println("Status changed to", status)
}

func main() {
	changeStatus(processing)
}
