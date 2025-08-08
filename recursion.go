package main

import "fmt"

// recursion - breaking down bigger/complex problems into smaller ones

func main() {
	fmt.Println(factorial(5)) // 120
	fmt.Println(factorial(3)) // 6

	fmt.Println(sumOfDigits(7)) // 7
	fmt.Println(sumOfDigits(234)) // 9

	fibonacci(6) // 0 1 1 2 3 5
	fibonacci(9) // 0 1 1 2 3 5 8 13 21 

}

func factorial(n int)int{
	// Base case: factorial of 0 is 1

	if n==0{
		return 1
	}
	
	// Recursive case: factorial of n is n*factorial(n-1)
	return n*factorial(n-1)
}

// Another example
func sumOfDigits(n int)int{
	// Base case
	if n<10{
		return n
	}

	// Recursive case
	return n%10 + sumOfDigits(n/10)
}

// Function to print Fibonacci series up to 'n' terms
func fibonacci(n int) {
	a, b := 0, 1

	for range n {
		fmt.Print(a, " ")
		a, b = b, a+b // move to next numbers
	}
}