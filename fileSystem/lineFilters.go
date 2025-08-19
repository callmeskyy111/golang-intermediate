package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./fileSystem/output.txt")
	if err!=nil{
		fmt.Println("🔴ERROR. opening file:",err)
		return
	}

	defer file.Close()

	scanner:= bufio.NewScanner(file)

	lineNumber:= 1

	// keyword to filter lines
	keyword:="important"

	// Read and filter lines
	for scanner.Scan(){
		line:=scanner.Text()
		if strings.Contains(line, keyword){
			updatedLine:=strings.ReplaceAll(line,keyword,"necessary")
			fmt.Printf("\n%d Filtered line: %v",lineNumber,line)
			//lineNumber++
			fmt.Printf("\n%d Updated line: %v",lineNumber,updatedLine)
			//lineNumber++
		}
	}

	err = scanner.Err() // always use with scanner
	if err!=nil{
		fmt.Println("🔴ERROR. scanning file:",err)
		return
	}


}