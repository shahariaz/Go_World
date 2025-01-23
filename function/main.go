package main

import (
	"fmt"
)

func split(sum int)(x,y int){
	x = sum * 4 / 9
	y = sum - x
	return
}
func add( x int, y int) int {
	return x + y
}
func swap (x,y int) (int, int) {
	return y,x
}

func main(){
	fmt.Println(split(13))
  fmt.Println(add(3,4))
  a, b := swap(5, 10)
  fmt.Println(a, b)
}