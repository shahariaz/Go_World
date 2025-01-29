package main

import (
	"fmt"
)

func printSlice[T any] (items []T) {

	for _, item := range items {
		fmt.Println(item)
	}

}
//func compPrint[T comparable,V string](items []T) {
	//for _, item := range items {
		//fmt.Println(item)
	//}
//}
func printSliceIntString[T string | int](items []T){
	for _, item := range items{
		fmt.Println(item)
	}
}
type Stack[T any] struct{
	elements []T
}
func main() {
	myStack := Stack[int] {
		elements: []int{1,2,3,4,5},
	}
	fmt.Println(myStack.elements)
  printSlice([]int{1,2,3,4,5})
	printSliceIntString([]string{"Hello", "World", "!"})
}