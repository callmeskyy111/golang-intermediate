package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader:=bufio.NewReader(strings.NewReader("Hello, bufio packagebb!\nHow"))
	
	// Reading byte slice
	data:= make([]byte,20)
	n,err:=reader.Read(data)
	if err!=nil{
		fmt.Println("🔴ERROR reading:",err)
		return
	}
	fmt.Printf("Read %d bytes: %s",n,data[:n]) // start from 0, terminate at n-1

	line,err:=reader.ReadString('\n')

	if err!=nil{
		fmt.Println("🔴ERROR reading string:",err)
		return
	}
	fmt.Println("Read str:",line)

	fmt.Println("-------------------------------------")

	// Now WRITE()
	writer:= bufio.NewWriter(os.Stdout) // Stdout is an acceptable writer ✅

	// Writing byte slice
	data1:=[]byte("Hello, WRITING to the bufio package!\n")
	nn,err:=writer.Write(data1)
	if err!=nil{
		fmt.Println("🔴ERROR writing string:",err)
		return
	}

	fmt.Printf("Wrote %d bytes!\n",nn)

	// To get the output (data1), we have to flush the buffer
	// Any writer doesn't automatically flushes the writer, we have to do it

	// Flushing the buffer, ensuring all data is written to os.Stdout - 
	err=writer.Flush()
	if err!=nil{
		fmt.Println("🔴ERROR flushing writer:",err)
		return
	}

	fmt.Println("-------------------------------------")

	// After writing bytes, writing strs:
	str:="This is a random string!\n"
	nn,err=writer.WriteString(str)
	if err!=nil{
		fmt.Println("🔴ERROR writing string:",err)
		return
	}

	fmt.Printf("Wrote %d bytes of string\n",nn)

	// Flushing
	err=writer.Flush()
	if err!=nil{
		fmt.Println("🔴ERROR flushing str-writer:",err)
		return
	}

	fmt.Println("-------------------------------------")



	
}