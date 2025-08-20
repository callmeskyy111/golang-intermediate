package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Reusable func to handle the errors
func checkErr(err error){
	if err!=nil {
		panic(err)
	}
}

func main() {
	// os.Mkdir("subdir",0755)
	// checkErr(err)
	// Or
	//checkErr(os.Mkdir("subdir1",0755))

	//defer os.RemoveAll("subdir1") // removal
	os.WriteFile("subdir/file1.txt",[]byte(""),0755)

	// Make multiple directories
	//checkErr((os.MkdirAll("subdir/parent/child",0755)))
	checkErr((os.MkdirAll("subdir/parent/child1",0755)))
	checkErr((os.MkdirAll("subdir/parent/child2",0755)))
	checkErr((os.MkdirAll("subdir/parent/child3",0755)))
	checkErr(os.WriteFile("subdir/parent/file2.txt",[]byte(""),0755))
	checkErr(os.WriteFile("subdir/parent/child/file2.txt",[]byte(""),0755))

	res,err:=os.ReadDir("subdir/parent")
	checkErr(err)
	for _,entry:= range res {
		fmt.Println(entry.Name(), entry.IsDir(), entry.Type())
	}

	checkErr((os.Chdir("subdir/parent/child")))
	res,err=os.ReadDir(".")
	checkErr(err)
	fmt.Println("---------- ****** ----------")
	for _,entry:= range res {
		fmt.Println(entry.Name(), entry.IsDir(), entry.Type())
	}

	// Goin up 1 direcetory/level
	checkErr(os.Chdir("../../.."))
	dir, err:= os.Getwd()
	checkErr(err)
	fmt.Println("dir:",dir)

fmt.Println("########### WALK-DIR #############")	
	// filepath.Walk and filepath.WalkDir
	pathfile:= "subdir/parent/child"
	err= filepath.WalkDir(pathfile,func(path string, d os.DirEntry, err error) error{
		if err!=nil{
			fmt.Println("🔴Error:",err)
		}
	fmt.Println(path)
	return nil
	})
	checkErr(err)

	// removing everything
	checkErr(os.RemoveAll("subdir"))

}