package main

import (
	"fmt"
	"maps"
)

func main() {

	m := make(map[string]string)
	m["name"] = "golang"
	m["age"] = "25"
	fmt.Println(m["name"])
  delete(m,"price")
	fmt.Println(m)
	//for fully clear map
	clear(m)

	k:= map[string]int{"price":40, "age":25}
// first value is the value and 2nd one is a boolean value . if value exists then true and not then false
	_,ok := k["price"]
	if ok {
		fmt.Println("Price is there")
	} else{
		fmt.Println("Price is not there")
	}
	n := make(map[string]string)
// cheaking the 2 map 
isE := maps.Equal(m,n)
fmt.Println(isE)


}