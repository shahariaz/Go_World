package main

import "os"

func main() {
 //create file
 f,err := os.Create("example.txt")
 if err != nil {
	 panic(err)
	  }
defer f.Close()

f.WriteString("hello world\n")


	// f,err:=os.Open("text.txt")
	// if err!=nil{
  //   panic(err)
  // }

	// fileInfo,err:=f.Stat()
	// if err!=nil{
	// 	panic(err)
  // }
	// println("file name: ",fileInfo.Name())
	
	//Reading a file
	//  f,err:=os.Open("text.txt")
  //   if err!=nil{
	// 		panic(err)
	// 	}
	// 	defer f.Close()
	// 	buf:=make([]byte,1000)
	// 	d,err := f.Read(buf)
	// 	if err != nil{
	// 		panic(err)
	// 	}
	// 	println("bytes read: ",d)
	// 	println("buffer: ",buf)
	// 	println("content: ",string(buf))
	}