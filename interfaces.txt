package main

import (
	"fmt"
	"math"
)

// BEHAVIVOUR - Polymorphism
// Set of method-signatures that a type must implement/ adhere to

type Geometry interface{
	area() float64
	perim() float64
}

// structs
type Rect struct{
	wd,ht float64
}

type Rect1 struct{
	wd,ht float64
}

type Circle struct{
	rad float64
}

// methods
func (r Rect1) area()float64{
	return r.ht * r.wd 
}

func (r Rect1) perim()float64{
	return r.ht * r.wd 
}

func (r Rect) area()float64{
	return r.ht * r.wd 
}

func (c Circle)area()float64{
	return math.Pi*c.rad*c.rad
}

func (c Circle) perim()float64{
	return 2*math.Pi*c.rad
}

func (r Rect) perim()float64{
	return 2*(r.ht + r.wd)
}

func (c Circle) diameter()float64{
	return 2*c.rad
}

func measure(g Geometry){
fmt.Println(g)
fmt.Println(g.area())
fmt.Println(g.perim())
}

// main f(x)
func main() {
	rect1:= Rect{wd: 3,ht: 4}
	circ1:= Circle{rad:5}
	fmt.Println(circ1.diameter())
	rect2:=Rect1{wd: 7,ht: 8}
	measure(rect1)
	measure(rect2)
	measure(circ1)

	// random vals of interface{}/any type
	myPrinter("myPrinter",45.9,true)

	// random val-types too..
	printValType("john")
	printValType(rect1)
	printValType(false)
}

// More usages with variadic params.
func myPrinter(i ...interface{}){
	for _,val:= range i{
		fmt.Println(val)
	}
}

//! ⚠️ interface{ } == any

func printValType(j any){
	switch j.(type){
	case int:
		fmt.Println("Type: int")
	case string:
		fmt.Println("Type: string")	
	case bool:
		fmt.Println("Type: bool")	
	default:
		fmt.Println("Type: UNKNOWN TYPE")	
	}
}