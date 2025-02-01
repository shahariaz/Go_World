package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

 // read and write file(streaming fashion)

 sourceFile,err:= os.Open("text.txt")
 if err!=nil{
	 panic(err)
	  }
defer sourceFile.Close()

 destinationFile,err:=os.Create("copy_text.txt")
 if err!=nil{
	 panic(err)
	  }
	defer destinationFile.Close()	
	reader := bufio.NewReader(sourceFile)
	writer:=bufio.NewWriter(destinationFile)
	for{
		b,err:=reader.ReadByte()
		if err!=nil{
			if err.Error() !="EOF"{
	  panic(err)
			}
			break
			
		
		}
		e:=writer.WriteByte(b)
		if e!=nil{
			panic(e)
		}


	}
	writer.Flush()
	fmt.Println("file copied")


 //create file
//  f,err := os.Create("example.txt")
//  if err != nil {
// 	 panic(err)
// 	  }
// defer f.Close()

// f.WriteString("hello world\n")


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