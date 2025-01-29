package main

import "fmt"

// 1st Way of Writing Enum

// type OrderStaus int

// const (
// 	Received OrderStaus = iota
// 	Processing
// 	Shipped
// 	Delivered
// 	Returned
// 	Cancelled
// 	Refunded
// 	PartiallyRefunded
// 	AwaitingPayment
// 	Failed
// 	Declined
// 	Completed
// 	Archived
// 	Unknown
// )

// 2nd Way of Writing Enum

type OrderStaus string

const (
	Received OrderStaus = "Received"
	Processing OrderStaus = "Processing"
	Shipped OrderStaus = "Shipped"
	Delivered OrderStaus = "Delivered"
	Returned OrderStaus = "Returned"
	Cancelled OrderStaus = "Cancelled"
	Refunded OrderStaus = "Refunded"
	PartiallyRefunded OrderStaus = "PartiallyRefunded"
	AwaitingPayment OrderStaus = "AwaitingPayment"
	Failed OrderStaus = "Failed"
	Declined OrderStaus = "Declined"
	Completed OrderStaus = "Completed"
	Archived OrderStaus = "Archived"
	Unknown OrderStaus = "Unknown"
	//... more statuses
)


func changeStatus(status OrderStaus) {
	fmt.Println("Status changed to", status)
}

func main() {
 changeStatus(Received)
}