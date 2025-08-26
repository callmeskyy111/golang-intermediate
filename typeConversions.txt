package main

import "fmt"

func main() {
	var a int = 32
	b := int32(a)
	c:= float64(b)
	// d:=bool("correct")

	e:=3.14
	f:=int(e)
	fmt.Println(f,c)

	// Syntax - Type(value)

	g:="Hello @ 🔥"
	gByte := []byte(g)
	fmt.Println(string(gByte))

	i:=[]byte{255,255}
	j:=string(i)
	fmt.Println(j)
}