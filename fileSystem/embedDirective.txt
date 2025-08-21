package main

import (
	//_ "embed" // not using directly, just the side-effect

	"embed"
	"fmt"
	"io/fs"
	"log"
)

//go:embed embedExample.txt
var content string

//go:embed embedFolder
var folderContent embed.FS


func main() {
	fmt.Println("Embedded File Content:",content) // file content
	content,err:=folderContent.ReadFile("embedFolder/hello.txt")
	if err!=nil{
		fmt.Println("🔴ERROR. reading file:",err)
		return
	}
	fmt.Println("Embedded Folder Content:",string(content)) // folder content

	err=fs.WalkDir(folderContent,"embedFolder",func(path string, d fs.DirEntry, err error)error{
		if err!=nil{
			fmt.Println("🔴ERR:",err)
			return err
		}
		fmt.Println(path)
		return nil
	})
	if err!=nil{
		log.Fatal(err)
	}
}