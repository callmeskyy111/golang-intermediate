package main

import (
	"errors"
	"fmt"
)

func main() {
// res,err:=exp(2)
// if err!=nil{
// 	fmt.Println("🔴 ERR:",err)
// 	return
// }
// fmt.Println(res)

// res1,err1:=exp(-11)
// if err1!=nil{
// 	fmt.Println("🔴 ERR:",err1)
// 	return
// }
// fmt.Println(res1)

//data:= []byte{}

// err2:= processErr(data)
// if err2!=nil{
// 	fmt.Println("🔴 ERR:",err2)
// 	return
// }
// fmt.Println("Data Processed Successfully ✅")

// err3:= eProcess()

// err3:= eProcess()
// if err3!=nil{
// 	fmt.Println(err3)
// 	return
//}

if err4:=readData();err4!=nil{
	fmt.Println(err4)
	return
}
fmt.Println("Data Processed Successfully ✅")
}

func exp(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("num cannot be negative")
	}
	return x*x,nil
}

func processErr(data []byte)error{
	if len(data)==0{
		return errors.New("empty-data")
	}

	// Process data
	return nil
}

// Custom Errors
type myErr struct{
	errMessage string
}

func (mE *myErr) Error()string{
	return fmt.Sprintf("🔴Error: %s",mE.errMessage)
}

func eProcess()error{
	return &myErr{"custom Err. Message"}
	// return &myErr{errMessage: "custom Err. Message"} // same $hit
}

func readData()error{
	err:= readConfig()

	if err!=nil{
    return fmt.Errorf("🔴 readData: %w",err)
	}

	return nil
}

func readConfig()error{
	return errors.New("config-error")
}