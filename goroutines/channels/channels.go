package main

import (
	"fmt"
	"time"
)

// func processNum(numChan chan int) {
// 	for num := range numChan {
// 		fmt.Println("processing number: ", num)
// 		time.Sleep(time.Second)
// 	}
// func sum(result chan int,num int,num2 int){
// 	res := num + num2
// 	result <- res
// }
// }
func emailSender( emailChan chan string,donechan chan bool){
	defer func ()  {
		donechan <- true
		
	} ()
	for email := range emailChan {
		fmt.Println("sending email to: ", email)
		time.Sleep(time.Second)
	}
}
func main() {

doneChan := make(chan bool)

emailChan := make(chan string, 100)

go emailSender(emailChan,doneChan)
for i:=0;i<8;i++ {
	emailChan <- fmt.Sprintf("%d@gmail.com",i)
}

fmt.Println("done sending email")


//this is important to close the channel
close(emailChan)

<-doneChan

	// resultChan :=make(chan int)
	// go sum(resultChan,10,20)
	// result := <- resultChan
	// println(result)
	// numChan := make(chan int)
	// messageChan := make(chan string)
	// messageChan <- "ping..."
	// msg := <- messageChan
	// fmt.Println(msg)
	// go processNum(numChan)
	// numChan <- 10
  // for{
	// 	numChan <- rand.Intn(100)
	// }

	// messageChan := make(chan string)
	// messageChan <- "ping"
	// msg := <-messageChan
	// println(msg)

}
