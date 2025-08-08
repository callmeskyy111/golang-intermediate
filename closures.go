package main

import "fmt"

// closures in Golang

func main() {
sequence:= adder()

fmt.Println(sequence())
fmt.Println(sequence())
fmt.Println(sequence())
fmt.Println(sequence())

fmt.Println("-------------------------")

sequence2:=adder()
fmt.Println(sequence2())
fmt.Println(sequence2())
fmt.Println(sequence2())
fmt.Println(sequence2())

fmt.Println("-------------------------")

// Another example (inside main f(x))
multiplier:=func() func(int)int{
	product:=2
	return func(x int)int{
		product*=x
		return product
	}
}()

// Using closer double
fmt.Println(multiplier(5))
fmt.Println(multiplier(3))

}

func adder() func() int {
	i := 0
	fmt.Println("Previous value of i:",i)

	return func() int {
	i++
	fmt.Println("Adding 1 to i")
	return i
	}
	
}