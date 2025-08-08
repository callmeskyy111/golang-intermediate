package main

import "fmt"

// var ptr *int

func main() {
	a := 10
	aPtr := &a

	fmt.Println(a) // 10
	fmt.Println(aPtr) // 0xc00000a0f8
	fmt.Println(*aPtr) // 10 // Dereferencing a *ptr

	// if aPtr != nil{
	// 	fmt.Println("*ptr is not nil:",aPtr)
	// }

	modifyVal(aPtr)
	fmt.Println(a) // 11

}

func modifyVal(ptr *int){
	*ptr++
}