package main

import "fmt"

// Struct Methods

type Rectangle struct{
	length float64
	width float64
}

// Mathod with value-receiver
func (r Rectangle) Area()float64{
	return r.length*r.width
}

// Method with *ptr receiver
func (r *Rectangle) Scale(factor float64){
	r.length*=factor
	r.width*=factor
}

// Method with embedded-structs
type Shape struct{
	Rectangle
}

type MyInt int
	// Method() on user-defined types
	func (m MyInt) IsPositive()bool{
		return m>0
}


func main() {
	rect:= Rectangle{length: 10, width: 9}
	area:= rect.Area()
	fmt.Println("Area before scaling:",area)
	rect.Scale(2)
	area = rect.Area()
	fmt.Println("Area after scaling:",area)

	num := MyInt(-5)
	num1 := MyInt(12)

	s:=Shape{Rectangle: Rectangle{length: 6, width: 5}}
	fmt.Println("Embedded Struct Area:",s.Area())
	fmt.Println("Embedded Struct Area INDIRECTLY:",s.Rectangle.Area())

	fmt.Println(num.IsPositive())
	fmt.Println(num1.IsPositive())
	fmt.Println(num.welcomeMsg())

}

func (MyInt) welcomeMsg()string{
	return "Welcome to MyInt-Type!"
}
