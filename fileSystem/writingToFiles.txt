package main

import (
	"fmt"
	"os"
)

func main() {
	file, err:=os.Create("./fileSystem/output.txt")
	if err!=nil{
		fmt.Println("🔴 ERROR creating file:",err)
		return
	}
	defer file.Close()

	// write data to file
	data:= []byte("Writing BYTE [ ]SLICE to output.txt\n\n")
	_,err=file.Write(data) // the no. of bytes == not needed
	if err!=nil{
		fmt.Println("🔴 ERROR writing to file:",err)
		return
	}
	fmt.Println("Data written to file successfully✅")

	// Another example - writing "strings" instead of []bytes
	file,err = os.Create("./fileSystem/writeString.txt")
	if err!=nil{
		fmt.Println("🔴 ERROR creating file:",err)
		return
	}
	defer file.Close()

	_,err = file.WriteString("Writing strings to a file using 'WriteString'!\n\n")
	if err!=nil{
		fmt.Println("🔴 ERROR writing strings to file:",err)
		return
	}

	fmt.Println("Strings written to file successfully☑️")


}