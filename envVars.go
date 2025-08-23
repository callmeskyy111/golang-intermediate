package main

import (
	"fmt"
	"os"
	"strings"
)

// In GoLang, env vars are accessed through the 'os' package.
// Key-Value pairs

func main() {
	user:= os.Getenv("USER")
	home:= os.Getenv("HOME")

	fmt.Println("User env. var is:",user)
	fmt.Println("Home env. var is:",home)

	// set env. var
	err:= os.Setenv("POKEMON","UMBREON")
	if err!=nil{
		fmt.Println("🔴ERROR setting env. variables:",err)
	}

	fmt.Println("POKEMON env. variable:",os.Getenv("POKEMON"))

	for _,e:=range os.Environ(){
		kvPair:=strings.SplitN(e,"=",2)
		fmt.Println("🟢 Key:",kvPair[0]," 🔵 Value:",kvPair[1])
	}

	// unset env. vars
	err = os.Unsetenv("POKEMON")
	if err!=nil{
		fmt.Println("🔴ERROR unsetting env. variable:",err)
		return
	}

	// confirming existence
	fmt.Println("POKEMON env. variable:",os.Getenv("POKEMON"))

	fmt.Println("--------------- **** strings.SplitN() **** --------------")

	str:= "a=b=c=d=e"
	fmt.Println(strings.SplitN(str, "=",-1)) // [a b c d e]
	fmt.Println(strings.SplitN(str, "=",0))  // []
	fmt.Println(strings.SplitN(str, "=",1))  // [a=b=c=d=e]
	fmt.Println(strings.SplitN(str, "=",2))	 // [a b=c=d=e]
	fmt.Println(strings.SplitN(str, "=",3))	 // [a b c=d=e]
	fmt.Println(strings.SplitN(str, "=",4))	 // [a b c d=e]
	fmt.Println(strings.SplitN(str, "=",5))	 // [a b c d e]
}
