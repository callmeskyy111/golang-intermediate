package main

import (
	"flag"
	"fmt"
)

func main() {
	// fmt.Println("Command:",os.Args[0])
	// for idx,arg:=range os.Args{
	// 	fmt.Println("Argument",idx, ":",arg)
	// }
	//fmt.Println("Argument 1:",os.Args[1])

	// Define flags 🏳️
	var name string
	var age int
	var male bool

	flag.StringVar(&name, "name", "John Doe","Name of the user")
	flag.IntVar(&age, "age",108,"Age of the user")
	flag.BoolVar(&male, "male", true,"Is user male?")

	flag.Parse()

	fmt.Println("Name:",name)
	fmt.Println("Age:",age)
	fmt.Println("IsMale:",male)
}