package main

import "fmt"

func main() {

	i := 15.5
	str := "Bruce Wayne"
	age := 29

	fmt.Println("")
	fmt.Println("------------ FLOATS ------------")
	fmt.Printf("%v\n", i)     // 15.5 → default format
	fmt.Printf("%#v\n", i)    // 15.5 → Go-syntax representation
	fmt.Printf("%T\n", i)     // float64 → type of the value
	fmt.Printf("%v%%\n", i)   // 15.5% → prints % after the value

	fmt.Println("")
	fmt.Println("------------ STRINGS ------------")
	fmt.Printf("%v\n", str)   // Bruce Wayne → default format
	fmt.Printf("%#v\n", str)  // "Bruce Wayne" → Go-syntax string literal
	fmt.Printf("%T\n", str)   // string → type

	fmt.Println("")
	fmt.Println("------------ INTEGERS ------------")
	fmt.Printf("%b\n", age)   // 11101 → binary
	fmt.Printf("%d\n", age)   // 29 → decimal
	fmt.Printf("%+d\n", age)  // +29 → always show sign
	fmt.Printf("%o\n", age)   // 35 → octal
	fmt.Printf("%O\n", age)   // 0o35 → octal with prefix
	fmt.Printf("%x\n", age)   // 1d → hex lowercase
	fmt.Printf("%#x\n", age)  // 0x1d → hex with prefix
	fmt.Printf("%4d\n", age)  // "  29" → right aligned in width 4
	fmt.Printf("%-4d\n", age) // "29  " → left aligned in width 4
	fmt.Printf("%04d\n", age) // 0029 → padded with zeros

	txt := "City"

	fmt.Println("")
	fmt.Println("------------ MORE STRINGS ------------")
	fmt.Printf("%s\n", txt)   // City → plain string
	fmt.Printf("%q\n", txt)   // "City" → quoted string
	fmt.Printf("%8s\n", txt)  // "    City" → right aligned in width 8
	fmt.Printf("%x\n", txt)   // 43697479 → hex (ASCII values)
	fmt.Printf("% x\n", txt)  // 43 69 74 79 → hex with spaces

	t := true
	f := false

	fmt.Println("")
	fmt.Println("------------ BOOLEANS ------------")
	fmt.Printf("%T\n", t)     // bool → type
	fmt.Printf("%T\n", f)     // bool → type
	fmt.Printf("%v\n", f)     // false → default format

	flt := 9.187

	fmt.Println("")
	fmt.Println("------------ MORE FLOATS ------------")
	fmt.Printf("%e\n", flt)     // 9.187000e+00 → scientific notation
	fmt.Printf("%f\n", flt)     // 9.187000 → fixed decimal
	fmt.Printf("%.2f\n", flt)   // 9.19 → fixed decimal, 2 precision
	fmt.Printf("%6.2f\n", flt)  // "  9.19" → width 6, precision 2
	fmt.Printf("%g\n", flt)     // 9.187 → compact representation
}
