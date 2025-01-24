package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < -1; i+=10 {
		fmt.Println(i)
		sum += i
	}
	fmt.Println(sum)
}