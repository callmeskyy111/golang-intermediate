package main

import (
	"fmt"
	"os"
)

// Temporary Files / Directories

// Reusable func to handle the errors
func checkErr(err error){
	if err!=nil {
		panic(err)
	}
}
func main() {
	tempFileName:="temporaryFile"
	tempDirName:="temporaryDir"
	
	tempFile,err:=os.CreateTemp("",tempFileName)
	checkErr(err)

	fmt.Println("Temporary file created✅", tempFile.Name())

	defer tempFile.Close()
	defer os.Remove(tempFile.Name())

	tempDir, err:=os.MkdirTemp("",tempDirName)
	checkErr(err)

	defer os.Remove(tempDir)
	
}