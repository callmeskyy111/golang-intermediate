package main

import (
	"fmt"
	"strconv"
)

// Number Parsing - Process of converting textual representations of numbers into their corresponding numeric vals.

func main() {
	numStr:= "12345"
	res,err:=strconv.Atoi(numStr)
	if err!=nil{
		fmt.Println("🔴ERR. parsing val!:",err)
	}
	fmt.Println(res)
	fmt.Println(res+1) // confirmation.

	numIStr, err:= strconv.ParseInt(numStr,10,32)
	if err!=nil{
		fmt.Println("🔴ERR. parsing val!:",err)
	}

	fmt.Println("Parsed integer:",numIStr)

	// Now ParseFloat()
	floatStr:= "5.88"
	floatVal,err:= strconv.ParseFloat(floatStr,64)
	if err!=nil{
		fmt.Println("🔴ERR. parsing val!:",err)
	}
	fmt.Println("Parsed float:",floatVal)

	// Binary
	binaryStr:="1010"
	decimal,err:=strconv.ParseInt(binaryStr,2,64)
		if err!=nil{
		fmt.Println("🔴ERR. parsing binary-val!:",err)
	}
	fmt.Println("Parsed binary-value:",decimal)

	// Hexadecimal
	hexStr:="FF"
	hex,err:=strconv.ParseInt(hexStr,16,64)
		if err!=nil{
		fmt.Println("🔴ERR. parsing hex-value!:",err)
	}
	fmt.Println("Parsed hex-value:",hex)
}