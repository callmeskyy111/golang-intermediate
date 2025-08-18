package main

import (
	"bufio"
	"fmt"
	"os"
)

// 1. opening the file
// 2. reading the content of the opened file (in []byte slice)

func main() {
	// open file
	file, err := os.Open("./output.txt")
	if err!= nil{
		fmt.Println("🔴ERROR opening file:",err)
		return
	}

	defer func () {
		fmt.Println("Closing opened file ")
		file.Close()
	}()

	fmt.Println("File opened successfully ✅")

	// read the content
	data:=make([]byte,1024) // buffer to read data into
	_,err=file.Read(data)
	if err!= nil{
		fmt.Println("🔴ERROR reading content:",err)
		return
	}
		fmt.Println("File Content:",string(data))
	
	// read a file line-by-line	

	// reset cursor before scanning
	file.Seek(0, 0)

	// create SCANNER
	scanner:=bufio.NewScanner(file)
	for scanner.Scan(){
		line:=scanner.Text()
		fmt.Println("Line:",line)
	}

	err = scanner.Err() // checking errors
	if err!= nil{
		fmt.Println("🔴ERROR reading content:",err)
		return
	}
	
}
