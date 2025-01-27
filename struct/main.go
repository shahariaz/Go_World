package main

import (
	"fmt"
	"time"
)

type Vertex struct {
	x int
	y int
}
type Order struct{
	OrderID string
	Amount float32
	Status string
	CratedAt time.Time // nanosecond precision
}
func NewOrder(id string,amount float32,status string)*Order{
	myorder:= Order{
		OrderID: id,
		Amount: amount,
		Status: status,
	}
	return &myorder
}
// Receiver type
func(o * Order) changeStatus(status string){
	o.Status = status
}
func main() {
	fmt.Println(Vertex{1,2})
	v := Vertex{1,2}

	order := Order{
		OrderID: "1234",
		Amount: 42.27,
		Status: "pending",
    CratedAt: time.Now(), 

	}
	myOrder := NewOrder("1234",42.27,"pending")
	fmt.Println(myOrder)
	order.changeStatus("processing")
	fmt.Println(order)
 fmt.Println(v.x)
	
}