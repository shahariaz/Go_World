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
	Processing  = "Processing"
	Shipped  = "Shipped"
	Delivered  = "Delivered"
	Returned  = "Returned"
	Cancelled  = "Cancelled"
	Refunded  = "Refunded"
	PartiallyRefunded  = "PartiallyRefunded"
	AwaitingPayment  = "AwaitingPayment"
	Failed  = "Failed"
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