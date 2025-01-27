package main

import "fmt"

func changeName ( num * int){
	*num = 5
	fmt.Println(*num)
}
func main() {
	i, j := 42, 2701
	p := &i   // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21 // set i through the pointer
	fmt.Println(i) 
	p = &j
	*p = *p / 37
	fmt.Println(j)
	
 num:=1
 changeName(&num)
 
}