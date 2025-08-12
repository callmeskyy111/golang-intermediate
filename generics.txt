package main

import "fmt"

// Creating STACK with GENERICs
// Stack - LIFO ✅
type Stack[T any] struct{
	elements []T
}

func (s *Stack[T]) push(element T){
	s.elements=append(s.elements, element)
}

func (s *Stack[T]) pop()(T,bool){
	if len(s.elements) == 0{
		var zero T
		return zero,false
	}
	element:= s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true
}

func (s *Stack[T]) isEmpty() bool{
	return len(s.elements) == 0
}

func (s Stack[T]) printAll(){
	if len(s.elements)==0{
		fmt.Println("The stack is empty ")
		return
	}
	fmt.Print("🟢Stack Elements: ")
	for _,el :=range s.elements{
		fmt.Print(el)
	}
	fmt.Println()
}

func main() {
	x, y := 1, 2
	x1 , y1 := "Skyy","Banerjee"

	x, y = swap(x, y)
	x1, y1 = swap(x1,y1)

	
	fmt.Println(x,y)
	fmt.Println(x1,y1)
	// int stack
	intStack:= Stack[int]{}
	intStack.push(1)
	intStack.push(34)
	intStack.push(18)
	intStack.push(29)
	intStack.push(69)

	println("-------------------------------------")


	intStack.printAll()

	fmt.Println(intStack.pop())
	fmt.Println(intStack.pop())
	intStack.printAll()
	fmt.Println("Is the int stack empty? :",intStack.isEmpty())

	// str stack
	strStack:=Stack[string]{}
	strStack.push("🎃")
	strStack.push("🧛🏻‍♂️")
	strStack.push("📖")
	strStack.push("🍁")
	strStack.push("🕯️")
	strStack.push("☕")

	println("---------------------------------------")

	strStack.printAll()
	strStack.pop()
	fmt.Println(strStack.pop())
	strStack.printAll()
	fmt.Println("Is the str stack empty? :",strStack.isEmpty())



}

func swap[T any](a, b T) (T, T) {
	return b, a
}
