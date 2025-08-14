package main

import "fmt"

func main() {
	num:= 69788
	fmt.Printf("%05d\n",num)

	message:="Halloween 🎃 erdtfyguhijofgsdfbn"
	fmt.Printf("|%10s|\n",message)
	fmt.Printf("|%-10s|\n",message)

	message1:="Thank \nYou"
	message2:=`Thank \nYou`

	fmt.Println(message1)
	fmt.Println(message2)
}