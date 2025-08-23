package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	strFlag:= flag.String("user","Guest","Name of the user")
	flag.Parse()
	fmt.Println(strFlag)
	
	subCommand1 := flag.NewFlagSet("firstSubC", flag.ExitOnError)
	subCommand2 := flag.NewFlagSet("secondSubC", flag.ExitOnError)

	firstFlag:=subCommand1.Bool("processing",false,"Command processing status")
	secondFlag:=subCommand1.Int("bytes",1024,"Byte [] length of result")

	flagsc2:= subCommand2.String("language","Golang","Enter your language")

	if len(os.Args)<2{
		fmt.Println("This program requires additional commands!")
		os.Exit(1)
	}

	switch os.Args[1]{
	case "firstSubC":
		subCommand1.Parse(os.Args[2:])
		fmt.Println("subCommand1:")
		fmt.Println("processing:",*firstFlag)
		fmt.Println("bytes:",*secondFlag)
	case "secondSubC":	
	subCommand1.Parse(os.Args[2:])
	fmt.Println("subCommand2:")
		fmt.Println("language:",*flagsc2)
	default:
		fmt.Println("No subCommand entered!")
		os.Exit(1)
	}
}

// terminal

// go run clSubCommands.go firstSubC 
// -processing=true -bytes=256
// subCommand1:
// processing: true
// bytes: 256

