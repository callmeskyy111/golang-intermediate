package main

import (
	"errors"
	"fmt"
)

type customErr struct {
	code   int
	errMsg string
	er error
}

func main() {
	err:=doSomething()
	if err !=nil{
		fmt.Println(err)
		return
	}
	fmt.Println("Operation completed successfully ✅")
}

// Error returns the error message. Implementing the Error() method of error interface
func (e *customErr) Error() string {
	return fmt.Sprintf("🔴ERROR %d: %s, %v\n",e.code, e.errMsg,e.er)
}

// Functions that return a custom error
// func doSomething()error{
// 	return &customErr{code:500,errMsg: "Something went wronng!"}
// }

func doSomething()error{
	err:=doSomethingElse()
	if err!=nil{
		return &customErr{
			code:500,errMsg: "Something went wronng!",er:err,
		}
	}
	return nil
}

func doSomethingElse()error{
	return errors.New("internal-error")
}