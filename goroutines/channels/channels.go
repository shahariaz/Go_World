package main

import (
	"fmt"
	"math/rand"
	"time"
)

func processNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("processing number: ", num)
		time.Sleep(time.Second)
	}
	
}
func main() {
	numChan := make(chan int)
	go processNum(numChan)
	numChan <- 10
  for{
		numChan <- rand.Intn(100)
	}

	// messageChan := make(chan string)
	// messageChan <- "ping"
	// msg := <-messageChan
	// println(msg)

}
