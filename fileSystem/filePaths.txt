package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

// filepath = offers platform-independence

func main() {
	 relativePath:="./data/file.txt"
	 absolutePath:="/home/user/docs/file.txt"

	// Join paths using filepath.join
	joinedPath := filepath.Join("home","Documents","downloads","file.zip")
	fmt.Println("Joined Path:",joinedPath)

	normalizedPath:= filepath.Clean("./data/../data/file.txt") // normalizes
	fmt.Println("Normalized Path:",normalizedPath)

	//Extraction
	dir, file:=filepath.Split("/home/user/docs/file.txt")
	fmt.Println("File:",file)
	fmt.Println("Dir:",dir)
	fmt.Println(filepath.Base("/home/user/docs/file.txt"))
	fmt.Println(filepath.Base("/home/user/docs"))

	// Absolute path or not
	fmt.Println("Is absolute path?:",filepath.IsAbs(relativePath))
	fmt.Println("Is absolute path?:",filepath.IsAbs(absolutePath))

	fmt.Println(filepath.Ext(file))
	fmt.Println(strings.TrimSuffix(file,filepath.Ext(file)))

	// Find a path. relative to another
	rel,err:=filepath.Rel("a/b","a/b/t/file")
	if err!=nil{
		panic(err)
	}

	fmt.Println(rel)

	relCmplx,err:=filepath.Rel("a/c","a/b/t/file")
	if err!=nil{
		panic(err)
	}
	fmt.Println(relCmplx)

	// Coverting relative path -> absolute path
	absPath, err:=filepath.Abs(relativePath)
	if err!=nil{
		fmt.Println("ERROR:",err)
	}else{
		fmt.Println("Absolute Path:",absPath)
	}
}