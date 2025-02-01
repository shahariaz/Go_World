package main

import "os"

func main() {
	f,err:=os.Open("text.txt")
	if err!=nil{
    panic(err)
  }

	fileInfo,err:=f.Stat()
	if err!=nil{
		panic(err)
  }
	println("file name: ",fileInfo.Name())
}